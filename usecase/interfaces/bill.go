package interfaces

import (
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/enigma-laundry-clean/model/dto"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/enigma-laundry-clean/model/entities"
)

type BillUsecase interface {
	RegisterNewBill(payload dto.BillRequestDto) (entities.TblBill, error)
}
