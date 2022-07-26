package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	var server Server
	newTodo := todos{
		ID:   12,
		Task: "task",
	}
	server.InitializeDB()
	jsonValue, _ := json.Marshal(newTodo)

	t.Run("create new todo", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(jsonValue))
		fmt.Println("here1")
		response := httptest.NewRecorder()
		fmt.Println("here2")
		fmt.Println(server.DB)
		fmt.Println("here3")

		server.CreateTodo(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		fmt.Println("here4")

	})
}

// func TestGetTodo(t *testing.T) {
// 	var server Server
// 	newTodo := todos{
// 		ID:   19,
// 		Task: "task3",
// 	}
// 	jsonValue, _ := json.Marshal(newTodo)

// 	t.Run("create new todo", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(jsonValue))
// 		response := httptest.NewRecorder()

// 		server.Gettodo(response, request)

// 		assertStatus(t, response.Code, http.StatusOK)
// 	})
// }

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
