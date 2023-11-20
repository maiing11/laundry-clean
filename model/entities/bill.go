package entities

import (
	"time"
)

type TblBill struct {
	ID           int32        `json:"id"`
	DateIn       time.Time    `json:"date_in"`
	DateFinished time.Time    `json:"date_finished"`
	Customer     TblCustomer  `json:"customer"`
	Services     []TblService `json:"services"`
	BillAmount   int64        `json:"bill_amount"`
}