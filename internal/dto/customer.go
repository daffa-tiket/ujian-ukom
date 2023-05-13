package dto

type Customer struct {
	CustomerID string `json:"customerID"`
	FirstName  string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	Age        int    `json:"age" validate:"required"`
	IsActive   bool   `json:"isActive" validate:"required"`
}
