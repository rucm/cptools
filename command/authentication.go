package command

import (
	"fmt"

	controller "github.com/rucm/cptools/atcoder/controller"
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
			cli.StringFlag{
				Name: "password, p",
			},
		},
		Action: func(c *cli.Context) error {

			user := c.String("user")
			password := c.String("password")
			auth := controller.NewAuthenticationController()

			if user == "" {
				fmt.Print("user -> ")
				fmt.Scan(&user)
			}

			if password == "" {
				fmt.Print("password -> ")
				fmt.Scan(&password)
			}

			auth.Login(user, password)

			return nil
		},
	})
}
