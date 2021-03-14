package core

import (
	"github.com/deissh/highloadcup-goldenrush/client"
	"time"
)

type Core struct {
	client *client.CupClient

	explorer    *Explorer
	licensePool *LicensePool
	wallet      *Wallet
	digger      *Digger
}

func New(client *client.CupClient) *Core {
	e := NewExplorer(client, 50)
	l := NewLicensePool(client, 20)
	w := NewWallet(client, 10)

	d := NewDigger(client, e, w, l, 100)

	return &Core{client, e, l, w, d}
}

func (g *Core) Start() error {
	go g.explorer.Start()
	go g.licensePool.Start()
	go g.wallet.Start()

	timer := time.Tick(time.Second * 30)
	go func() {
		for {
			<- timer
			g.digger.metrics.Print()
			g.licensePool.metrics.Print()
		}
	}()

	g.digger.Start()
	return nil
}
