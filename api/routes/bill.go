package routes

import (
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/api/controllers"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/config"
)

// Routes
type BillRoutes struct {
	handler        config.RequestHandler
	billController controllers.BillController
}

func (s BillRoutes) Setup() {
	api := s.handler.Gin.Group("/api/v1/bills")

	// bills
	{
		api.POST("/", s.billController.RegisterNewBill)
	}
}

// NewBillRoutes create new bill controller
func NewBillRoutes(
	handler config.RequestHandler,
	billController controllers.BillController,
) BillRoutes {
	return BillRoutes{
		handler:        handler,
		billController: billController,
	}
}
