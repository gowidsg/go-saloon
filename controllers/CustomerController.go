package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gowidsg/go-saloon/models"
	"github.com/gowidsg/go-saloon/services"
	"github.com/mitchellh/mapstructure"
)

type CustomerController struct {
	CustSRV *services.CustomerServices
}

func NewCustomerController(custsrv *services.CustomerServices) *CustomerController {
	return &CustomerController{CustSRV: custsrv}
}

func (custController *CustomerController) CustomerRoute(route *mux.Router) {
	route.HandleFunc("customer", custController.createCustomer).Methods("POST")
}

func (custController *CustomerController) createCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var custStruct models.Customer
	err := mapstructure.Decode(params, &custStruct)
	if err != nil {
		log.Fatal("Cannot Decode the param values ", err)
	}
	err = custController.CustSRV.AddCustomer(&custStruct)
	if err != nil {
		log.Fatal("Customer not added ", err)
	}
	json.NewEncoder(w).Encode(params)
}
