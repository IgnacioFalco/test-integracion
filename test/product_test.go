package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ignaciofalco/test-integracion/cmd/server/handler"
	"github.com/ignaciofalco/test-integracion/internal/products"
	"github.com/ignaciofalco/test-integracion/pkg/store"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "../products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_SaveProduct_OK(t *testing.T) {
	objReq := struct {
		Code string           `json:"code"`
		Data products.Product `json:"data"`
	}{}
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/products/", `{
        "nombre": "Tester","tipo": "Funcional","cantidad": 10,"precio": 99.99
    }`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objReq)
	assert.Nil(t, err)
	assert.Equal(t, objReq.Code, "200")
	assert.Equal(t, objReq.Data.Price, 99.99)
	assert.Equal(t, objReq.Data.Count, 10)
	assert.Equal(t, objReq.Data.Type, "Funcional")
	assert.Equal(t, objReq.Data.Name, "Tester")
}

func Test_GetProduct_OK(t *testing.T) {
	objReq := struct {
		Code string             `json:"code"`
		Data []products.Product `json:"data"`
	}{}
	r := createServer()

	req, rr := createRequestTest(http.MethodGet, "/products/", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objReq)
	assert.Nil(t, err)
	assert.Equal(t, objReq.Code, "200")
	assert.Equal(t, len(objReq.Data) > 0, true)
}
