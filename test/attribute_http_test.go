package test

import (
	// "bytes"
	"encoding/json"
	"fmt"
	sw "github.com/ttaki/go-clean-architecture-sample/internal/api/http"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	// "github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

func TestPingAttributeRoute(t *testing.T) {
	entity := SetEntityTestData()

	router := sw.NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/entities/"+strconv.Itoa(entity.ID)+"/attributes", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateAttributeInvalidRoute(t *testing.T) {
	entity := SetEntityTestData()
	router := sw.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/entities/"+strconv.Itoa(entity.ID)+"/attributes", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

// test deleteAttribute
func TestDeleteAttributeRoute(t *testing.T) {
	attribute := SetAttributeTestData()
	//update test
	router := sw.NewRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/v1/attributes/"+strconv.Itoa(attribute.ID), nil)
	req.Header.Set("authorization", "")
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}

// test FindByID
func TestFindByIDAttributeRoute(t *testing.T) {
	attribute := SetAttributeTestData()
	attributeID := attribute.ID

	router := sw.NewRouter()
	w := httptest.NewRecorder()

	fmt.Println("/v1/attributes/"+strconv.Itoa(attribute.ID), "GET")
	req, _ := http.NewRequest("GET", "/v1/attributes/"+strconv.Itoa(attribute.ID), nil)

	router.ServeHTTP(w, req)
	response := Attribute{}

	fmt.Println(w.Body.String())
	b := []byte(w.Body.String())
	json.Unmarshal(b, &response)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, attributeID, response.ID)
	assert.Equal(t, attribute.Body, response.Body)
}
