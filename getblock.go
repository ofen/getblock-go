package getblock

import (
	"context"

	"github.com/ybbus/jsonrpc/v3"
)

const authorizationHeaderKey = "x-api-key"

// New creates Client.
func New(token string, endpoint string) *Client {
	opts := &jsonrpc.RPCClientOpts{}
	if token != "" {
		opts.CustomHeaders = map[string]string{authorizationHeaderKey: token}
	}

	return &Client{jsonrpc.NewClientWithOpts(endpoint, opts)}
}

// Client is common JSON-RPC client.
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
