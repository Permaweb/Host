package host

import "os"

// Relative paths for the whole application.
// Use them with `rootConfig` and `rootCache`.
const (
	DirApp    = "/gipns"
	DirBadger = DirApp + "/badger"
	DirGit    = DirApp + "/git"
)

// Permissions
const (
	PermPrivateDirectory os.FileMode = 0700
	PermPrivateFile      os.FileMode = 0600
)

// Calculations
const (
	DownloadSpeed   = 10 * 1024 * 1024
	DownloadSeconds = 60
)

// PublicGateway is a public IPFS Gateway that supports `_dnslink`.
type PublicGateway string

// Public Gateways
const (
	PublicGatewayPermaweb PublicGateway = "/dnsaddr/permaweb.io"
	PublicGatewayLibp2p   PublicGateway = "/dnsaddr/bootstrap.libp2p.io"
)

// GitHost is a Git host.
type GitHost string

// Services are third-party Git hosts with oAuth2.
const (
	GitHostGitHub GitHost = "GitHub"
	GitHostGitLab GitHost = "GitLab"
)
