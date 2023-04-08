package cli

import "github.com/urfave/cli"

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
		},
	}
}

// namespace create
func nodeCreate() *cli.Command {
	return &cli.Command {
		Name: "create",
		Usage: "Use for adding node to k2edge",
		Description: "Use 'namespace create --name=<name>' to create namespace",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "the name of namespace",
				Required: true,
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: 
		func(ctx *cli.Context) error { 
			server := ctx.App.Metadata["config-server"].(string)
			ctx.Args()
			masterCli := client.NewClient(server)
			name := ctx.String("name")
			err := masterCli.Namespace.NamespaceCreate(context.Background() , types.CreateNamespaceRequest{
				Name: name,
			})
			if err != nil {
				return err
			}
			
			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("create namespace '%s' success", name)
			return nil
		},
	}
}