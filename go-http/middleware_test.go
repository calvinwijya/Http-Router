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

type LogMiddleware struct {
	Handler http.Handler
}

func (L *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Print("receive request")
	L.Handler.ServeHTTP(w, r)
}

func TestMiddleWare(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Middleware")
	})
	middleware := &LogMiddleware{Handler: router}
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recoder := httptest.NewRecorder()

	middleware.ServeHTTP(recoder, request)
	response := recoder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Middleware", string(body))
}
