package entity

type Auth struct {
	UserID       string `json:"userId"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	RefreshToken string `json:"refreshToken"`
	LastLogin    string `json:"lastLogin"`
}
