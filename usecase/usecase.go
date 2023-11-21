package usecase

import (
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/domains"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/repository"
)

type Domains interface {
	domains.ServiceDetailsUC
	domains.CustomerUsecase
	domains.BillUsecase
}

type Usecase struct {
	repo repository.Queries
}

func NewUsecase(d Domains, repo repository.Queries) *Usecase {
	return &Usecase{repo: repo}
}
