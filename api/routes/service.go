package routes

import (
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/api/controllers"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/config"
)

// ServiceRoutes struct
type ServiceRoutes struct {
	handler           config.RequestHandler
	serviceController controllers.ServiceController
}

// Setup service routes
func (s ServiceRoutes) Setup() {
	api := s.handler.Gin.Group("/api")

	// services
	{
		api.GET("/services/", s.serviceController.FetchServices)
		api.GET("/services/:id", s.serviceController.FindServiceById)
		api.POST("services", s.serviceController.RegisterService)
		api.PUT("/services", s.serviceController.EditService)
		api.DELETE("/services", s.serviceController.DeleteService)
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
