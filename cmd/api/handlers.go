package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"projects/internal/dto"
	"projects/internal/response"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getCustomers(w http.ResponseWriter, r *http.Request) {
	var err error
	data, err := app.GetAllCustomer()

	err = response.JSONCustom(w, data, err)
}

func (app *application) getCustomerByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	data, err := app.GetCustomerByID(ps.ByName("id"))

	err = response.JSONCustom(w, data, err)
}

func (app *application) addCustomer(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var customer, res dto.Customer
	var err error
	json.Unmarshal(reqBody, &customer)

	err = app.validator.Struct(customer)

	if (err == nil){
		res, err = app.PostCustomer(customer)
	}

	err = response.JSONCustom(w, res, err)
}

func (app *application) editCustomer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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


