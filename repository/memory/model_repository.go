package memory

import (
	"gobackend/model"
	"sort"
	"time"
	// "gobackend/model"
)

type Model struct {
	Index    int
	entities map[int]model.Model
}

func NewModel() *Model {
	return &Model{
		Index:    0,
		entities: make(map[int]model.Model),
	}
}

func (repo *Model) Store(model model.Model) (id int, err error) {
	// todo refactoring
	id = repo.Index + 1
	repo.Index = id
	model.ID = id
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	repo.entities[id] = model

	return
}

func (repo *Model) Update(model model.Model) (id int, err error) {
	model.UpdatedAt = time.Now()
	repo.entities[model.ID] = model

	id = model.ID
	return
}

func (repo *Model) FindByID(id int) (model model.Model, err error) {

	model, ok := repo.entities[id]

	if !ok {
		err = domain.ErrNotFound
	}
	return
}

func (repo *Model) FindAll() (entities []model.Model, err error) {
	//sort map key
	var keys []int
	for k := range repo.entities {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		entities = append(entities, repo.entities[k])
	}
	return
}

//Delete
func (repo *Model) Delete(id int) (err error) {
	_, ok := repo.entities[id]

	if !ok {
		err = domain.ErrNotFound
		return
	}

	delete(repo.entities, id)
	return
}
