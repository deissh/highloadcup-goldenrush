package core

import (
	"github.com/deissh/highloadcup-goldenrush/client"
	"github.com/deissh/highloadcup-goldenrush/logger"
	"github.com/deissh/highloadcup-goldenrush/models"
)

type Game struct {
	client *client.CupClient

	explorer    *Explorer
	licensePool *LicensePool
}

func New(client *client.CupClient) *Game {
	e := NewExplorer(client, 100)
	l := NewLicensePool(client, 1)

	return &Game{client, e, l}
}

func (g *Game) Start() error {
	g.explorer.Init()
	g.explorer.Start()

	g.licensePool.Init()
	g.licensePool.Start()

	for report := range g.explorer.reportChan {
		// TODO: goroutinize it
		// TODO: write to chan task
		left := report.Amount
		for depth := uint8(1); depth <= PlayFieldDepth; depth++ {
			lic := g.licensePool.GetLicense()

			treasures, err := g.Dig(report.Area.PosX, report.Area.PosY, depth, lic.ID)
			if err != nil {
				continue
			}

			_ = g.CashTreasures(treasures)

			left -= uint64(len(treasures))
			if left <= 0 {
				break
			}
		}
	}

	return nil
}

func (g Game) CashTreasures(list models.TreasureList) error {
	for _, id := range list {
		data, err := g.client.Cash(id)
		if err != nil {
			logger.Warn.Println(id, " not cashed")
			return err
		}

		logger.Info.Println(id, " cashed with", data)
	}

	return nil
}

func (g Game) Dig(x, y uint16, depth uint8, license uint64) (models.TreasureList, error) {
	result, err := g.client.Dig(&models.Dig{
		Depth:     depth,
		LicenseID: license,
		PosX:      x,
		PosY:      y,
	})
	if err != nil {
		return nil, err
	}

	return *result, nil
}
