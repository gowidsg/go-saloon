package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gowidsg/go-saloon/controllers"
	"github.com/gowidsg/go-saloon/repository"
	"github.com/gowidsg/go-saloon/services"
	"github.com/jinzhu/gorm"
)

func main() {
	conn := MySQLConn()
	m := mux.NewRouter()
	route := m.PathPrefix("/api/go").Subrouter()
	repo := repository.NewRepositorySRV()
	header := handlers.AllowedHeaders([]string{"Content-Type", "application/json"})
	srv := &http.Server{
		Handler:      handlers.CORS(header)(route),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         ":9000",
	}
	fmt.Print("Hi above ini")
	m.HandleFunc("/", Index)
	initiateController(conn, route, repo)
	fmt.Print("above listen")
	log.Fatal(srv.ListenAndServe())
	fmt.Print("After listen")
	defer func() {
		conn.Close()
	}()
}

func initiateController(conn *gorm.DB, route *mux.Router, repo *repository.RepositorySRV) {
	custSrv := services.NewCustomerServices(conn, repo)
	custController := controllers.NewCustomerController(custSrv)
	custController.CustomerRoute(route)
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Print("HHHHHHHHHHHHHH")
	w.Write([]byte("Welcome Customer"))
}
