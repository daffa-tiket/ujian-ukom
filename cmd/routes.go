package cmd

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {
	mux := httprouter.New()

	mux.NotFound = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	//customer
	mux.HandlerFunc("GET", "/customers", app.getCustomers)
	mux.GET("/customer/:id", app.getCustomerByID)
	mux.HandlerFunc("POST", "/customer", app.addCustomer)
	mux.PUT("/customer/:id", app.editCustomer)

	//product
	mux.HandlerFunc("GET", "/products", app.getProducts)
	mux.GET("/product/:id", app.getProductByID)
	mux.HandlerFunc("POST", "/product", app.addProduct)
	mux.PUT("/product/:id", app.editProduct)

	//transaction
	mux.HandlerFunc("GET", "/transactions", app.getTransactions)
	mux.GET("/transaction/:id", app.getTransactionByID)
	mux.HandlerFunc("POST", "/transaction", app.addTransaction)
	mux.PUT("/transaction/:id", app.editTransaction)
	mux.DELETE("/transaction/:id", app.removeTransaction)

	return app.recoverPanic(mux)
}
