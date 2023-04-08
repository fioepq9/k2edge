package cli

import (
	"context"
	"fmt"
	"k2edge/master/client"
	"k2edge/master/internal/types"

	"github.com/pterm/pterm"
	"github.com/samber/lo"
	"github.com/urfave/cli"
)

func Node() *cli.Command {
	return &cli.Command {
		Name: "node",
		Usage: "Use for node management",
		Description: "Use 'node <command>' to manage node",
		Before: func(ctx *cli.Context) error {
			etcd := ctx.App.Metadata["config-etcd"].(string)
			server := getServer(etcd)
			ctx.App.Metadata = map[string]interface{}{
				"config-server": server,
				"config-etcd": etcd,
			}
			return nil
		},
		Subcommands: cli.Commands{
			*nodeCreate(),
		},
	}
}

// namespace register
func nodeCreate() *cli.Command {
	return &cli.Command {
		Name: "create",
		Usage: "Use for adding node to k2edge",
		Description: "Use 'namespace create --name=<name>' to create namespace",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of node, no repetition allowed",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name: "roles",
				Usage: "the roles of node, master/worker",
				Required: true,
			},
			&cli.StringFlag{
				Name: "masterurl",
				Usage: "the url of master",
			},
			&cli.StringFlag{
				Name: "workerurl",
				Usage: "the url of worker",
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: 
		func(ctx *cli.Context) error { 
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)
			name := ctx.String("name")
			roles := ctx.StringSlice("roles")
			masterurl := ctx.String("masterurl")
			workerurl := ctx.String("workerurl")
			if lo.Contains(roles, "master") && masterurl == "" {
				return fmt.Errorf("masterurl is empty")
			}
			if lo.Contains(roles, "worker") && workerurl == "" {
				return fmt.Errorf("workerurl is empty")
			}

			err := masterCli.Node.Register(context.Background() , &types.RegisterRequest{
				Name: name,
				Roles: roles,
				BaseURL: types.NodeURL{
					MasterURL: workerurl,
					WorkerURL: masterurl,
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