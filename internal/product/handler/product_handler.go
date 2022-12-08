package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/middleware"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/utils"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/product"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
)

type ProductHandler struct {
	r  *gin.RouterGroup
	ts product.ProductService
}

func NewProductHandler(r *gin.RouterGroup, ts product.ProductService) *gin.RouterGroup {
	delivery := ProductHandler{
		r:  r,
		ts: ts,
	}
	productRoute := delivery.r.Group("/products")
	productProtectedRoute := delivery.r.Group("/products", middleware.AuthMiddleware())
	{
		productProtectedRoute.Handle(http.MethodPost, "/", delivery.createProduct)
		productProtectedRoute.Handle(http.MethodGet, "/", delivery.viewProduct)
		productProtectedRoute.Handle(http.MethodPut, "/:productId", delivery.updateProduct)
	}
	return productRoute
}

func (p *ProductHandler) createProduct(c *gin.Context) {
	var requestBody dto.CreateProductRequest
	err := c.BindJSON(&requestBody)
	if err != nil {
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["userId"].(float64))

	res, err := p.ts.CreateProduct(c, &requestBody, userID)
	if err != nil {
		log.Printf("[createProduct] failed to create user, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusCreated, res)
	c.JSON(http.StatusCreated, response)
}

func (p *ProductHandler) viewProduct(c *gin.Context) {
	res, err := p.ts.ViewProduct(c)
	if err != nil {
		log.Printf("[viewProduct] failed to view product, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}

func (p *ProductHandler) updateProduct(c *gin.Context) {
	data := dto.EditProductRequest{}

	err := c.BindJSON(&data)
	if err != nil {
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["userId"].(float64))
	productID := c.Param("productId")
	ProductIDConv, _ := strconv.ParseUint(productID, 10, 64)

	res, err := p.ts.UpdateProduct(c, ProductIDConv, userID, &data)
	if err != nil {
		log.Printf("[UpdateProduct] failed to update product, id: %v, err: %v", ProductIDConv, err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}
