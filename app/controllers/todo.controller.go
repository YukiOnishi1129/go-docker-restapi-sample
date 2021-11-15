package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"myapp/models"
	"myapp/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DeleteTodoResponse struct {
    Id string `json:"id"`
}


func fetchAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var todos []models.Todo
    services.GetAllTodos(&todos)
    responseBody, err := json.Marshal(todos)
    if err != nil {
        log.Fatal(err)
    }
    w.Write(responseBody)
}

func fetchTodoById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    vars := mux.Vars(r)
    id := vars["id"] 

    var todo models.Todo
    services.GetTodoById(&todo, id)

    responseBody, err := json.Marshal(todo)
    if err != nil {
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "application/json")
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


func deleteItem(w http.ResponseWriter, r *http.Request) {
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
	router.HandleFunc("/todo", fetchAllTodos).Methods("GET")
    router.HandleFunc("/todo/{id}", fetchTodoById).Methods("GET")

    router.HandleFunc("/todo", createTodo).Methods("POST")
    router.HandleFunc("/todo/{id}", deleteItem).Methods("DELETE")
    router.HandleFunc("/todo/{id}", updateTodo).Methods("PUT")
}