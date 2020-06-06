package memory

import (
	"go-clean-architecture/domain"
	"sort"
	"time"
	// "go-clean-architecture/domain"
)

type EntityRepository struct {
	Index    int
	entities map[int]domain.Entity
}

func NewEntityRepository() *EntityRepository {
	return &EntityRepository{
		Index:    0,
		entities: make(map[int]domain.Entity),
	}
}

func (repo *EntityRepository) Store(entity domain.Entity) (id int, err error) {
	// todo refactoring
	id = repo.Index + 1
	repo.Index = id
	entity.ID = id
	entity.CreatedAt = time.Now()
	entity.UpdatedAt = time.Now()
	repo.entities[id] = entity

	return
}

func (repo *EntityRepository) Update(entity domain.Entity) (id int, err error) {
	entity.UpdatedAt = time.Now()
	repo.entities[entity.ID] = entity

	id = entity.ID
	return
}

func (repo *EntityRepository) GetByID(identifier int) (entity domain.Entity, err error) {

	entity, ok := repo.entities[identifier]

	if !ok {
		err = domain.ErrNotFound
	}
	return
}

func (repo *EntityRepository) GetAll() (entities []domain.Entity, err error) {
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
func (repo *EntityRepository) Delete(identifier int) (err error) {
	_, ok := repo.entities[identifier]

	if !ok {
		err = domain.ErrNotFound
		return
	}

	delete(repo.entities, identifier)
	return
}
