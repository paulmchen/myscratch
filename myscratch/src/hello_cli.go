package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli"
	app.Usage = "Print hello world"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "World",
			Usage: "Who to say hello to.",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		// arglist := c.Args()[0]
		// fmt.Printf("argument passed in: %s\n", arglist)
		fmt.Printf("Hello %s!\n", name)
		return nil
	}

	app.Run(os.Args)
}
