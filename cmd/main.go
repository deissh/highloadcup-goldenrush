package main

import (
	"fmt"
	"github.com/deissh/highloadcup-goldenrush/client"
	"github.com/deissh/highloadcup-goldenrush/core"
	"log"
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
	log.Println("Starting")

	go http.ListenAndServe("localhost:2233", nil)

	host := fmt.Sprintf("http://%s:%d", os.Getenv("ADDRESS"), 8000)
	api := client.New(&client.TransportConfig{
		BaseUrl: host,
	})

	log.Println("Wait server")
	WaitReady(api)

	game := core.New(api)

	if err := game.Start(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
