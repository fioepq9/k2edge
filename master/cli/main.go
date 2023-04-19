package main

import (
	"fmt"
	"io"
	cmd "k2edge/master/cli/command"
	"os"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/fatih/color"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"

	"gopkg.in/yaml.v2"
)

func main() {
	cli.HelpPrinterCustom = func(out io.Writer, templ string, data interface{}, customFuncs map[string]interface{}) {
		funcMap := template.FuncMap{
			"sprint": fmt.Sprint,
			"join":   strings.Join,
			"blue":   color.HiBlueString,
			"cyan":   color.HiCyanString,
			"cyanFlag": func(flag cli.Flag) string {
				str := strings.SplitN(flag.String(), "\t", 2)
				if len(str) == 1 {
					return color.HiCyanString(str[0])
				}
				return color.HiCyanString(str[0]) + "\t" + str[1]
			},
		}
		for key, value := range customFuncs {
			funcMap[key] = value
		}

		w := tabwriter.NewWriter(out, 1, 8, 2, ' ', 0)
		t := template.Must(template.New("help").Funcs(funcMap).Parse(templ))
		err := t.Execute(w, data)
		if err != nil {
			// If the writer is closed, t.Execute will fail, and there's nothing
			// we can do to recover.
			if os.Getenv("CLI_TEMPLATE_ERROR_DEBUG") != "" {
				_, _ = fmt.Fprintf(os.Stderr, "CLI TEMPLATE ERROR: %#v\n", err)
			}
			return
		}
		_ = w.Flush()
	}
	cli.AppHelpTemplate = AppHelpTemplate
	cli.CommandHelpTemplate = CommandHelpTemplate
	cli.SubcommandHelpTemplate = SubcommandHelpTemplate

	app := cli.NewApp()
	app.Name = "k2e-ctl"
	app.Version = "v1.0.1"
	app.Usage = "a control panel of k2edge"
	app.Description = "Use for managing K2edge's resource"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Usage: "the configuration of k2edge",
			Value: "/etc/master-api.yaml",
		},
	}
	app.Before = func(ctx *cli.Context) error {
		data, err := os.ReadFile(ctx.String("config"))
		if err != nil {
			return fmt.Errorf("k2e get configuration failed")
		}

		config := new(config)
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			return fmt.Errorf("k2e get configuration failed")
		}

		if len(config.Etcd.Endpoints) == 0 {
			panic("k2e initial: cannot found etcd")
		}

		ctx.App.Metadata = map[string]interface{}{
			"config-etcd": config.Etcd.Endpoints[0],
		}

		return nil
	}

	app.CommandNotFound = func(ctx *cli.Context, s string) {
		pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgRed)).Printfln("cannot find the command '%s', input -h for help", s)
	}

	app.Commands = []cli.Command{
		cmd.Namespace(),
		cmd.Node(),
		cmd.Container(),
		cmd.Deployment(),
		cmd.Job(),
	}

	err := app.Run(os.Args)
	if err != nil {
		pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgRed)).Printfln(err.Error())
	}
}

type config struct {
	Etcd Etcd `yaml:"Etcd"`
}

type Etcd struct {
	Endpoints   []string `yaml:"Endpoints"`
	DialTimeout int      `yaml:"DialTimeout"`
}

var AppHelpTemplate = `{{blue "NAME:"}}
   {{.Name}}{{if .Usage}} - {{.Usage}}{{end}}
{{blue "USAGE:"}}
   {{if .Description}}{{.Description}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if len .Authors}}
{{blue "AUTHOR"}}{{with $length := len .Authors}}{{if ne 1 $length}}{{blue "S"}}{{end}}{{end}}:
   {{range $index, $author := .Authors}}{{if $index}}
   {{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}{{if .Copyright}}
{{blue "COPYRIGHT:"}}
   {{cyan .Copyright}}{{end}}{{if .Version}}{{if not .HideVersion}}
{{blue "VERSION:"}}
   {{.Version}}{{end}}{{end}}

{{blue "COMMANDS:"}}{{range .VisibleCategories}}{{if .Name}}

   {{.Name}}:{{range .VisibleCommands}}{{$name1 := join .Names ", "}}
     {{cyan $name1}}{{"\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}{{$name2 := join .Names ", "}}
   {{cyan $name2}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

{{blue "GLOBAL OPTIONS:"}}
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{cyanFlag $option}}{{end}}{{end}}

`

var SubcommandHelpTemplate = `{{blue "NAME:"}}
   {{.HelpName}} - {{.Usage}}
{{blue "USAGE:"}}
   {{if .Description}}{{.Description}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}

{{blue "COMMANDS:"}}{{range .VisibleCategories}}{{if .Name}}

   {{.Name}}:{{range .VisibleCommands}}{{$name1 := join .Names ", "}}
     {{cyan $name1}}{{"\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}{{$name2 := join .Names ", "}}
   {{cyan $name2}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

{{blue "OPTIONS:"}}
   {{range .VisibleFlags}}{{cyanFlag .}}
   {{end}}{{end}}
`

var CommandHelpTemplate = `{{blue "NAME:"}}
   {{.HelpName}} - {{.Usage}}
{{blue "USAGE:"}}
   {{if .Description}}{{.Description}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Category}}
{{blue "CATEGORY:"}}
   {{.Category}}{{end}}{{if .VisibleFlags}}

{{blue "OPTIONS:"}}
   {{range .VisibleFlags}}{{cyanFlag .}}
   {{end}}{{end}}
`
