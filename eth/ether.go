package eth

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
	//节点
	ENDPOINT = "https://ropsten.infura.io/v3/e2a64621539843ebbae17402e672d210"
	//代币
	TOKEN_ADDRESS = "0x07aa03e058c3c1dbf6d19e2751cc8cd60148719f"
)

//ConnectToRPC 建立RPC连接
func ConnectToRPC() (*ethclient.Client, error) {
	client, err := rpc.Dial(ENDPOINT)
	if err != nil {
		return nil, err
	}
	conn := ethclient.NewClient(client)
	return conn, nil
}

//GetBalance 获取余额
func GetBalance(addressHex string) (float64, error) {
	client, err := ConnectToRPC()
	if err != nil {
		return -1, err
	}
	balance, err := client.BalanceAt(context.TODO(), common.HexToAddress(addressHex), nil)
	return float64(balance.Int64()) * math.Pow(10, -18), nil

}

//GetBlock 获取区块
func GetBlock(blockNumberInt int64) (*types.Block, error) {
	client, err := ConnectToRPC()
	if err != nil {
		return nil, err
	}
	blockNumber := big.NewInt(blockNumberInt)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}
	return block, nil
}
