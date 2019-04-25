package main

import (
	// "fmt"
	"github.com/djumpen/console-probe/commands"
	_ "github.com/djumpen/console-probe/config"
	"gopkg.in/urfave/cli.v2"
	"os"
	// "strings"
	// survey "gopkg.in/AlecAivazis/survey.v1"
)

func main() {

	// TODO: check config

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			var sub string
			if c.NArg() > 0 {
				sub = c.Args().Get(0)
				commands.Run(sub, c.Args().Slice()[1:])
			}
			// if c.String("lang") == "spanish" {
			// 	fmt.Println("Hola", name)
			// }
			return nil
		},
	}

	app.Run(os.Args)

}
