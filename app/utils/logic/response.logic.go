package logic

import (
	"encoding/json"
	"net/http"
)

type ResponseLogic interface {
	SendResponse(w http.ResponseWriter, response []byte, code int)
	SendNotBodyResponse(w http.ResponseWriter)
	CreateErrorResponse(err error) []byte
	CreateErrorStringResponse(errMessage string) []byte
}

type responseLogic struct {}

func NewResponseLogic() ResponseLogic {
	return &responseLogic{}
}

/*
 APIレスポンス送信処理
*/
func (rl *responseLogic) SendResponse(w http.ResponseWriter, response []byte, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

/*
 APIレスポンス送信処理 (レスポンスBodyなし)
*/
func (rl *responseLogic) SendNotBodyResponse(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

/*
 エラーレスポンス作成
*/
func (rl *responseLogic) CreateErrorResponse(err error) []byte {
	response := map[string]interface{}{
		"error": err,
	}
	responseBody, _ := json.Marshal(response)

	return responseBody
}

/*
 エラーレスポンス作成 (エラーメッセージはstring)
*/
func (rl *responseLogic) CreateErrorStringResponse(errMessage string) []byte {
	response := map[string]interface{}{
		"error": errMessage,
	}
	responseBody, _ := json.Marshal(response)

	return responseBody
}

