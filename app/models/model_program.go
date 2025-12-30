package models

type Program struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Owner    string `json:"owner"`
	Donasi   int64  `json:"donasi"`
	Priority bool   `json:"priority"`
	Category string `json:"category"`
	Image    string `json:"image"`
}
