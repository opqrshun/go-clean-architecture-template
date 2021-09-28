package test

import (
	repo "github.com/ttaki/go-clean-architecture-template/internal/repository/database"
	"github.com/ttaki/go-clean-architecture-template/model"

	"github.com/brianvoe/gofakeit"
)

//SetEntityTestDataWithUsername
func SetEntityTestData() (r model.Entity) {
	//set entity
	repository := repo.NewEntity()

	entity := model.Entity{
		Body: gofakeit.Sentence(1),
	}
	id, _ := repository.Store(&entity)
	r, _ = repository.FindFullByID(id)
	return
}

func SetAttributeTestData() (rAttribute model.Attribute) {
	//set entity
	repository := repo.NewEntity()

	entity := model.Entity{
		Body: gofakeit.Sentence(1),
	}
	id, _ := repository.Store(&entity)
	r, _ := repository.FindFullByID(id)

	//set attribute
	repositoryAttribute := repo.NewAttribute()

	attribute := model.Attribute{
		Body:     gofakeit.Sentence(10),
		EntityID: r.ID,
	}

	attributeID, _ := repositoryAttribute.Store(&attribute)
	rAttribute, _ = repositoryAttribute.FindByID(attributeID)
	return
}
