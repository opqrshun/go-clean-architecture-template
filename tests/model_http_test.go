package tests

import (
	// "bytes"
	"strconv"
	"encoding/json"
	"fmt"
	sw "gobackend/api/http"
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"

	"gobackend/tests"
)

func TestPingModelRoute(t *testing.T) {
	parent := tests.SetParentTestData()

	router := sw.NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/parents/"+strconv.Itoa(parent.ID)+"/models", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateModelInvalidRoute(t *testing.T) {
	parent := tests.SetParentTestData()
	router := sw.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/parents/"+strconv.Itoa(parent.ID)+"/models", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}


// test deleteModel
func TestDeleteModelRoute(t *testing.T) {
	model := SetModelTestData()
	//update test
	router := sw.NewRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/v1/models/"+strconv.Itoa(model.ID), nil)
	req.Header.Set("authorization", "")
	router.ServeHTTP(w, req)

	//model
	response := tests.Parent{}

	fmt.Println(w.Body.String())
	b := []byte(w.Body.String())
	// a, _ := w.Body.ReadBytes(10)
	json.Unmarshal(b, &response)

	assert.Equal(t, 204, w.Code)
}

// test FindByID
func TestFindByIDModelRoute(t *testing.T) {
	model := SetModelTestData()
	modelID := model.ID

	router := sw.NewRouter()
	w := httptest.NewRecorder()

	fmt.Println("/v1/models/"+strconv.Itoa(model.ID), "GET")
	req, _ := http.NewRequest("GET", "/v1/models/"+strconv.Itoa(model.ID), nil)

	router.ServeHTTP(w, req)
	response := tests.Model{}

	fmt.Println(w.Body.String())
	b := []byte(w.Body.String())
	json.Unmarshal(b, &response)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, modelID, response.ID)
	assert.Equal(t, model.Body, response.Body)
}

