package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"BankTellerAPI/handlers"
	"BankTellerAPI/middlewares"
)

// RegisterRoutes initializes and returns a router with all defined routes
func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", handlers.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/logout", middlewares.AuthMiddleware(handlers.LogoutHandler)).Methods(http.MethodPost)
	router.HandleFunc("/payment", middlewares.AuthMiddleware(handlers.CreatePaymentHandler)).Methods(http.MethodPost)
	router.HandleFunc("/payment/{id}", middlewares.AuthMiddleware(handlers.UpdatePaymentStatusHandler)).Methods(http.MethodPatch)

	return router
}
