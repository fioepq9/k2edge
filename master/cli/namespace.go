package cli

import (
	"context"
	"fmt"
	"k2edge/master/client"
	"k2edge/master/internal/types"

	"github.com/urfave/cli"
)

func namespace() *cli.Command {
	
	return &cli.Command {
		Name: "namespace",
		Aliases: []string{"ns"},
		Usage: "Use for namespace management",
		UsageText: "Use 'namespace <command>' to manage namespace. \nUse 'namespace namespace help' for a list of global command-line options.",
		Subcommands: cli.Commands{
			*namespaceCreate(),
		},
	}
}

// namespace create
func namespaceCreate() *cli.Command {
	return &cli.Command {
		Name: "namespace",
		Aliases: []string{"ns"},
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

			config := ctx.App.Metadata["config"].(map[string]string)
			server := config["server"]

			masterCli := client.NewClient(server)
			masterCli.Namespace.NamespaceCreate(context.Background() , types.CreateNamespaceRequest{
				Name: ctx.String("name"),
			})
			return nil
		},
	}
}