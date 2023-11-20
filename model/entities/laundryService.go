package entities

type TblService struct {
	ID        int   `json:"id"`
	ServiceNo int   `json:"service_no"`
	Quantity  int   `json:"quantity"`
	Subtotal  int64 `json:"subtotal"`
}

type TblServiceDetail struct {
	ServiceNo   int    `json:"serviceNo" form:"serviceNo"`
	ServiceName string `json:"serviceName" form:"serviceName"`
	Units       string `json:"units" form:"units"`
	UnitPrice   int64  `json:"unitPrice" form:"unitPrice"`
}
