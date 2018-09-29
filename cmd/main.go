package cmd

import (
	"fmt"
	"github.com/nnitquan/sghcrypto/cmd/crypto"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func InitApp() error {
	app := cli.NewApp()
	app.Name = "sghcrypto"
	app.Usage = "crypto for important data!"
	app.Version = "1.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Sanguohot",
			Email: "hw535431@163.com",
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Printf("%s-%s", app.Name, app.Version)
		fmt.Printf("\n%s", app.Usage)
		return nil
	}
	app.Commands = crypto.CryptoCommands

	return app.Run(os.Args)
}

func Init() {
	err := InitApp()
	if err != nil {
		log.Fatal(err)
	}
}