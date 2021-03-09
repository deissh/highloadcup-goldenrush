package models

// Report report
type Report struct {

	// amount
	// Required: true
	Amount uint64 `json:"amount"`

	// area
	// Required: true
	Area *Area `json:"area"`
}
