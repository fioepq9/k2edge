package cli

import (
	"context"
	"fmt"
	"k2edge/master/client"
	"k2edge/master/internal/types"

	"github.com/pterm/pterm"

	"github.com/urfave/cli"
)

func Namespace() *cli.Command {
	return &cli.Command {
		Name: "namespace",
		Aliases: []string{"ns"},
		Usage: "Use for namespace management",
		Description: "Use 'namespace <command>' to manage namespace",
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
			*namespaceCreate(),
			*namespaceGet(),
			*namespaceList(),
			*namespaceDelete(),
		},
	}
}

// namespace create
func namespaceCreate() *cli.Command {
	return &cli.Command {
		Name: "create",
		Usage: "Use for creating namespace ",
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

// namespace get
func namespaceGet() *cli.Command {
	return &cli.Command {
		Name: "get",
		Usage: "Use for get namespace info",
		Description: "Use 'namespace get --name=<name>' to get namespace",
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

			masterCli := client.NewClient(server)
			name := ctx.String("name")
			resp, err := masterCli.Namespace.NamespaceGet(context.Background() , types.GetNamespaceRequest{
				Name: name,
			})
			if err != nil {
				return err
			}

			tableData := [][]string{
				{"name", "status", "age"},
				{resp.Name, resp.Status, resp.Age},
			}
			pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()

			return nil
		},
	}
}

// namespace list
func namespaceList() *cli.Command {
	return &cli.Command {
		Name: "list",
		Aliases: []string{"l"},
		Usage: "Use for list namespace info",
		Description: "Use 'namespace list --all=<All>' to list namespace",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "all",
				Usage: "slow all namespace(option)",
				Value: "true",
			},
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return fmt.Errorf(err.Error())
		},

		Action: 
		func(ctx *cli.Context) error { 
			server := ctx.App.Metadata["config-server"].(string)
			masterCli := client.NewClient(server)

			req := types.ListNamespaceRequest{}
			if ctx.String("all") == "true" {
				req.All = true
			} else {
				req.All = false
			}

			resp, err := masterCli.Namespace.NamespaceList(context.Background(), req)
			if err != nil {
				return err
			}

			tableData := [][]string{
				{"name", "status", "age"},
			}
			for _, n := range resp.Namespaces {
				tableData = append(tableData, []string{n.Name, n.Status, n.Age})
			}

			pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
			return nil
		},
	}
}

// namespace delete
func namespaceDelete() *cli.Command {
	return &cli.Command {
		Name: "delete",
		Usage: "Use for deleting namespace",
		Description: "Use 'namespace delete --name=<name>' to delete namespace",
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
			masterCli := client.NewClient(server)

			name := ctx.String("name")
			err := masterCli.Namespace.NamespaceDelete(context.Background(), types.DeleteNamespaceRequest{
				Name: name,
			})
			if err != nil {
				return err
			}

			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Printfln("delete namespace '%s' success", name)
			return nil
		},
	}
}

