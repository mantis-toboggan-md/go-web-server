package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETTodos(t *testing.T) {
	t.Run("returns counting todo", func(t *testing.T) {
		// NewRequest takes (method, path,body) makes req
		request, _ := http.NewRequest(http.MethodGet, "/todos/123", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		got := response.Body.String()
		want := "Count to three."

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns satan todo", func(t *testing.T) {
		// NewRequest takes (method, path,body) makes req
		request, _ := http.NewRequest(http.MethodGet, "/todos/666", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		got := response.Body.String()
		want := "Hail satan!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
