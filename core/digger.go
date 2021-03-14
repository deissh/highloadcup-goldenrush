package core

import (
	"github.com/deissh/highloadcup-goldenrush/client"
	"github.com/deissh/highloadcup-goldenrush/models"
	"log"
	"sync"
	"time"
)

type Digger struct {
	client      *client.CupClient
	workerCount int

	explorer    *Explorer
	wallet      *Wallet
	licensePool *LicensePool

	metrics DiggerMetrics
}

type DiggerMetrics struct {
	sync.Mutex

	DigsDone uint32
	TreasuresExchanged uint32
	TreasuresExchangedValue uint32

	LicenseTimeTotal time.Duration
	WalletTimeTotal time.Duration
	ReportTimeTotal time.Duration
}

func (d *DiggerMetrics) Print() {
	log.Println("*** DIGGER REPORT ***")
	log.Println("Digs total", d.DigsDone)
	log.Println("Digger treasures total", d.TreasuresExchanged)
	log.Println("Digger treasures value", d.TreasuresExchangedValue)
	log.Println("Digger wait for cashier total", d.WalletTimeTotal.String())
	log.Println("Digger wait for license total", d.LicenseTimeTotal.String())
	log.Println("Digger wait for report total", d.ReportTimeTotal.String())
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
	var cashierTime time.Time
	var licenseTime time.Time
	var reportTime time.Time

	for {
		if IsDebug {
			reportTime = time.Now()
		}
		report := <-d.explorer.reportChan
		if IsDebug {
			d.metrics.Lock()
			d.metrics.ReportTimeTotal += time.Now().Sub(reportTime)
			d.metrics.Unlock()
		}

		left := report.Amount
		for depth := uint8(1); depth <= PlayFieldDepth; depth++ {

			if IsDebug {
				licenseTime = time.Now()
			}
			lic := d.licensePool.GetLicense()
			if IsDebug {
				d.metrics.Lock()
				d.metrics.LicenseTimeTotal += time.Now().Sub(licenseTime)
				d.metrics.Unlock()
			}

			var treasures []string

			err := d.client.Dig(&models.Dig{
				Depth:     depth,
				LicenseID: lic.ID,
				PosX:      report.Area.PosX,
				PosY:      report.Area.PosY,
			}, &treasures)

			if IsDebug {
				d.metrics.Lock()
				d.metrics.DigsDone += 1
				d.metrics.Unlock()
			}

			if err != nil {
				continue
			}

			if IsDebug {
				cashierTime = time.Now()
			}
			d.wallet.CashTreasures(treasures)
			if IsDebug {
				d.metrics.Lock()
				d.metrics.TreasuresExchanged += 1
				d.metrics.TreasuresExchangedValue += uint32(len(treasures))
				d.metrics.WalletTimeTotal += time.Now().Sub(cashierTime)
				d.metrics.Unlock()
			}

			left -= uint64(len(treasures))
			if left <= 0 {
				break
			}
		}

		d.explorer.ReleaseReport(report)
	}
}
