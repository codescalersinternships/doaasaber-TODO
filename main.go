package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/mattn/go-sqlite3"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DBName = "todos.db"

var DB, err = gorm.Open(sqlite.Open(DBName), &gorm.Config{})

type todos struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var new todos
	json.NewDecoder(r.Body).Decode(&new)
	DB.Create(&new)
	json.NewEncoder(w).Encode(new)
	w.WriteHeader(http.StatusOK)

}
func Gettodo(w http.ResponseWriter, r *http.Request) {
	_, err := DB.Model(&todos{}).Where("ID > ?", "0").Rows()

	if err != nil {
		panic("error parsing data")
	}
	w.Header().Set("Content-Type", "application/json")
	var new []todos
	DB.Find(&new)
	json.NewEncoder(w).Encode(new)
	w.WriteHeader(http.StatusOK)

}

func Gettodobyid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var res todos
	DB.First(&res, id)
	data, err := json.Marshal(res)
	if err != nil {
		panic("error parsing data")
	}
	w.Write(data)
	w.Write([]byte("\n"))

}

func UpadateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var new todos
	DB.First(&new, param["ID"])
	json.NewDecoder(r.Body).Decode(&new)
	DB.Save(&new)
	json.NewEncoder(w).Encode(&new)
	w.WriteHeader(http.StatusOK)

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["taskId"])
	var new = todos{ID: id}
	DB.Delete(&new)
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("200 - Task deleted"))
}

func main() {

	if err != nil {
		panic("couldn't connect")
	}
	DB.AutoMigrate(&todos{})
	//DB.AutoMigrate(&todos{})
	router := mux.NewRouter()

	router.HandleFunc("/gettodo", Gettodo).Methods("GET")
	router.HandleFunc("/gettobyid/{id}", Gettodobyid).Methods("GET")
	router.HandleFunc("/createtodo", CreateTodo).Methods("POST")
	router.HandleFunc("/updatetodo/{ID}", UpadateTodo).Methods("PUT")
	router.HandleFunc("/deletetodo/{taskId}", DeleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", router))

}
