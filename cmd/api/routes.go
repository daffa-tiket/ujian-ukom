package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()

	mux.NotFound = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	mux.HandlerFunc("GET", "/customers", app.getCustomers)
	mux.GET("/customer/:id", app.getCustomerByID)
	mux.HandlerFunc("POST", "/customer", app.addCustomer)
	mux.PUT("/customer/:id", app.editCustomer)

	return app.recoverPanic(mux)
}
