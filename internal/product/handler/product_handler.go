package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/middleware"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/utils"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/product"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
	"go.uber.org/zap"
)

type ProductHandler struct {
	r      *gin.RouterGroup
	ts     product.ProductService
	logger *zap.Logger
}

func NewProductHandler(r *gin.RouterGroup, ts product.ProductService, logger *zap.Logger) *gin.RouterGroup {
	delivery := ProductHandler{
		r:      r,
		ts:     ts,
		logger: logger,
	}
	productRoute := delivery.r.Group("/products")
	productProtectedRoute := delivery.r.Group("/products", middleware.AuthMiddleware())
	{
		productProtectedRoute.Handle(http.MethodPost, "/", delivery.createProduct)
		productProtectedRoute.Handle(http.MethodGet, "/", delivery.viewProduct)
		productProtectedRoute.Handle(http.MethodPut, "/:productId", delivery.updateProduct)
		productProtectedRoute.Handle(http.MethodDelete, "/:productId", delivery.deleteProduct)
	}
	return productRoute
}

func (p *ProductHandler) createProduct(c *gin.Context) {
	var requestBody dto.CreateProductRequest
	errDecodeRequest := c.BindJSON(&requestBody)
	if errDecodeRequest != nil {
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["userId"].(float64))

	res, errCreate := p.ts.CreateProduct(c, &requestBody, userID)
	if errCreate != nil {
		p.logger.Sugar().Errorf("[createProduct] failed to create user, err: %v", zap.Error(errCreate))
		errResponse := utils.NewErrorResponse(c.Writer, errCreate)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusCreated, res)
	c.JSON(http.StatusCreated, response)
}

func (p *ProductHandler) viewProduct(c *gin.Context) {
	res, errView := p.ts.ViewProduct(c)
	if errView != nil {
		p.logger.Sugar().Errorf("[viewProduct] failed to view product, err: %v", zap.Error(errView))
		errResponse := utils.NewErrorResponse(c.Writer, errView)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}

func (p *ProductHandler) updateProduct(c *gin.Context) {
	data := dto.EditProductRequest{}

	errDecodeRequest := c.BindJSON(&data)
	if errDecodeRequest != nil {
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["userId"].(float64))
	productID := c.Param("productId")
	ProductIDConv, _ := strconv.ParseUint(productID, 10, 64)

	res, errUpdate := p.ts.UpdateProduct(c, ProductIDConv, userID, &data)
	if errUpdate != nil {
		p.logger.Sugar().Errorf("[UpdateProduct] failed to update product, id: %v, err: %v", ProductIDConv, zap.Error(errUpdate))
		errResponse := utils.NewErrorResponse(c.Writer, errUpdate)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}

func (p *ProductHandler) deleteProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["userId"].(float64))
	taskID := c.Param("productId")
	taskIDConv, _ := strconv.ParseUint(taskID, 10, 64)

	res, errDelete := p.ts.DeleteProduct(c, taskIDConv, userID)
	if errDelete != nil {
		p.logger.Sugar().Errorf("[deleteProduct] failed to delete product, id: %v, err: %v", taskID, zap.Error(errDelete))
		errResponse := utils.NewErrorResponse(c.Writer, errDelete)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusCreated, res)
	c.JSON(http.StatusOK, response)
}
