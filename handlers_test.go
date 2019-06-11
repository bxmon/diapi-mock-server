package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	engine := SetUpEngine()

	t.Run("AddNewUser", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/users/register", nil)
		engine.ServeHTTP(rec, req)

		assert.Equal(t, 200, rec.Code)
	})
}
