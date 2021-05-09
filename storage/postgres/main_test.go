package postgres_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/TemurMannonov/buy_event/config"
	"github.com/TemurMannonov/buy_event/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/manveru/faker"
	"github.com/stretchr/testify/assert"
)

var (
	postgresConn *sqlx.DB
	err          error
	cfg          config.Config
	strg         storage.StorageI
	fakeData     *faker.Faker
)

func createRandomId(t *testing.T) string {
	id, err := uuid.NewRandom()
	assert.NoError(t, err)
	return id.String()
}

func TestMain(m *testing.M) {
	cfg = config.Load()

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
		"disable",
	)

	fakeData, _ = faker.New("en")
	postgresConn, err = sqlx.Open("postgres", conStr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(postgresConn)

	strg = storage.NewStoragePg(postgresConn)

	os.Exit(m.Run())
}
