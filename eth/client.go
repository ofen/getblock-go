package eth

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ofen/getblock-go"
)

const (
	// Wei is the smallest denomination of ether.
	Wei = 1
	// KWei (or "babbage") is equal to 1.000 Wei.
	KWei = 1e3
	// MWei (or "lovelace") is equal to 1.000.000 Wei.
	MWei = 1e6
	// GWei (or "shannon") is equal to 1.000.000.000 Wei.
	GWei = 1e9
	// TWei (or "szabo" or "microether") is equal to 1.000.000.000.000 Wei.
	TWei = 1e12
	// Pwei (or "finney" or "milliether") is equal to 1.000.000.000.000.000 Wei.
	Pwei = 1e15
	// Ether is the main transaction token in ethereum network. It equals to 1.000.000.000.000.000.000 Wei.
	Ether = 1e18
)

// Endpoint is default endpoint for etherium mainnet.
const Endpoint = "https://eth.getblock.io/mainnet/"

// New creates JSON-RPC client for https://eth.getblock.io/mainnet/.
func New(token string) *Client {
	return &Client{getblock.New(token, Endpoint)}
}

// Client is JSON-RPC client
type Client struct {
	Client *getblock.Client
}

// BlockNumber returns the index corresponding to the block number of the current chain head
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_blockNumber/.
func (c *Client) BlockNumber(ctx context.Context) (*big.Int, error) {
	r, err := c.Client.Call(ctx, "eth_blockNumber")
	if err != nil {
		return nil, err
	}

	v, ok := r.Result.(string)
	if !ok {
		return nil, fmt.Errorf("value is not string: %v", r.Result)
	}

	return hex2int(v), nil
}

// GetBlockByNumber returns information about a block by block number
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getBlockByNumber/.
func (c *Client) GetBlockByNumber(ctx context.Context, blockNumber *big.Int, detailedTransactions bool) (*Block, error) {
	bn := int2hex(blockNumber)
	r, err := c.Client.Call(ctx, "eth_getBlockByNumber", bn, detailedTransactions)
	if err != nil {
		return nil, err
	}

	v := &Block{}
	err = r.GetObject(&v)

	return v, err

}

// Accounts returns a list of account addresses a client owns.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_accounts
func (c *Client) Accounts() {}

// Call invokes a contract function locally and does not change the state of the blockchain.
//
// You can interact with contracts using eth_sendRawTransaction or eth_call.
//
// If revert reason is enabled with --revert-reason-enabled, the eth_call error response will include the revert reason.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_call
func (c *Client) Call() {}

// ChainID Returns the chain ID.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_chainId
func (c *Client) ChainID() {}

// Coinbase returns the client coinbase address. The coinbase address is the account to pay mining rewards to.
//
//To set a coinbase address, start Besu with the --miner-coinbase option set to a valid Ethereum account address. You can get the Ethereum account address from a client such as MetaMask or Etherscan.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_coinbase
func (c *Client) Coinbase() {}

// EstimateGas returns an estimate of the gas required for a transaction to complete. The estimation process does not use gas and the transaction is not added to the blockchain. The resulting estimate can be greater than the amount of gas the transaction ends up using, for reasons including EVM mechanics and node performance.
//
// The eth_estimateGas call does not send a transaction. You must call eth_sendRawTransaction to execute the transaction.
//
// If revert reason is enabled with --revert-reason-enabled, the eth_estimateGas error response will include the revert reason.
//
//https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_estimateGas
func (c *Client) EstimateGas() {}

// GasPrice returns a percentile gas unit price for the most recent blocks, in Wei. By default, the last 100 blocks are examined and the 50th percentile gas unit price (that is, the median value) is returned.
//
// If there are no blocks, the value for --min-gas-price is returned. The value returned is restricted to values between --min-gas-price and --api-gas-price-max. By default, 1000 Wei and 500GWei.
// Use the --api-gas-price-blocks, --api-gas-price-percentile , and --api-gas-price-max command line options to configure the eth_gasPrice default values.
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_gasPrice
func (c *Client) GasPrice() {}

// GetBalance returns the account balance of the specified address.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getBalance
func (c *Client) GetBalance() {}

// GetBlockByHash returns information about the block by hash.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getBlockByHash
func (c *Client) GetBlockByHash() {}

// GetBlockTransactionCountByHash returns the number of transactions in the block matching the given block hash.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getBlockTransactionCountByHash
func (c *Client) GetBlockTransactionCountByHash() {}

// GetBlockTransactionCountByNumber returns the number of transactions in a block matching the specified block number.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getBlockTransactionCountByNumber
func (c *Client) GetBlockTransactionCountByNumber() {}

// GetCode returns the code of the smart contract at the specified address. Besu stores compiled smart contract code as a hexadecimal value.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getCode
func (c *Client) GetCode() {}

// GetFilterChanges polls the specified filter and returns an array of changes that have occurred since the last poll.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getFilterChanges
func (c *Client) GetFilterChanges() {}

// GetFilterLogs returns an array of logs for the specified filter.
// Leave the --auto-log-bloom-caching-enabled command line option at the default value of true to improve log retrieval performance.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getFilterLogs
func (c *Client) GetFilterLogs() {}

// GetLogs returns an array of logs matching a specified filter object.
//
// Leave the --auto-log-bloom-caching-enabled command line option at the default value of true to improve log retrieval performance.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getLogs
func (c *Client) GetLogs() {}

// GetMinerDataByBlockHash returns miner data for the specified block.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getMinerDataByBlockHash
func (c *Client) GetMinerDataByBlockHash() {}

// GetMinerDataByBlockNumber returns miner data for the specified block.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getMinerDataByBlockNumber
func (c *Client) GetMinerDataByBlockNumber() {}

// GetProof returns the account and storage values of the specified account, including the Merkle proof.
//
// The API allows IoT devices or mobile apps which are unable to run light clients to verify responses from untrusted sources, by using a trusted block hash.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getProof
func (c *Client) GetProof() {}

// GetStorageAt returns the value of a storage position at a specified address.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getStorageAt
func (c *Client) GetStorageAt() {}

// GetTransactionByBlockHashAndIndex returns transaction information for the specified block hash and transaction index position.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getTransactionByBlockHashAndIndex
func (c *Client) GetTransactionByBlockHashAndIndex() {}

// GetTransactionByBlockNumberAndIndex returns transaction information for the specified block number and transaction index position.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getTransactionByBlockNumberAndIndex
func (c *Client) GetTransactionByBlockNumberAndIndex() {}

// GetTransactionByHash returns transaction information for the specified transaction hash.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getTransactionByHash
func (c *Client) GetTransactionByHash() {}

// GetTransactionCount returns the number of transactions sent from a specified address. Use the pending tag to get the next account nonce not used by any pending transactions.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getTransactionCount
func (c *Client) GetTransactionCount() {}

// GetTransactionReceipt returns the receipt of a transaction by transaction hash. Receipts for pending transactions are not available.
//
// If you enabled revert reason, the receipt includes available revert reasons in the response.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getTransactionReceipt
func (c *Client) GetTransactionReceipt() {}

// GetUncleByBlockHashAndIndex returns uncle specified by block hash and index.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getUncleByBlockHashAndIndex
func (c *Client) GetUncleByBlockHashAndIndex() {}

// GetUncleByBlockNumberAndIndex returns uncle specified by block number and index.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getUncleByBlockNumberAndIndex
func (c *Client) GetUncleByBlockNumberAndIndex() {}

// GetUncleCountByBlockHash returns the number of uncles in a block from a block matching the given block hash.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getUncleCountByBlockHash
func (c *Client) GetUncleCountByBlockHash() {}

// GetUncleCountByBlockNumber returns the number of uncles in a block matching the specified block number.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getUncleCountByBlockNumber
func (c *Client) GetUncleCountByBlockNumber() {}

// GetWork returns the hash of the current block, the seed hash, and the required target boundary condition.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_getWork
func (c *Client) GetWork() {}

// Hashrate returns the number of hashes per second with which the node is mining.
//
// When the stratum server is enabled, this method returns the cumulative hashrate of all sealers reporting their hashrate.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_hashrate
func (c *Client) Hashrate() {}

// Mining returns whether the client is actively mining new blocks. Besu pauses mining while the client synchronizes with the network regardless of command settings or methods called.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_mining
func (c *Client) Mining() {}

// NewBlockFilter creates a filter to retrieve new block hashes. To poll for new blocks, use eth_getFilterChanges.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_newBlockFilter
func (c *Client) NewBlockFilter() {}

// NewFilter creates a log filter. To poll for logs associated with the created filter, use eth_getFilterChanges. To get all logs associated with the filter, use eth_getFilterLogs.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_newFilter
func (c *Client) NewFilter() {}

// NewPendingTransactionFilter creates a filter to retrieve new pending transactions hashes. To poll for new pending transactions, use eth_getFilterChanges.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_newPendingTransactionFilter
func (c *Client) NewPendingTransactionFilter() {}

// ProtocolVersion returns current Ethereum protocol version.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_protocolVersion
func (c *Client) ProtocolVersion() {}

// SendRawTransaction sends a signed transaction. A transaction can send ether, deploy a contract, or interact with a contract. Set the maximum transaction fee for transactions using the --rpc-tx-feecap CLI option.
//
// You can interact with contracts using eth_sendRawTransaction or eth_call.
//
// To avoid exposing your private key, create signed transactions offline and send the signed transaction data using eth_sendRawTransaction.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_sendRawTransaction
func (c *Client) SendRawTransaction() {}

// SubmitHashrate submits the mining hashrate.
//
// Used by mining software such as Ethminer.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_submitHashrate
func (c *Client) SubmitHashrate() {}

// SubmitWork submits a Proof of Work (Ethash) solution.
//
// Used by mining software such as Ethminer.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_submitWork
func (c *Client) SubmitWork() {}

// Syncing returns an object with data about the synchronization status, or false if not synchronizing.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_syncing
func (c *Client) Syncing() {}

// UninstallFilter uninstalls a filter with the specified ID. When a filter is no longer required, call this method.
//
// Filters time out when not requested by eth_getFilterChanges or eth_getFilterLogs for 10 minutes.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/eth_uninstallFilter
func (c *Client) UninstallFilter() {}

// Enode returns the enode URL.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/net_enode
func (c *Client) Enode() {}

// Listening returns whether the client is actively listening for network connections.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/net_listening
func (c *Client) Listening() {}

// PeerCount returns the number of peers currently connected to the client.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/net_peerCount
func (c *Client) PeerCount() {}

// Services returns enabled services (for example, jsonrpc) and the host and port for each service.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/net_services
func (c *Client) Services() {}

// Version returns the network ID.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/net_version
func (c *Client) Version() {}

// Block provides transaction processing of type trace for the specified block.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/trace_block
func (c *Client) Block() {}

// ReplayBlockTransactions provides transaction processing tracing per block.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/trace_replayBlockTransactions
func (c *Client) ReplayBlockTransactions() {}

// Transaction provides transaction processing of type trace for the specified transaction.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/trace_transaction
func (c *Client) Transaction() {}

// BesuPendingTransactions lists pending transactions that match the supplied filter conditions.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/txpool_besuPendingTransactions
func (c *Client) BesuPendingTransactions() {}

// BesuStatistics lists statistics about the node transaction pool.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/txpool_besuStatistics
func (c *Client) BesuStatistics() {}

// BesuTransactions lists transactions in the node transaction pool.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/txpool_besuTransactions
func (c *Client) BesuTransactions() {}

// ClientVersion returns the current client version.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/web3_clientVersion
func (c *Client) ClientVersion() {}

// SHA3 returns a SHA3 hash of the specified data. The result value is a Keccak-256 hash, not the standardized SHA3-256.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/web3_sha3
func (c *Client) SHA3() {}

// RPCModules lists enabled APIs and the version of each.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/rpc_modules
func (c *Client) RPCModules() {}

// ReloadPluginConfig reloads specified plugin configuration.
//
// https://getblock.io/docs/available-nodes-methods/ETH/JSON-RPC/plugins_reloadPluginConfig
func (c *Client) ReloadPluginConfig() {}

// Wei2ether converts wei to ether
func Wei2ether(val *big.Int) *big.Float {
	f := new(big.Float).SetInt(val)
	return f.Quo(f, big.NewFloat(Ether))
}

func hex2int(s string) *big.Int {
	i := new(big.Int)
	if s == "" {
		return i
	}

	_, ok := i.SetString(s, 0)
	if !ok {
		panic(fmt.Sprintf("valid is not big integer: %s", s))
	}

	return i
}

func int2hex(i *big.Int) string {
	return fmt.Sprintf("%#x", i)
}
