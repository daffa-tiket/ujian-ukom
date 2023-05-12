package main

import (
	"projects/internal/dto"

	"github.com/google/uuid"
	//"strconv"
)

func (app *application) GetAllCustomer() ([]dto.Customer, error){
	var err error
	rows, _ := app.db.Query("select * from customer")
	defer rows.Close()

    var customers []dto.Customer

    for rows.Next() {
        var cus dto.Customer
        if err := rows.Scan(&cus.CustomerID, &cus.FirstName, &cus.LastName,
            &cus.Age); err != nil {
            return customers, err
        }
        customers = append(customers, cus)
    }
    if err = rows.Err(); err != nil {
        return customers, err
    }
    return customers, nil
}

func (app *application) GetCustomerByID(id string) (dto.Customer, error){
	//var err error
	customer := dto.Customer{}
	row := app.db.QueryRow(`SELECT * FROM customer WHERE id=$1;`, id)
	if err := row.Scan(&customer.CustomerID, &customer.FirstName, &customer.LastName, &customer.Age); err != nil {
        return customer, err
    }
    return customer, nil
} 

func (app *application) PostCustomer(data dto.Customer) (dto.Customer, error){

	data.CustomerID = uuid.New().String()
	_, err := app.db.Exec(`INSERT INTO customer (id, firstname, lastname, age) VALUES ($1, $2, $3, $4)`, data.CustomerID, data.FirstName, data.LastName, data.Age)
    if err != nil {
        return data, err
    }

    // Return the new album's ID.
    return data, nil
}

func (app *application) UpdateCustomer(data dto.Customer, id string) (dto.Customer, error){
	_, err := app.db.Exec(`UPDATE customer
	SET firstname=$1, lastname=$2, age=$3
	WHERE id=$4;`, data.FirstName, data.LastName, data.Age, id)
    if err != nil {
        return data, err
    }

    // Return the new album's ID.
	data.CustomerID = id
    return data, nil
}
