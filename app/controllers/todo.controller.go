package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"myapp/models"
	"myapp/services"
	"myapp/utils/logic"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DeleteTodoResponse struct {
    Id string `json:"id"`
}

/*
 Todoリスト取得
*/
func fetchAllTodos(w http.ResponseWriter, r *http.Request) {
	// トークンからuserIdを取得
	userId, err := logic.GetUserIdFromContext(r)
	if err != nil {
		// レスポンスデータ作成
		response := map[string]interface{}{
			"err": "認証エラー",
		}
		responseBody, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(responseBody)
	}

	var todos []models.Todo
    services.GetAllTodos(&todos, userId)

	// レスポンスデータ作成
	response := map[string]interface{}{
		"todos": todos,
	}
    responseBody, err := json.Marshal(response)
    if err != nil {
        log.Fatal(err)
    }
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK) // ステータスコード
    w.Write(responseBody)
}


/*
 idに紐づくTodoを取得
*/
func fetchTodoById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    // トークンからuserIdを取得
	userId, err := logic.GetUserIdFromContext(r)
	if err != nil {
		// レスポンスデータ作成
		response := map[string]interface{}{
			"err": "認証エラー",
		}
		responseBody, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(responseBody)
	}

    vars := mux.Vars(r)
    id := vars["id"] 

    var todo models.Todo
    services.GetTodoById(&todo, id, userId)

    if todo.ID == 0 {
        // レスポンスデータ作成
		response := map[string]interface{}{
			"err": "データがありません。",
		}
		responseBody, _ := json.Marshal(response)
        w.WriteHeader(http.StatusBadRequest)
		w.Write(responseBody)
        return
    }

    responseBody, err := json.Marshal(todo)
    if err != nil {
        log.Fatal(err)
    }

    w.WriteHeader(http.StatusOK) // ステータスコード
    w.Write(responseBody)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
    // ioutil: ioに特化したパッケージ
    reqBody,_ := ioutil.ReadAll(r.Body)
    var todo models.Todo
    // json.Unmarshal()
    // 第１引数で与えたjsonデータを、第二引数に指定した値にマッピングする
    // 返り値はerrorで、エラーが発生しない場合はnilになる
    if err := json.Unmarshal(reqBody, &todo); err != nil {
        log.Fatal(err)
    }

    services.InsertTodo(&todo)

    responseBody, err := json.Marshal(todo)
    if err != nil {
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(responseBody)
}


func deleteTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    services.DeleteTodo(id)
    responseBody, err := json.Marshal(DeleteResponse{Id: id})
    if err != nil {
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(responseBody)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    reqBody, _ := ioutil.ReadAll(r.Body)
    var updateTodo models.Todo
    if err := json.Unmarshal(reqBody, &updateTodo); err != nil {
        log.Fatal(err)
    }

    services.UpdateTodo(&updateTodo, id)
    convertUnitId, _ := strconv.ParseUint(id, 10, 64)
    updateTodo.BaseModel.ID = uint(convertUnitId)

    responseBody, err := json.Marshal(updateTodo)
    if err != nil {
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(responseBody)
}


func SetTodoRouting(router *mux.Router) {
	router.Handle("/todo", logic.JwtMiddleware.Handler(http.HandlerFunc(fetchAllTodos))).Methods("GET")
    router.Handle("/todo/{id}", logic.JwtMiddleware.Handler(http.HandlerFunc(fetchTodoById))).Methods("GET")

    router.HandleFunc("/todo", createTodo).Methods("POST")
    router.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
    router.HandleFunc("/todo/{id}", updateTodo).Methods("PUT")
}