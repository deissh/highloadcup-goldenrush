package client

import (
	"gopkg.in/h2non/gentleman-retry.v2"
	"gopkg.in/h2non/gentleman.v2"
)

type CupClient struct {
	Service
}

// New creates a new high load cup2021 HTTP client,
// using a customizable transport config.
func New(cfg *TransportConfig) *CupClient {
	cli := CupClient{}
	cli.Service = newService(
		gentleman.
			New().
			BaseURL(cfg.Host).
			Use(retry.New(retry.ConstantBackoff)),
	)

	return &cli
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}
