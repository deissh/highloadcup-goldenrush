package core

import (
	"github.com/deissh/highloadcup-goldenrush/client"
	"github.com/deissh/highloadcup-goldenrush/logger"
	"github.com/deissh/highloadcup-goldenrush/models"
	"go.uber.org/ratelimit"
	"sync"
)

type Explorer struct {
	client      *client.CupClient
	workerCount int

	reportChan chan *models.Report
	pointChan  chan *models.Area

	pool sync.Pool
}

func NewExplorer(client *client.CupClient, workerCount int) *Explorer {
	e := Explorer{client: client, workerCount: workerCount}

	e.reportChan = make(chan *models.Report, PlayFieldX*PlayFieldY)
	e.pointChan = make(chan *models.Area, PlayFieldX*PlayFieldY)

	e.pool = sync.Pool{
		New: func() interface{} {
			return &models.Area{PosX: 0, PosY: 0, SizeX: 1, SizeY: 1}
		},
	}

	return &e
}

func (e *Explorer) Init() {
	for x := 0; x < PlayFieldX; x++ {
		for y := 0; y < PlayFieldY; y++ {
			area := e.pool.Get().(*models.Area)

			area.PosX = uint16(x)
			area.PosY = uint16(y)

			e.pointChan <- area
		}
	}
}

func (e *Explorer) Start() {
	wg := &sync.WaitGroup{}
	rl := ratelimit.New(1000)

	wg.Add(e.workerCount)
	for i := 1; i <= e.workerCount; i++ {
		go e.explore(wg, rl)
	}
}

func (e *Explorer) explore(wg *sync.WaitGroup, rl ratelimit.Limiter) {
	defer wg.Done()

	for point := range e.pointChan {
		rl.Take()

		var report models.Report

		err := e.client.ExploreArea(point, &report)
		e.pool.Put(point)

		if err != nil {
			logger.Error.Println(err)
			continue
		}

		if report.Amount == 0 {
			continue
		}

		// todo: pool
		e.reportChan <- &report
	}
}
