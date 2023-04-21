package models

type Author struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Rating int8   `json:"rating"`
	Info   string `json:"info"`
}
