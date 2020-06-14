package usecase

import "go-clean-architecture/domain"

type EntityUsecase struct {
	EntityRepository EntityRepository
}

func (usecase *EntityUsecase) Store(e domain.Entity) (r domain.Entity, err error) {
	identifier, err := usecase.EntityRepository.Store(e)
	if err != nil {
		return
	}
	r, err = usecase.EntityRepository.GetById(identifier)
	return
}

func (usecase *EntityUsecase) Update(e domain.Entity) (r domain.Entity, err error) {
	identifier, err := usecase.EntityRepository.Update(e)
	if err != nil {
		return
	}
	r, err = usecase.EntityRepository.GetById(identifier)
	return
}

func (usecase *EntityUsecase) GetById(identifier int) (r domain.Entity, err error) {
	r, err = usecase.EntityRepository.GetById(identifier)
	return
}

func (usecase *EntityUsecase) Search(query string) (r []domain.Entity, err error) {
	r, err = usecase.EntityRepository.Search(query)
	return
}

func (usecase *EntityUsecase) GetAll() (r []domain.Entity, err error) {
	r, err = usecase.EntityRepository.GetAll()
	return
}

func (usecase *EntityUsecase) Delete(identifier int) (err error) {
	err = usecase.EntityRepository.Delete(identifier)
	return
}
