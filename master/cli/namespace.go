package cli
import "github.com/urfave/cli"

func namespace() (cli.Command, error) {
	namespaceCmd := cli.Command {
		Name: "namespace",
		Aliases: []string{"ns"},
		Usage: "Use 'kubectl options' for a list of global command-line options (applies to all commands).",
	}

	
}