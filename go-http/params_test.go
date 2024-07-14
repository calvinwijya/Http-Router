package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		text := "product " + id
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/products/1", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "product 1", string(body))
}
