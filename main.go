package main

import (
	"BankTellerAPI/database"
	"BankTellerAPI/handlers"
	"BankTellerAPI/routes"
	"os"

	"log"
	"net/http"
)

func main() {

	secretKeyCheck := os.Getenv("SECRET_KEY")
	if secretKeyCheck == "" {
		log.Fatal("Missing SECRET_KEY environment variable")
		return
	}

	log.Println("Starting application, adding initial user if not exists")
	err := database.EnsureFileExists(database.JWTDBPath)
	if err != nil {
		log.Fatal(err)
	}
	err = database.EnsureFileExists(database.UserDBPath)
	if err != nil {
		log.Fatal(err)
	}
	err = database.EnsureFileExists(database.PaymentDBPath)
	if err != nil {
		log.Fatal(err)
	}

	err = handlers.AddUserHandler("admin", "password")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("use username 'admin' and password 'password' to login")

	router := routes.RegisterRoutes()

	log.Println("Server running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", router))

}
