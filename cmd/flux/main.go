package main

import (
	"fmt"
	"os"

	"github.com/ehotinger/flux/version"
	"github.com/urfave/cli"
)

const usage = `a toy for tinkering with criu
_____.__                
_/ ____\  |  __ _____  ___
\   __\|  | |  |  \  \/  /
 |  |  |  |_|  |  />    < 
 |__|  |____/____//__/\_ \
                        \/
`

func init() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Name, version.Package, c.App.Version, version.Revision)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "flux"
	app.Version = version.Version
	app.Usage = usage

	app.Commands = []cli.Command{
		createCmd,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
