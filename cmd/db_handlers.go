package cmd

import (
	"projects/internal/dto"

	"github.com/google/uuid"
)

// Customer
func (app *Application) GetAllCustomer() ([]dto.Customer, error) {
	var err error
	rows, _ := app.DB.Query("select * from customer")
	defer rows.Close()

	var customers []dto.Customer

	if err = rows.Err(); err != nil {
		return customers, err
	}

	for rows.Next() {
		var cus dto.Customer
		if err := rows.Scan(&cus.CustomerID, &cus.FirstName, &cus.LastName,
			&cus.Age, &cus.IsActive); err != nil {
			return customers, err
		}
		customers = append(customers, cus)
	}

	return customers, nil
}

func (app *Application) GetCustomerByID(id string) (dto.Customer, error) {
	//var err error
	customer := dto.Customer{}
	row := app.DB.QueryRow(`SELECT * FROM customer WHERE id=$1;`, id)
	if err := row.Scan(&customer.CustomerID, &customer.FirstName, &customer.LastName, &customer.Age, &customer.IsActive); err != nil {
		return customer, err
	}
	return customer, nil
}

func (app *Application) PostCustomer(data dto.Customer) (dto.Customer, error) {

	data.CustomerID = uuid.New().String()
	_, err := app.DB.Exec(`INSERT INTO customer (id, firstname, lastname, age, is_active) VALUES ($1, $2, $3, $4, $5)`, data.CustomerID, data.FirstName, data.LastName, data.Age, true)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (app *Application) UpdateCustomer(data dto.Customer, id string) (dto.Customer, error) {
	_, err := app.DB.Exec(`UPDATE customer
	SET firstname=$1, lastname=$2, age=$3, is_active=$4
	WHERE id=$5;`, data.FirstName, data.LastName, data.Age, data.IsActive, id)
	if err != nil {
		return data, err
	}

	data.CustomerID = id
	return data, nil
}

// Product
func (app *Application) GetAllProduct() ([]dto.Product, error) {
	var err error
	rows, _ := app.DB.Query("select * from product")
	defer rows.Close()

	var products []dto.Product

	if err = rows.Err(); err != nil {
		return products, err
	}

	for rows.Next() {
		var product dto.Product
		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.Price,
			&product.Stock, &product.IsActive); err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (app *Application) GetProductByID(id string) (dto.Product, error) {
	//var err error
	product := dto.Product{}
	row := app.DB.QueryRow(`SELECT * FROM product WHERE id=$1;`, id)
	if err := row.Scan(&product.ProductID, &product.ProductName, &product.Price, &product.Stock, &product.IsActive); err != nil {
		return product, err
	}
	return product, nil
}

func (app *Application) PostProduct(data dto.Product) (dto.Product, error) {

	data.ProductID = uuid.New().String()
	_, err := app.DB.Exec(`INSERT INTO product (id, productname, stock, price, is_active) VALUES ($1, $2, $3, $4, $5)`, data.ProductID, data.ProductName, data.Price, data.Stock, true)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (app *Application) UpdateProduct(data dto.Product, id string) (dto.Product, error) {
	_, err := app.DB.Exec(`UPDATE product
	SET productname=$1, price=$2, stock=$3, is_active=$4
	WHERE id=$5;`, data.ProductName, data.Price, data.Stock, data.IsActive, id)
	if err != nil {
		return data, err
	}

	data.ProductID = id
	return data, nil
}

//transaction

func (app *Application) GetAllTransaction() ([]dto.TransactionResponse, error) {
	var err error
	rows, _ := app.DB.Query("select * from transaction left join customer on customer.id = transaction.customerid left join product on product.id = transaction.productid")
	defer rows.Close()

	var transactions []dto.TransactionResponse

	if err = rows.Err(); err != nil {
		return transactions, err
	}

	for rows.Next() {
		var transaction dto.TransactionResponse
		if err := rows.Scan(&transaction.TransactionID, &transaction.Amount, &transaction.Total,
			&transaction.CustomerID, &transaction.ProductID,
			&transaction.Customer.CustomerID, &transaction.Customer.FirstName, &transaction.Customer.LastName,
			&transaction.Customer.Age, &transaction.Customer.IsActive,
			&transaction.Product.ProductID, &transaction.Product.ProductName,
			&transaction.Product.Stock, &transaction.Product.Price, &transaction.Product.IsActive); err != nil {
			return transactions, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (app *Application) GetTransactionByID(id string) (dto.TransactionResponse, error) {
	//var err error
	transaction := dto.TransactionResponse{}
	row := app.DB.QueryRow(`SELECT * FROM transaction WHERE id=$1;`, id)
	if err := row.Scan(&transaction.TransactionID, &transaction.Amount, &transaction.Total, &transaction.CustomerID, &transaction.ProductID); err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (app *Application) PostTransaction(data dto.Transaction) (dto.Transaction, error) {

	data.TransactionID = uuid.New().String()
	_, err := app.DB.Exec(`INSERT INTO transaction (id, customerid, productid, amount, total) VALUES ($1, $2, $3, $4, $5)`, data.TransactionID, data.CustomerID, data.ProductID, data.Amount, data.Total)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (app *Application) UpdateTransaction(data dto.Transaction, id string) (dto.Transaction, error) {
	_, err := app.DB.Exec(`UPDATE transaction
	SET amount=$1, total=$2
	WHERE id=$5;`, data.Amount, data.Total, id)
	if err != nil {
		return data, err
	}

	data.ProductID = id
	return data, nil
}

func (app *Application) DeleteTransaction(id string) error {
	_, err := app.DB.Exec(`DELETE FROM transaction
	WHERE id = $1;`, id)
	if err != nil {
		return err
	}

	return nil
}
