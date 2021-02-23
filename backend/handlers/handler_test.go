package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/vasuvanka/todo-app/backend/handlers"
)

func TestCheck(t *testing.T){
	router := httprouter.New()
    router.GET("/api", handlers.Check)

    req, _ := http.NewRequest("GET", "/api", nil)
    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)
	assert.Equal(t,rr.Code, http.StatusOK," Failed to get ok status")
}