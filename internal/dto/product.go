package dto

type Product struct {
	ProductID   string  `json:"productID"`
	ProductName string  `json:"productName" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	IsActive    bool    `json:"isActive" validate:"required"`
}
