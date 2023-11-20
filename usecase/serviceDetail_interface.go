package usecase

import (
	"context"
	"fmt"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/domains"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/entities"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/repository"
)

type ServiceDetailUC struct {
	repo repository.Queries
}

func NewServiceDetailUC(repo repository.Queries) domains.ServiceDetailsUC {
	return ServiceDetailUC{repo: repo}
}

func (uc ServiceDetailUC) FindServiceById(ctx context.Context, id int) (entities.TblServiceDetail, error) {
	service, err := uc.repo.GetServiceDetails(ctx, id)
	if err != nil {
		return entities.TblServiceDetail{}, fmt.Errorf("failed to find by id:%v", err)
	}
	return service, nil
}

func (uc ServiceDetailUC) RegisterNewLaundryService(ctx context.Context, arg entities.TblServiceDetail) (entities.TblServiceDetail, error) {
	service, err := uc.repo.CreateServiceDetails(ctx, arg)
	if err != nil {
		return entities.TblServiceDetail{}, fmt.Errorf("register failed because :%v", err)
	}
	return service, nil
}

func (uc ServiceDetailUC) EditLaundryService(ctx context.Context, arg entities.TblServiceDetail) error {
	err := uc.repo.UpdateServiceDetails(ctx, arg)
	if err != nil {
		return fmt.Errorf("update service failed! :%v", err)
	}
	return nil
}

func (uc ServiceDetailUC) FetchingAllServices(ctx context.Context) ([]entities.TblServiceDetail, error) {
	service, err := uc.repo.ListServiceDetails(ctx)
	if err != nil {
		return []entities.TblServiceDetail{}, fmt.Errorf("failed to fetching all services! :%v", err)
	}
	return service, nil
}

func (uc ServiceDetailUC) DeleteLaundryService(ctx context.Context, id int) error {
	err := uc.repo.DeleteServiceDetail(ctx, id)
	if err != nil {
		return fmt.Errorf("the service deletion process failed! :%v", err)
	}
	return nil
}
