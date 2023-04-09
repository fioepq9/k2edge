package cli

import (
	"context"
	"fmt"
	"k2edge/master/client"
	"k2edge/master/internal/types"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func Container() cli.Command {
	return cli.Command{
		Name:        "container",
		Usage:       "Use for container management",
		Description: "Use 'container <command>' to manage container",
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
			containerCreate(),
			containerGet(),
			containerList(),
			containerDelete(),
		},
	}
}

// container create
func containerCreate() cli.Command {
	return cli.Command{
		Name:        "create",
		Usage:       "Use for adding container on node",
		Description: "Use 'container create --name=<name> [args...]' to create container",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "namespace",
				Usage:    "the namespace of container",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of container. If not set, the name of the contianer is random",
			},
			&cli.StringFlag{
				Name:     "image",
				Usage:    "the image of container",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "nodeName",
				Usage: "specify the location where the container is created",
			},
			&cli.StringFlag{
				Name:  "command",
				Usage: "the command executed by the container",
			},
			&cli.StringSliceFlag{
				Name:  "args",
				Usage: "command parameters",
			},
			&cli.StringSliceFlag{
				Name:  "expose",
				Usage: "ports exposed by the container, example: port,protocol,hostport",
			},
			&cli.StringSliceFlag{
				Name:  "env",
				Usage: "container environment configuration",
			},
			&cli.StringFlag{
				Name:  "limit",
				Usage: "limited resources, contain CPU and Memory, example: cpu,mempry",
			},
			&cli.StringFlag{
				Name:  "request",
				Usage: "requested resources, contain CPU and Memory, example: cpu,mempry",
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			namespace := ctx.String("namespace")
			name := ctx.String("name")
			image := ctx.String("image")
			nodeName := ctx.String("nodeName")
			command := ctx.String("command")
			args := ctx.StringSlice("args")
			exposeStr := ctx.StringSlice("expose")
			expose := []types.ExposedPort{}
			for _, e := range exposeStr {
				es := strings.Split(e, ",")
				if len(es) != 3 {
					return fmt.Errorf("'%s' is in the wrong format", e)
				}
				port, err := strconv.Atoi(es[0])
				if err != nil {
					return fmt.Errorf("'%s' is in the wrong format", e)
				}
				hostPort, err := strconv.Atoi(es[2])
				if err != nil {
					return fmt.Errorf("'%s' is in the wrong format", e)
				}
				expose = append(expose, types.ExposedPort{
					Port:     int64(port),
					Protocol: es[1],
					HostPort: int64(hostPort),
				})
			}

			env := ctx.StringSlice("env")
			limitStr := ctx.String("limit")
			limit := types.ContainerLimit{}
			if ctx.IsSet("limit") {
				ls := strings.Split(limitStr, ",")
				if len(ls) != 2 {
					return fmt.Errorf("'%s' is in the wrong format", ls)
				}

				cpul, err := strconv.Atoi(ls[0])
				if err != nil {
					return err
				}
				memoryl, err := strconv.Atoi(ls[1])
				if err != nil {
					return err
				}

				limit.CPU = int64(cpul)
				limit.Memory = int64(memoryl)
			}

			requestStr := ctx.String("request")
			request := types.ContainerRequest{}
			if ctx.IsSet("request") {
				rs := strings.Split(requestStr, ",")
				if len(rs) != 2 {
					return fmt.Errorf("'%s' is in the wrong format", rs)
				}

				cpur, err := strconv.Atoi(rs[0])
				if err != nil {
					return err
				}
				memoryr, err := strconv.Atoi(rs[1])
				if err != nil {
					return err
				}

				request.CPU = int64(cpur)
				request.Memory = int64(memoryr)
			}

			err := masterCli.Container.Create(context.Background(), types.CreateContainerRequest{
				Container: types.Container{
					Metadata: types.Metadata{
						Namespace: namespace,
						Kind:      "container",
						Name:      name,
					},
					ContainerConfig: types.ContainerConfig{
						Image:    image,
						NodeName: nodeName,
						Command:  command,
						Args:     args,
						Expose:   expose,
						Env:      env,
						Limit:    limit,
						Request:  request,
					},
				},
			})

			if err != nil {
				return err
			}

			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("create container '%s' successfully", name)
			return nil
		},
	}
}

// container get
func containerGet() cli.Command {
	return cli.Command{
		Name:        "get",
		Usage:       "Use to get container info",
		Description: "Use 'node get --namespace=<namespace> --name=<name>' to get node",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "namespace",
				Usage: "the namespace of container",
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of container.",
			},
			&cli.BoolFlag{
				Name: "detail",
				Usage: "for more detail",
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			namespace := ctx.String("namespace")
			name := ctx.String("name")

			resp, err := masterCli.Container.Get(context.Background(), types.GetContainerRequest{
				Namespace: namespace,
				Name:      name,
			})

			if err != nil {
				return err
			}

			command := resp.Container.ContainerConfig.Command
			command += " " + strings.Join(resp.Container.ContainerConfig.Args, " ")

			info := ""
			info += color.BlueString("namespace:        ") + fmt.Sprintf("%s\n", resp.Container.Metadata.Namespace)
			info += color.BlueString("name:             ") + fmt.Sprintf("%s\n", resp.Container.Metadata.Name)
			info += color.CyanString("info:\n")
			info += color.BlueString("status:           ") + fmt.Sprintf("%s\n", resp.Container.ContainerStatus.Status)
			info += color.BlueString("located node:     ") + fmt.Sprintf("%s\n", resp.Container.ContainerStatus.Node)
			info += color.BlueString("container id:     ") + fmt.Sprintf("%s\n", resp.Container.ContainerStatus.ContainerID)
			
			info += color.CyanString("spec:\n")
			info += color.BlueString("image:            ") + fmt.Sprintf("%s\n", resp.Container.ContainerConfig.Image)
			if resp.Container.ContainerConfig.NodeName != "" {
				info += color.BlueString("node name:        ") + fmt.Sprintf("%s\n", resp.Container.ContainerConfig.NodeName)
			}
			if resp.Container.ContainerConfig.Command != "" {
				info += color.BlueString("command:          ") + fmt.Sprintf("%s\n", command)
			}
			if len(resp.Container.ContainerConfig.Expose) > 0 {
				info += color.BlueString("expose ports:          ")
				for _, e := range resp.Container.ContainerConfig.Expose {
					info += fmt.Sprintf("port: %d  protocol: %s  hostPort: %d\n", e.Port, e.Protocol, e.HostPort)
				}
			}
			if len(resp.Container.ContainerConfig.Env) > 0 {
				info += color.BlueString("env:          ")
				for _, e := range resp.Container.ContainerConfig.Env {
					info += fmt.Sprintf("%s\n", e)
				}
			}
			info += color.BlueString("limit:            ") + fmt.Sprintf("CPU: %d  memory: %d\n", resp.Container.ContainerConfig.Limit.CPU, resp.Container.ContainerConfig.Limit.Memory)
			info += color.BlueString("request:          ") + fmt.Sprintf("CPU: %d  memory: %d\n", resp.Container.ContainerConfig.Request.CPU, resp.Container.ContainerConfig.Request.Memory)
			if ctx.Bool("detail") {
				info += color.CyanString("detail:\n")
				r := resp.Container.ContainerStatus.Info.(map[string]interface{})
				for _, i := range r {
					info += fmt.Sprint(i)
				}
			}

			fmt.Println(info)
			return nil
		},
	}
}

//container list
func containerList() cli.Command {
	return cli.Command{
		Name:        "list",
		Usage:       "Use to list container info",
		Description: "Use 'node list --namespace=<namespace> [args...]' to get node",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "namespace",
				Usage: "the namespace of container",
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			namespace := ctx.String("namespace")

			resp, err := masterCli.Container.List(context.Background(), types.ListContainerRequest{
				Namespace: namespace,
			})

			if err != nil {
				return err
			}

			tableData := [][]string{
				{"name", "namespace", "status", "node"},
			}

			for _, c := range resp.ContainerSimpleInfo {
				tableData = append(tableData, []string{c.Name, c.Namespace, c.Status, c.Node})
			}

			pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
			return nil
		},
	}
}

//container delete
func containerDelete() cli.Command {
	return cli.Command{
		Name:        "delete",
		Usage:       "Use to delete container",
		Description: "Use 'container delete --namespace=<namespace>' to delete container",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "namespace",
				Usage: "the namespace of container",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of container",
				Required: true,
			},
			&cli.BoolFlag{
				Name:  "removeVolumnes",
				Usage: "whether to remove volumnes",
			},&cli.BoolFlag{
				Name:  "removeLinks",
				Usage: "whether to remove links",
			},
			&cli.BoolFlag{
				Name:  "force",
				Usage: "whether to force delete",
			},&cli.IntFlag{
				Name:  "timeout",
				Usage: "how many seconds to set the timeout",
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			namespace := ctx.String("namespace")
			name := ctx.String("name")
			removeVolumnes := ctx.Bool("removeVolumnes")
			removeLinks := ctx.Bool("removeLinks")
			force := ctx.Bool("force")
			timeout := ctx.Int("timeout")

			err := masterCli.Container.Delete(context.Background(), types.DeleteContainerRequest{
				Namespace: namespace,
				Name: name,
				RemoveVolumnes: removeVolumnes,
				RemoveLinks:  removeLinks,
				Force: force,
				Timeout: timeout,
			})

			if err != nil {
				return err
			}

			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("delete container '%s' successfully", name)
			return nil
		},
	}
}