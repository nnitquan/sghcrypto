package util

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"errors"
	"crypto/rand"
	"fmt"
)

func EciesEncrypt(dataByte []byte, publicKeyByte []byte) (encryptDataByte []byte, err error) {
	if len(dataByte)<=0 || len(publicKeyByte)<=0 {
		return nil, errors.New("params error")
	}
	publicKeyEcdsa, err := crypto.UnmarshalPubkey(publicKeyByte)
	fmt.Println("publicKeyByte", publicKeyByte)
	if err != nil {
		return nil, err
	}
	publicKeyEcies := ecies.ImportECDSAPublic(publicKeyEcdsa)
	encryptDataByte, err = ecies.Encrypt(rand.Reader, publicKeyEcies, dataByte, nil, nil)
	if err != nil {
		return nil, err
	}
	return encryptDataByte, nil
}

func EciesDecrypt(encryptDataByte []byte, privateKeyByte []byte) (decryptDataByte []byte, err error) {
	if len(encryptDataByte)<=0 || len(privateKeyByte)<=0 {
		return nil, errors.New("params error")
	}
	privateKeyEcdsa, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		return nil, err
	}
	privateKeyEcies := ecies.ImportECDSA(privateKeyEcdsa)
	decryptDataByte, err = privateKeyEcies.Decrypt(encryptDataByte, nil, nil)
	if err != nil {
		return nil, err
	}
	return decryptDataByte, nil
}

