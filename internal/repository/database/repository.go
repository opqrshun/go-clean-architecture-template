package database

import (
	"gobackend/pkg/errors"
)

type Repository struct {
	*SQLHandler
}

func New() *Repository {
	return &Repository{
		SQLHandler: NewSQLHandler(),
	}
}

//Store
func (repo *Repository) Store(m Base) (int, error) {
	if err := repo.db.Create(m).Error; err != nil {
		return 0, errors.Wrapf(err, "database error, method: Store, ID : %d", m.GetID()).DatabaseError()
	}

	return m.GetID(), nil
}

//Update
func (repo *Repository) Update(m Base) (int, error) {
	m.SetCurrentTimeUpdatedAt()
	if err := repo.db.Save(m).Error; err != nil {
		return 0, errors.Wrapf(err, "database error, method: Update, ID : %d", m.GetID()).DatabaseError()
	}
	return m.GetID(), nil
}

//Delete
func (repo *Repository) Delete(m Base) error {
	if err := repo.db.Delete(m).Error; err != nil {
		return errors.Wrapf(err, "database error, method: Delete, ID : %d", m.GetID()).DatabaseError()
	}
	return nil
}

//FindBaseByID
func (repo *Repository) FindBaseByID(m interface{}, id int) error {
	err := repo.db.First(m, "id = ?", id).Error
	return BuildDBError(err, "repo.FindBaseByID")
}

func (repo *Entity) FindBaseByIDs(m interface{}, ids []int) error {
	err := repo.db.Where("id IN ?", ids).Find(m).Error
	return BuildDBError(err, "repo.FindBaseByIDs")
}

//GetOffset
func (repo *Repository) GetOffsetLimit(page int, pageSize int) (offset int, limit int) {
	p := page
	s := pageSize
	if p == 0 {
		p = 1
	}
	switch {
	case s > 100:
		s = 100
	case s <= 0:
		s = 20
	}
	offset = (p - 1) * s
	limit = s
	return
}
