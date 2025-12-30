package web

type RegistrationRequest struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	NoHp     string `json:"noHp"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SessionResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
