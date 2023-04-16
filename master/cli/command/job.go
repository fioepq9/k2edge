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

func Job() cli.Command {
	return cli.Command{
		Name:        "job",
		Usage:       "Use for job management",
		Description: "Use 'job <command>' to manage job",
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
			jobCreate(),
			jobGet(),
			jobList(),
			jobDelete(),
		},
	}
}

// job create
func jobCreate() cli.Command {
	return cli.Command{
		Name:        "create",
		Usage:       "Use for creating job",
		Description: "Use 'job create --namespace=<namespace> --name=<name> [args...]' to create job",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "namespace",
				Usage:    "the namespace of job",
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of job",
			},
			&cli.IntFlag{
				Name:  "completions",
				Usage: "job completion times",
			},
			&cli.StringFlag{
				Name:  "schedule",
				Usage: "job schedule time",
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
				if !ctx.IsSet("completions") {
					notSetFlags = append(notSetFlags, "completions")
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
			completions := ctx.Int("completions")
			schedule := ctx.String("schedule")
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

			arg := types.CreateJobRequest{
				Job: types.Job{
					Metadata: types.Metadata{
						Namespace: namespace,
						Kind:      "job",
						Name:      name,
					},
					Config: types.JobConfig{
						Completions: completions,
						Schedule: schedule,
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
				config, err :=  yaml2args[types.Job](file)
				if err != nil {
					return err
				}
				arg.Job.Metadata = config.Metadata
				arg.Job.Config = config.Config
				name = arg.Job.Metadata.Name
				arg.Job.Metadata.Kind = "job"
			}

			err := masterCli.Job.Create(context.Background(), arg)

			if err != nil {
				return err
			}

			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("create job '%s' successfully", name)
			return nil
		},
	}
}

// job get
func jobGet() cli.Command {
	return cli.Command{
		Name:        "get",
		Usage:       "Use to get job info",
		Description: "Use 'job get --namespace=<namespace> --name=<name>' to get job",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "namespace",
				Usage: "the namespace of job",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of job",
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

			resp, err := masterCli.Job.Get(context.Background(), types.GetJobRequest{
				Namespace: namespace,
				Name:      name,
			})

			if err != nil {
				return err
			}

			command := resp.Job.Config.Template.Command
			command += " " + strings.Join(resp.Job.Config.Template.Args, " ")
			info := ""
			info += color.BlueString("namespace:        ") + fmt.Sprintf("%s\n", resp.Job.Metadata.Namespace)
			info += color.BlueString("name:             ") + fmt.Sprintf("%s\n", resp.Job.Metadata.Name)
			info += color.BlueString("age:              ") + fmt.Sprintf("%s\n",  time.Since(time.Unix(resp.Job.Config.CreateTime, 0)).Round(time.Second).String())
			info += color.CyanString("\nconfig\n")
			info += color.BlueString("completions:      ") + fmt.Sprintf("%d\n", resp.Job.Config.Completions)
			if resp.Job.Config.Schedule != "" {
				info += color.BlueString("schedule:         ") + fmt.Sprintf("%s\n", resp.Job.Config.Schedule)
			}

			info += color.CyanString("template\n")
			if resp.Job.Config.Template.Name != "" {
				info += color.BlueString("  container name: ") + fmt.Sprintf("%s\n", resp.Job.Config.Template.Name)
			}
			info += color.BlueString("  image:          ") + fmt.Sprintf("%s\n", resp.Job.Config.Template.Image)
			if resp.Job.Config.Template.NodeName != "" {
				info += color.BlueString("  node:           ") + fmt.Sprintf("%s\n", resp.Job.Config.Template.NodeName)
			}
			if resp.Job.Config.Template.Command != "" {
				info += color.BlueString("  command:        ") + fmt.Sprintf("%s\n", command)
			}
			if len(resp.Job.Config.Template.Expose) > 0 {
				info += color.BlueString("  expose ports:   ")
				for _, e := range resp.Job.Config.Template.Expose {
					info += fmt.Sprintf("port: %d  protocol: %s  hostPort: %d\n", e.Port, e.Protocol, e.HostPort)
				}
			}
			if len(resp.Job.Config.Template.Env) > 0 {
			info += color.BlueString("  env:            ")
				for _, e := range resp.Job.Config.Template.Env {
					info += fmt.Sprintf("  %s\n", e)
				}
			}
			info += color.BlueString("  limit:          ") + fmt.Sprintf("CPU: %d  memory: %d\n", resp.Job.Config.Template.Limit.CPU, resp.Job.Config.Template.Limit.Memory)
			info += color.BlueString("  request:        ") + fmt.Sprintf("CPU: %d  memory: %d\n", resp.Job.Config.Template.Request.CPU, resp.Job.Config.Template.Request.Memory)
			
			info += color.CyanString("status\n")
			info += color.BlueString("succeeded:        ") + fmt.Sprintf("%d/%d\n", resp.Job.Succeeded, resp.Job.Config.Completions)
			
			fmt.Println(info)
			return nil
		},
	}
}

// job list
func jobList() cli.Command {
	return cli.Command{
		Name:        "list",
		Aliases: 	 []string{"ls"},
		Usage:       "Use to list job",
		Description: "Use 'job list --namespace=<namespace>' to list job",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "namespace",
				Usage: "the namespace of job",
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return err
		},

		Action: func(ctx *cli.Context) error {
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			namespace := ctx.String("namespace")

			resp, err := masterCli.Job.List(context.Background(), types.ListJobRequest{
				Namespace: namespace,
			})

			if err != nil {
				return err
			}

			tableData := [][]string{
				{"name", "namespace", "age", "succeeded", "schedule"},
			}

			for _, i := range resp.Info {
				tableData = append(tableData, []string{i.Name, i.Namespace, time.Since(time.Unix(i.CreateTime, 0)).Round(time.Second).String(), fmt.Sprintf("%d/%d", i.Succeeded, i.Completions), })
			}

			pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
			return nil
		},
	}
}

// job delete
func jobDelete() cli.Command {
	return cli.Command{
		Name:        "delete",
		Usage:       "Use to delete job",
		Description: "Use 'job delete --namespace=<namespace> --name=<name>' to delete job",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "namespace",
				Usage: "the namespace of job",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of job",
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

			err := masterCli.Job.Delete(context.Background(), types.DeleteJobRequest{
				Namespace: namespace,
				Name: name,
			})

			if err != nil {
				return err
			}

			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("delete job '%s' successfully", name)
			return nil
		},
	}
}