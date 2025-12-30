package web

type ProgramUpsertRequest struct {
	Title    string `json:"title"`
	Owner    string `json:"owner"`
	Priority bool   `json:"priority"`
	Category string `json:"category"`
	Image    string `json:"image"`
}
