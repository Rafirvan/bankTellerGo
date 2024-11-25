package handlers

import (
	"BankTellerAPI/database"
	"BankTellerAPI/middlewares"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// CreatePaymentHandler handles the creation of a new payment for a given user
func CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("firing CreatePaymentHandler")
	userID := middlewares.GetUserID(r)

	payment, err := database.AddPayment(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("error adding payment: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"success": payment,
	})
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdatePaymentStatusHandler updates the status of a payment by payment ID
func UpdatePaymentStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("firing UpdatePaymentStatusHandler")
	vars := mux.Vars(r)
	paymentIDStr := vars["id"]
	if paymentIDStr == "" {
		http.Error(w, "paymentID is required", http.StatusBadRequest)
		return
	}
	userID := middlewares.GetUserID(r)

	updatedPayment, err := database.UpdatePaymentStatus(paymentIDStr, userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating payment status: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"success": updatedPayment,
	})
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
