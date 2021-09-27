package tests

import (
	"gobackend/model"
	repo "gobackend/infrastructure/database"

	"github.com/brianvoe/gofakeit"
)

//SetParentTestDataWithUsername
func SetParentTestData() (r model.Parent) {
	//set parent
	repository := repo.NewParent()

	parent := model.Parent{
		Body:  gofakeit.Sentence(1),
	}
	id, _ := repository.Store(&parent)
	r, _ = repository.FindFullByID(id)
	return
}

func SetModelTestData() (rModel model.Model) {
	//set parent
	repository := repo.NewParent()

	parent := model.Parent{
		Body:  gofakeit.Sentence(1),
	}
	id, _ := repository.Store(&parent)
  r, _ := repository.FindFullByID(id)

	//set model
	repositoryModel := repo.NewModel()

	model := model.Model{
		Body:      gofakeit.Sentence(10),
		ParentID: r.ID,
	}

	modelID, _ := repositoryModel.Store(&model)
	rModel, _ = repositoryModel.FindByID(modelID)
	return
}
