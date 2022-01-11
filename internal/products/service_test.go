package products

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/ignaciofalco/test-integracion/pkg/store"
	"github.com/stretchr/testify/assert"
)

type DummyRepo struct{}

func (dr *DummyRepo) GetAll() ([]Product, error) {
	return []Product{}, nil
}
func (dr *DummyRepo) Store(id int, name, productType string, count int, price float64) (Product, error) {
	return Product{}, nil
}
func (dr *DummyRepo) LastID() (int, error) {
	return 0, nil
}
func (dr *DummyRepo) UpdateName(id int, name string) (Product, error) {
	return Product{}, nil
}
func (dr *DummyRepo) Update(id int, name, productType string, count int, price float64) (Product, error) {
	return Product{}, nil
}
func (dr *DummyRepo) Delete(id int) error {
	return nil
}

func TestServiceGetAll(t *testing.T) {
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
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, input, result)
	assert.Nil(t, err)
}

func TestServiceGetAllError(t *testing.T) {
	// Initializing Input/output
	expectedError := errors.New("error for GetAll")
	dbMock := store.Mock{
		Err: expectedError,
	}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
}

func TestStore(t *testing.T) {
	testProduct := Product{
		Name:  "CellPhone",
		Type:  "Tech",
		Count: 3,
		Price: 52.0,
	}
	dbMock := store.Mock{}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, _ := myService.Store(testProduct.Name, testProduct.Type, testProduct.Count, testProduct.Price)

	assert.Equal(t, testProduct.Name, result.Name)
	assert.Equal(t, testProduct.Type, result.Type)
	assert.Equal(t, testProduct.Price, result.Price)
	assert.Equal(t, 1, result.ID)
}

func TestStoreError(t *testing.T) {
	testProduct := Product{
		Name:  "CellPhone",
		Type:  "Tech",
		Count: 3,
		Price: 52.0,
	}
	expectedError := errors.New("error for Storage")
	dbMock := store.Mock{
		Err: expectedError,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.Store(testProduct.Name, testProduct.Type, testProduct.Count, testProduct.Price)

	assert.Equal(t, expectedError, err)
	assert.Equal(t, Product{}, result)
}

func TestSum(t *testing.T) {
	expectedResult := float64(6)
	myDummyRepo := DummyRepo{}
	myService := NewService(&myDummyRepo)

	result := myService.Sum(1, 2, 3)

	assert.Equal(t, expectedResult, result)
}
