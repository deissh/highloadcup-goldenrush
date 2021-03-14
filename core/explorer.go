package core

import (
	"github.com/deissh/highloadcup-goldenrush/client"
	"github.com/deissh/highloadcup-goldenrush/models"
	"log"
	"sync"
)

type Explorer struct {
	client      *client.CupClient
	workerCount int

	reportChan chan *models.Report
	pointChan  chan *models.Area

	pointPool sync.Pool
	reportPool sync.Pool
}

func NewExplorer(client *client.CupClient, workerCount int) *Explorer {
	e := Explorer{client: client, workerCount: workerCount}

	e.reportChan = make(chan *models.Report, workerCount * 50)
	e.pointChan = make(chan *models.Area, PlayFieldX*PlayFieldY)

	e.pointPool = sync.Pool{
		New: func() interface{} {
			return &models.Area{PosX: 0, PosY: 0, SizeX: 1, SizeY: 1}
		},
	}
	e.reportPool = sync.Pool{
		New: func() interface{} {
			return &models.Report{}
		},
	}

	return &e
}

func (e *Explorer) Init() {
	for x := 0; x < PlayFieldX; x++ {
		for y := 0; y < PlayFieldY; y++ {
			area := e.pointPool.Get().(*models.Area)

			area.PosX = uint16(x)
			area.PosY = uint16(y)

			e.pointChan <- area
		}
	}
}

func (e *Explorer) Start() {
	wg := &sync.WaitGroup{}

	e.Init()

	wg.Add(e.workerCount)
	for i := 1; i <= e.workerCount; i++ {
		go e.explore(wg)
	}

	wg.Wait()
}

func (e *Explorer) ReleaseReport(ptr interface{}) {
	e.reportPool.Put(ptr)
}

func (e *Explorer) explore(wg *sync.WaitGroup) {
	defer wg.Done()

	for point := range e.pointChan {
		report := e.reportPool.Get().(*models.Report)

		err := e.client.ExploreArea(point, report)
		e.pointPool.Put(point)

		if err != nil {
			log.Println(err)
			continue
		}

		if report.Amount == 0 {
			continue
		}

		e.reportChan <- report
	}
}
