package cli

import (
	"context"
	"fmt"
	"k2edge/master/client"
	"k2edge/master/internal/types"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func Node() cli.Command {
	return cli.Command{
		Name:        "node",
		Usage:       "Use for node management",
		Description: "Use 'node <command>' to manage node",
		Before: func(ctx *cli.Context) error {
			etcd := ctx.App.Metadata["config-etcd"].(string)
			server := getServer(etcd)
			ctx.App.Metadata = map[string]interface{}{
				"config-server": server,
				"config-etcd":   etcd,
			}
			return nil
		},
		Subcommands: cli.Commands{
			nodeCreate(),
			nodeList(),
			nodeTop(),
			nodeCordon(),
			nodeUncordon(),
			nodeDrain(),
			nodeDelete(),
		},
	}
}

// node register
func nodeCreate() cli.Command {
	return cli.Command{
		Name:        "create",
		Usage:       "Use for adding node to k2edge",
		Description: "Use 'node create --name=<name> [args...]' to create node",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Usage:    "the name of node, no repetition allowed",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     "roles",
				Usage:    "master/worker OR both, exp: --roles worker --roles master",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "masterurl",
				Usage: "the url of master",
			},
			&cli.StringFlag{
				Name:  "workerurl",
				Usage: "the url of worker",
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)
			name := ctx.String("name")
			roles := ctx.StringSlice("roles")
			masterurl := ctx.String("masterurl")
			workerurl := ctx.String("workerurl")
			err := masterCli.Node.Register(context.Background(), types.RegisterRequest{
				Name:  name,
				Roles: roles,
				BaseURL: types.NodeURL{
					MasterURL: masterurl,
					WorkerURL: workerurl,
				},
			})
			if err != nil {
				return err
			}

			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("create node '%s' successfully", name)
			return nil
		},
	}
}

// node list
func nodeList() cli.Command {
	return cli.Command{
		Name:        "list",
		Aliases:     []string{"ls"},
		Usage:       "Use for listing node",
		Description: "Use 'node list [usable]' to list node",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "usable",
				Usage: "only show good node(option)",
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			resp, err := masterCli.Node.List(context.Background(), types.NodeListRequest{
				All: !ctx.Bool("usable"),
			})
			if err != nil {
				return err
			}

			tableData := [][]string{
				{"name", "age", "status", "roles", "url"},
			}

			sort.Slice(resp.NodeList, func(i, j int) bool {
				return resp.NodeList[i].RegisterTime > resp.NodeList[j].RegisterTime
			})
			for _, n := range resp.NodeList {
				urls := []string{}
				if strings.Contains(n.Roles, "master") {
					urls = append(urls, n.URL.MasterURL)
				}
				if strings.Contains(n.Roles, "worker") {
					urls = append(urls, n.URL.WorkerURL)
				}
				url := strings.Join(urls, " | ")
				age := time.Since(time.Unix(n.RegisterTime, 0)).Round(time.Second).String()
				tableData = append(tableData, []string{n.Name, age, n.Status, n.Roles, url})
			}

			pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
			return nil
		},
	}
}

// node top
func nodeTop() cli.Command {
	return cli.Command{
		Name:        "top",
		Usage:       "Use for showing node's top",
		Description: "Use 'node top --name=<name>' to show node's top",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Usage:    "the name of node",
				Required: true,
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			resp, err := masterCli.Node.Top(context.Background(), types.NodeTopRequest{
				Name: ctx.String("name"),
			})

			if err != nil {
				return err
			}

			info := ""
			image := strings.Join(resp.Images, "\n")
			info += color.BlueString("image:\n") + fmt.Sprintf("%s\n", image)
			info += color.BlueString("CPU%:             ") + fmt.Sprintf("%f\n", resp.CPUUsedPercent)
			info += color.BlueString("CPUUsed:          ") + fmt.Sprintf("%f\n", resp.CPUUsed)
			info += color.BlueString("CPUFree:          ") + fmt.Sprintf("%f\n", resp.CPUFree)
			info += color.BlueString("CPUTotal:         ") + fmt.Sprintf("%f\n", resp.CPUTotal)
			info += color.BlueString("Memory%:          ") + fmt.Sprintf("%f\n", resp.MemoryUsedPercent)
			info += color.BlueString("MemoryUsed:       ") + fmt.Sprintf("%d\n", resp.MemoryUsed)
			info += color.BlueString("MemoryAvailable:  ") + fmt.Sprintf("%d\n", resp.MemoryAvailable)
			info += color.BlueString("MemoryTotal:      ") + fmt.Sprintf("%d\n", resp.MemoryTotal)
			info += color.BlueString("Disk%:            ") + fmt.Sprintf("%f\n", resp.DiskUsedPercent)
			info += color.BlueString("DiskUsed:         ") + fmt.Sprintf("%d\n", resp.DiskUsed)
			info += color.BlueString("DiskFree:         ") + fmt.Sprintf("%d\n", resp.DiskFree)
			info += color.BlueString("DiskTotal:        ") + fmt.Sprintf("%d\n", resp.DiskTotal)

			fmt.Println(info)
			return nil
		},
	}
}

// node cordon
func nodeCordon() cli.Command {
	return cli.Command{
		Name:        "cordon",
		Usage:       "Used to set whether the node is unschedulable",
		Description: "Use 'node cordon --name=<name>' to set node to be unschedulable",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Usage:    "the name of node",
				Required: true,
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			name := ctx.String("name")
			err := masterCli.Node.Cordon(context.Background(), types.CordonRequest{
				Name: name,
			})

			if err != nil {
				return err
			}
			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("set node '%s' unschedulable successfully", name)
			return nil
		},
	}
}

// node cordon
func nodeUncordon() cli.Command {
	return cli.Command{
		Name:        "uncordon",
		Usage:       "Used to set whether the node is schedulable",
		Description: "Use 'node uncordon --name=<name>' to set node to be schedulable",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Usage:    "the name of node",
				Required: true,
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			name := ctx.String("name")
			err := masterCli.Node.Uncordon(context.Background(), types.UncordonRequest{
				Name: name,
			})

			if err != nil {
				return err
			}
			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("set node '%s' schedulable successfully", name)
			return nil
		},
	}
}

// node cordon
func nodeDrain() cli.Command {
	return cli.Command{
		Name:        "drain",
		Usage:       "Used to set whether the node is unschedulable and migrate node",
		Description: "Use 'node drain --name=<name>' to drain node",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Usage:    "the name of node",
				Required: true,
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			name := ctx.String("name")
			err := masterCli.Node.Drain(context.Background(), types.DrainRequest{
				Name: name,
			})

			if err != nil {
				return err
			}
			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("drain node '%s' successfully", name)
			return nil
		},
	}
}

// node delete
func nodeDelete() cli.Command {
	return cli.Command{
		Name:        "delete",
		Usage:       "Used to delete the node",
		Description: "Use 'node delete --name=<name>' to delete node",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Usage:    "the name of node",
				Required: true,
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			name := ctx.String("name")
			err := masterCli.Node.Delete(context.Background(), types.DeleteRequest{
				Name: name,
			})

			if err != nil {
				return err
			}
			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("delete node '%s' successfully", name)
			return nil
		},
	}
}