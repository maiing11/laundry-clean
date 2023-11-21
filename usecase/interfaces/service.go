package interfaces

import (
	"context"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/entities"
)

type ServiceDetailsUC interface {
	FindServiceById(ctx context.Context, id int) (entities.TblServiceDetail, error)
	RegisterNewLaundryService(ctx context.Context, arg entities.TblServiceDetail) (entities.TblServiceDetail, error)
	EditLaundryService(ctx context.Context, arg entities.TblServiceDetail) error
	FetchingAllServices(ctx context.Context) ([]entities.TblServiceDetail, error)
	DeleteLaundryService(ctx context.Context, id int) error
}
