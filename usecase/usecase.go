package usecase

import "go.uber.org/fx"

// Module export usecases present
var Module = fx.Options(
	fx.Provide(NewCustomerUsecase),
	fx.Provide(NewServiceDetailUC),
	fx.Provide(NewBillUsecase),
)
