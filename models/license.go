package models

// License License for digging.
type License struct {

	// dig allowed
	// Required: true
	DigAllowed Amount `json:"digAllowed"`

	// dig used
	// Required: true
	DigUsed Amount `json:"digUsed"`

	// id
	ID uint64 `json:"id"`
}
