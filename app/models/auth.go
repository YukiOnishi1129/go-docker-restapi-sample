package models

// SignInRequest ログインパラメータ
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUpRequest 会員登録パラメータ
type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	BaseModel
	Name  string `gorm:"size:255" json:"name,omitempty"`
	Email string `gorm:"size:255;not null;unique" json:"email,omitempty"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
