package memory

import (
	"go-clean-architecture/domain"
	"sort"
	"strings"
	"time"
	// "go-clean-architecture/domain"
)

type EntityRepository struct {
	Index int
	data  map[int]domain.Entity
}

func NewEntityRepository() *EntityRepository {
	return &EntityRepository{
		Index: 0,
		data:  make(map[int]domain.Entity),
	}
}

func (repo *EntityRepository) Store(entity domain.Entity) (id int, err error) {
	// todo refactoring
	id = repo.Index + 1
	repo.Index = id
	entity.Id = id
	entity.CreatedAt = time.Now()
	entity.UpdatedAt = time.Now()
	repo.data[id] = entity

	return
}

func (repo *EntityRepository) Update(entity domain.Entity) (id int, err error) {
	entity.UpdatedAt = time.Now()
	repo.data[entity.Id] = entity

	id = entity.Id
	return
}

func (repo *EntityRepository) GetById(identifier int) (r domain.Entity, err error) {

	r, ok := repo.data[identifier]

	if !ok {
		err = domain.ErrNotFound
	}
	return
}

func (repo *EntityRepository) GetAll() (r []domain.Entity, err error) {
	//sort map key
	var keys []int
	for k := range repo.data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		r = append(r, repo.data[k])
	}
	return
}

//Search
func (repo *EntityRepository) Search(query string) (r []domain.Entity, err error) {
	for _, j := range repo.data {
		if strings.Contains(strings.ToLower(j.Text), query) {
			r = append(r, j)
		}
	}
	return
}

//Delete
func (repo *EntityRepository) Delete(identifier int) (err error) {
	_, ok := repo.data[identifier]

	if !ok {
		err = domain.ErrNotFound
		return
	}

	delete(repo.data, identifier)
	return
}
