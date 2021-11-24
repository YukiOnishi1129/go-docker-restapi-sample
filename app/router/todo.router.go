package router

import (
	"myapp/controllers"
	"myapp/utils/logic"
	"net/http"

	"github.com/gorilla/mux"
)

type TodoRouter interface {
	SetTodoRouting(router *mux.Router)
}

type todoRouter struct {
	tc controllers.TodoController
}

func NewTodoRouter(tc controllers.TodoController) TodoRouter {
	return &todoRouter{tc}
}


func (tr *todoRouter) SetTodoRouting(router *mux.Router) {
	router.Handle("/api/v1/todo", logic.JwtMiddleware.Handler(http.HandlerFunc(tr.tc.FetchAllTodos))).Methods("GET")
    router.Handle("/api/v1/todo/{id}", logic.JwtMiddleware.Handler(http.HandlerFunc(tr.tc.FetchTodoById))).Methods("GET")

    router.Handle("/api/v1/todo", logic.JwtMiddleware.Handler(http.HandlerFunc(tr.tc.CreateTodo))).Methods("POST")
    router.Handle("/api/v1/todo/{id}", logic.JwtMiddleware.Handler(http.HandlerFunc(tr.tc.DeleteTodo))).Methods("DELETE")
    router.Handle("/api/v1/todo/{id}", logic.JwtMiddleware.Handler(http.HandlerFunc(tr.tc.UpdateTodo))).Methods("PUT")
}