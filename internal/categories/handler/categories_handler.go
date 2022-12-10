package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/categories"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/middleware"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/utils"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
)

type CategoriesHandler struct {
	r   *gin.RouterGroup
	ctg categories.CategoriesService
}

func NewCategoriesHandler(r *gin.RouterGroup, ctg categories.CategoriesService) *gin.RouterGroup {
	delivery := CategoriesHandler{
		r:   r,
		ctg: ctg,
	}
	categoriesRoute := delivery.r.Group("/categories")

	categoriesProtectedRoute := delivery.r.Group("/categories", middleware.AuthMiddleware())
	{
		categoriesProtectedRoute.Handle(http.MethodPost, "/", delivery.createCategories)
		categoriesProtectedRoute.Handle(http.MethodGet, "/", delivery.viewCategories)
		categoriesProtectedRoute.Handle(http.MethodPatch, "/:categoryId", delivery.updateCategories)
		categoriesProtectedRoute.Handle(http.MethodDelete, "/:categoryId", delivery.deleteCategories)
	}
	return categoriesRoute
}

func (ctgh *CategoriesHandler) createCategories(c *gin.Context) {
	var requestBody dto.CreateCategoriesRequest
	err := c.BindJSON(&requestBody)
	if err != nil {
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	role, _ := userData["user_role"].(string)

	res, err := ctgh.ctg.CreateCategories(c, &requestBody, role)
	if err != nil {
		log.Printf("[createCategory] failed to create user, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusCreated, res)
	c.JSON(http.StatusCreated, response)
}

func (ctgh *CategoriesHandler) viewCategories(c *gin.Context) {
	res, err := ctgh.ctg.ViewCategories(c)
	if err != nil {
		log.Printf("[viewCategory] failed to view category, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}

func (ctgh *CategoriesHandler) updateCategories(c *gin.Context) {
	data := dto.EditCategoriesRequest{}

	err := c.BindJSON(&data)
	if err != nil {
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	role, _ := userData["user_role"].(string)
	categoryID := c.Param("categoryId")
	categoryIDConv, _ := strconv.ParseUint(categoryID, 10, 64)

	res, err := ctgh.ctg.UpdateCategories(c, categoryIDConv, role, &data)
	if err != nil {
		log.Printf("[updateTask] failed to update task, id: %v, err: %v", categoryIDConv, err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}

func (ctgh *CategoriesHandler) deleteCategories(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	role, _ := userData["user_role"].(string)
	categoryID := c.Param("categoryId")
	categoryIDConv, _ := strconv.ParseUint(categoryID, 10, 64)

	res, err := ctgh.ctg.DeleteCategories(c, categoryIDConv, role)
	if err != nil {
		log.Printf("[deleteCategory] failed to delete category, id: %v, err: %v", categoryID, err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusCreated, res)
	c.JSON(http.StatusOK, response)
}
