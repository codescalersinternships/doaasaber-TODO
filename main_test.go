package todo

import (
	"bytes"
	"encoding/json"
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
		response := httptest.NewRecorder()

		server.CreateTodo(response, request)

		assertStatus(t, response.Code, http.StatusOK)

	})
}

func TestGetTodo(t *testing.T) {
	var server Server
	server.InitializeDB()

	newTodo := todos{
		ID:   9,
		Task: "task3",
	}
	jsonValue, _ := json.Marshal(newTodo)

	t.Run("get all todo", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todo", bytes.NewBuffer(jsonValue))
		response := httptest.NewRecorder()

		server.Gettodo(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}
func TestGETTodoById(t *testing.T) {
	var server Server
	server.InitializeDB()

	newTodo := todos{
		ID:   19,
		Task: "task3",
	}
	jsonValue, _ := json.Marshal(newTodo)

	t.Run("get single todo", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todo/19", bytes.NewBuffer(jsonValue))
		response := httptest.NewRecorder()

		server.Gettodobyid(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})

}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
