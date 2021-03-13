package core

import (
	"github.com/deissh/highloadcup-goldenrush/client"
	"github.com/deissh/highloadcup-goldenrush/models"
	"sync"
)

type LicensePool struct {
	client      *client.CupClient
	workerCount int

	licenseChan chan *models.License
	issueChan   chan struct{}
	pool        sync.Pool
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
	license := <-l.licenseChan

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

func (l *LicensePool) Init() {
	// request first MaxLicenses licenses
	for i := 0; i < MaxLicenses; i++ {
		l.issueChan <- struct{}{}
	}
}

func (l *LicensePool) Start() {
	wg := &sync.WaitGroup{}

	wg.Add(l.workerCount)
	for i := 1; i <= l.workerCount; i++ {
		go l.issueLicense(wg)
	}
}

func (l *LicensePool) issueLicense(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		<- l.issueChan

		data, err := l.client.IssueLicense([]uint64{})
		if err != nil {
			continue
		}

		lic := l.pool.Get().(*models.License)
		lic.ID = data.ID
		lic.DigAllowed = data.DigAllowed
		lic.DigUsed = data.DigUsed

		l.licenseChan <- lic
	}
}
