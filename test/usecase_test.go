package main

import (
	"testing"

	"go-clean-architecture/domain"
	repo "go-clean-architecture/infrastructure/memory"
	"go-clean-architecture/usecase"
)

func TestAddEntitySuccess(t *testing.T) {

	usecase := usecase.Usecase{
		EntityRepository: repo.NewEntityRepository(),
	}

	Entity := domain.Entity{}
	r, err := usecase.Store(Entity)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if r.Id != 1 {
		t.Fatalf("failed test %v", r.Id)
	}

	entities, err := usecase.GetAll()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if entities[0].Id != 1 {
		t.Fatal("failed test")
	}
}
