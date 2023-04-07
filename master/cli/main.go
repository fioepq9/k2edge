package main

import (
	"fmt"
	cmd "k2edge/master/cli/command"
	"os"

	"github.com/pterm/pterm"
	"github.com/urfave/cli"

	"gopkg.in/yaml.v2"
)

func main() {
	var filePath string = "../etc/master-api.yaml"

	app := cli.NewApp()
	app.Name = "k2e-ctl"
	app.Version = "v1.0.1"
	app.UsageText = "Use for managing K2edge's resource"
	app.Before = func(ctx *cli.Context) error {
		data, err := os.ReadFile(filePath)
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

	app.Commands = []cli.Command{*cmd.Namespace()}

	err := app.Run(os.Args)
    if err != nil {
        pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgRed)).Printfln(err.Error())
    }
}

type config struct {
	Etcd Etcd `yaml:"Etcd"`
}

type Etcd struct {	
	Endpoints []string `yaml:"Endpoints"`
	DialTimeout int		`yaml:"DialTimeout"`
}