package controllers

import (
	"net/http"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/dto"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/usecase/interfaces"
	"github.com/gin-gonic/gin"
)

type BillController struct {
	uc interfaces.BillUsecase
}

func NewBillController(uc interfaces.BillUsecase) BillController {
	return BillController{uc: uc}
}

func (b BillController) RegisterNewBill(ctx *gin.Context) {
	var payload dto.BillRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	payloadResponse, err := b.uc.RegisterNewBill(payload)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}
	dto.SendSingleResponse(ctx, http.StatusCreated, "Nice 1", payloadResponse)
}
