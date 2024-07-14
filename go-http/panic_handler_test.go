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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = (func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "panic bang: ", i)
	})
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("ups")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "panic bang: ups", string(body))
}
