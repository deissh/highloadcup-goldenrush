package core

import (
	"github.com/deissh/highloadcup-goldenrush/client"
	"github.com/deissh/highloadcup-goldenrush/models"
	"sync"
)

type Digger struct {
	client      *client.CupClient
	workerCount int

	explorer    *Explorer
	wallet      *Wallet
	licensePool *LicensePool
}

func NewDigger(
	client *client.CupClient,
	explorer *Explorer,
	wallet *Wallet,
	licensePool *LicensePool,
	workerCount int,
) *Digger {
	d := Digger{
		client:      client,
		workerCount: workerCount,
		explorer:    explorer,
		wallet:      wallet,
		licensePool: licensePool,
	}

	return &d
}

func (d *Digger) Init() {}

func (d *Digger) Start() {
	wg := &sync.WaitGroup{}

	wg.Add(d.workerCount)
	for i := 1; i <= d.workerCount; i++ {
		go d.dig(wg)
	}

	wg.Wait()
}

func (d *Digger) dig(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		report := <-d.explorer.reportChan

		left := report.Amount
		for depth := uint8(1); depth <= PlayFieldDepth; depth++ {
			lic := d.licensePool.GetLicense()

			var treasures []string

			err := d.client.Dig(&models.Dig{
				Depth:     depth,
				LicenseID: lic.ID,
				PosX:      report.Area.PosX,
				PosY:      report.Area.PosY,
			}, &treasures)
			if err != nil {
				continue
			}

			d.wallet.CashTreasures(treasures)

			left -= uint64(len(treasures))
			if left <= 0 {
				break
			}
		}
	}
}
