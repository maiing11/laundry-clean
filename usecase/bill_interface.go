package usecase

import (
	"context"
	"fmt"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/dto"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/entities"
	"github.com/jackc/pgx/v5"
)

// type BillUsecase struct {
// 	repo       repository.Queries
// 	customerUC domains.CustomerUsecase
// 	serviceUC  domains.ServiceDetailsUC
// }

// func NewBillUsecase(
// 	repo repository.Queries,
// 	customerUC domains.CustomerUsecase,
// 	serviceUc domains.ServiceDetailsUC,
// ) domains.BillUsecase {
// 	return &BillUsecase{
// 		repo:       repo,
// 		customerUC: customerUC,
// 		serviceUC:  serviceUc,
// 	}
// }

func (b *Usecase) RegisterNewBill(payload dto.BillRequestDto) (entities.TblBill, error) {
	var newServices []entities.TblService
	var db *pgx.Conn
	var ctx context.Context
	customer, err := b.FindById(ctx, payload.CustomerId)
	if err != nil {
		return entities.TblBill{}, err
	}

	for _, v := range payload.Services {
		service, err := b.FindServiceById(ctx, v.ID)
		if err != nil {
			return entities.TblBill{}, err
		}
		newServices = append(newServices, entities.TblService{ServiceNo: service.ServiceNo})
	}
	newBill := entities.TblBill{
		Customer: customer,
		Services: newServices,
	}

	// create
	bill, err := b.repo.CreateBill(ctx, db, newBill)
	if err != nil {
		return entities.TblBill{}, fmt.Errorf("failed to register bill :%v", err)
	}
	return bill, nil
}
