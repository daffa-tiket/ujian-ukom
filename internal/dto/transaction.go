package dto

type Transaction struct {
	TransactionID string  `json:"transactionID"`
	ProductID     string  `json:"productID" validate:"required"`
	CustomerID    string  `json:"customerID" validate:"required"`
	Amount        int     `json:"amount" validate:"required"`
	Total         float64 `json:"total" validate:"required"`
}

type TransactionResponse struct {
	TransactionID string  `json:"transactionID"`
	ProductID     string  `json:"productID" validate:"required"`
	CustomerID    string  `json:"customerID" validate:"required"`
	Amount        int     `json:"amount" validate:"required"`
	Total         float64 `json:"total" validate:"required"`
	Customer Customer
	Product Product
}
