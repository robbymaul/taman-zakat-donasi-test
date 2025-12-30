package models

type News struct {
	Id        string `json:"id"`
	Owner     string `json:"owner"`
	HeadLine  string `json:"headLine"`
	Lead      string `json:"lead"`
	Body      string `json:"body"`
	Tail      string `json:"tail"`
	CreatedAt string `json:"createdAt"`
	Image     string `json:"image"`
}
