package web

type NewsUpsertRequest struct {
	Owner    string `json:"owner"`
	HeadLine string `json:"headLine"`
	Lead     string `json:"lead"`
	Body     string `json:"body"`
	Tail     string `json:"tail"`
	Image    string `json:"image"`
}
