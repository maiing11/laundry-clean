package manager

import (
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/usecase"
)

type UsecaseManager interface {
	Usecase() usecase.Usecase
}

type usecaseManager struct {
	repo RepoManager
	d    usecase.Domains
}

func (u *usecaseManager) Usecase() usecase.Usecase {
	return *usecase.NewUsecase(u.d, *u.repo.Repository())
}

func NewUsecaseManager(repo RepoManager) UsecaseManager {
	return &usecaseManager{repo: repo}
}
