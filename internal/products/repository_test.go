package products

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/ignaciofalco/test-integracion/pkg/store"
	"github.com/stretchr/testify/assert"
)

const errorGetAll = "error for GetAll"

func TestGetAllError(t *testing.T) {
	// Initializing Input/output
	expectedError := errors.New(errorGetAll)
	dbMock := store.Mock{
		Data: nil,
		Err:  expectedError,
	}
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeMocked)

	_, err := myRepo.GetAll()

	assert.Equal(t, err, expectedError)
}

func TestGetAll(t *testing.T) {
	// Initializing Input/output
	input := []Product{
		{
			ID:    1,
			Name:  "CellPhone",
			Type:  "Tech",
			Count: 3,
			Price: 250,
		}, {
			ID:    2,
			Name:  "Notebook",
			Type:  "Tech",
			Count: 10,
			Price: 1750.5,
		},
	}
	dataJson, _ := json.Marshal(input)
	dbStub := store.Mock{
		Data: dataJson,
		Err:  nil,
	}
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbStub,
	}
	myRepo := NewRepository(&storeMocked)

	resp, _ := myRepo.GetAll()

	assert.Equal(t, resp, input)
}
