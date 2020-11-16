package rpc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const URL = "http://123.56.170.101:8545"

//普通返回值
type Reply struct {
	JsonRpc string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Result  interface{} `json:"result"`
}

//区块信息返回值
type Reply1 struct {
	JsonRpc string  `json:"jsonrpc"`
	Id      int     `json:"id"`
	Result  *RBlock `json:"result"`
}

//签名信息返回值
type Reply2 struct {
	JsonRpc string   `json:"jsonrpc"`
	Id      int      `json:"id"`
	Result  *RawInfo `json:"result"`
}

//签名信息返回值
type Reply3 struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      int           `json:"id"`
	Result  *RTransaction `json:"result"`
}

//签名信息
type RawInfo struct {
	Raw string `json:"raw"`
}

//参数
type Param struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Data     string `json:"data"`
	Value    string `json:"value"`
}

//区块信息
type RBlock struct {
	Number           string          `json:"number"`
	Timestamp        string          `json:"timestamp"`
	Transactions     []*RTransaction `json:"transactions"`
	Mined            string          `json:"miner"`
	MinedTime        string          `json:"mined_time"`
	BlockReward      string          `json:"block_reward"`
	UnclesReward     string          `json:"uncles_reward"`
	Difficulty       string          `json:"difficulty"`
	TotalDifficulty  string          `json:"totalDifficulty"`
	Size             string          `json:"size"`
	ExtraData        string          `json:"extraData"`
	Hash             string          `json:"hash"`
	MixHash          string          `json:"mixHash"`
	ParentHash       string          `json:"parentHash"`
	Sha3uncles       string          `json:"sha3Uncles"`
	StateRoot        string          `json:"stateRoot"`
	ReceiptsRoot     string          `json:"receiptsRoot"`
	TransactionsRoot string          `json:"transactionsRoot"`
	Nonce            string          `json:"nonce"`
	GasUsed          string          `json:"gasUsed"`
	GasLimit         string          `json:"gasLimit"`
}

//交易信息
type RTransaction struct {
	Hash             string `json:"hash"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	Timestamp        string `json:"timestamp"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Gas              string `json:"gas"`
	V                string `json:"v"`
	GasPrice         string `json:"gasPrice"`
	TransactionIndex string `json:"transactionIndex"`
	Nonce            string `json:"nonce"`
	Input            string `json:"input"`
}

func EthSign(param1, param2 string) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_sign", "id": 1, "params": []interface{}{param1, param2}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthGetCode(param1, param2 string) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_getCode", "id": 1, "params": []interface{}{param1, param2}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthGetStorageAt(param1, param2, param3 string) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_getStorageAt", "id": 1, "params": []interface{}{param1, param2, param3}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthEstimateGas(param1 interface{}) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_estimateGas", "id": 1, "params": []interface{}{param1}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthSendTransaction(param1 interface{}) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_sendTransaction", "id": 1, "params": []interface{}{param1}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthSignTransaction(param1 interface{}) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_signTransaction", "id": 1, "params": []interface{}{param1}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply2
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthSendRawTransaction(param1 interface{}) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_sendRawTransaction", "id": 1, "params": []interface{}{param1}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthGetTransactionReceipt(param1 interface{}) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_getTransactionReceipt", "id": 1, "params": []interface{}{param1}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply3
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthCall(param1 interface{}, param2 string) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_call", "id": 1, "params": []interface{}{param1, param2}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthSyncing() (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_syncing", "id": 67, "params": []interface{}{}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthGasPrice() (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_gasPrice", "id": 67, "params": []interface{}{}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthAccounts() (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_accounts", "id": 67, "params": []interface{}{}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthGetBalance(param1, param2 string) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_getBalance", "id": 1, "params": []interface{}{param1, param2}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthBlockNumber() (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_blockNumber", "id": 1, "params": []interface{}{}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthGetBlockByHash(param string) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_getBlockByHash", "id": 1, "params": []interface{}{param, false}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply1
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthGetBlockByNumber(param string) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_getBlockByNumber", "id": 1, "params": []interface{}{param, false}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply1
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthGetTransactionCount(param, PARAM1 string) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_getTransactionCount", "id": 1, "params": []interface{}{param, PARAM1}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthGetTransactionByHash(param string) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_getTransactionByHash", "id": 1, "params": []interface{}{param}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply3
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthGetTransactionByBlockHashAndIndex(param, param1 string) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_getTransactionByBlockHashAndIndex", "id": 1, "params": []interface{}{param, param1}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply3
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EthGetTransactionByBlockNumberAndIndex(param, param1 string) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "eth_getTransactionByBlockNumberAndIndex", "id": 1, "params": []interface{}{param, param1}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply3
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Web3ClientVersion() (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "web3_clientVersion", "id": 1, "params": []interface{}{}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result.Result, nil
}

func NetVersion() (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "net_version", "id": 1, "params": []interface{}{}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result.Result, nil
}

func Web3Sha3(param string) (interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0", "method": "web3_sha3", "id": 1, "params": []interface{}{param}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(URL, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Reply
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
