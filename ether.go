package main

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	//测试节点
	TEST_ENDPOINT = "https://ropsten.infura.io/v3/e2a64621539843ebbae17402e672d210"
	//测试合约
	TEST_CONTRACT = "0x33075eDc32474D89BCd1aD23E180A1E96A45FeA2"
	//测试钱包
	TEST_ADDRESS1 = "0xE7bc6d2F28B68626106391332fEdFD31A3725bBb"
	TEST_ADDRESS2 = "0x0698c06FC0c46f57CA561E561b03E2b42522455f"
)

//ConnectToRPC 建立RPC连接
func ConnectToRPC() (*ethclient.Client, error) {
	client, err := rpc.Dial(TEST_ENDPOINT)
	if err != nil {
		return nil, err
	}
	conn := ethclient.NewClient(client)
	return conn, nil
}

//GetBalance 获取地址余额
func GetBalance(address string) (float64, error) {
	client, err := ConnectToRPC()
	if err != nil {
		return -1, err
	}
	balance, err := client.BalanceAt(context.TODO(), common.HexToAddress(address), nil)
	return float64(balance.Int64()) * math.Pow(10, -18), nil

}

//GetBlock 获取区块
func GetBlock(num int64) (*types.Block, error) {
	client, err := ConnectToRPC()
	if err != nil {
		return nil, err
	}
	blockNumber := big.NewInt(num)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}
	return block, nil
}
