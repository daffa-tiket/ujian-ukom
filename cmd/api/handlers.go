package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"projects/internal/dto"
	"projects/internal/response"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetAllCustomer()

	err = response.JSONCustom(w, data, err)
}

func (app *application) getCustomerByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetCustomerByID(ps.ByName("id"))

	err = response.JSONCustom(w, data, err)
}

func (app *application) addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var customer, res dto.Customer
	var err error
	json.Unmarshal(reqBody, &customer)

	err = app.validator.Struct(customer)

	if err == nil {
		res, err = app.PostCustomer(customer)
	}

	err = response.JSONCustom(w, res, err)
}

func (app *application) editCustomer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var customer, res dto.Customer
	var err error
	json.Unmarshal(reqBody, &customer)

	err = app.validator.Struct(customer)

	if err == nil {
		res, err = app.UpdateCustomer(customer, ps.ByName("id"))
	}

	err = response.JSONCustom(w, res, err)
}

//Product
func (app *application) getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetAllProduct()

	err = response.JSONCustom(w, data, err)
}

func (app *application) getProductByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetProductByID(ps.ByName("id"))

	err = response.JSONCustom(w, data, err)
}

func (app *application) addProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product, res dto.Product
	var err error
	json.Unmarshal(reqBody, &product)

	err = app.validator.Struct(product)

	if err == nil {
		res, err = app.PostProduct(product)
	}

	err = response.JSONCustom(w, res, err)
}

func (app *application) editProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product, res dto.Product
	var err error
	json.Unmarshal(reqBody, &product)

	err = app.validator.Struct(product)

	if err == nil {
		res, err = app.UpdateProduct(product, ps.ByName("id"))
	}

	err = response.JSONCustom(w, res, err)
}


//Transaction
func (app *application) getTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetAllTransaction()

	err = response.JSONCustom(w, data, err)
}

func (app *application) getTransactionByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetTransactionByID(ps.ByName("id"))

	err = response.JSONCustom(w, data, err)
}

func (app *application) addTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var transaction, res dto.Transaction
	var err error
	json.Unmarshal(reqBody, &transaction)

	err = app.validator.Struct(transaction)

	fmt.Println(transaction)

	if err == nil {
		res, err = app.PostTransaction(transaction)
	}

	err = response.JSONCustom(w, res, err)
}

func (app *application) editTransaction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var transaction, res dto.Transaction
	var err error
	json.Unmarshal(reqBody, &transaction)

	err = app.validator.Struct(transaction)

	if err == nil {
		res, err = app.UpdateTransaction(transaction, ps.ByName("id"))
	}

	err = response.JSONCustom(w, res, err)
}

func (app *application) removeTransaction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	id := ps.ByName("id")
	if id != "" {
		err = app.DeleteTransaction(ps.ByName("id"))
	}

	err = response.JSONCustom(w, nil, err)
}