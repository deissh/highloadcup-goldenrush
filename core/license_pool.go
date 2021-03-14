package core

import (
	"github.com/deissh/highloadcup-goldenrush/client"
	"github.com/deissh/highloadcup-goldenrush/models"
	"log"
	"sync"
)

type LicensePool struct {
	client      *client.CupClient
	workerCount int

	licenseChan chan *models.License
	issueChan   chan struct{}
	pool        sync.Pool

	metrics LicensePoolMetrics
}

type LicensePoolMetrics struct {
	sync.Mutex

	FreeLicenses uint32
	PaidLicenses uint32
	FreeLicensesIssued uint32
	PaidLicensesIssued uint32
}

func (l *LicensePoolMetrics) Print() {
	log.Println("*** LICENSE POOL REPORT ***")
	log.Println("Free licenses", l.FreeLicenses, "total", l.FreeLicensesIssued)
	log.Println("Paid licenses", l.PaidLicenses, "total", l.FreeLicensesIssued)
}

func NewLicensePool(client *client.CupClient, workerCount int) *LicensePool {
	l := LicensePool{client: client, workerCount: workerCount}

	l.licenseChan = make(chan *models.License, MaxLicenses)
	l.issueChan = make(chan struct{}, MaxLicenses)
	l.pool = sync.Pool{
		New: func() interface{} {
			return &models.License{}
		},
	}

	return &l
}

func (l *LicensePool) GetLicense() *models.License {
	select {
	case license := <- l.licenseChan: {
		if license.DigUsed == license.DigAllowed {
			// remove old license and request new
			l.pool.Put(license)
			l.issueChan <- struct{}{}

			return l.GetLicense()
		}

		license.DigUsed += 1
		l.licenseChan <- license

		return license
	}
	default:
		l.issueChan <- struct{}{}
	}

	return l.GetLicense()
}

func (l *LicensePool) Init() {}

func (l *LicensePool) Start() {
	wg := &sync.WaitGroup{}

	l.Init()

	wg.Add(l.workerCount)
	for i := 1; i <= l.workerCount; i++ {
		go l.issueLicense(wg)
	}

	wg.Wait()
}

func (l *LicensePool) issueLicense(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		<-l.issueChan

		lic := l.pool.Get().(*models.License)

		err := l.client.IssueLicense([]uint64{}, lic)
		if err != nil {
			continue
		}

		l.licenseChan <- lic

		if IsDebug {
			l.metrics.Lock()
			l.metrics.FreeLicenses = uint32(len(l.licenseChan))
			l.metrics.PaidLicenses = uint32(len(l.licenseChan))
			l.metrics.FreeLicensesIssued += 1
			l.metrics.PaidLicensesIssued += 1
			l.metrics.Unlock()
		}
	}
}
