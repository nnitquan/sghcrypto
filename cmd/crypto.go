package cmd

import (
	"github.com/urfave/cli"
	"github.com/manifoldco/promptui"
	"sghcrypto/util"
	"fmt"
	"errors"
)

func checkArgs(c *cli.Context) error {
	if len(c.Args()) <= 0 {
		return  errors.New("argument required")
	}

	return nil
}

func getCryptoKey() (string, error) {
	if CRYPTO_KEY == "" {
		//fmt.Printf("Crypto Key(16 bytes): ")
		//keyByte, err := gopass.GetPasswd()
		//if err != nil {
		//	return "", err
		//}

		prompt := promptui.Prompt{
			Label:    "Crypto Key(16 bytes)",
			Mask:     '*',
		}

		key, err := prompt.Run()
		if err != nil {
			return "", err
		}

		CRYPTO_KEY = key
	}
	if len(CRYPTO_KEY) != 16 {
		return "", errors.New("Crypto Key must be equal to 16 characters")
	}
	return CRYPTO_KEY, nil
}

var crypto  = &[]cli.Command {
	{
		Name:  "encrypt",
		Aliases:     []string{"e", "en"},
		Usage: "encrypt a message",
		Action: func(c *cli.Context) error {
			err := checkArgs(c)
			if err != nil {
				return err
			}
			key, err := getCryptoKey()
			if err != nil {
				return err
			}
			arg := c.Args().First()
			content,err :=util.Encrypt([]byte(key), arg)
			if err != nil {
				return err
			}
			fmt.Printf("encrypt %s ===> %s\n", arg, content)
			return nil
		},
		//Category:    "crypto",
	},
	{
		Name:  "decrypt",
		Usage: "decrypt a message",
		Aliases:     []string{"d", "de"},
		Action: func(c *cli.Context) error {
			err := checkArgs(c)
			if err != nil {
				return err
			}
			key, err := getCryptoKey()
			if err != nil {
				return err
			}
			arg := c.Args().First()
			content,err :=util.Decrypt([]byte(key), arg)
			if err != nil {
				return err
			}
			fmt.Printf("decrypt %s ===> %s\n", arg, content)
			return nil
		},
		//Category:    "crypto",
	},
}