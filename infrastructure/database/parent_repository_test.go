package database

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"

	"gobackend/domain"
	repo "gobackend/infrastructure/database"
	"gobackend/request"
)

func TestStore(t *testing.T) {
	repo := repo.NewParent()
	_parent := domain.Parent{
		Body:  gofakeit.Sentence(1),
	}

	id, _ := repo.Store(&_parent)
	parent, _ := repo.FindByID(id)
	assert.Equal(t, id, parent.ID, "Should be equal")
	assert.Equal(t, _parent.Body, parent.Body, "Should be equal")
}


func TestUpdateParentRepo(t *testing.T) {
	repository := repo.NewParent()
	parent := domain.Parent{
		Body:  gofakeit.Sentence(1),
	}

	id, _ := repository.Store(&parent)
	createdParent, _ := repository.FindByID(id)
	assert.Equal(t, createdParent.Body, parent.Body)

	changedBody := gofakeit.Sentence(2)
	createdParent.Body = changedBody

	//Save again for testing increment
	id2, _ := repository.Update(&createdParent)
	updatedParent, _ := repository.FindByID(id2)
	assert.Equal(t, updatedParent.ID, createdParent.ID)
	assert.Equal(t, updatedParent.Body, changedBody)

}

func TestSearchParents(t *testing.T) {
	r := repo.NewParent()
	parent := domain.Parent{
		Body:   "search target body",
	}

	r.Store(&parent)

	parent = domain.Parent{
		Body:   "search no target body",
	}

	r.Store(&parent)
	q := request.ParentQuery{
		Query: "search target",
	}

	parents, _ := r.FindAllFull(&q)
	assert.Contains(t, parents[0].Body, "search target")
}



