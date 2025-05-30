package utils_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MakeRequest(T *testing.T, app *TestApp, metod, url string, body interface{}, token string) (*httptest.ResponseRecorder, *http.Request) {
	var reqBody io.Reader

	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			T.Fatalf("Failed to marshal request body: %v", err)
		}

		reqBody = bytes.NewBuffer(jsonData)
	}

	fmt.Println("Making request to URL:", url)

	req, err := http.NewRequest(metod, url, reqBody)

	if err != nil {
		T.Fatalf("Failed to create request: %v", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	w := httptest.NewRecorder()

	app.Router.ServeHTTP(w, req)

	return w, req
}
