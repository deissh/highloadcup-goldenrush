// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"errors"
	"github.com/deissh/highloadcup-goldenrush/models"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

// New creates a new operations API client.
func newService(transport *gentleman.Client) Service {
	return &client{transport: transport}
}

type client struct {
	transport *gentleman.Client
}

// Service is the interface for client methods
type Service interface {
	Cash(id string) (*[]uint32, error)

	Dig(params *models.Dig) (*models.TreasureList, error)

	ExploreArea(params *models.Area) (*models.Report, error)

	GetBalance() (*models.Balance, error)

	HealthCheck() error

	IssueLicense(coins []uint64) (*models.License, error)

	ListLicenses() (*[]models.License, error)
}

func (c client) Cash(id string) (*[]uint32, error) {
	req := c.transport.Request()
	req.Path("/cash")
	req.Method("POST")
	req.SetHeader("accept", "application/json")
	req.SetHeader("Content-Type", "application/json")
	req.BodyString("\"" + id + "\"")

	res, err := req.Send()
	if err != nil {
		return nil, err
	}

	if !res.Ok {
		var errRes models.Error
		res.JSON(&errRes)
		return nil, errors.New("req err: " + errRes.Message)
	}

	data := make([]uint32, 0)
	if err = res.JSON(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (c client) Dig(params *models.Dig) (*models.TreasureList, error) {
	req := c.transport.Request()
	req.Path("/dig")
	req.Method("POST")
	req.Use(body.JSON(params))

	res, err := req.Send()
	if err != nil {
		return nil, err
	}

	if !res.Ok {
		var errRes models.Error
		res.JSON(&errRes)
		return nil, errors.New("req err: " + errRes.Message)
	}

	var data models.TreasureList
	if err = res.JSON(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (c client) ExploreArea(params *models.Area) (*models.Report, error) {
	req := c.transport.Request()
	req.Path("/explore")
	req.Method("POST")
	req.Use(body.JSON(params))

	res, err := req.Send()
	if err != nil {
		return nil, err
	}

	if !res.Ok {
		var errRes models.Error
		res.JSON(&errRes)
		return nil, errors.New("req err: " + errRes.Message)
	}

	var data models.Report
	if err = res.JSON(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (c client) GetBalance() (*models.Balance, error) {
	req := c.transport.Request()
	req.Path("/balance")
	req.Method("GET")

	res, err := req.Send()
	if err != nil {
		return nil, err
	}

	if !res.Ok {
		var errRes models.Error
		res.JSON(&errRes)
		return nil, errors.New("req err: " + errRes.Message)
	}

	var data models.Balance
	if err = res.JSON(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (c client) HealthCheck() error {
	req := c.transport.Request()
	req.Path("/health-check")
	req.Method("GET")

	res, err := req.Send()
	if err != nil {
		return err
	}

	if !res.Ok {
		var errRes models.Error
		res.JSON(&errRes)
		return errors.New(errRes.Message)
	}

	if res.StatusCode != 200 {
		return errors.New("not ready")
	}

	return nil
}

func (c client) IssueLicense(coins []uint64) (*models.License, error) {
	req := c.transport.Request()
	req.Path("/licenses")
	req.Method("POST")
	req.Use(body.JSON(coins))

	res, err := req.Send()
	if err != nil {
		return nil, err
	}

	if !res.Ok {
		var errRes models.Error
		res.JSON(&errRes)
		return nil, errors.New("req err: " + errRes.Message)
	}

	var data models.License
	if err = res.JSON(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (c client) ListLicenses() (*[]models.License, error) {
	req := c.transport.Request()
	req.Path("/licenses")
	req.Method("GET")

	res, err := req.Send()
	if err != nil {
		return nil, err
	}

	if !res.Ok {
		var errRes models.Error
		res.JSON(&errRes)
		return nil, errors.New("req err: " + errRes.Message)
	}

	data := make([]models.License, 0)
	if err = res.JSON(&data); err != nil {
		return nil, err
	}

	return &data, nil
}