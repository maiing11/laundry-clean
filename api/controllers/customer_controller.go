package controllers

import (
	"net/http"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/domains"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/dto"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/entities"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	uc domains.CustomerUsecase
}

func NewCustomerController(uc domains.CustomerUsecase) CustomerController {
	return CustomerController{uc: uc}
}

func (c CustomerController) FindCustomerById(ctx *gin.Context) {
	id := ctx.Param("id")

	customer, err := c.uc.FindById(ctx, id)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}
	dto.SendSingleResponse(ctx, http.StatusOK, "ok", customer)

}

// /customers
// create/ register customer
func (c CustomerController) RegisterCustomer(ctx *gin.Context) {
	payload := entities.TblCustomer{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	payloadResponse, err := c.uc.RegisterCustomer(ctx, payload)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	dto.SendSingleResponse(ctx, http.StatusCreated, "Successfully created", payloadResponse)
}

func (c CustomerController) EditCustomer(ctx *gin.Context) {
	payload := entities.TblCustomer{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	if err := c.uc.EditCustomer(ctx, payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Customer updated", nil)

}

func (c CustomerController) FetchAllCustomers(ctx *gin.Context) {
	customer, err := c.uc.FetchingAllCustomers(ctx)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Feched customers.", customer)
}

func (c CustomerController) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.uc.DeleteCustomer(ctx, id); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}
	dto.SendSingleResponse(ctx, http.StatusOK, "Deleted Customer.", nil)
}
