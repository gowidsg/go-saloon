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
	var custStruct models.Customer

	//Using mux.Vars to get value but we are getting empty map
	// params := mux.Vars(&r)
	// fmt.Println(params)\
	// err := mapstructure.Decode(params, &custStruct)
	// if err != nil {
	// 	log.Fatal("Cannot Decode the param values ", err)
	// }
	// fmt.Println(custStruct)

	//using json to get input value from body
	json.NewDecoder(r.Body).Decode(&custStruct)

	err := custController.CustSRV.AddCustomer(&custStruct)
	if err != nil {
		log.Fatal("Customer not added ", err)
	}
	json.NewEncoder(w).Encode(custStruct)
}
