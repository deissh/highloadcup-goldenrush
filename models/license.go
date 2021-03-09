package models

// License License for digging.
type License struct {

	// dig allowed
	// Required: true
	DigAllowed uint16 `json:"digAllowed"`

	// dig used
	// Required: true
	DigUsed uint16 `json:"digUsed"`

	// id
	ID uint64 `json:"id"`
}
