package models

/*
 ログインパラメータ
*/
type SingInRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}