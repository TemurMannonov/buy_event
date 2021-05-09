package postgres_test

import (
	"testing"

	"github.com/TemurMannonov/buy_event/models"
	"github.com/stretchr/testify/assert"
)

func createCustomer(t *testing.T) *models.Customer {
	customer := &models.Customer{
		ID:          createRandomId(t),
		PhoneNumber: fakeData.PhoneNumber(),
		Email:       fakeData.Email(),
	}

	err := strg.Customer().Create(customer)
	assert.NoError(t, err)

	return customer
}

func deleteCustomer(t *testing.T, id string) {
	err := strg.Customer().Delete(id)
	assert.NoError(t, err)
}

func TestCreateCustomer(t *testing.T) {
	customer := createCustomer(t)
	deleteCustomer(t, customer.ID)
}

func TestUpdateCustomer(t *testing.T) {
	customer := createCustomer(t)

	customer.PhoneNumber = fakeData.PhoneNumber()
	customer.Email = fakeData.Email()

	err := strg.Customer().Update(customer)
	assert.NoError(t, err)

	deleteCustomer(t, customer.ID)
}

func TestCustomerGet(t *testing.T) {
	customer := createCustomer(t)

	_, err := strg.Customer().Get(customer.ID)

	assert.NoError(t, err)

	deleteCustomer(t, customer.ID)
}

func TestCustomerGetAll(t *testing.T) {
	_, err := strg.Customer().GetAll()
	assert.NoError(t, err)
}

func TestCustomerDelete(t *testing.T) {
	customer := createCustomer(t)
	deleteCustomer(t, customer.ID)
}
