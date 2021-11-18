package controllers

import (
	"log"
	"myapp/services"
	"myapp/utils/logic"
	"net/http"

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

    // レスポンス送信処理
    services.SendTodoResponse(w, &responseTodo)
}

/*
 Todo新規登録
*/
func createTodo(w http.ResponseWriter, r *http.Request) {
    // トークンからuserIdを取得
	userId, err := services.GetUserIdFromToken(w,r)
    if userId == 0 || err != nil {
        return
    }

    // todoデータ取得処理
    responseTodo, err := services.CreateTodo(w , r, userId)
    if err !=nil {
        return
    }

    // レスポンス送信処理
    services.SendCreateTodoResponse(w, &responseTodo)
}

/*
 削除処理
*/
func deleteTodo(w http.ResponseWriter, r *http.Request) {
    // トークンからuserIdを取得
    userId, err := services.GetUserIdFromToken(w,r)
    if userId == 0 || err != nil {
        return
    }

    // データ削除処理
    if err := services.DeleteTodo(w, r, userId); err != nil {
        return
    }

    // レスポンス送信処理
    services.SendDeleteTodoResponse(w)
    log.Fatal(err)
}

/*
 Todo更新処理
*/
func updateTodo(w http.ResponseWriter, r *http.Request) {
	// tokenからuserIdを所得
    userId, err := services.GetUserIdFromToken(w,r)
    if userId == 0 || err != nil {
        return
    }

    // todo更新処理
    responseTodo, err := services.UpdateTodo(w , r, userId)
    if err !=nil {
        return
    }

    // レスポンス送信処理
    services.SendCreateTodoResponse(w, &responseTodo)
}


func SetTodoRouting(router *mux.Router) {
	router.Handle("/todo", logic.JwtMiddleware.Handler(http.HandlerFunc(fetchAllTodos))).Methods("GET")
    router.Handle("/todo/{id}", logic.JwtMiddleware.Handler(http.HandlerFunc(fetchTodoById))).Methods("GET")

    router.Handle("/todo", logic.JwtMiddleware.Handler(http.HandlerFunc(createTodo))).Methods("POST")
    router.Handle("/todo/{id}", logic.JwtMiddleware.Handler(http.HandlerFunc(deleteTodo))).Methods("DELETE")
    router.Handle("/todo/{id}", logic.JwtMiddleware.Handler(http.HandlerFunc(updateTodo))).Methods("PUT")
}