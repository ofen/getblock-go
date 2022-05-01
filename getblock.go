package getblock

import (
	"context"

	"github.com/ybbus/jsonrpc/v3"
)

// Endpoint is default endpoint for etherium mainnet.
const Endpoint = "https://eth.getblock.io/mainnet/"

const authorizationHeaderKey = "x-api-key"

type ClientOptions struct {
	Endpoint string
}

// New creates JSON-RPC client for https://eth.getblock.io/mainnet/.
func New(token string, options *ClientOptions) *Client {
	endpoint := Endpoint
	if options != nil {
		if options.Endpoint != "" {
			endpoint = options.Endpoint
		}
	}
	return &Client{
		Client: jsonrpc.NewClientWithOpts(endpoint, &jsonrpc.RPCClientOpts{
			CustomHeaders: map[string]string{authorizationHeaderKey: token},
		}),
	}
}

// Client is JSON-RPC client
type Client struct {
	Client jsonrpc.RPCClient
}

// Call sends request to JSON-RPC endpoint.
// Repeats request on 5xx error up to 5 times.
func (c *Client) Call(ctx context.Context, method string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	var r *jsonrpc.RPCResponse
	var err error
	for i := 0; i < 5; i++ {
		r, err = c.Client.Call(ctx, method, params)
		if err == nil {
			break
		}

		e, ok := err.(*jsonrpc.HTTPError)
		if ok {
			if e.Code < 500 {
				break
			}
		}
	}

	return r, err
}
