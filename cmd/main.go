package main

import (
	"fmt"
	"github.com/deissh/highloadcup-goldenrush/client"
	"github.com/deissh/highloadcup-goldenrush/core"
	"github.com/deissh/highloadcup-goldenrush/logger"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func WaitReady(apiClient *client.CupClient) {
	for {
		err := apiClient.HealthCheck()
		if err == nil {
			break
		}

		time.Sleep(time.Second)
	}
}

func main() {
	logger.Info.Println("Starting")

	go http.ListenAndServe("localhost:2233", nil)

	host := fmt.Sprintf("%s:%d", os.Getenv("ADDRESS"), 8000)
	api := client.New(&client.TransportConfig{
		Host: host,
	})

	logger.Info.Println("Wait server")
	WaitReady(api)

	game := core.New(api)

	if err := game.Start(); err != nil {
		logger.Error.Println(err)
		os.Exit(1)
	}
}
