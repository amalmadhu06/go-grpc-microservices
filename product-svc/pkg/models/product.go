package models

type Product struct {
	Id                int64            `json:"id"`
	Name              string           `json:"name"`
	Stock             int64            `json:"stock"`
	Price             int64            `json:"price"`
	StockDecreaseLogs StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
}
