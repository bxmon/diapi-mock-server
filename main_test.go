package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxmon/diapi-mock-server/controller"
	"github.com/bxmon/diapi-mock-server/service"
	"github.com/bxmon/diapi-mock-server/storage"
	"github.com/stretchr/testify/assert"
)

func TestController(t *testing.T) {
	db := storage.NewStorage("account.db", "accountbucket")
	service := service.NewService(db)
	controller := controller.NewController(service)
	engine := SetUpEngine(controller)

	t.Run("AddNewUser", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/users/register", nil)
		engine.ServeHTTP(rec, req)

		assert.Equal(t, 200, rec.Code)
	})
}
