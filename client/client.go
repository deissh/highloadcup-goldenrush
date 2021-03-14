package client

import (
	"github.com/valyala/fasthttp"
)

type CupClient struct {
	Service
}

// New creates a new high load cup2021 HTTP client,
// using a customizable transport config.
func New(cfg *TransportConfig) *CupClient {
	cli := CupClient{}
	cli.Service = newService(
		&fasthttp.Client{
			NoDefaultUserAgentHeader: true,
			DisablePathNormalizing: false,
			DisableHeaderNamesNormalizing: false,
			MaxIdemponentCallAttempts: 3,
			RetryIf:                   func(req *fasthttp.Request) bool { return true },
		},
		cfg.BaseUrl,
	)

	return &cli
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	BaseUrl string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.BaseUrl = host
	return cfg
}
