package models

// Report report
type Report struct {

	// amount
	// Required: true
	Amount *Amount `json:"amount"`

	// area
	// Required: true
	Area *Area `json:"area"`
}
