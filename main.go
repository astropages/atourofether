package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	/*
		创建ETH交易
	*/

	client, err := ConnectToRPC()
	if err != nil {
		log.Fatal(err)
	}

	//交易目标地址
	account, err := ImportKeystore(KS_DIR+"UTC--2019-11-04T07-33-27.805299499Z--2c8ce8efc5d3cf7a4a3833764fcc307ba98a3067", "123456")
	if err != nil {
		log.Fatal(err)
	}
	toAddress := account.Address

	//加载私钥
	privateKey, err := crypto.HexToECDSA("47C629B5130B6E8BBDE4BB0B72898A5464500C9ADB482F05D4173F599239A426")
	if err != nil {
		log.Fatal(err)
	}
	//交易发起地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("publickKey error")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//读取随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	//设置交易金额
	value := big.NewInt(1000000000000000000) //wei(1ETH)
	//设置燃油上限
	gasLimit := uint64(21000)
	//根据前区块获得平均燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//创建交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	//使用私钥对交易进行签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	//广播交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	//打印交易哈希
	fmt.Println(signedTx.Hash().Hex())

}
