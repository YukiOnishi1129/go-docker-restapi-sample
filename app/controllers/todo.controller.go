package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"myapp/models"
	"myapp/services"
	"myapp/utils/logic"
	"myapp/utils/validation"
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
    // tokenからuserIdを所得
    userId, err := services.GetUserIdFromToken(w,r)
    if userId == 0 || err != nil {
        return
    }

    // todoリスト取得処理
    alltodo, err := services.GetAllTodos(w, userId)
    if err !=nil {
        return
    }

    // レスポンス送信処理
    services.SendAllTodoResponse(w, &alltodo)
}


/*
 idに紐づくTodoを取得
*/
func fetchTodoById(w http.ResponseWriter, r *http.Request) {
	// tokenからuserIdを所得
    userId, err := services.GetUserIdFromToken(w,r)
    if userId == 0 || err != nil {
        return
    }
    // todoデータ取得処理
    responseTodo, err := services.GetTodoById(w , r, userId)
    if err !=nil {
        return
    }
    services.SendTodoResponse(w, &responseTodo)
}

/*
 Todo新規登録
*/
func createTodo(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
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
    // ioutil: ioに特化したパッケージ
    reqBody,_ := ioutil.ReadAll(r.Body)

    // バリデーション
    var mutationTodoRequest models.MutationTodoRequest
    if err := json.Unmarshal(reqBody, &mutationTodoRequest); err != nil {
        log.Fatal(err)
    }
    if err := validation.MutationTodoValidate(mutationTodoRequest); err != nil {
        response := map[string]interface{}{
            "error": err,
        }
        responseBody, _ := json.Marshal(response)
        w.WriteHeader(http.StatusBadRequest) // ステータスコード
        w.Write(responseBody)
        return
    }

    var todo models.Todo
    // json.Unmarshal()
    // 第１引数で与えたjsonデータを、第二引数に指定した値にマッピングする
    // 返り値はerrorで、エラーが発生しない場合はnilになる
    if err := json.Unmarshal(reqBody, &todo); err != nil {
        log.Fatal(err)
    }

    todo.UserId = userId

    services.InsertTodo(&todo)

    responseBody, err := json.Marshal(todo)
    if err != nil {
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(responseBody)
}

/*
 削除処理
*/
func deleteTodo(w http.ResponseWriter, r *http.Request) {
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
        return
	}
    vars := mux.Vars(r)
    id := vars["id"]

    // var todo models.Todo

    // 削除データの確認
    // services.GetTodoById(&todo, id, userId)
    // if todo.ID == 0 {
    //     // レスポンスデータ作成
	// 	response := map[string]interface{}{
	// 		"err": "データがありません。",
	// 	}
	// 	responseBody, _ := json.Marshal(response)
    //     w.WriteHeader(http.StatusBadRequest)
	// 	w.Write(responseBody)
    //     return
    // }

    services.DeleteTodo(id, userId)
    // responseBody, err := json.Marshal(DeleteResponse{Id: id})
    // responseBody, err := json.Marshal(todo)
    // if err != nil {
    //     log.Fatal(err)
    // }

    w.WriteHeader(http.StatusNoContent)
    // w.Write(responseBody)
}

/*
 Todo更新処理
*/
func updateTodo(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
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

    reqBody, _ := ioutil.ReadAll(r.Body)
    // バリデーション
    var mutationTodoRequest models.MutationTodoRequest
    if err := json.Unmarshal(reqBody, &mutationTodoRequest); err != nil {
        log.Fatal(err)
    }
    if err := validation.MutationTodoValidate(mutationTodoRequest); err != nil {
        response := map[string]interface{}{
            "error": err,
        }
        responseBody, _ := json.Marshal(response)
        w.WriteHeader(http.StatusBadRequest) // ステータスコード
        w.Write(responseBody)
        return
    }

    // 更新データの有無確認用
    var todo models.Todo
    if err := json.Unmarshal(reqBody, &todo); err != nil {
        log.Fatal(err)
    }
    // データ更新処理用
    var updateTodo models.Todo
    if err := json.Unmarshal(reqBody, &updateTodo); err != nil {
        log.Fatal(err)
    }

    updateTodo.UserId = userId

    // 更新データの有無確認
    // services.GetTodoById(&todo, id, userId)
    // if todo.ID == 0 {
    //     // レスポンスデータ作成
	// 	response := map[string]interface{}{
	// 		"err": "データがありません。",
	// 	}
	// 	responseBody, _ := json.Marshal(response)
    //     w.WriteHeader(http.StatusBadRequest)
	// 	w.Write(responseBody)
    //     return
    // }

    // データ更新
    services.UpdateTodo(&updateTodo, id)
    convertUnitId, _ := strconv.ParseUint(id, 10, 64)
    updateTodo.BaseModel.ID = uint(convertUnitId)

    response := map[string]interface{}{
        "todo": updateTodo,
    }
    responseBody, err := json.Marshal(response)
    if err != nil {
        log.Fatal(err)
    }

    w.WriteHeader(http.StatusOK) // ステータスコード
    w.Write(responseBody)
}


func SetTodoRouting(router *mux.Router) {
	router.Handle("/todo", logic.JwtMiddleware.Handler(http.HandlerFunc(fetchAllTodos))).Methods("GET")
    router.Handle("/todo/{id}", logic.JwtMiddleware.Handler(http.HandlerFunc(fetchTodoById))).Methods("GET")

    router.Handle("/todo", logic.JwtMiddleware.Handler(http.HandlerFunc(createTodo))).Methods("POST")
    router.Handle("/todo/{id}", logic.JwtMiddleware.Handler(http.HandlerFunc(deleteTodo))).Methods("DELETE")
    router.Handle("/todo/{id}", logic.JwtMiddleware.Handler(http.HandlerFunc(updateTodo))).Methods("PUT")
}