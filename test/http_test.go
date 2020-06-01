package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	sw "go-clean-architecture/api/openapi"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := sw.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/entities", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "null", w.Body.String())
}
