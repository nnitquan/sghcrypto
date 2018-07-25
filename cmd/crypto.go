package cmd

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

var (
	ENCRYPT = "encrypt"
	DECRYPT = "decrypt"
	AES     = "aes"
	ECIES   = "ecies"
)

func checkArgs(c *cli.Context) error {
	if len(c.Args()) <= 0 {
		return  errors.New("argument required")
	}

	return nil
}

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

func getCryptoAlgorithm(c *cli.Context) (string, error) {
	var crptoAlgorithm string
	if c.String("algorithm") != "" {
		crptoAlgorithm = c.String("algorithm")
	}else if os.Getenv("CRYPTO_ALG") != "" {
		crptoAlgorithm = os.Getenv("CRYPTO_ALG")
	}else {
		prompt := promptui.Select{
			Label: "Select Algorithm",
			Items: []string{AES, ECIES},
		}
		_, result, err := prompt.Run()

		if err != nil {
			return "", err
		}
		crptoAlgorithm = result
	}
	if crptoAlgorithm == "" {
		return "", errors.New("not found crypto algorithm")
	}
	if crptoAlgorithm!=AES && crptoAlgorithm!=ECIES {
		return "", errors.New("unsupported crypto algorithm")
	}
	return crptoAlgorithm, nil
}

var cryptoFlags = []cli.Flag{
	cli.StringFlag{
		Name: "aeskey",
		Value: "",
		Usage: "aes cryto key",
	},
	cli.StringFlag{
		Name: "pubkey",
		Value: "",
		Usage: "ecies public key",
	},
	cli.StringFlag{
		Name: "prikey",
		Value: "",
		Usage: "ecies private key",
	},
	cli.StringFlag{
		Name: "algorithm,alg",
		Value: "",
		Usage: "cryto algorithm",
	},
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

func cryotoAction(c *cli.Context, action string)  error {
	err := checkArgs(c)
	if err != nil {
		return err
	}
	alg, err := getCryptoAlgorithm(c)
	if err != nil {
		return err
	}
	if alg == AES {
		return aesCryptoAction(c, action)
	}else if alg == ECIES {
		return eciesCryptoAction(c, action)
	}
	return errors.New("unsupported crypto algorithm")
}

var crypto  = &[]cli.Command {
	{
		Name:  ENCRYPT,
		Aliases:     []string{"e", "en"},
		Usage: "encrypt a message",
		Flags: cryptoFlags,
		Action: func(c *cli.Context) error {
			return cryotoAction(c, ENCRYPT)
		},
		//Category:    "crypto",
	},
	{
		Name:  DECRYPT,
		Usage: "decrypt a message",
		Aliases:     []string{"d", "de"},
		Flags: cryptoFlags,
		Action: func(c *cli.Context) error {
			return cryotoAction(c, DECRYPT)
		},
	},
}