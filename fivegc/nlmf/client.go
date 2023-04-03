package nlmf

import "github.com/5GCoreNet/5GCoreNetSDK/fivegc"

type Client struct {
	*BroadcastClient
	*LocationClient
}

// NewClient returns a new client for an NLMF service.
func NewClient(config fivegc.ClientConfiguration) *Client {
	return &Client{
		BroadcastClient: NewBroadcastClient(config),
		LocationClient:  NewLocationClient(config),
	}
}
