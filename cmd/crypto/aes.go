package crypto

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/nnitquan/sghcrypto/util"
	"github.com/urfave/cli"
	"os"
)

func getAesKey(c *cli.Context) (string, error) {
	var (
		key string
		err error
	)
	if c.String("aeskey") != "" {
		key = c.String("aeskey")
	}else if os.Getenv("AES_KEY") != "" {
		key = os.Getenv("AES_KEY")
	}else {
		prompt := promptui.Prompt{
			Label:    "AES Key(16 bytes)",
			Mask:     '*',
		}
		key, err = prompt.Run()
		if err != nil {
			return "", err
		}
	}
	if len(key) != 16 {
		return "", errors.New("AES Key must be equal to 16 characters")
	}
	return key, nil
}

func aesCryptoAction(c *cli.Context, action string) error {
	key, err := getAesKey(c)
	if err != nil {
		return err
	}
	arg := c.Args().First()
	var content string
	if action == ENCRYPT {
		content, err = util.AesEncrypt([]byte(key), arg)
	}else if action == DECRYPT {
		content, err = util.AesDecrypt([]byte(key), arg)
	}else{
		err = errors.New("unknown action")
	}
	if err != nil {
		return err
	}
	fmt.Printf("aes %s %s ===> %s\n", action, arg, content)
	return nil
}