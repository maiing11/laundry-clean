package interfaces

import (
	"context"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/entities"
)

type CustomerUsecase interface {
	FindById(ctx context.Context, id string) (entities.TblCustomer, error)
	RegisterCustomer(ctx context.Context, arg entities.TblCustomer) (entities.TblCustomer, error)
	EditCustomer(ctx context.Context, arg entities.TblCustomer) error
	FetchingAllCustomers(ctx context.Context) ([]entities.TblCustomer, error)
	DeleteCustomer(ctx context.Context, id string) error
}
