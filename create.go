package main

import "github.com/urfave/cli"

var createCmd = cli.Command{
	Name:  "create",
	Usage: "create a container",
	// Flags: ,
	Action: func(clix *cli.Context) error {
		return nil
	},
}
