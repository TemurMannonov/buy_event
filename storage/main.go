package storage

import (
	"github.com/TemurMannonov/buy_event/storage/postgres"
	"github.com/TemurMannonov/buy_event/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Customer() repo.CustomerStorageI
	Order() repo.OrderStorageI
	Log() repo.LogStorageI
}

type storagePg struct {
	customerRepo repo.CustomerStorageI
	orderRepo    repo.OrderStorageI
	logRepo      repo.LogStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		customerRepo: postgres.NewCustomer(db),
		orderRepo:    postgres.NewOrder(db),
		logRepo:      postgres.NewLog(db),
	}
}

func (s *storagePg) Customer() repo.CustomerStorageI {
	return s.customerRepo
}

func (s *storagePg) Order() repo.OrderStorageI {
	return s.orderRepo
}

func (s *storagePg) Log() repo.LogStorageI {
	return s.logRepo
}
