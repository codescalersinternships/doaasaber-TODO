package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	newTodo := todos{
		ID:   12,
		Task: "task",
	}
	jsonValue, _ := json.Marshal(newTodo)

	t.Run("create new todo", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(jsonValue))
		response := httptest.NewRecorder()

		CreateTodo(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
