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