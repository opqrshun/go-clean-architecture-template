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

	// "github.com/ttaki/go-clean-architecture-template/pkg/identifier"
	sw "github.com/ttaki/go-clean-architecture-template/internal/api/http"
)

func TestPingEntityRoute(t *testing.T) {
	router := sw.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/entities", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateEntityInvalidRoute(t *testing.T) {
	router := sw.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/entities", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func BuildCreateEntity() EntityCreate {
	entityCreate := EntityCreate{}
	gofakeit.Struct(&entityCreate)
	return entityCreate
}

// test create
func TestCreateEntityRoute(t *testing.T) {

	router := sw.NewRouter()
	w := httptest.NewRecorder()
	entityCreate := BuildCreateEntity()

	fmt.Printf("%+v\n", entityCreate)

	request, _ := json.Marshal(entityCreate)
	fmt.Println(string(request), "json")

	req, _ := http.NewRequest("POST", "/v1/entities", bytes.NewBuffer(request))
	req.Header.Set("Entity-Type", "application/json")
	req.Header.Set("authorization", "")
	router.ServeHTTP(w, req)

	//  for bind json
	res := Entity{}

	fmt.Println(w.Body)
	b := []byte(w.Body.String())
	// a, _ := w.Body.ReadBytes(10)
	json.Unmarshal(b, &res)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, entityCreate.Body, res.Body)
}

//test update
func TestUpdateEntityRoute(t *testing.T) {
	entity := SetEntityTestData()

	//update test
	router := sw.NewRouter()
	w := httptest.NewRecorder()

	entityUpdate := EntityUpdate{}
	gofakeit.Struct(&entityUpdate)
	entityUpdate.ID = entity.ID

	fmt.Printf("%+v\n", entityUpdate)
	request, _ := json.Marshal(entityUpdate)
	fmt.Println("/v1/entities/"+strconv.Itoa(entityUpdate.ID), "PATCH")

	req, _ := http.NewRequest("PATCH", "/v1/entities/"+strconv.Itoa(entity.ID), bytes.NewBuffer(request))
	req.Header.Set("Entity-Type", "application/json")
	req.Header.Set("authorization", "")
	router.ServeHTTP(w, req)

	//res
	res := Entity{}

	fmt.Println(w.Body)
	b := []byte(w.Body.String())
	// a, _ := w.Body.ReadBytes(10)
	json.Unmarshal(b, &res)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, entity.ID, res.ID)
	assert.Equal(t, entityUpdate.Body, res.Body)

}

// test FindByID
func TestFindByIDEntityRoute(t *testing.T) {
	entity := SetEntityTestData()

	//test
	router := sw.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/entities/"+strconv.Itoa(entity.ID), nil)

	router.ServeHTTP(w, req)
	res := Entity{}

	fmt.Println(w.Body)
	b := []byte(w.Body.String())
	json.Unmarshal(b, &res)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, entity.ID, res.ID)
	assert.Equal(t, entity.Body, res.Body)
	assert.Equal(t, entity.Body, res.Body)
}
