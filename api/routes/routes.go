package routes

type route struct {
}

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
