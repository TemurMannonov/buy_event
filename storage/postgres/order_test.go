package postgres_test

import (
	"testing"

	"github.com/TemurMannonov/buy_event/models"
	"github.com/stretchr/testify/assert"
)

func createOrder(t *testing.T) *models.Order {
	order := &models.Order{
		ID:         createRandomId(t),
		CustomerID: createCustomer(t).ID,
		Products:   fakeData.Name(),
		TotalPrice: 1000,
	}

	err := strg.Order().Create(order)
	assert.NoError(t, err)

	return order
}

func deleteOrder(t *testing.T, id string) {
	err := strg.Order().Delete(id)
	assert.NoError(t, err)
}

func TestCreateOrder(t *testing.T) {
	order := createOrder(t)
	deleteOrder(t, order.ID)
	deleteCustomer(t, order.CustomerID)
}

func TestUpdateOrder(t *testing.T) {
	order := createOrder(t)

	order.Products = fakeData.Name()
	order.TotalPrice = 2000

	err := strg.Order().Update(order)
	assert.NoError(t, err)

	deleteOrder(t, order.ID)
	deleteCustomer(t, order.CustomerID)
}

func TestOrderGet(t *testing.T) {
	order := createOrder(t)

	_, err := strg.Order().Get(order.ID)

	assert.NoError(t, err)

	deleteOrder(t, order.ID)
	deleteCustomer(t, order.CustomerID)
}

func TestOrderGetAll(t *testing.T) {
	_, err := strg.Order().GetAll()
	assert.NoError(t, err)
}

func TestOrderDelete(t *testing.T) {
	order := createOrder(t)

	deleteOrder(t, order.ID)
	deleteCustomer(t, order.CustomerID)
}
