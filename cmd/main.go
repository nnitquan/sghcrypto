package cmd

import (
	"github.com/urfave/cli"
	"log"
	"fmt"
	"encoding/hex"
	"sghcrypto/util"
	"os"
	"time"
)
var DEFAULT_KEY_HEX = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
var DEFAULT_KEY_BYTE []byte
var CRYPTO_KEY string
func InitKey() error {
	DEFAULT_KEY_BYTE, err := hex.DecodeString(DEFAULT_KEY_HEX)
	if err != nil {
		return err
	}
	cryptoKeyEnv := os.Getenv("CRYPTO_KEY")
	if cryptoKeyEnv=="" {
		fmt.Println("Not found CRYPTO_KEY in env, you have to input an 16 bytes key!")
		return nil
	}
	CRYPTO_KEY, err = util.Decrypt(DEFAULT_KEY_BYTE, cryptoKeyEnv)
	if err != nil {
		return err
	}
	return nil
}

func InitApp() error {
	app := cli.NewApp()
	app.Name = "sghcrypto"
	app.Usage = "crypto for important data!"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Sanguohot",
			Email: "hw535431@163.com",
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println("hello world!")
		return nil
	}
	app.Commands = *crypto

	return app.Run(os.Args)
}

func Init() {
	err := InitKey()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = InitApp()
	if err != nil {
		log.Fatal(err)
	}
}