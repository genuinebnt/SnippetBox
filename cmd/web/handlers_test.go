package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateSnippets(t *testing.T) {
	t.Run("Test create snippets endpoint with get", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/snippets/create", nil)
		response := httptest.NewRecorder()

		CreateSnippets(response, request)

		got := response.Code
		want := 405

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Testing id's", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/snippets/create", nil)
		response := httptest.NewRecorder()

		CreateSnippets(response, request)

		got := response.Body.String()
		want := "Create a new snippet"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
