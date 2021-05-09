package repo

import "github.com/TemurMannonov/buy_event/models"

type CustomerStorageI interface {
	Create(customer *models.Customer) error
	Update(customer *models.Customer) error
	Get(id string) (*models.Customer, error)
	GetAll() ([]*models.Customer, error)
	Delete(id string) error
}
