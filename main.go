package main

import (
	"fmt"
	"log"
	"net/http"
	_ "time"

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
	header := handlers.AllowedHeaders([]string{"Content-Type", "token"})
	methods := handlers.AllowedMethods([]string{"PUT", "DELETE", "GET", "POST", "OPTION"})
	origin := handlers.AllowedOrigins([]string{"*"})
	srv := &http.Server{
		Handler: handlers.CORS(header, methods, origin)(route),
		// WriteTimeout: 150 * time.Second,
		// ReadTimeout:  150 * time.Second,
		Addr: ":9000",
	}

	route.HandleFunc("/index", index)
	initiateController(conn, route, repo)
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

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Customer"))
}
