package core

import (
	"github.com/deissh/highloadcup-goldenrush/client"
)

type Core struct {
	client *client.CupClient

	explorer    *Explorer
	licensePool *LicensePool
	wallet      *Wallet
	digger      *Digger
}

func New(client *client.CupClient) *Core {
	e := NewExplorer(client, 100)
	l := NewLicensePool(client, 5)
	w := NewWallet(client, 5)

	d := NewDigger(client, e, w, l, 20)

	return &Core{client, e, l, w, d}
}

func (g *Core) Start() error {
	go g.explorer.Start()
	go g.licensePool.Start()
	go g.wallet.Start()

	g.digger.Start()
	return nil
}
