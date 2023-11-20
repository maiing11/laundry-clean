package dto

import (
	"time"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/entities"
)

type BillRequestDto struct {
	DateIn       time.Time             `json:"date_in"`
	DateFinished time.Time             `json:"date_finished"`
	CustomerId   string                `json:"customer"`
	Services     []entities.TblService `json:"services"`
}
