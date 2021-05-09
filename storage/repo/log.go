package repo

import "github.com/TemurMannonov/buy_event/models"

type LogStorageI interface {
	Create(log *models.Log) error
	GetAll() ([]*models.Log, error)
	DeleteAll() error
}
