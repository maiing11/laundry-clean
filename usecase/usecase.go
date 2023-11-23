package usecase

import (
	"go.uber.org/fx"
)

// Module exports usecases present
var Module = fx.Options(
	fx.Provide(NewCustomerUC),
	fx.Provide(NewServiceUC),
	fx.Provide(NewBillUsecase),
)
