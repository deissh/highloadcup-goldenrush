package core

import (
	"github.com/deissh/highloadcup-goldenrush/client"
	"github.com/deissh/highloadcup-goldenrush/logger"
	"github.com/deissh/highloadcup-goldenrush/models"
	"sync"
)

type Explorer struct {
	client      *client.CupClient
	workerCount int

	reportChan chan *models.Report
	pointChan  chan *models.Area
}

func NewExplorer(client *client.CupClient, workerCount int) *Explorer {
	e := Explorer{client: client, workerCount: workerCount}

	e.reportChan = make(chan *models.Report)
	e.pointChan = make(chan *models.Area)

	return &e
}

func (e *Explorer) Init() {
	// todo
	zone := &models.Area{
		PosX:  0,
		PosY:  0,
		SizeX: PlayFieldX,
		SizeY: PlayFieldY,
	}

	for x := zone.PosX; x <= zone.PosX+zone.SizeX; x++ {
		for y := zone.PosY; x <= zone.PosY+zone.SizeY; y++ {
			e.pointChan <- &models.Area{
				PosX:  x,
				PosY:  y,
				SizeX: 1,
				SizeY: 1,
			}
		}
	}
}

func (e *Explorer) Start() {
	wg := &sync.WaitGroup{}

	wg.Add(e.workerCount)
	for i := 0; i < e.workerCount; i++ {
		go e.explore(wg)
	}
}

func (e *Explorer) explore(wg *sync.WaitGroup) {
	defer wg.Done()

	for point := range e.pointChan {
		report, err := e.client.ExploreArea(point)
		if err != nil {
			logger.Error.Println(err)
			continue
		}

		if *report.Amount == 0 {
			continue
		}
		e.reportChan <- report
	}
}
