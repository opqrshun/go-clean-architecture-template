package database

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"

	"gobackend/model"
)

func TestStore(t *testing.T) {
	repo := NewEntity()
	_entity := model.Entity{
		Body: gofakeit.Sentence(1),
	}

	id, _ := repo.Store(&_entity)
	entity, _ := repo.FindByID(id)
	assert.Equal(t, id, entity.ID, "Should be equal")
	assert.Equal(t, _entity.Body, entity.Body, "Should be equal")
}

func TestUpdateEntityRepo(t *testing.T) {
	repository := NewEntity()
	entity := model.Entity{
		Body: gofakeit.Sentence(1),
	}

	id, _ := repository.Store(&entity)
	createdEntity, _ := repository.FindByID(id)
	assert.Equal(t, createdEntity.Body, entity.Body)

	changedBody := gofakeit.Sentence(2)
	createdEntity.Body = changedBody

	//Save again for testing increment
	id2, _ := repository.Update(&createdEntity)
	updatedEntity, _ := repository.FindByID(id2)
	assert.Equal(t, updatedEntity.ID, createdEntity.ID)
	assert.Equal(t, updatedEntity.Body, changedBody)

}

func TestSearchEntities(t *testing.T) {
	r := NewEntity()
	entity := model.Entity{
		Body: "search target body",
	}

	r.Store(&entity)

	entity = model.Entity{
		Body: "search no target body",
	}

	r.Store(&entity)
	q := model.EntityQuery{
		Query: "search target",
	}

	entities, _ := r.FindAllFull(&q)
	assert.Contains(t, entities[0].Body, "search target")
}
