package cmd

import (
	"github.com/urfave/cli"
	"sghcrypto/cmd/crypto"
	"log"
	"fmt"
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
		fmt.Println("A command-line applications to encrypt or decrypt your important data")
		fmt.Println("Very easy to use, support interactive prompt, and as simple as possible")
		return nil
	}
	app.Commands = *crypto.CryptoCommands

	return app.Run(os.Args)
}

func Init() {
	err := InitApp()
	if err != nil {
		log.Fatal(err)
	}
}