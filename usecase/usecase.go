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
	entity, err = usecase.EntityRepository.GetByID(identifier)
	return
}

func (usecase *Usecase) GetAll() (entities []domain.Entity, err error) {
	entities, err = usecase.EntityRepository.GetAll()
	return
}

func (usecase *Usecase) GetByID(identifier int) (entity domain.Entity, err error) {
	entity, err = usecase.EntityRepository.GetByID(identifier)
	return
}

func (usecase *Usecase) Delete(identifier int) (err error) {
	err = usecase.EntityRepository.Delete(identifier)
	return
}
