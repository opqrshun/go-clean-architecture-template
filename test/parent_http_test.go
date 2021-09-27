package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/suite"

	// "gobackend/pkg/identifier"
	sw "gobackend/internal/api/http"
	"gobackend/test"
)

func TestPingParentRoute(t *testing.T) {
	router := sw.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/parents", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateParentInvalidRoute(t *testing.T) {
	router := sw.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/parents", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func BuildCreateParent() test.ParentCreate {
	parentCreate := test.ParentCreate{}
	gofakeit.Struct(&parentCreate)
	return parentCreate
}

// test create
func TestCreateParentRoute(t *testing.T) {

	router := sw.NewRouter()
	w := httptest.NewRecorder()
	parentCreate := BuildCreateParent()

	fmt.Printf("%+v\n", parentCreate)

	request, _ := json.Marshal(parentCreate)
	fmt.Println(string(request), "json")

	req, _ := http.NewRequest("POST", "/v1/parents", bytes.NewBuffer(request))
	req.Header.Set("Parent-Type", "application/json")
	req.Header.Set("authorization", "")
	router.ServeHTTP(w, req)

	//  for bind json
	res := test.Parent{}

	fmt.Println(w.Body)
	b := []byte(w.Body.String())
	// a, _ := w.Body.ReadBytes(10)
	json.Unmarshal(b, &res)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, parentCreate.Body, res.Body)
}

//test update
func TestUpdateParentRoute(t *testing.T) {
	parent := SetParentTestData()

	//update test
	router := sw.NewRouter()
	w := httptest.NewRecorder()

	parentUpdate := test.ParentUpdate{}
	gofakeit.Struct(&parentUpdate)
	parentUpdate.ID = parent.ID

	fmt.Printf("%+v\n", parentUpdate)
	request, _ := json.Marshal(parentUpdate)
	fmt.Println("/v1/parents/"+strconv.Itoa(parentUpdate.ID), "PATCH")

	req, _ := http.NewRequest("PATCH", "/v1/parents/"+strconv.Itoa(parent.ID), bytes.NewBuffer(request))
	req.Header.Set("Parent-Type", "application/json")
	req.Header.Set("authorization", "")
	router.ServeHTTP(w, req)

	//res
	res := test.Parent{}

	fmt.Println(w.Body)
	b := []byte(w.Body.String())
	// a, _ := w.Body.ReadBytes(10)
	json.Unmarshal(b, &res)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, parent.ID, res.ID)
	assert.Equal(t, parentUpdate.Body, res.Body)

}

// test FindByID
func TestFindByIDParentRoute(t *testing.T) {
	parent := SetParentTestData()

	//test
	router := sw.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/parents/"+strconv.Itoa(parent.ID), nil)

	router.ServeHTTP(w, req)
	res := test.Parent{}

	fmt.Println(w.Body)
	b := []byte(w.Body.String())
	json.Unmarshal(b, &res)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, parent.ID, res.ID)
	assert.Equal(t, parent.Body, res.Body)
	assert.Equal(t, parent.Body, res.Body)
}

