package usecase

import "go-clean-architecture/domain"

type EntityUsecase struct {
	EntityRepository EntityRepository
}

func (usecase *EntityUsecase) Store(e domain.Entity) (entity domain.Entity, err error) {
	identifier, err := usecase.EntityRepository.Store(e)
	if err != nil {
		return
	}
	entity, err = usecase.EntityRepository.GetByID(identifier)
	return
}

func (usecase *EntityUsecase) Update(e domain.Entity) (entity domain.Entity, err error) {
	identifier, err := usecase.EntityRepository.Update(e)
	if err != nil {
		return
	}
	entity, err = usecase.EntityRepository.GetByID(identifier)
	return
}

func (usecase *EntityUsecase) GetAll() (entities []domain.Entity, err error) {
	entities, err = usecase.EntityRepository.GetAll()
	return
}

func (usecase *EntityUsecase) GetByID(identifier int) (entity domain.Entity, err error) {
	entity, err = usecase.EntityRepository.GetByID(identifier)
	return
}

func (usecase *EntityUsecase) Delete(identifier int) (err error) {
	err = usecase.EntityRepository.Delete(identifier)
	return
}
