package memory

import (
	"gobackend/domain"
	"sort"
	"time"
	// "gobackend/domain"
)

type Model struct {
	Index    int
	entities map[int]domain.Model
}

func NewModel() *Model {
	return &Model{
		Index:    0,
		entities: make(map[int]domain.Model),
	}
}

func (repo *Model) Store(model domain.Model) (id int, err error) {
	// todo refactoring
	id = repo.Index + 1
	repo.Index = id
	model.ID = id
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	repo.entities[id] = model

	return
}

func (repo *Model) Update(model domain.Model) (id int, err error) {
	model.UpdatedAt = time.Now()
	repo.entities[model.ID] = model

	id = model.ID
	return
}

func (repo *Model) FindByID(id int) (model domain.Model, err error) {

	model, ok := repo.entities[id]

	if !ok {
		err = domain.ErrNotFound
	}
	return
}

func (repo *Model) FindAll() (entities []domain.Model, err error) {
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
