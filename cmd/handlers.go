package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"projects/internal/dto"
	"projects/internal/response"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetAllCustomer()

	err = response.JSONCustom(w, data, err)
}

func (app *Application) getCustomerByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetCustomerByID(ps.ByName("id"))

	err = response.JSONCustom(w, data, err)
}

func (app *Application) addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var customer, res dto.Customer
	var err error
	json.Unmarshal(reqBody, &customer)

	err = app.Validator.Struct(customer)

	if err == nil {
		res, err = app.PostCustomer(customer)
	}

	err = response.JSONCustom(w, res, err)
}

func (app *Application) editCustomer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var customer, res dto.Customer
	var err error
	json.Unmarshal(reqBody, &customer)

	err = app.Validator.Struct(customer)

	if err == nil {
		res, err = app.UpdateCustomer(customer, ps.ByName("id"))
	}

	err = response.JSONCustom(w, res, err)
}

// Product
func (app *Application) getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetAllProduct()

	err = response.JSONCustom(w, data, err)
}

func (app *Application) getProductByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetProductByID(ps.ByName("id"))

	err = response.JSONCustom(w, data, err)
}

func (app *Application) addProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product, res dto.Product
	var err error
	json.Unmarshal(reqBody, &product)

	err = app.Validator.Struct(product)

	if err == nil {
		res, err = app.PostProduct(product)
	}

	err = response.JSONCustom(w, res, err)
}

func (app *Application) editProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product, res dto.Product
	var err error
	json.Unmarshal(reqBody, &product)

	err = app.Validator.Struct(product)

	if err == nil {
		res, err = app.UpdateProduct(product, ps.ByName("id"))
	}

	err = response.JSONCustom(w, res, err)
}

// Transaction
func (app *Application) getTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetAllTransaction()

	err = response.JSONCustom(w, data, err)
}

func (app *Application) getTransactionByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetTransactionByID(ps.ByName("id"))

	err = response.JSONCustom(w, data, err)
}

func (app *Application) addTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var transaction, res dto.Transaction
	var err error
	json.Unmarshal(reqBody, &transaction)

	err = app.Validator.Struct(transaction)

	fmt.Println(transaction)

	if err == nil {
		res, err = app.PostTransaction(transaction)
	}

	err = response.JSONCustom(w, res, err)
}

func (app *Application) editTransaction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var transaction, res dto.Transaction
	var err error
	json.Unmarshal(reqBody, &transaction)

	err = app.Validator.Struct(transaction)

	if err == nil {
		res, err = app.UpdateTransaction(transaction, ps.ByName("id"))
	}

	err = response.JSONCustom(w, res, err)
}

func (app *Application) removeTransaction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	id := ps.ByName("id")
	if id != "" {
		err = app.DeleteTransaction(ps.ByName("id"))
	}

	err = response.JSONCustom(w, nil, err)
}
