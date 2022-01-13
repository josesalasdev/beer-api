package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		mockPing := &pingMessage{"pong"}

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/v1/ping", Ping)

		request, err := http.NewRequest(http.MethodGet, "/v1/ping", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(mockPing)
		assert.NoError(t, err)

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
	})
}
