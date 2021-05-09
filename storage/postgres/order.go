package postgres

import (
	"database/sql"

	"github.com/TemurMannonov/buy_event/models"
	"github.com/TemurMannonov/buy_event/storage/repo"
	"github.com/jmoiron/sqlx"
)

type orderRepo struct {
	db *sqlx.DB
}

func NewOrder(db *sqlx.DB) repo.OrderStorageI {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) Create(order *models.Order) error {
	query := `
		INSERT INTO "order"(
			id,
			customer_id,
			products,
			total_price
		) VALUES ($1, $2, $3, $4);`

	_, err := o.db.Exec(
		query,
		order.ID,
		order.CustomerID,
		order.Products,
		order.TotalPrice,
	)

	if err != nil {
		return err
	}

	return nil
}

func (o *orderRepo) Get(id string) (*models.Order, error) {
	var order models.Order

	query := `
		SELECT
			id,
			customer_id,
			products,
			total_price
		FROM "order" 
		WHERE id=$1`

	err := o.db.Get(&order, query, id)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *orderRepo) GetAll() ([]*models.Order, error) {
	var orders []*models.Order

	query := `
		SELECT
			id,
			customer_id,
			products,
			total_price
		FROM "order"`

	err := o.db.Select(&orders, query)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *orderRepo) Update(order *models.Order) error {
	query := `
		UPDATE "order" SET 
			products = $1,
			total_price = $2
		WHERE id = $3`

	res, err := o.db.Exec(
		query,
		order.Products,
		order.TotalPrice,
		order.ID,
	)

	if err != nil {
		return err
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (o *orderRepo) Delete(id string) error {
	query := `DELETE FROM "order" where id=$1`

	res, err := o.db.Exec(query, id)
	if err != nil {
		return err
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
