package interfaces

import (
	"context"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/entities"
)

type ServiceRepository interface {
	CreateServiceDetails(ctx context.Context, arg entities.TblServiceDetail) (entities.TblServiceDetail, error)
	DeleteServiceDetail(ctx context.Context, serviceNo int) error
	GetServiceDetails(ctx context.Context, serviceNo int) (entities.TblServiceDetail, error)
	ListServiceDetails(ctx context.Context) ([]entities.TblServiceDetail, error)
	UpdateServiceDetails(ctx context.Context, arg entities.TblServiceDetail) error
}
