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
		},
	}
}

// namespace register
func nodeCreate() cli.Command {
	return cli.Command{
		Name:        "create",
		Usage:       "Use for adding node to k2edge",
		Description: "Use 'node create --name=<name> ...' to create namespace",
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
			err := masterCli.Node.Register(context.Background(), &types.RegisterRequest{
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

			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("create namespace '%s' success", name)
			return nil
		},
	}
}

// namespace list
func nodeList() cli.Command {
	return cli.Command{
		Name:        "list",
		Aliases:     []string{"l"},
		Usage:       "Use for listing node",
		Description: "Use 'namespace list [usable]' to list node",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "usable",
				Usage: "only show good namespace(option)",
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

// namespace top
func nodeTop() cli.Command {
	return cli.Command{
		Name:        "top",
		Usage:       "Use for showing node's top",
		Description: "Use 'namespace top --name=<name>' to show node's top",
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
