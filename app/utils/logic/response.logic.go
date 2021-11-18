package logic

import (
	"encoding/json"
	"net/http"
)

/*
 APIレスポンス送信処理
*/
func SendResponse(w http.ResponseWriter, response []byte, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

/*
 APIレスポンス送信処理 (レスポンスBodyなし)
*/
func SendNotBodyResponse(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

/*
 エラーレスポンス作成
*/
func CreateErrorResponse(err error) []byte {
	response := map[string]interface{}{
		"error": err,
	}
	responseBody, _ := json.Marshal(response)

	return responseBody
}

/*
 エラーレスポンス作成 (エラーメッセージはstring)
*/
func CreateErrorStringResponse(errMessage string) []byte {
	response := map[string]interface{}{
		"error": errMessage,
	}
	responseBody, _ := json.Marshal(response)

	return responseBody
}

