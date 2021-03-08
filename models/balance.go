package models

// Balance Current balance and wallet with up to 1000 coins.
type Balance struct {

	// balance
	Balance uint32   `json:"balance"`
	Wallet  []uint32 `json:"wallet"`
}
