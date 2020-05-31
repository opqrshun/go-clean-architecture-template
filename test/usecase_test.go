package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go-clean-architecture/domain"
	repo "go-clean-architecture/infrastructure/memory"
	"go-clean-architecture/usecase"
)

func TestStore(t *testing.T) {

	usecase := usecase.Usecase{
		EntityRepository: repo.NewEntityRepository(),
	}

	Entity := domain.Entity{}
	r, _ := usecase.Store(Entity)
	assert.Equal(t, r.Id, 1)

}

func TestGetAll(t *testing.T) {

	usecase := usecase.Usecase{
		EntityRepository: repo.NewEntityRepository(),
	}

	Entity := domain.Entity{}
	Entity2 := domain.Entity{}
	usecase.Store(Entity)
	usecase.Store(Entity2)

	entities, _ := usecase.GetAll()
	assert.Equal(t, len(entities), 2)
}

func TestDelete(t *testing.T) {

	usecase := usecase.Usecase{
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
