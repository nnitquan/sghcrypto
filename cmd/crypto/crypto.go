package crypto

import (
	"github.com/urfave/cli"
	"github.com/manifoldco/promptui"
	"errors"
	"os"
)

var (
	ENCRYPT = "encrypt"
	DECRYPT = "decrypt"
	AES     = "aes"
	ECIES   = "ecies"
	algList = []string{AES, ECIES}
)

func checkArgs(c *cli.Context) error {
	if len(c.Args()) <= 0 {
		return  errors.New("argument required")
	}

	return nil
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
			Items: algList,
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
	isSupportAlg := false
	for _, alg:=range algList {
		if alg == crptoAlgorithm {
			isSupportAlg = true
			break
		}
	}
	if !isSupportAlg {
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
		Usage: "cryto algorithm, currently support ase and ecies",
	},
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
	switch alg {
		case ECIES:
			err = eciesCryptoAction(c, action)
			break
		case AES:
			err = aesCryptoAction(c, action)
			break
		default:
			err = errors.New("unsupported crypto algorithm")
			break
	}
	return err
}

var CryptoCommands  = []cli.Command {
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