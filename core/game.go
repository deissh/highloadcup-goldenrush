package core

import (
	"context"
	"github.com/deissh/highloadcup-goldenrush/client"
	"github.com/deissh/highloadcup-goldenrush/client/operations"
	"github.com/deissh/highloadcup-goldenrush/logger"
	"github.com/deissh/highloadcup-goldenrush/models"
)

type Area struct {
	x     int64
	y     int64
	depth int64
}

type Game struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	client    *client.HighLoadCup2021

	area *Area
}

func New(client *client.HighLoadCup2021) *Game {
	ctx, cancel := context.WithCancel(context.Background())

	return &Game{
		ctx,
		cancel,
		client,
		&Area{
			x:     3500,
			y:     3500,
			depth: 10,
		},
	}
}

func (g Game) Start() error {
	for x := int64(0); x < g.area.x; x++ {
		for y := int64(0); y < g.area.x; y++ {
			amount, err := g.Explore(x, y)
			if err != nil || amount == nil {
				logger.Error.Println(err)
				continue
			}

			left := int64(*amount)

			license, err := g.GetLicense()
			if err != nil {
				logger.Error.Println(err)
				continue
			}

			for depth := int64(1); depth <= 10; depth++ {
				treasures, err := g.Dig(x, y, depth, *license.ID)
				if err != nil {
					logger.Error.Println(err)
					continue
				}

				if err = g.CashTreasures(treasures); err != nil {
					logger.Error.Println(err)
				}

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

func (g Game) Explore(x, y int64) (*models.Amount, error) {
	area := models.Area{
		PosX:  &x,
		PosY:  &y,
		SizeX: 1,
		SizeY: 1,
	}

	result, err := g.client.Operations.ExploreArea(
		operations.NewExploreAreaParamsWithContext(g.ctx).
			WithArgs(&area),
	)
	if err != nil {
		return nil, err
	}

	return result.GetPayload().Amount, nil
}

func (g Game) CashTreasures(list models.TreasureList) error {
	for _, id := range list {
		_, err := g.client.Operations.Cash(
			operations.NewCashParamsWithContext(g.ctx).WithArgs(id),
		)
		if err != nil {
			logger.Warn.Println(id, " not cashed")
		}
	}

	return nil
}

func (g Game) GetLicense() (*models.License, error) {
	// todo: получение платной лицензии
	coins := models.Wallet{}

	result, err := g.client.Operations.IssueLicense(
		operations.NewIssueLicenseParamsWithContext(g.ctx).
			WithArgs(coins),
	)
	if err != nil {
		return nil, err
	}

	return result.GetPayload(), err
}

func (g Game) Dig(x, y, depth, license int64) (models.TreasureList, error) {
	dig := models.Dig{
		Depth:     &depth,
		LicenseID: &license,
		PosX:      &x,
		PosY:      &y,
	}

	result, err := g.client.Operations.Dig(
		operations.NewDigParamsWithContext(g.ctx).
			WithArgs(&dig),
	)
	if err != nil {
		return nil, err
	}

	return result.GetPayload(), nil
}
