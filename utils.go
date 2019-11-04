package main

import (
	"context"
	"errors"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

//AddressIsContract 校验地址
func AddressIsContract(str string) (bool, error) {
	client, err := ConnectToRPC()
	if err != nil {
		return false, err
	}
	//校验地址格式
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(str) {
		return false, errors.New("Bad Format")
	}
	//校验地址类型(地址上没有字节码时表示他是一个合约)
	address := common.HexToAddress(str)
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		return false, err
	}
	isContract := len(bytecode) > 0
	return isContract, nil

}
