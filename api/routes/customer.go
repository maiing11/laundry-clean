package routes

import (
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/api/controllers"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/config"
)

// CustomerRoutes struct
type CustomerRoutes struct {
	handler        config.RequestHandler
	custController controllers.CustomerController
}

// set up  routes
func (s CustomerRoutes) Setup() {
	api := s.handler.Gin.Group("api/v1")

	// customers
	{
		api.GET("/customers", s.custController.FetchAllCustomers)
		api.GET("/customers/:id", s.custController.FindCustomerById)
		api.POST("/customers", s.custController.RegisterCustomer)
		api.PUT("/customers", s.custController.EditCustomer)
		api.DELETE("/customers/id:", s.custController.DeleteCustomer)
	}

}

// NewCustomerRoutes creates new customer controller
func NewCustomerRoutes(
	handler config.RequestHandler,
	custController controllers.CustomerController,

) CustomerRoutes {
	return CustomerRoutes{
		handler:        handler,
		custController: custController,
	}
}
