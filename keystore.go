package main

import (
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

const (
	//Keystore文件目录
	KS_DIR = "./wallets/"
)

//NewKeystore keystore创建账户
func NewKeystore(password string) (*accounts.Account, error) {
	ks := keystore.NewKeyStore(KS_DIR, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

//ImportKeystore keystore导入账户
func ImportKeystore(file, password string) (*accounts.Account, error) {
	ks := keystore.NewKeyStore(KS_DIR, keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		return nil, err
	}
	if err := os.Remove(file); err != nil {
		return nil, err
	}
	return &account, nil
}
