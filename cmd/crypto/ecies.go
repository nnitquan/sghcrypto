package crypto

import (
	"github.com/urfave/cli"
	"github.com/manifoldco/promptui"
	"sghcrypto/util"
	"fmt"
	"errors"
	"os"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func getPublicKey(c *cli.Context) ([]byte, error) {
	var key string
	if c.String("pubkey") != "" {
		key = c.String("pubkey")
	}else if os.Getenv("PUBLIC_KEY") != "" {
		key = os.Getenv("PUBLIC_KEY")
	}else {
		prompt := promptui.Prompt{
			Label:    "Public Key(64 bytes)",
			Mask:     '*',
		}
		result, err := prompt.Run()
		if err != nil {
			return nil, err
		}
		key = result
	}
	// We strip off the 0x and the first 2 characters 04 which is always the EC prefix and is not required in public key
	key = fmt.Sprintf("0x04%s", key)
	keyByte, err := hexutil.Decode(key)
	if err != nil {
		return nil, err
	}
	additionByte, err := hexutil.Decode("0x04")
	if err != nil {
		return nil, err
	}
	additionByteLen := len(additionByte)
	if len(keyByte) != (64+additionByteLen) {
		return nil, errors.New("Public Key must be equal to 64 bytes")
	}
	return keyByte, nil
}

func getPrivateKey(c *cli.Context) ([]byte, error) {
	var key string
	if c.String("prikey") != "" {
		key = c.String("prikey")
	}else if os.Getenv("PRIVATE_KEY") != "" {
		key = os.Getenv("PRIVATE_KEY")
	}else {
		prompt := promptui.Prompt{
			Label:    "Private Key(32 bytes)",
			Mask:     '*',
		}
		result, err := prompt.Run()
		if err != nil {
			return nil, err
		}
		key = result
	}
	keyByte, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	if len(keyByte) != 32 {
		return nil, errors.New("Private Key must be equal to 32 bytes")
	}
	return keyByte, nil
}

func eciesEncryptAction(c *cli.Context, data string) (string, error) {
	key, err := getPublicKey(c)
	if err != nil {
		return "", err
	}
	encryptDataByte, err := util.EciesEncrypt([]byte(data), key)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(encryptDataByte), nil
}

func eciesDecryptAction(c *cli.Context, data string) (string, error)  {
	key, err := getPrivateKey(c)
	if err != nil {
		return "", err
	}
	encryptDataByte, err := hexutil.Decode(data)
	if err != nil {
		return "", err
	}
	dataByte, err := util.EciesDecrypt(encryptDataByte, key)
	if err != nil {
		return "", err
	}
	return string(dataByte), nil
}

func eciesCryptoAction(c *cli.Context, action string) error {
	arg := c.Args().First()
	var (
		content string
		err error
	)
	if action == ENCRYPT {
		content, err = eciesEncryptAction(c, arg)
	}else if action == DECRYPT {
		content, err = eciesDecryptAction(c, arg)
	}else{
		err = errors.New("unknown action")
	}
	if err != nil {
		return err
	}
	fmt.Printf("ecies %s %s ===> %s\n", action, arg, content)
	return nil
}