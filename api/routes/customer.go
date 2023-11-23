package routes

import (
	"log"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/enigma-laundry-clean/api/controllers"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/enigma-laundry-clean/config"
)

// CustomerRoutes struct
type CustomerRoutes struct {
	handler        config.RequestHandler
	custController controllers.CustomerController
}

// set up  routes
func (s CustomerRoutes) Setup() {
	api := s.handler.Gin.Group("/api/v1/customers")
	log.Println("setting up routes")
	// customers
	{
		api.GET("/", s.custController.FetchAllCustomers)
		api.GET("/:id", s.custController.FindCustomerById)
		api.POST("/", s.custController.RegisterCustomer)
		api.PUT("/", s.custController.EditCustomer)
		api.DELETE("/:id", s.custController.DeleteCustomer)
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
