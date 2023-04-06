package cli

import (
	"context"
	"fmt"
	"k2edge/master/client"
	"k2edge/master/internal/types"

	"github.com/urfave/cli"
)

func Namespace() *cli.Command {

	cli := &cli.Command {
		Name: "namespace",
		Aliases: []string{"ns"},
		Usage: "Use for namespace management",
		UsageText: "Use 'namespace <command>' to manage namespace. \nUse 'namespace namespace help' for a list of global command-line options.",
		Before: func(ctx *cli.Context) error {
			etcd := ctx.App.Metadata["config-etcd"].(string)
			server := getServer(etcd)
			ctx.App.Metadata = map[string]interface{}{
				"config-server": server,
			}
			return nil
		},
		Subcommands: cli.Commands{
			*namespaceCreate(),
		},
	}
	return cli
}

// namespace create
func namespaceCreate() *cli.Command {
	return &cli.Command {
		Name: "create",
		Aliases: []string{"c"},
		Usage: "Use for creating namespace ",
		UsageText: "Use 'namespace create <name>' to create namespace. \n",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of namespace",
				Required: true,
			},
		},

		Action: 
		func(ctx *cli.Context) error { 
			if !ctx.IsSet("name") {
				return fmt.Errorf("missing required parameter --name")
			}

			server := ctx.App.Metadata["config-server"].(string)

			masterCli := client.NewClient(server)
			name := ctx.String("name")
			err := masterCli.Namespace.NamespaceCreate(context.Background() , types.CreateNamespaceRequest{
				Name: name,
			})
			if err != nil {
				return err
			}
			fmt.Printf("creat namespace '%s' success\n", name)
			return nil
		},
	}
}

// namespace get
func namespaceGet() *cli.Command {
	return &cli.Command {
		Name: "get",
		Aliases: []string{"g"},
		Usage: "Use for get namespace info",
		UsageText: "Use 'namespace get <name>' to get namespace. \n",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of namespace",
				Required: true,
			},
		},

		Action: 
		func(ctx *cli.Context) error { 
			if !ctx.IsSet("name") {
				return fmt.Errorf("missing required parameter --name")
			}

			config := ctx.App.Metadata["config"].(map[string]string)
			server := config["server"]

			masterCli := client.NewClient(server)
			name := ctx.String("name")
			err := masterCli.Namespace.NamespaceCreate(context.Background() , types.CreateNamespaceRequest{
				Name: name,
			})
			if err != nil {
				return err
			}
			fmt.Printf("creat namespace '%s' success\n", name)
			return nil
		},
	}
}