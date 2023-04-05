package models

type Comment struct {
	Id     string `json:"id"`
	Author int8
	Title  string `json:"title"`
	Book   int8
}
