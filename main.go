package todo

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DBName = "todos.db"

type todos struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}
type Server struct {
	DB *gorm.DB
}

func (t *Server) CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var new todos
	json.NewDecoder(r.Body).Decode(&new)
	t.DB.Create(&new)
	json.NewEncoder(w).Encode(new)
	w.WriteHeader(http.StatusCreated)

}
func (t *Server) Gettodo(w http.ResponseWriter, r *http.Request) {
	_, err := t.DB.Model(&todos{}).Where("ID > ?", "0").Rows()

	if err != nil {
		panic("error parsing data")
	}
	w.Header().Set("Content-Type", "application/json")
	var new []todos
	t.DB.Find(&new)
	json.NewEncoder(w).Encode(new)
	w.WriteHeader(http.StatusOK)

}

func (t *Server) Gettodobyid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var res todos
	out := t.DB.First(&res, id)
	if out.Error != nil {
		http.Error(w, out.Error.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - can't find this task"))
		return
	}
	data, err := json.Marshal(res)
	if err != nil {
		http.Error(w, out.Error.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Error can't find this task"))
		return
	}
	w.Write(data)
	w.Write([]byte("\n"))

}
func (t *Server) UpadateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var new todos
	res := t.DB.First(&new, param["ID"])
	json.NewDecoder(r.Body).Decode(&new)

	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - can't find this task"))
		return
	}

	t.DB.Save(&new)
	json.NewEncoder(w).Encode(&new)
	w.WriteHeader(http.StatusOK)

}

func (t *Server) DeleteTodo(w http.ResponseWriter, r *http.Request) { //DELETE
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["taskId"])

	var new = todos{ID: id}
	t.DB.Find(&new)
	res := t.DB.First(&new, "ID")
	t.DB.Delete(&new)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - can't find this task"))
		return
	} else {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("200 - Task deleted successfully"))
	}
}
func (t *Server) InitializeDB() {
	var err error
	t.DB, err = gorm.Open(sqlite.Open(DBName), &gorm.Config{})
	if !(t.DB.Migrator().HasTable(&todos{})) {
		log.Println("table { todos } created")
		t.DB.Migrator().CreateTable(&todos{})
	}
	if err != nil {
		panic("can't connect to DB")
	}
}

func main() {
	t := Server{}
	t.InitializeDB()
	// var err error
	// t.DB, err = gorm.Open(sqlite.Open(DBName), &gorm.Config{})
	// if !(t.DB.Migrator().HasTable(&todos{})) {
	// 	log.Println("table { todos } created")
	// 	t.DB.Migrator().CreateTable(&todos{})
	// }
	// if err != nil {
	// 	panic("can't connect to DB")
	// }
	router := mux.NewRouter()
	router.HandleFunc("/todo", t.Gettodo).Methods("GET")
	router.HandleFunc("/todo/{id}", t.Gettodobyid).Methods("GET")
	router.HandleFunc("/todo", t.CreateTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", t.UpadateTodo).Methods("PUT")
	router.HandleFunc("/todo/{taskId}", t.DeleteTodo).Methods("DELETE")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":9000", router))
}
