package eth

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
)

//Wallet 钱包
type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublickKey *ecdsa.PublicKey
}

//NewWalletKeyPair 创建钱包密钥对
func NewWalletKeyPair() (*Wallet, error) {
	//随机私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	//公钥生成
	publickKey := privateKey.Public()
	publickKeyECDSA, ok := publickKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("publickKey error")
	}

	wallet := Wallet{privateKey, publickKeyECDSA}
	return &wallet, nil
}

//AddressHex 钱包地址
func (w *Wallet) AddressHex() string {
	return crypto.PubkeyToAddress(*w.PublickKey).Hex()
}
