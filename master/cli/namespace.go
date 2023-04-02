package cli

import (
	"fmt"

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
		func(c *cli.Context) error { 
			if !c.IsSet("name") {
				return fmt.Errorf("missing required parameter --name")
			}


			return nil
		},
	}
}