package rpc

import (
	"fmt"
	"testing"
)

const (
	MYKEY    = "0xe0Df5ba2C23B2EdDCaB7DF15255F12474680B8E2" //基地址
	NEW      = "0x474b5eeDfE398398C04e26082102C7Bddc6beb5d" //新创建的地址
	CONTRACT = "0x09884521c53a4e8687c2d67a4f38e414e08afeb0" //合约地址
	//合约编译后的code
	CODE = "0x608060405234801561001057600080fd5b5061302e60008190555061dbd6600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506101c08061006f6000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063555516691461005c578063853b2d2914610087578063bcdfe0d5146100b4575b600080fd5b34801561006857600080fd5b50610071610144565b6040518082815260200191505060405180910390f35b34801561009357600080fd5b506100b26004803603810190808035906020019092919050505061014d565b005b3480156100c057600080fd5b506100c9610157565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101095780820151818401526020810190506100ee565b50505050905090810190601f1680156101365780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60008054905090565b8060008190555050565b60606040805190810160405280600581526020017f48656c6c6f0000000000000000000000000000000000000000000000000000008152509050905600a165627a7a72305820eea9bb99be157b2c85d66eeb1a5de7628fb52f6260f3b2cda129051c97de003a0029"
)

func TestEthSign(t *testing.T) {
	param2 := "0xdeadbeaf" //待签名的数据
	result, err := EthSign(MYKEY, param2)
	if err != nil {
		fmt.Errorf("test eth_sign method error." + err.Error())
		return
	}
	fmt.Println("test eth_sign method result.", result.(Reply))
}

func TestEthCall(t *testing.T) {
	p := Param{From: MYKEY, To: CONTRACT, Gas: "0x1000", GasPrice: "0x9184e72a000", Data: "0xbcdfe0d5"}
	result, err := EthCall(p, "latest")
	if err != nil {
		fmt.Errorf("test eth_call method error." + err.Error())
		return
	}
	fmt.Println("test eth_call method result.", result.(Reply))
}

func TestEthSyncing(t *testing.T) {
	result, err := EthSyncing()
	if err != nil {
		fmt.Errorf("test eth_syncing method error." + err.Error())
		return
	}
	fmt.Println("test eth_syncing method result.", result)
}

func TestEthGasPrice(t *testing.T) {
	result, err := EthGasPrice()
	if err != nil {
		fmt.Errorf("test eth_gasPrice method error." + err.Error())
		return
	}
	fmt.Println("test eth_gasPrice method result.", result)
}

func TestEthAccounts(t *testing.T) {
	result, err := EthAccounts()
	if err != nil {
		fmt.Errorf("test eth_accounts method error." + err.Error())
		return
	}
	fmt.Println("test eth_accounts method result.", result)
}

func TestEthGetBalance(t *testing.T) {
	result, err := EthGetBalance(MYKEY, "latest")
	if err != nil {
		fmt.Errorf("test eth_getbalance method error." + err.Error())
		return
	}
	fmt.Println("test eth_getbalance method result.", result)
}

func TestEthBlockNumber(t *testing.T) {
	result, err := EthBlockNumber()
	if err != nil {
		fmt.Errorf("test EthBlockNumber method error." + err.Error())
		return
	}
	fmt.Println("test EthBlockNumber method result.", result)
}

func TestEthGetBlockByNumber(t *testing.T) {
	result, err := EthGetBlockByNumber("0x1")
	if err != nil {
		fmt.Errorf("test EthGetBlockByNumber method error." + err.Error())
		return
	}
	fmt.Println("test EthGetBlockByNumber method result.", result.(Reply1).Result)
}

func TestEthGetBlockByHash(t *testing.T) {
	result, err := EthGetBlockByHash("0x7e90d041513ce3eb12adf33a1e4c904ee7b7f21d97c432a94c3f363ddf83d272")
	if err != nil {
		fmt.Errorf("test EthGetBlockByHash method error." + err.Error())
		return
	}
	fmt.Println("test EthGetBlockByHash method result.", result.(Reply1).Result)
}

func TestEthGetTransactionByBlockHashAndIndex(t *testing.T) {
	result, err := EthGetTransactionByBlockHashAndIndex("0x7e90d041513ce3eb12adf33a1e4c904ee7b7f21d97c432a94c3f363ddf83d272", "0x0")
	if err != nil {
		fmt.Errorf("test EthGetTransactionByBlockHashAndIndex method error." + err.Error())
		return
	}
	fmt.Println("test EthGetTransactionByBlockHashAndIndex method result.", result.(Reply3).Result)
}

func TestEthGetTransactionCount(t *testing.T) {
	result, err := EthGetTransactionCount(MYKEY, "latest")
	if err != nil {
		fmt.Errorf("test EthGetTransactionCount method error." + err.Error())
		return
	}
	fmt.Println("test EthGetTransactionCount method result.", result)
}

func TestEthSendTransaction(t *testing.T) {
	param := Param{From: MYKEY, To: NEW, Gas: "0x10000", GasPrice: "0x14", Value: "0x10000000"}
	result, err := EthSendTransaction(param)
	if err != nil {
		fmt.Errorf("test EthSendTransaction method error." + err.Error())
		return
	}
	fmt.Println("test EthSendTransaction method result.", result)
}

func TestEthGetTransactionByHash(t *testing.T) {
	result, err := EthGetTransactionByHash("0x0f06c9ef36373d4c703f499e5ed91b77368513aa5a9005d1a3b3692b5b02f906")
	if err != nil {
		fmt.Errorf("test EthGetTransactionByHash method error." + err.Error())
		return
	}
	fmt.Println("test EthGetTransactionByHash method result.", result.(Reply3).Result)
}

func TestEthGetTransactionByBlockNumberAndIndex(t *testing.T) {
	result, err := EthGetTransactionByBlockNumberAndIndex("0x1", "0x0")
	if err != nil {
		fmt.Errorf("test EthGetTransactionByBlockNumberAndIndex method error." + err.Error())
		return
	}
	fmt.Println("test EthGetTransactionByBlockNumberAndIndex method result.", result.(Reply3).Result)
}

func TestWeb3Sha3(t *testing.T) {
	result, err := Web3Sha3("0x1110")
	if err != nil {
		fmt.Errorf("test Web3Sha3 method error." + err.Error())
		return
	}
	fmt.Println("test Web3Sha3 method result.", result)
}

func TestWeb3ClientVersion(t *testing.T) {
	result, err := Web3ClientVersion()
	if err != nil {
		fmt.Errorf("test Web3ClientVersion method error." + err.Error())
		return
	}
	fmt.Println("test Web3ClientVersion method result.", result)
}

func TestNetVersion(t *testing.T) {
	result, err := NetVersion()
	if err != nil {
		fmt.Errorf("test NetVersion method error." + err.Error())
		return
	}
	fmt.Println("test NetVersion method result.", result)
}

func TestEthGetCode(t *testing.T) {
	result, err := EthGetCode(CONTRACT, "latest")
	if err != nil {
		fmt.Errorf("test EthGetCode method error." + err.Error())
		return
	}
	fmt.Println("test EthGetCode method result.", result)
}

func TestEthGetStorageAt(t *testing.T) {
	result, err := EthGetStorageAt(CONTRACT, "0x0", "latest")
	if err != nil {
		fmt.Errorf("test EthGetStorageAt method error." + err.Error())
		return
	}
	fmt.Println("test EthGetStorageAt method result.", result)
}

func TestEthEstimateGas(t *testing.T) {
	param := Param{From: MYKEY, To: "", Gas: "0x47b760", GasPrice: "0x14", Data: CODE}
	result, err := EthEstimateGas(param)
	if err != nil {
		fmt.Errorf("test EthEstimateGas method error." + err.Error())
		return
	}
	fmt.Println("test EthEstimateGas method result.", result)
}

func TestEthSignTransaction(t *testing.T) {
	param := Param{From: MYKEY, To: NEW, Gas: "0x10000", GasPrice: "0x14", Value: "0x10000000"}
	result, err := EthSignTransaction(param)
	if err != nil {
		fmt.Errorf("test EthSignTransaction method error." + err.Error())
		return
	}
	fmt.Println("test EthSignTransaction method result.", result)
}

func TestEthSendRawTransaction(t *testing.T) {
	result, err := EthSendRawTransaction("0xf8662c148301000094474b5eedfe398398c04e26082102c7bddc6beb5d841000000080824593a08d4c2221e5953af60aea444dd176a8b2eabc41f6e8aa6b45bf804d5f177c7395a00a363d3ebb134973e6e407fa1744d4d79ef2d00a0afd111012cd8a7199add16b")
	if err != nil {
		fmt.Errorf("test EthSendRawTransaction method error." + err.Error())
		return
	}
	fmt.Println("test EthSendRawTransaction method result.", result)
}

func TestEthGetTransactionReceipt(t *testing.T) {
	result, err := EthGetTransactionReceipt("0xeec6fd571dfa303a552567bfe197186c178ec986b5379af6534c985dcc908cd7")
	if err != nil {
		fmt.Errorf("test EthGetTransactionReceipt method error." + err.Error())
		return
	}
	fmt.Println("test EthGetTransactionReceipt method result.", result.(Reply3).Result)
}
