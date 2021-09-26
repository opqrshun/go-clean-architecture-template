package tests

import (
	"gobackend/domain"
	"gobackend/response"
	repo "gobackend/infrastructure/database"

	"github.com/brianvoe/gofakeit"
)

//SetParentTestDataWithUsername
func SetParentTestData() (r response.Parent) {
	//set parent
	repository := repo.NewParent()

	parent := domain.Parent{
		Body:  gofakeit.Sentence(1),
	}
	id, _ := repository.Store(&parent)
	r, _ = repository.FindFullByID(id)
	return
}

func SetModelTestData() (rModel domain.Model) {
	//set parent
	repository := repo.NewParent()

	parent := domain.Parent{
		Body:  gofakeit.Sentence(1),
	}
	id, _ := repository.Store(&parent)
  r, _ := repository.FindFullByID(id)

	//set model
	repositoryModel := repo.NewModel()

	model := domain.Model{
		Body:      gofakeit.Sentence(10),
		ParentID: r.ID,
	}

	modelID, _ := repositoryModel.Store(&model)
	rModel, _ = repositoryModel.FindByID(modelID)
	return
}
