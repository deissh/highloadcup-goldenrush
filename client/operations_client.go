package client

import (
	"encoding/json"
	"errors"
	"github.com/deissh/highloadcup-goldenrush/logger"
	"github.com/deissh/highloadcup-goldenrush/models"
	"github.com/valyala/fasthttp"
)

// New creates a new operations API client.
func newService(transport *fasthttp.Client, baseUrl string) Service {
	return &client{transport, baseUrl}
}

type client struct {
	transport *fasthttp.Client
	baseUrl   string
}

// Service is the interface for client methods
type Service interface {
	Cash(id string, target *[]uint) error

	Dig(params *models.Dig, target *[]string) error

	ExploreArea(params *models.Area, target *models.Report) error

	GetBalance(target *models.Balance) error

	HealthCheck() error

	IssueLicense(coins []uint64, target *models.License) error

	ListLicenses(target *[]models.License) error
}

func (c client) Cash(id string, target *[]uint) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(c.baseUrl + "/cash")
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	req.SetBody([]byte("\"" + id + "\""))

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	err := c.transport.Do(req, resp)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return errors.New("not ok")
	}

	if target == nil {
		return nil
	}
	if err = json.Unmarshal(resp.Body(), target); err != nil {
		logger.Error.Println("unmarshal err: ", err)
		return err
	}

	return nil
}

func (c client) Dig(params *models.Dig, target *[]string) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(c.baseUrl + "/dig")
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	body, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req.SetBody(body)

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err = c.transport.Do(req, resp); err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return errors.New("not ok")
	}

	if err = json.Unmarshal(resp.Body(), target); err != nil {
		logger.Error.Println("unmarshal err: ", err)
		return err
	}

	return nil
}

func (c client) ExploreArea(params *models.Area, target *models.Report) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(c.baseUrl + "/explore")
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	body, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req.SetBody(body)

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err = c.transport.Do(req, resp); err != nil {
		return err
	}

	if err = json.Unmarshal(resp.Body(), target); err != nil {
		logger.Error.Println("unmarshal err: ", err)
		return err
	}

	return nil
}

func (c client) GetBalance(target *models.Balance) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(c.baseUrl + "/balance")
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err := c.transport.Do(req, resp); err != nil {
		return err
	}

	if err := json.Unmarshal(resp.Body(), target); err != nil {
		logger.Error.Println("unmarshal err: ", err)
		return err
	}

	return nil
}

func (c client) HealthCheck() error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(c.baseUrl + "/health-check")
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err := c.transport.Do(req, resp); err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return errors.New("not ready")
	}

	return nil
}

func (c client) IssueLicense(coins []uint64, target *models.License) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(c.baseUrl + "/licenses")
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	body, err := json.Marshal(coins)
	if err != nil {
		return err
	}
	req.SetBody(body)

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err = c.transport.Do(req, resp); err != nil {
		return err
	}

	if err = json.Unmarshal(resp.Body(), target); err != nil {
		logger.Error.Println("unmarshal err: ", err)
		return err
	}

	return nil
}

func (c client) ListLicenses(target *[]models.License) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(c.baseUrl + "/licenses")
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err := c.transport.Do(req, resp); err != nil {
		return err
	}

	if err := json.Unmarshal(resp.Body(), target); err != nil {
		logger.Error.Println("unmarshal err: ", err)
		return err
	}

	return nil
}
