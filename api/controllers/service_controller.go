package controllers

import (
	"net/http"
	"strconv"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/domains"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/dto"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/entities"
	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	uc domains.ServiceDetailsUC
}

func NewServiceController(uc domains.ServiceDetailsUC) ServiceController {
	return ServiceController{uc: uc}
}

func (c ServiceController) FindServiceById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	service, err := c.uc.FindServiceById(ctx, id)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}
	dto.SendSingleResponse(ctx, http.StatusOK, "ok", service)

}

// /services
// create/ register service
func (c ServiceController) RegisterService(ctx *gin.Context) {
	payload := entities.TblServiceDetail{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	payloadResponse, err := c.uc.RegisterNewLaundryService(ctx, payload)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	dto.SendSingleResponse(ctx, http.StatusCreated, "Successfully created", payloadResponse)
}

func (c ServiceController) EditService(ctx *gin.Context) {
	payload := entities.TblServiceDetail{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	if err := c.uc.EditLaundryService(ctx, payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Service updated", nil)

}

func (c ServiceController) FetchServices(ctx *gin.Context) {
	service, err := c.uc.FetchingAllServices(ctx)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Feched Services.", service)
}

func (c ServiceController) DeleteService(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	if err := c.uc.DeleteLaundryService(ctx, id); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}
	dto.SendSingleResponse(ctx, http.StatusOK, "Deleted Service.", nil)
}
