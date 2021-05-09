package repo

import "github.com/TemurMannonov/buy_event/models"

type OrderStorageI interface {
	Create(order *models.Order) error
	Update(order *models.Order) error
	Get(id string) (*models.Order, error)
	GetAll() ([]*models.Order, error)
	Delete(id string) error
}
