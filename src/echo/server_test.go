package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestHomePage( t *testing.T) {
	t.Run("returns my Headers", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		request.Header.Add("x-forwarded-for", "myself")
		response := httptest.NewRecorder()

		homePage(response, request)

		got := response.Body.String()
		want := "X-Forwarded-For: myself\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
