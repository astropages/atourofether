package eth

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

//EtherTransaction 创建ETH交易(发送者私钥，接收者地址，以太比)
func EtherTransaction(privateKeyHex string, toAddressHex string, ether float64) (*types.Transaction, error) {
	//连接节点
	client, err := ConnectToRPC()
	if err != nil {
		return nil, err
	}
	//接收者地址
	toAddress := common.HexToAddress(toAddressHex)

	//读取发送者私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	//燃气价格：根据前区块从节点获得平均燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	//推导发送者地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	//根据发送者地址从节点自动读取随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	//交易量（1 ether = 10**18 wei）
	value := big.NewInt(int64(ether * math.Pow(10, 18)))
	//燃气上限（固定）
	gasLimit := uint64(21000)

	//创建交易(随机数，接收者，单位为wei的交易量，固定的燃气上限，自动获取的燃气价格)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	//使用私钥对交易进行签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}
	//广播交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	//返回交易
	return signedTx, nil
}

//TokenTransaction 创建代币交易
func TokenTransaction(privateKeyHex string, toAddressHex string, token float64) (*types.Transaction, error) {
	//连接节点
	client, err := ConnectToRPC()
	if err != nil {
		return nil, err
	}
	//接收者地址
	toAddress := common.HexToAddress(toAddressHex)
	//合约地址
	tokenAddress := common.HexToAddress(TOKEN_ADDRESS)
	//检索合约方法
	transferFnSignature := []byte("transfer(address,uint256)")
	//获取方法ID
	hash := crypto.Keccak256Hash(transferFnSignature)
	methodID := hash[:4]
	//填充接收者地址
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	//发送代币数量
	amount := big.NewInt(int64(token * math.Pow(10, 18)))
	//填充代币
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	//拼接数据
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//读取发送者私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	//燃气价格：根据前区块从节点获得平均燃气价格
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	gasPrice := big.NewInt(int64(100 * math.Pow(10, 9))) //10 Gwei

	//推导发送者地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	//根据发送者地址从节点自动读取随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	//交易量（代币交易不需要ETH）
	value := big.NewInt(0)

	//燃气上限（自动估算）
	// gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	// 	To:   &toAddress,
	// 	Data: data,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	gasLimit := uint64(344740)

	//创建交易
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	//使用私钥对交易进行签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}
	//广播交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	//返回交易
	return signedTx, nil
}
