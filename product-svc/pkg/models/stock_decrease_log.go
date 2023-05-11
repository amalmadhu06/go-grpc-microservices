package models

type StockDecreaseLog struct {
	Id           int64 `json:"id"`
	OrderId      int64 `json:"orderId"`
	ProductRefer int64 `json:"productRefer"`
}
