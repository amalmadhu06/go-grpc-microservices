package models

type Order struct {
	Id        int64 `json:"id" gorm:"primaryKey"`
	Price     int64 `json:"price"`
	ProductId int64 `json:"product_id"`
	UserId    int64 `json:"user_id"`
}
