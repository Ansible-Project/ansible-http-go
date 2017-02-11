package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mkideal/cli"
)

var (
	version string
	config  *Config
)

type opt struct {
	Help    bool   `cli:"h,help" usage:"display help"`
	Version bool   `cli:"v,version" usage:"display version and revision"`
	Config  string `cli:"c,config" usage:"set path to config file"`
}

func main() {
	var configPath string

	cli.Run(&opt{}, func(ctx *cli.Context) error {
		argv := ctx.Argv().(*opt)
		if argv.Help {
			ctx.String(ctx.Usage())
			os.Exit(0)
		}
		if argv.Version {
			ctx.String(fmt.Sprintf("%s\n", version))
			os.Exit(0)
		}
		if argv.Config == "" {
			ctx.String(ctx.Usage())
		}
		configPath = argv.Config

		return nil
	})
	if configPath == "" {
		log.Fatal("--config is required.")
	}
	c, err := loadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	config = c

	newServer(c)
}
