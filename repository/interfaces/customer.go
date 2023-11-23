package interfaces

import (
	"context"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/enigma-laundry-clean/model/entities"
)

type CustomerRepository interface {
	CreateCustomers(ctx context.Context, arg entities.TblCustomer) (entities.TblCustomer, error)
	DeleteCustomer(ctx context.Context, id string) error
	GetCustomer(ctx context.Context, id string) (entities.TblCustomer, error)
	ListCustomers(ctx context.Context) ([]entities.TblCustomer, error)
	UpdateCustomer(ctx context.Context, arg entities.TblCustomer) error
}
