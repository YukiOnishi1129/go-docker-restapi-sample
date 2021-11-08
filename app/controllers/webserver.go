package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ItemParams struct {
    Id           string    `json:"id"`
    JanCode      string    `json:"jan_code,omitempty"`
    ItemName     string    `json:"item_name,omitempty"`
    Price        int       `json:"price,omitempty"`
    CategoryId   int       `json:"category_id,omitempty"`
    SeriesId     int       `json:"series_id,omitempty"`
    Stock        int       `json:"stock,omitempty"`
    Discontinued bool      `json:"discontinued"`
    ReleaseDate  time.Time `json:"release_date,omitempty"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
    DeletedAt    time.Time `json:"deleted_at"`
}

var items []*ItemParams

/*
* ルートAPI
*/
func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Api Server")
    fmt.Println("Root endpoint is hooked!")
}

/*
* 全てのItemを取得
*/
func fetchAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
    //json.NewEncoder()
    // Go言語のデータ型からjsonに変換する
	json.NewEncoder(w).Encode(items)
}

/*
* IDに紐づくItemを取得
*/
func fetchSingleItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    // mux.Vars(): パスパラメータの取得
    vars := mux.Vars(r)
    key := vars["id"] // getパラメーターを取得

    // range: 配列を展開
    for _, item := range items {
        if item.Id == key {
            json.NewEncoder(w).Encode(item)
        }
    }
}
/*
* 新規作成
*/
func createItem(w http.ResponseWriter, r *http.Request) {
    // ioutil: ioに特化したパッケージ
    reqBody,_ := ioutil.ReadAll(r.Body)
    var item ItemParams
    // json.Unmarshal()
    // 第１引数で与えたjsonデータを、第二引数に指定した値にマッピングする
    // 返り値はerrorで、エラーが発生しない場合はnilになる
    if err := json.Unmarshal(reqBody, &item); err != nil {
        log.Fatal(err)
    }

    items = append(items, &item)
    json.NewEncoder(w).Encode(item)
}

/*
* 削除
*/
func deleteItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, item := range items {
        if item.Id == id {
            items = append(items[:index], items[index+1:]...)
        }
    }
}


/*
* 更新
*/
func updateItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    reqBody, _ := ioutil.ReadAll(r.Body)
    var updateItem ItemParams
    if err := json.Unmarshal(reqBody, &updateItem); err != nil {
        log.Fatal(err)
    }

    for index, item := range items {
        if item.Id == id {
            items[index] = &ItemParams{
                Id:           item.Id,
                JanCode:      updateItem.JanCode,
                ItemName:     updateItem.ItemName,
                Price:        updateItem.Price,
                CategoryId:   updateItem.CategoryId,
                SeriesId:     updateItem.SeriesId,
                Stock:        updateItem.Stock,
                Discontinued: updateItem.Discontinued,
                ReleaseDate:  updateItem.ReleaseDate,
                CreatedAt:    item.CreatedAt,
                UpdatedAt:    updateItem.UpdatedAt,
                DeletedAt:    item.DeletedAt,
            }
        }
    }
}

// 先頭を「大文字」にすると外部ファイルから読み込めるようになります。（export）
/*
* ルーティングを定義
*/
func StartWebServer() error {
    fmt.Println("Rest API with Mux Routers")
    router := mux.NewRouter().StrictSlash(true)

    // router.HandleFunc({ エンドポイント }, { レスポンス関数 }).Methods({ リクエストメソッド（複数可能） })
    router.HandleFunc("/", rootPage)
    router.HandleFunc("/items", fetchAllItems).Methods("GET")
    router.HandleFunc("/item/{id}", fetchSingleItem).Methods("GET")

    router.HandleFunc("/item", createItem).Methods("POST")
    router.HandleFunc("/item/{id}", deleteItem).Methods("DELETE")
    router.HandleFunc("/item/{id}", updateItem).Methods("PUT")

    return http.ListenAndServe(fmt.Sprintf(":%d", 3000), router)
}

// モックデータを初期値として読み込みます
func init() {
    items = []*ItemParams{
        &ItemParams{
            Id:           "1",
            JanCode:      "327390283080",
            ItemName:     "item_1",
            Price:        2500,
            CategoryId:   1,
            SeriesId:     1,
            Stock:        100,
            Discontinued: false,
            ReleaseDate:  time.Now(),
            CreatedAt:    time.Now(),
            UpdatedAt:    time.Now(),
            DeletedAt:    time.Now(),
        },
        &ItemParams{
            Id:           "2",
            JanCode:      "3273902878656",
            ItemName:     "item_2",
            Price:        1200,
            CategoryId:   2,
            SeriesId:     2,
            Stock:        200,
            Discontinued: false,
            ReleaseDate:  time.Now(),
            CreatedAt:    time.Now(),
            UpdatedAt:    time.Now(),
            DeletedAt:    time.Now(),
        },
    }
}