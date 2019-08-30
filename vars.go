package host

import "errors"

// Paths
var (
	RootConfig string
	RootCache  string
)

// Errors
var (
	ErrNoURL = errors.New("url is empty")
)

// PublicGateways is a list of public IPFS gateways that can be used with `ipfs swarm connect`.
var PublicGateways = []PublicGateway{
	PublicGatewayPermaweb,
	PublicGatewayLibp2p,
}
