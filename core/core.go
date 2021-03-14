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

func (c *Core) Start() error {
	go c.explorer.Start()
	go c.licensePool.Start()
	go c.wallet.Start()

	if IsDebug {
		go c.debug()
	}

	c.digger.Start()
	return nil
}

func (c Core) debug() {
	timer := time.Tick(time.Second * 30)
	go func() {
		for {
			<-timer
			c.digger.metrics.Print()
			c.licensePool.metrics.Print()
		}
	}()
}
