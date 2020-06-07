package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go-clean-architecture/domain"
	repo "go-clean-architecture/infrastructure/memory"
	"go-clean-architecture/usecase"
)

func TestStoreEntity(t *testing.T) {

	usecase := usecase.EntityUsecase{
		EntityRepository: repo.NewEntityRepository(),
	}

	Entity := domain.Entity{}
	r, _ := usecase.Store(Entity)
	assert.Equal(t, r.ID, 1)

}

func TestUpdateEntity(t *testing.T) {

	usecase := usecase.EntityUsecase{
		EntityRepository: repo.NewEntityRepository(),
	}
	entityTestData := domain.Entity{Text: "test text"}
	entity, _ := usecase.Store(entityTestData)

	//change Text
	entity.Text = "updated text"
	r, _ := usecase.Update(entity)

	assert.Equal(t, r.ID, 1)
	assert.Equal(t, r.Text, "updated text")

}

func TestGetAllEntity(t *testing.T) {

	usecase := usecase.EntityUsecase{
		EntityRepository: repo.NewEntityRepository(),
	}

	Entity := domain.Entity{}
	Entity2 := domain.Entity{}
	usecase.Store(Entity)
	usecase.Store(Entity2)

	r, _ := usecase.GetAll()
	assert.Equal(t, len(r), 2)
}

func TestDeleteEntity(t *testing.T) {

	usecase := usecase.EntityUsecase{
		EntityRepository: repo.NewEntityRepository(),
	}

	Entity := domain.Entity{}
	Entity2 := domain.Entity{}
	usecase.Store(Entity)
	usecase.Store(Entity2)

	err := usecase.Delete(2)
	assert.Nil(t, err)

	err = usecase.Delete(2)
	assert.Equal(t, domain.ErrNotFound, err)

}
