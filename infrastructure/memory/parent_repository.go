package memory

import (
	"gobackend/domain"
	"sort"
	"time"
	// "gobackend/domain"
)

type Parent struct {
	Index    int
	entities map[int]domain.Parent
}

func NewParent() *Parent {
	return &Parent{
		Index:    0,
		entities: make(map[int]domain.Parent),
	}
}

func (repo *Parent) Store(parent domain.Parent) (id int, err error) {
	// todo refactoring
	id = repo.Index + 1
	repo.Index = id
	parent.ID = id
	parent.CreatedAt = time.Now()
	parent.UpdatedAt = time.Now()
	repo.entities[id] = parent

	return
}

func (repo *Parent) Update(parent domain.Parent) (id int, err error) {
	parent.UpdatedAt = time.Now()
	repo.entities[parent.ID] = parent

	id = parent.ID
	return
}

func (repo *Parent) FindByID(id int) (parent domain.Parent, err error) {

	parent, ok := repo.entities[id]

	if !ok {
		err = domain.ErrNotFound
	}
	return
}

func (repo *Parent) FindAll() (entities []domain.Parent, err error) {
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
func (repo *Parent) Delete(id int) (err error) {
	_, ok := repo.entities[id]

	if !ok {
		err = domain.ErrNotFound
		return
	}

	delete(repo.entities, id)
	return
}
