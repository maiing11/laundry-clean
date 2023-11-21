package manager

import "git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/repository"

type RepoManager interface {
	Repository() *repository.Queries
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) Repository() *repository.Queries {
	return repository.NewRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
