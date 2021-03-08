package models

// Area area
type Area struct {

	// pos x
	// Required: true
	// Minimum: 0
	PosX uint16 `json:"posX"`

	// pos y
	// Required: true
	// Minimum: 0
	PosY uint16 `json:"posY"`

	// size x
	// Minimum: 1
	SizeX uint16 `json:"sizeX,omitempty"`

	// size y
	// Minimum: 1
	SizeY uint16 `json:"sizeY,omitempty"`
}
