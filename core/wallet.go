package core

import (
	"github.com/deissh/highloadcup-goldenrush/client"
	"sync"
)

// Wallet cash treasures and store coins
// todo: add coin holder for paid licenses
type Wallet struct {
	client      *client.CupClient
	workerCount int

	treasuresChan chan string
}

func NewWallet(client *client.CupClient, workerCount int) *Wallet {
	w := Wallet{client: client, workerCount: workerCount}

	w.treasuresChan = make(chan string, 1000) // change me

	return &w
}

func (w *Wallet) CashTreasures(ids []string) {
	for _, id := range ids {
		w.treasuresChan <- id
	}
}

func (w *Wallet) Init() {}

func (w *Wallet) Start() {
	wg := &sync.WaitGroup{}

	wg.Add(w.workerCount)
	for i := 1; i <= w.workerCount; i++ {
		go w.cashier(wg)
	}

	wg.Wait()
}

func (w *Wallet) cashier(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		treasure := <-w.treasuresChan

		err := w.client.Cash(treasure, nil)
		if err != nil {
			// try again
			w.treasuresChan <- treasure
		}
	}
}
