package models

// Dig dig
type Dig struct {

	// depth
	// Required: true
	// Maximum: 100
	// Minimum: 1
	Depth uint8 `json:"depth"`

	// ID of the license this request is attached to.
	// Required: true
	LicenseID uint64 `json:"licenseID"`

	// pos x
	// Required: true
	// Minimum: 0
	PosX uint16 `json:"posX"`

	// pos y
	// Required: true
	// Minimum: 0
	PosY uint16 `json:"posY"`
}
