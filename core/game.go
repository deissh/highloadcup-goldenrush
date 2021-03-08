package core

import (
	"context"
	"github.com/deissh/highloadcup-goldenrush/client"
	"github.com/deissh/highloadcup-goldenrush/logger"
	"github.com/deissh/highloadcup-goldenrush/models"
)

const (
	PlayFieldX     = 3500
	PlayFieldY     = 3500
	PlayFieldDepth = 10
)

type Game struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	client    *client.CupClient
}

func New(client *client.CupClient) *Game {
	ctx, cancel := context.WithCancel(context.Background())

	return &Game{
		ctx,
		cancel,
		client,
	}
}

func (g *Game) Start() error {
	// TODO: License pool
	license, err := g.GetLicense()
	if err != nil {
		logger.Error.Println(err)
		return err
	}

	for x := uint16(0); x < PlayFieldX; x++ {
		for y := uint16(0); y < PlayFieldY; y++ {
			amount, err := g.Explore(x, y)
			if err != nil {
				logger.Error.Println(err)
				continue
			}

			left := int64(*amount)
			if left == 0 {
				continue
			}

			// TODO: goroutinize it
			// TODO: write to chan task
			for depth := uint8(1); depth <= PlayFieldDepth; depth++ {
				if license.DigUsed >= license.DigAllowed {
					data, err := g.GetLicense()
					if err != nil {
						logger.Error.Println(err)
						continue
					}

					*license = *data
				}

				treasures, err := g.Dig(x, y, depth, license.ID)
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
	}

	return nil
}

func (g Game) Stop() error {
	defer g.ctxCancel()
	logger.Info.Println("Stopping context and all goroutines")

	return nil
}

func (g Game) Explore(x, y uint16) (*models.Amount, error) {
	result, err := g.client.ExploreArea(&models.Area{
		PosX:  x,
		PosY:  y,
		SizeX: 1,
		SizeY: 1,
	})
	if err != nil {
		return nil, err
	}

	return result.Amount, nil
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
