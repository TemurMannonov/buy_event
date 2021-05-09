package postgres_test

import (
	"testing"

	"github.com/TemurMannonov/buy_event/models"
	"github.com/stretchr/testify/assert"
)

func createLog(t *testing.T) *models.Log {
	log := &models.Log{
		ID:      createRandomId(t),
		Message: fakeData.Sentence(10, true),
	}

	err := strg.Log().Create(log)
	assert.NoError(t, err)

	return log
}

func deleteLogs(t *testing.T) {
	err := strg.Log().DeleteAll()
	assert.NoError(t, err)
}

func TestCreateLog(t *testing.T) {
	createLog(t)
	deleteLogs(t)
}

func TestLogGetAll(t *testing.T) {
	_, err := strg.Log().GetAll()
	assert.NoError(t, err)
}

func TestLogDelete(t *testing.T) {
	createLog(t)
	deleteLogs(t)
}
