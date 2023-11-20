package controllers

import "go.uber.org/fx"

// Modules exported for initializing application
var Module = fx.Options(
	fx.Provide(NewCustomerController),
	fx.Provide(NewServiceController),
	fx.Provide(NewBillController),
)
