package eth

import (
	"encoding/json"
	"math/big"
	"time"
)

// Block is block representations
type Block struct {
	BaseFeePerGas    *big.Int      `json:"baseFeePerGas"`
	Difficulty       *big.Int      `json:"difficulty"`
	ExtraData        string        `json:"extraData"`
	GasLimit         *big.Int      `json:"gasLimit"`
	GasUsed          *big.Int      `json:"gasUsed"`
	Hash             string        `json:"hash"`
	LogsBloom        string        `json:"logsBloom"`
	Miner            string        `json:"miner"`
	MixHash          string        `json:"mixHash"`
	Nonce            string        `json:"nonce"`
	Number           *big.Int      `json:"number"`
	ParentHash       string        `json:"parentHash"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	Sha3Uncles       string        `json:"sha3Uncles"`
	Size             *big.Int      `json:"size"`
	StateRoot        string        `json:"stateRoot"`
	Timestamp        time.Time     `json:"timestamp"`
	TotalDifficulty  *big.Int      `json:"totalDifficulty"`
	Transactions     []Transaction `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []interface{} `json:"uncles"`
}

func (t *Block) UnmarshalJSON(data []byte) error {
	type alias Block

	aux := &struct {
		BaseFeePerGas   string `json:"baseFeePerGas"`
		Difficulty      string `json:"difficulty"`
		GasLimit        string `json:"gasLimit"`
		GasUsed         string `json:"gasUsed"`
		Number          string `json:"number"`
		Size            string `json:"size"`
		Timestamp       string `json:"timestamp"`
		TotalDifficulty string `json:"totalDifficulty"`
		*alias
	}{
		alias: (*alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t.BaseFeePerGas = hex2int(aux.BaseFeePerGas)
	t.Difficulty = hex2int(aux.Difficulty)
	t.GasLimit = hex2int(aux.GasLimit)
	t.GasUsed = hex2int(aux.GasUsed)
	t.Number = hex2int(aux.Number)
	t.Size = hex2int(aux.Size)
	t.Timestamp = time.Unix(hex2int(aux.Timestamp).Int64(), 0)
	t.TotalDifficulty = hex2int(aux.TotalDifficulty)

	return nil
}

// Transaction is transaction representations
type Transaction struct {
	BlockHash            string        `json:"blockHash"`
	BlockNumber          *big.Int      `json:"blockNumber"`
	From                 string        `json:"from"`
	Gas                  *big.Int      `json:"gas"`
	GasPrice             *big.Int      `json:"gasPrice"`
	Hash                 string        `json:"hash"`
	Input                string        `json:"input"`
	Nonce                *big.Int      `json:"nonce"`
	To                   string        `json:"to"`
	TransactionIndex     *big.Int      `json:"transactionIndex"`
	Value                *big.Int      `json:"value"`
	Type                 *big.Int      `json:"type"`
	V                    *big.Int      `json:"v"`
	R                    string        `json:"r"`
	S                    string        `json:"s"`
	MaxFeePerGas         *big.Int      `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int      `json:"maxPriorityFeePerGas"`
	AccessList           []interface{} `json:"accessList"`
	ChainID              *big.Int      `json:"chainId"`
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	type alias Transaction

	aux := &struct {
		BlockNumber          string `json:"blockNumber"`
		Gas                  string `json:"gas"`
		GasPrice             string `json:"gasPrice"`
		Nonce                string `json:"nonce"`
		TransactionIndex     string `json:"transactionIndex"`
		Value                string `json:"value"`
		Type                 string `json:"type"`
		MaxFeePerGas         string `json:"maxFeePerGas"`
		MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
		V                    string `json:"v"`
		ChainID              string `json:"chainId"`
		*alias
	}{
		alias: (*alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t.BlockNumber = hex2int(aux.BlockNumber)
	t.Gas = hex2int(aux.Gas)
	t.GasPrice = hex2int(aux.GasPrice)
	t.Nonce = hex2int(aux.Nonce)
	t.TransactionIndex = hex2int(aux.TransactionIndex)
	t.Value = hex2int(aux.Value)
	t.Type = hex2int(aux.Type)
	t.MaxFeePerGas = hex2int(aux.MaxFeePerGas)
	t.MaxPriorityFeePerGas = hex2int(aux.MaxPriorityFeePerGas)
	t.V = hex2int(aux.V)
	t.ChainID = hex2int(aux.ChainID)

	return nil
}
