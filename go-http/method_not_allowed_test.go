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

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "MethodNotAllowed")
	})
	router.POST("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(rw, "POST")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "MethodNotAllowed", string(body))

}
