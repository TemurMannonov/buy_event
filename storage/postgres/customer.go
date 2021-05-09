package postgres

import (
	"database/sql"

	"github.com/TemurMannonov/buy_event/models"
	"github.com/TemurMannonov/buy_event/storage/repo"
	"github.com/jmoiron/sqlx"
)

type customerRepo struct {
	db *sqlx.DB
}

func NewCustomer(db *sqlx.DB) repo.CustomerStorageI {
	return &customerRepo{
		db: db,
	}
}

func (c *customerRepo) Create(customer *models.Customer) error {
	query := `
		INSERT INTO customer(
			id,
			phone_number,
			email
		) VALUES ($1, $2, $3);`

	_, err := c.db.Exec(
		query,
		customer.ID,
		customer.PhoneNumber,
		customer.Email,
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepo) Get(id string) (*models.Customer, error) {
	var customer models.Customer

	query := `
		SELECT
			id,
			phone_number,
			email
		FROM customer 
		WHERE id=$1`

	err := c.db.Get(&customer, query, id)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (c *customerRepo) GetAll() ([]*models.Customer, error) {
	var customers []*models.Customer

	query := `
		SELECT
			id,
			phone_number,
			email
		FROM customer`

	err := c.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (c *customerRepo) Update(customer *models.Customer) error {
	query := `
		UPDATE customer set 
			phone_number = $1,
			email = $2
		WHERE id = $3`

	res, err := c.db.Exec(
		query,
		customer.PhoneNumber,
		customer.Email,
		customer.ID,
	)

	if err != nil {
		return err
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (c *customerRepo) Delete(id string) error {
	query := "DELETE FROM customer WHERE id=$1"

	res, err := c.db.Exec(query, id)
	if err != nil {
		return err
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
