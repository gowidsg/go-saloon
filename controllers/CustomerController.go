package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
	route.HandleFunc("/customer", custController.addCustomer).Methods("POST")
	route.HandleFunc("/customer", custController.getAllCustomer).Methods("GET")
	// route.HandleFunc("/customer/{id}", custController.getCustomerByID).Methods("POST")
	route.HandleFunc("/customer/{userid}", custController.getCustomerByUserID).Methods("GET")
	route.HandleFunc("/customer/{userid}", custController.deleteCustomer).Methods("DELETE")
	route.HandleFunc("/customer/{userid}", custController.updateCustomer).Methods("PUT")

}

func (custController *CustomerController) getCustomerByUserID(w http.ResponseWriter, r *http.Request) {
	customer := models.Customer{}
	customer.UserID = strings.TrimSpace(mux.Vars(r)["userid"])
	if len(customer.UserID) == 0 {
		w.Write([]byte("Provide valid userId"))
		return
	}
	err := custController.CustSRV.GetCustomerByUserIDService(&customer)
	// err := custController.CustSRV.GetCustomerByUserIDService(&customer, userId)
	if err != nil {
		w.Write([]byte("Customer Not Found!!! " + customer.UserID))
		return
	}
	json.NewEncoder(w).Encode(customer)
}

// func (custController *CustomerController) getCustomerByID(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(mux.Vars(r)["id"])
// 	if err != nil {
// 		w.Write([]byte("Pass integer value"))
// 		return
// 	}
// 	w.Write([]byte("Getting customer by id " + strconv.Itoa(id)))
// }

func (custController *CustomerController) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []models.Customer{}
	custController.CustSRV.GetAllCustomerService(&customers)
	if len(customers) == 0 {
		w.Write([]byte("Customer Table is empty"))
		return
	}
	json.NewEncoder(w).Encode(customers)
}

func (custController *CustomerController) addCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Data can't be handled"))
		return
	}
	if len(body) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No data is coming"))
		return
	}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to parse data"))
		return
	}

	//using json to get input value from body Method 1
	// json.NewDecoder(r.Body).Decode(&customer)

	err = custController.CustSRV.AddCustomerService(&customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Customer not added"))
	}
	json.NewEncoder(w).Encode(customer)
}

func (custController *CustomerController) deleteCustomer(w http.ResponseWriter, r *http.Request) {
	customer := models.Customer{}
	customer.UserID = strings.TrimSpace(mux.Vars(r)["userid"])
	if len(customer.UserID) == 0 {
		w.Write([]byte("Provide valid userId"))
		return
	}
	err := custController.CustSRV.DeleteCustomerService(&customer)
	// err := custController.CustSRV.GetCustomerByUserIDService(&customer, userId)
	if err != nil {
		w.Write([]byte("Customer Not Found!!! " + customer.UserID))
		fmt.Print(err)
		return
	}
	w.Write([]byte("Customer Deleted  " + customer.UserID))
}

func (custController *CustomerController) updateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer = models.Customer{}
	customer.UserID = strings.TrimSpace(mux.Vars(r)["userid"])
	json.NewDecoder(r.Body).Decode(&customer)
	err := custController.CustSRV.UpdateCustomerService(&customer)
	if err != nil {
		w.Write([]byte("Customer Not found!!! " + customer.UserID))
	}
	// fmt.Println(customer)
	w.Write([]byte("Customer Updated"))
	json.NewEncoder(w).Encode(customer)

}
