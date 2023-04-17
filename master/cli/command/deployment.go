package cli

import (
	"context"
	"fmt"
	"k2edge/master/client"
	"k2edge/master/internal/types"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func Deployment() cli.Command {
	return cli.Command{
		Name:        "deployment",
		Usage:       "Use for deployment management",
		Description: "Use 'deployment <command>' to manage deployment",
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
			deploymentCreate(),
			deploymentGet(),
			deploymentList(),
			deploymentDelete(),
			deploymentApply(),
			deploymentScale(),
		},
	}
}

// deployment create
func deploymentCreate() cli.Command {
	return cli.Command{
		Name:        "create",
		Usage:       "Use for creating deployment",
		Description: "Use 'deployment create --namespace=<namespace> --name=<name> [args...]' to create deployment",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "namespace",
				Usage:    "the namespace of deployment",
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of deployment",
			},
			&cli.IntFlag{
				Name:  "replicas",
				Usage: "the number of replicas in deployment",
			},
			&cli.StringFlag{
				Name:  "cname",
				Usage: "the name of container",
			},
			&cli.StringFlag{
				Name:     "image",
				Usage:    "the image of deployment",
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
			&cli.StringFlag{
				Name:  "f",
				Usage: "YAML configuration file",
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return err
		},
		Action: func(ctx *cli.Context) error {
			notSetFlags := []string{}
			if !ctx.IsSet("f") {
				if !ctx.IsSet("namespace") {
					notSetFlags = append(notSetFlags, "namespace")
				}
				if !ctx.IsSet("name") {
					notSetFlags = append(notSetFlags, "name")
				}
				if !ctx.IsSet("replicas") {
					notSetFlags = append(notSetFlags, "replicas")
				}
				if !ctx.IsSet("image") {
					notSetFlags = append(notSetFlags, "image")
				}
				if len(notSetFlags) != 0 {
					return fmt.Errorf("required flags %s not set", strings.Join(notSetFlags, ", "))
				}
			}

			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			namespace := ctx.String("namespace")
			name := ctx.String("name")
			replicas := ctx.Int("replicas")
			cname := ctx.String("cname")
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

			arg := types.CreateDeploymentRequest {
				Deployment: types.Deployment{
					Metadata: types.Metadata{
						Namespace: namespace,
						Kind:      "deployment",
						Name:      name,
					},
					Config: types.DeploymentConfig{
						Replicas: replicas,
						Template: types.ContainerTemplate{
							Name: cname,
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
				},
			}

			if ctx.IsSet("f") {
				file := ctx.String("f")
				config, err :=  yaml2args[types.Deployment](file)
				if err != nil {
					return err
				}
				arg.Deployment.Metadata = config.Metadata
				arg.Deployment.Config = config.Config
				name = arg.Deployment.Metadata.Name
				arg.Deployment.Metadata.Kind = "deployment"
			}

			resp, err := masterCli.Deployment.Create(context.Background(), arg)

			if err != nil {
				return err
			}

			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("create deployment '%s' successfully", name)
			for _, e := range resp.Err {
				pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgRed)).Printfln("container error occured: %s", e)
			}
			return nil
		},
	}
}

// deployment get
func deploymentGet() cli.Command {
	return cli.Command{
		Name:        "get",
		Usage:       "Use to get deployment info",
		Description: "Use 'deployment get --namespace=<namespace> --name=<name>' to get node",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "namespace",
				Usage: "the namespace of deployment",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of deployment",
				Required: true,
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return err
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			namespace := ctx.String("namespace")
			name := ctx.String("name")

			resp, err := masterCli.Deployment.Get(context.Background(), types.GetDeploymentRequest{
				Namespace: namespace,
				Name:      name,
			})

			if err != nil {
				return err
			}

			command := resp.Deployment.Config.Template.Command
			command += " " + strings.Join(resp.Deployment.Config.Template.Args, " ")
			info := ""
			info += color.BlueString("namespace:            ") + fmt.Sprintf("%s\n", resp.Deployment.Metadata.Namespace)
			info += color.BlueString("name:                 ") + fmt.Sprintf("%s\n", resp.Deployment.Metadata.Name)
			info += color.BlueString("age:                  ") + fmt.Sprintf("%s\n",  time.Since(time.Unix(resp.Deployment.Config.CreateTime, 0)).Round(time.Second).String())
			info += color.CyanString("\nconfig\n")
			info += color.BlueString("replicas:             ") + fmt.Sprintf("%d\n", resp.Deployment.Config.Replicas)
			
			info += color.CyanString("template\n")
			if resp.Deployment.Config.Template.NodeName != "" {
				info += color.BlueString("  container name:     ") + fmt.Sprintf("%s\n", resp.Deployment.Config.Template.Name)
			}
				info += color.BlueString("  image:              ") + fmt.Sprintf("%s\n", resp.Deployment.Config.Template.Image)
			if resp.Deployment.Config.Template.NodeName != "" {
				info += color.BlueString("  node:               ") + fmt.Sprintf("%s\n", resp.Deployment.Config.Template.NodeName)
			}
			if resp.Deployment.Config.Template.Command != "" {
				info += color.BlueString("  command:            ") + fmt.Sprintf("%s\n", command)
			}
			if len(resp.Deployment.Config.Template.Expose) > 0 {
				info += color.BlueString("  expose ports:       ")
				for _, e := range resp.Deployment.Config.Template.Expose {
					info += fmt.Sprintf("   port: %d  protocol: %s  hostPort: %d\n", e.Port, e.Protocol, e.HostPort)
				}
			}
			if len(resp.Deployment.Config.Template.Env) > 0 {
				info += color.BlueString("    env:              ")
				for _, e := range resp.Deployment.Config.Template.Env {
					info += fmt.Sprintf("  %s\n", e)
				}
			}
			info += color.BlueString("  limit:              ") + fmt.Sprintf("CPU: %d  memory: %d\n", resp.Deployment.Config.Template.Limit.CPU, resp.Deployment.Config.Template.Limit.Memory)
			info += color.BlueString("  request:            ") + fmt.Sprintf("CPU: %d  memory: %d\n", resp.Deployment.Config.Template.Request.CPU, resp.Deployment.Config.Template.Request.Memory)
			
			info += color.CyanString("\nstatus\n")
			info += color.BlueString("available:            ") + fmt.Sprintf("%d/%d\n", resp.Deployment.Status.AvailableReplicas, resp.Deployment.Config.Replicas)
			if len(resp.Deployment.Status.Containers) > 0 {
				info += color.BlueString("containers:\n")
				for _, c := range resp.Deployment.Status.Containers {
					info += fmt.Sprintf("name:                %s\n", c.Name)
					info += fmt.Sprintf("node:                %s\n", c.Node)
					info += fmt.Sprintf("container id:        %s\n\n", c.ContainerID)
				}
			}
			
			fmt.Println(info)
			return nil
		},
	}
}

// deployment list
func deploymentList() cli.Command {
	return cli.Command{
		Name:        "list",
		Aliases: 	 []string{"ls"},
		Usage:       "Use to list deployment",
		Description: "Use 'deployment list --namespace=<namespace>' to list deployment",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "namespace",
				Usage: "the namespace of deployment",
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return err
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			namespace := ctx.String("namespace")

			resp, err := masterCli.Deployment.List(context.Background(), types.ListDeploymentRequest{
				Namespace: namespace,
			})

			if err != nil {
				return err
			}

			tableData := [][]string{
				{"name", "namespace", "age", "available"},
			}

			for _, i := range resp.Info {
				tableData = append(tableData, []string{i.Name, i.Namespace, time.Since(time.Unix(i.CreateTime, 0)).Round(time.Second).String(), fmt.Sprintf("%d/%d", i.Available, i.Replicas)})
			}

			pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
			return nil
		},
	}
}

// deployment delete
func deploymentDelete() cli.Command {
	return cli.Command{
		Name:        "delete",
		Usage:       "Use to delete deployment",
		Description: "Use 'deployment delete --namespace=<namespace> --name=<name>' to delete deployment",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "namespace",
				Usage: "the namespace of deployment",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of deployment",
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
			return err
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

			resp, err := masterCli.Deployment.Delete(context.Background(), types.DeleteDeploymentRequest{
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

			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("delete deployment '%s' successfully", name)
			for _, e := range resp.Err {
				pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgRed)).Printfln("container error occured: %s", e)
			}
			return nil
		},
	}
}

// deployment apply
func deploymentApply() cli.Command {
	return cli.Command{
		Name:        "apply",
		Usage:       "apply the configuration of the deployment",
		Description: "Use 'deployment apply --namespace=<namespace> --name=<name> [args...]' to apply configuration",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "namespace",
				Usage:    "the namespace of deployment",
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of deployment",
			},
			&cli.IntFlag{
				Name:  "replicas",
				Usage: "the number of replicas in deployment",
			},
			&cli.StringFlag{
				Name:  "cname",
				Usage: "the name of container",
			},
			&cli.StringFlag{
				Name:     "image",
				Usage:    "the image of deployment",
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
			&cli.StringFlag{
				Name:  "f",
				Usage: "YAML configuration file",
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return err
		},
		Action: func(ctx *cli.Context) error {
			notSetFlags := []string{}
			if !ctx.IsSet("f") {
				if !ctx.IsSet("namespace") {
					notSetFlags = append(notSetFlags, "namespace")
				}
				if !ctx.IsSet("name") {
					notSetFlags = append(notSetFlags, "name")
				}
				if !ctx.IsSet("image") {
					notSetFlags = append(notSetFlags, "image")
				}
				if !ctx.IsSet("replicas") {
					notSetFlags = append(notSetFlags, "replicas")
				}
				if len(notSetFlags) != 0 {
					return fmt.Errorf("required flags %s not set", strings.Join(notSetFlags, ", "))
				}
			}

			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			namespace := ctx.String("namespace")
			name := ctx.String("name")
			replicas := ctx.Int("replicas")
			cname := ctx.String("cname")
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

			arg := types.CreateDeploymentRequest {
				Deployment: types.Deployment{
					Metadata: types.Metadata{
						Namespace: namespace,
						Kind:      "deployment",
						Name:      name,
					},
					Config: types.DeploymentConfig{
						Replicas: replicas,
						Template: types.ContainerTemplate{
							Name: cname,
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
				},
			}

			if ctx.IsSet("f") {
				file := ctx.String("f")
				config, err :=  yaml2args[types.Deployment](file)
				if err != nil {
					return err
				}
				arg.Deployment.Metadata = config.Metadata
				arg.Deployment.Config = config.Config
				name = arg.Deployment.Metadata.Name
				arg.Deployment.Metadata.Kind = "deployment"
			}

			resp, err := masterCli.Deployment.Apply(context.Background(), types.ApplyDeploymentRequest{
				Namespace: arg.Deployment.Metadata.Namespace,
				Name: arg.Deployment.Metadata.Name,
				Config: arg.Deployment.Config,
			})

			if err != nil {
				return err
			}

			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("apply the configuration of deployment'%s' successfully", name)
			for _, e := range resp.Err {
				pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgRed)).Printfln("container error occured: %s", e)
			}
			return nil
		},
	}
}

// deployment scale
func deploymentScale() cli.Command {
	return cli.Command{
		Name:        "scale",
		Usage:       "Use to adjust the replicas for deployment.",
		Description: "Use 'deployment scale --namespace=<namespace> --name=<name> --replicas=<replicas>' to adjust the replicas",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "namespace",
				Usage: "the namespace of deployment",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of deployment",
				Required: true,
			},
			&cli.IntFlag{
				Name:  "replicas",
				Usage: "the number of replicas in deployment",
				Required: true,
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return err
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			namespace := ctx.String("namespace")
			name := ctx.String("name")
			replicas := ctx.Int("replicas")

			err := masterCli.Deployment.Scale(context.Background(), types.ScaleRequest{
				Namespace: namespace,
				Name:      name,
				Replicas:	   replicas,
			})

			if err != nil {
				return err
			}

			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("adjusting the number of replica successfully")
			return nil
		},
	}
}