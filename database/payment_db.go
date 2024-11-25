package database

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"

	"BankTellerAPI/models"
)

const PaymentDBPath = "payments.json"

// WritePayments writes payments to the JSON file
func WritePayments(payments []models.Payment) error {
	log.Println("Writing payments to database")
	file, err := os.Create(PaymentDBPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		cerr := file.Close()
		if cerr != nil {
			fmt.Printf("Failed to close file: %v\n", cerr)
		}
	}(file)

	err = json.NewEncoder(file).Encode(payments)
	if err != nil {
		return err
	}

	return nil
}

// AddPayment adds a payment for a given user ID
func AddPayment(userID string) (models.Payment, error) {
	log.Println("Adding payment to database")
	payments, err := ReadPayments()
	if err != nil {
		return models.Payment{}, err
	}

	newPayment := models.Payment{
		ID:     uuid.NewString(),
		UserID: userID,
		Status: "unpaid",
	}
	payments = append(payments, newPayment)

	err = WritePayments(payments)
	if err != nil {
		return models.Payment{}, err
	}

	return newPayment, nil
}

func UpdatePaymentStatus(paymentID string, userID string) (models.Payment, error) {
	log.Println("Updating payment to database")
	payments, err := ReadPayments()
	if err != nil {
		return models.Payment{}, err
	}

	for i, payment := range payments {
		if payment.ID == paymentID {

			if payments[i].UserID != userID {
				return models.Payment{}, fmt.Errorf("you are not allowed to update this payment")
			}

			if payments[i].Status == "paid" {
				return models.Payment{}, fmt.Errorf("this payment is already paid")
			}
			payments[i].Status = "paid"

			err = WritePayments(payments)
			if err != nil {
				return models.Payment{}, err
			}
			return payments[i], nil
		}
	}

	return models.Payment{}, fmt.Errorf("payment not found")
}

// ReadPayments reads all payments from the JSON file
func ReadPayments() ([]models.Payment, error) {
	log.Println("Reading payments from database")
	var payments []models.Payment

	file, err := os.Open(PaymentDBPath)
	if err != nil {
		if os.IsNotExist(err) {
			return payments, nil
		}
		return nil, err
	}
	defer func(file *os.File) {
		cerr := file.Close()
		if cerr != nil {
			fmt.Printf("Failed to close file: %v\n", cerr)
		}
	}(file)

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if stat.Size() == 0 {
		return payments, nil
	}

	err = json.NewDecoder(file).Decode(&payments)
	if err != nil {
		return nil, err
	}

	return payments, nil
}
