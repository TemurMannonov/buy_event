package postgres

import (
	"github.com/TemurMannonov/buy_event/models"
	"github.com/TemurMannonov/buy_event/storage/repo"
	"github.com/jmoiron/sqlx"
)

type logRepo struct {
	db *sqlx.DB
}

func NewLog(db *sqlx.DB) repo.LogStorageI {
	return &logRepo{
		db: db,
	}
}

func (l *logRepo) Create(log *models.Log) error {
	query := `
		INSERT INTO log(
			id,
			message
		) VALUES ($1, $2);`

	_, err := l.db.Exec(
		query,
		log.ID,
		log.Message,
	)

	if err != nil {
		return err
	}

	return nil
}

func (l *logRepo) GetAll() ([]*models.Log, error) {
	var logs []*models.Log

	query := `
		SELECT
			id,
			message,
			created_at
		FROM log`

	err := l.db.Select(&logs, query)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (l *logRepo) DeleteAll() error {
	query := "DELETE FROM log"

	_, err := l.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
