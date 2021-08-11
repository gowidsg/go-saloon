package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gowidsg/go-saloon/models"
	"github.com/gowidsg/go-saloon/services"
)

type CustomerController struct {
	CustSRV *services.CustomerServices
}

func NewCustomerController(custsrv *services.CustomerServices) *CustomerController {
	return &CustomerController{CustSRV: custsrv}
}

func (custController *CustomerController) CustomerRoute(route *mux.Router) {
	fmt.Print("Hi Customer Route")
	route.HandleFunc("/customer", custController.createCustomer).Methods("POST")
}

func (custController *CustomerController) createCustomer(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(&r)

	// fmt.Println(params)

	// var custStruct models.Customer
	var xyz models.Customer

	json.NewDecoder(r.Body).Decode(&xyz)
	fmt.Println(xyz)
	// err := mapstructure.Decode(params, &custStruct)
	// if err != nil {
	// 	log.Fatal("Cannot Decode the param values ", err)
	// }
	// fmt.Println(custStruct)
	err := custController.CustSRV.AddCustomer(&xyz)
	if err != nil {
		log.Fatal("Customer not added ", err)
	}
	json.NewEncoder(w).Encode(xyz)
}
