package tests

import (
	"BankTellerAPI/database"
	"BankTellerAPI/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var testDBPath = "payments.json"

func cleanUpTestDB() {
	_ = os.Remove(testDBPath)
}

func TestWritePayments(t *testing.T) {
	cleanUpTestDB()

	payments := []models.Payment{
		{ID: "1", UserID: "user1", Status: "unpaid"},
		{ID: "2", UserID: "user2", Status: "unpaid"},
	}

	err := database.WritePayments(payments)

	require.NoError(t, err)
	_, err = os.Stat(testDBPath)
	require.NoError(t, err)

	cleanUpTestDB()
}

func TestReadPaymentsFileExists(t *testing.T) {
	testPayments := []models.Payment{
		{ID: "1", UserID: "user1", Status: "unpaid"},
	}
	_ = database.WritePayments(testPayments)

	payments, err := database.ReadPayments()

	require.NoError(t, err)
	assert.Len(t, payments, 1)
	assert.Equal(t, "user1", payments[0].UserID)

	cleanUpTestDB()
}

func TestReadPaymentsFileNotExist(t *testing.T) {
	payments, err := database.ReadPayments()

	require.NoError(t, err)
	assert.Len(t, payments, 0)
}

func TestAddPayment(t *testing.T) {
	cleanUpTestDB()

	userID := "user1"

	payment, err := database.AddPayment(userID)

	require.NoError(t, err)
	assert.Equal(t, userID, payment.UserID)
	assert.Equal(t, "unpaid", payment.Status)
	assert.NotEmpty(t, payment.ID)

	cleanUpTestDB()
}

func TestUpdatePaymentStatus(t *testing.T) {
	cleanUpTestDB()

	userID := "user1"
	payment, err := database.AddPayment(userID)
	require.NoError(t, err)

	updatedPayment, err := database.UpdatePaymentStatus(payment.ID, userID)

	require.NoError(t, err)
	assert.Equal(t, "paid", updatedPayment.Status)

	_, err = database.UpdatePaymentStatus(payment.ID, userID)
	assert.Error(t, err)

	cleanUpTestDB()
}

func TestUpdatePaymentStatusInvalidUser(t *testing.T) {
	cleanUpTestDB()

	userID := "user1"
	anotherUserID := "user2"
	payment, err := database.AddPayment(userID)
	require.NoError(t, err)

	_, err = database.UpdatePaymentStatus(payment.ID, anotherUserID)

	assert.Error(t, err)

	cleanUpTestDB()
}

func TestUpdatePaymentStatusPaymentNotFound(t *testing.T) {
	cleanUpTestDB()

	_, err := database.UpdatePaymentStatus("nonexistent_id", "user1")

	assert.Error(t, err)

	cleanUpTestDB()
}
