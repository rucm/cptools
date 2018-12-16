package command

import (
	"log"

	"github.com/urfave/cli"
)

func init() {
	commandList = append(commandList, cli.Command{
		Name:  "login",
		Usage: "",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "user, u",
			},
		},
		Action: func(c *cli.Context) error {

			test := c.String("user")
			if test == "" {
				log.Println(test)
			}

			return nil
		},
	})
}
