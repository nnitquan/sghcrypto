package cmd

import (
	"github.com/urfave/cli"
	"github.com/manifoldco/promptui"
	"sghcrypto/util"
	"fmt"
	"errors"
	"os"
)

var (
	ENCRYPT = "encrypt"
	DECRYPT = "decrypt"
)


func checkArgs(c *cli.Context) error {
	if len(c.Args()) <= 0 {
		return  errors.New("argument required")
	}

	return nil
}

func getCryptoKey(c *cli.Context) (string, error) {
	var crptoKey string
	if c.String("key") != "" {
		crptoKey = c.String("key")
	}else if os.Getenv("CRYPTO_KEY") != "" {
		crptoKey = os.Getenv("CRYPTO_KEY")
	}else {
		prompt := promptui.Prompt{
			Label:    "Crypto Key(16 bytes)",
			Mask:     '*',
		}
		key, err := prompt.Run()
		if err != nil {
			return "", err
		}
		crptoKey = key
	}
	if len(crptoKey) != 16 {
		return "", errors.New("Crypto Key must be equal to 16 characters")
	}
	return crptoKey, nil
}

var cryptoKeyFlag = cli.StringFlag{
	Name: "key,k",
	Value: "",
	Usage: "cryto key for your encrypt or decrypt action",
}

func cryotoAction(c *cli.Context, action string)  error {
	err := checkArgs(c)
	if err != nil {
		return err
	}
	key, err := getCryptoKey(c)
	if err != nil {
		return err
	}
	arg := c.Args().First()
	var content string
	if action == ENCRYPT {
		content,err =util.Encrypt([]byte(key), arg)
	}else if action == DECRYPT {
		content,err =util.Decrypt([]byte(key), arg)
	}else{
		err = errors.New("unknown action")
	}
	if err != nil {
		return err
	}
	fmt.Printf("%s %s ===> %s\n", action, arg, content)
	return nil
}

var crypto  = &[]cli.Command {
	{
		Name:  ENCRYPT,
		Aliases:     []string{"e", "en"},
		Usage: "encrypt a message",
		Flags: []cli.Flag{
			cryptoKeyFlag,
		},
		Action: func(c *cli.Context) error {
			return cryotoAction(c, ENCRYPT)
		},
		//Category:    "crypto",
	},
	{
		Name:  DECRYPT,
		Usage: "decrypt a message",
		Aliases:     []string{"d", "de"},
		Flags: []cli.Flag{
			cryptoKeyFlag,
		},
		Action: func(c *cli.Context) error {
			return cryotoAction(c, DECRYPT)
		},
	},
}