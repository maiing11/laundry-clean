package routes

import (
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/enigma-laundry-clean/api/controllers"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/enigma-laundry-clean/config"
)

// ServiceRoutes struct
type ServiceRoutes struct {
	handler           config.RequestHandler
	serviceController controllers.ServiceController
}

// Setup service routes
func (s ServiceRoutes) Setup() {
	api := s.handler.Gin.Group("/api/v1/services")

	// services
	{
		api.GET("/", s.serviceController.FetchServices)
		api.GET("/:id", s.serviceController.FindServiceById)
		api.POST("/", s.serviceController.RegisterService)
		api.PUT("/", s.serviceController.EditService)
		api.DELETE("/:id", s.serviceController.DeleteService)
	}
}

// NewServiceRoutes creates new service controller
func NewServiceRoutes(
	handler config.RequestHandler,
	serviceController controllers.ServiceController,
) ServiceRoutes {
	return ServiceRoutes{
		handler:           handler,
		serviceController: serviceController,
	}
}
