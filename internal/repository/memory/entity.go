package memory

import (
	"sort"
	"time"

	"github.com/ttaki/go-clean-architecture-sample/model"
	"github.com/ttaki/go-clean-architecture-sample/pkg/errors"
)

type Entity struct {
	Index    int
	entities map[int]model.Entity
}

func NewEntity() *Entity {
	return &Entity{
		Index:    0,
		entities: make(map[int]model.Entity),
	}
}

func (repo *Entity) Store(m model.Entity) (int, error) {
	// todo refactoring
	id := repo.Index + 1
	repo.Index = id
	m.ID = id
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	repo.entities[id] = m

	return id, nil
}

func (repo *Entity) Update(m model.Entity) (int, error) {
	_, ok := repo.entities[m.ID]
	if !ok {
		return 0, errors.Errorf("record is not found").NotFound()
	}

	m.UpdatedAt = time.Now()
	repo.entities[m.ID] = m

	return m.ID, nil
}

func (repo *Entity) FindByID(id int) (model.Entity, error) {

	m, ok := repo.entities[id]

	if !ok {
		return model.Entity{}, errors.Errorf("record is not found").NotFound()
	}
	return m, nil
}

func (repo *Entity) FindAll() ([]model.Entity, error) {
	var s []model.Entity
	//sort map key
	var keys []int
	for k := range repo.entities {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		s = append(s, repo.entities[k])
	}
	return s, nil
}

//Delete
func (repo *Entity) Delete(id int) error {
	_, ok := repo.entities[id]

	if !ok {
		return errors.Errorf("record is not found").NotFound()
	}

	delete(repo.entities, id)
	return nil
}
