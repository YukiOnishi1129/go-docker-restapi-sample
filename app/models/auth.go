package models

/*
 ログインパラメータ
*/
type SingInRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

/*
 会員登録パラメータ
*/
type SignUpRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	BaseModel
	Name   string `gorm:"size:255" json:"name,omitempty"`
	Email  string `gorm:"size:255;not null;unique" json:"email,omitempty"`
}

type SignUpResponse struct {
	Token string `json:"token"`
	User UserResponse
}