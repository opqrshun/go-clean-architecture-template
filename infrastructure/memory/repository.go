package memory

import (
	"errors"
	"go-clean-architecture/domain"
	"sort"
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
	id = repo.Index + 1
	repo.Index = id
	entity.Id = id
	repo.entities[id] = entity

	return
}

func (repo *EntityRepository) FindById(identifier int) (entity domain.Entity, err error) {

	entity, ok := repo.entities[identifier]

	if !ok {
		err = errors.New("not found")
	}
	return
}

func (repo *EntityRepository) FindAll() (entities []domain.Entity, err error) {
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
