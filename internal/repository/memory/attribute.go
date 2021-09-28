package memory

import (
	"sort"
	"time"

	"github.com/ttaki/go-clean-architecture-sample/model"
	"github.com/ttaki/go-clean-architecture-sample/pkg/errors"
)

type Attribute struct {
	Index      int
	attributes map[int]model.Attribute
}

func NewAttribute() *Attribute {
	return &Attribute{
		Index:      0,
		attributes: make(map[int]model.Attribute),
	}
}

func (repo *Attribute) Store(m model.Attribute) (id int, err error) {
	// todo refactoring
	id = repo.Index + 1
	repo.Index = id
	m.ID = id
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	repo.attributes[id] = m

	return
}

func (repo *Attribute) Update(m model.Attribute) (int, error) {
	_, ok := repo.attributes[m.ID]
	if !ok {
		return 0, errors.Errorf("record is not found").NotFound()
	}

	m.UpdatedAt = time.Now()
	repo.attributes[m.ID] = m
	return m.ID, nil
}

func (repo *Attribute) FindByID(id int) (model.Attribute, error) {

	m, ok := repo.attributes[id]

	if !ok {
		return model.Attribute{}, errors.Errorf("record is not found").NotFound()
	}
	return m, nil
}

func (repo *Attribute) FindAll() ([]model.Attribute, error) {
	//sort map key
	var s []model.Attribute

	var keys []int
	for k := range repo.attributes {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		s = append(s, repo.attributes[k])
	}
	return s, nil
}

//Delete
func (repo *Attribute) Delete(id int) error {
	_, ok := repo.attributes[id]

	if !ok {
		return errors.Errorf("record is not found").NotFound()
	}

	delete(repo.attributes, id)
	return nil
}
