package routes

import "go.uber.org/fx"

// Module exports depedency to container
var Module = fx.Options(
	fx.Provide(NewCustomerRoutes),
	fx.Provide(NewServiceRoutes),
	fx.Provide(NewBillRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	customerRoutes CustomerRoutes,
	serviceRoutes ServiceRoutes,
	billRoutes BillRoutes,
) Routes {
	return Routes{
		customerRoutes,
		serviceRoutes,
		billRoutes,
	}
}

// Setup all route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
