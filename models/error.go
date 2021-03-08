package models

type Error struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}
