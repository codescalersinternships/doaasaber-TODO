package main

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
	//"github.com/mattn/go-sqlite3"

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
	out := DB.First(&res, id)
	//	fmt.Println(err)
	if out.Error != nil {
		http.Error(w, out.Error.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - can't find this task"))
		return
		//panic("error parsing data")

	}
	data, err := json.Marshal(res)
	if err != nil {
		http.Error(w, out.Error.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - can't find this task"))
		return
	}
	w.Write(data)
	w.Write([]byte("\n"))

}

func UpadateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var new todos
	res := DB.First(&new, param["ID"])
	json.NewDecoder(r.Body).Decode(&new)

	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - can't find this task"))
		return
	}

	DB.Save(&new)
	json.NewEncoder(w).Encode(&new)
	w.WriteHeader(http.StatusOK)

}
func DeleteTodo(w http.ResponseWriter, r *http.Request) { //DELETE
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["taskId"])

	var new = todos{ID: id}
	DB.Find(&new)
	res := DB.First(&new, "ID")
	DB.Delete(&new)
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

/*func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["taskId"])
	var new = todos{ID: id}
	DB.Find(&new)
	DB.Delete(&new)

	//var id_deleted = 0
	if DB.RowsAffected > 0 {
		DB.Delete(DB.Find(&new))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - this task is already deleted successfuly"))
		return
	} else if DB.RowsAffected == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - can't find this task"))
		return
	}
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("200 - Task deleted successfuly"))
}*/

/*func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["taskId"])
	var new = todos{ID: id}
	DB.Delete(&new, param["ID"])
	if out.Error != nil {

		http.Error(w, out.Error.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - can't find this task"))
		return
	}
	var id_deleted = 0
	if DB.RowsAffected > 0 {
		DB.Delete(DB.Find(&id_deleted))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Task deleted successfuly"))
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - Task deleted"))
}*/

/*func DeleteTodo(w http.ResponseWriter, r *http.Request) { //DELETE
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["taskId"])

	var new = todos{ID: id}
	DB.Find(&new)
	DB.Delete(&new)
	var id_deleted = 0
	if DB.RowsAffected > 0 {
		DB.Delete(DB.Find(&id_deleted))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Task deleted successfuly"))
	} else {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("204 - Task already deleted"))
	}
}*/

func main() {

	if err != nil {
		panic("can't connect to the database")
	}
	DB.AutoMigrate(&todos{})
	router := mux.NewRouter()

	router.HandleFunc("/todo", Gettodo).Methods("GET")
	router.HandleFunc("/todo/{id}", Gettodobyid).Methods("GET")
	router.HandleFunc("/todo", CreateTodo).Methods("POST")
	router.HandleFunc("/todo/{ID}", UpadateTodo).Methods("PUT")
	router.HandleFunc("/todo/{taskId}", DeleteTodo).Methods("DELETE")
	//router.HandleFunc("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)).Methods("GET")
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":9000", router))

}
