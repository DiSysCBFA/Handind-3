package client

import "github.com/DiSysCBFA/Handind-3/api"

type Client struct {
	api.ChittyChatClient
	port      string
	name      string
	timestamp int
}

// NewClient creates a new client instance
func NewClient(name string, port string) *Client {
	return &Client{
		port:      port,
		name:      name,
		timestamp: 0,
	}
}

func (c *Client) Broadcast() {

}

func (c *Client) Join() {

}
