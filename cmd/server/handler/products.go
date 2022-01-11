package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ignaciofalco/test-integracion/internal/products"
	"github.com/ignaciofalco/test-integracion/pkg/web"
)

type request struct {
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if len(p) == 0 {
			ctx.JSON(404, web.NewResponse(404, nil, "No hay productos almacenados"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		fmt.Println(os.Getenv("TOKEN"))

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}

		if req.Type == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El tipo del producto es requerido"))
			return
		}

		if req.Count == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La cantidad es requerida"))
			return
		}

		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio es requerido"))
			return
		}

		p, err := c.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// StoreProducts godoc
// @Summary Update products
// @Tags Products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "Products ID"
// @Param product body request true "Product to update"
// @Success 200 {object} web.Response
// @Router /products/{id} [put]
func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		if req.Name == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}

		if req.Type == "" {
			ctx.JSON(400, gin.H{"error": "El tipo del producto es requerido"})
			return
		}

		if req.Count == 0 {
			ctx.JSON(400, gin.H{"error": "La cantidad es requerida"})
			return
		}

		if req.Price == 0 {
			ctx.JSON(400, gin.H{"error": "El precio es requerido"})
			return
		}

		p, err := c.service.Update(int(id), req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

// StoreProducts godoc
// @Summary Update products name
// @Tags Products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "Products ID"
// @Param product body request true "Product name to update"
// @Success 200 {object} web.Response
// @Router /products/{id} [patch]
func (c *Product) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		if req.Name == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}

		p, err := c.service.UpdateName(int(id), req.Name)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

// StoreProducts godoc
// @Summary Delete products
// @Tags Products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "Products ID"
// @Success 200 {object} web.Response
// @Router /products/{id} [delete]
func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
	}
}
