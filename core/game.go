package core

import (
	"github.com/deissh/highloadcup-goldenrush/client"
	"github.com/deissh/highloadcup-goldenrush/logger"
	"github.com/deissh/highloadcup-goldenrush/models"
)

type Game struct {
	client   *client.CupClient
	explorer *Explorer
}

func New(client *client.CupClient) *Game {
	e := NewExplorer(client, 100)

	return &Game{client, e}
}

func (g *Game) Start() error {
	g.explorer.Init()
	g.explorer.Start()

	// TODO: License pool
	license, err := g.GetLicense()
	if err != nil {
		logger.Error.Println(err)
		return err
	}

	for report := range g.explorer.reportChan {
		// TODO: goroutinize it
		// TODO: write to chan task
		left := *report.Amount
		for depth := uint8(1); depth <= PlayFieldDepth; depth++ {
			if license.DigUsed >= license.DigAllowed {
				data, err := g.GetLicense()
				if err != nil {
					logger.Error.Println(err)
					break
				}

				*license = *data
			}

			treasures, err := g.Dig(report.Area.PosX, report.Area.PosY, depth, license.ID)
			license.DigUsed++

			if err != nil {
				logger.Error.Println(err)
				continue
			}

			_ = g.CashTreasures(treasures)

			left--
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

func (g Game) GetLicense() (*models.License, error) {
	// todo: получение платной лицензии
	result, err := g.client.IssueLicense([]uint64{})
	if err != nil {
		return nil, err
	}

	return result, err
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
