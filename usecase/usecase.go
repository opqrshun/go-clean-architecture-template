package usecase

import "go-clean-architecture/domain"

type Usecase struct {
	EntityRepository EntityRepository
}

func (usecase *Usecase) Store(e domain.Entity) (entity domain.Entity, err error) {
	identifier, err := usecase.EntityRepository.Store(e)
	if err != nil {
		return
	}
	entity, err = usecase.EntityRepository.FindById(identifier)
	return
}

func (usecase *Usecase) GetAll() (entities []domain.Entity, err error) {
	entities, err = usecase.EntityRepository.FindAll()
	return
}

func (usecase *Usecase) GetById(identifier int) (entity domain.Entity, err error) {
	entity, err = usecase.EntityRepository.FindById(identifier)
	return
}
