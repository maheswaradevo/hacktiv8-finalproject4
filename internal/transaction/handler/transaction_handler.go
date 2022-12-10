package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/middleware"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/utils"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/transaction"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/constants"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
	"go.uber.org/zap"
)

type transactionHandler struct {
	r      *gin.RouterGroup
	ts     transaction.TransactionService
	logger *zap.Logger
}

func NewTransactionHandler(r *gin.RouterGroup, ts transaction.TransactionService, logger *zap.Logger) *gin.RouterGroup {
	delivery := transactionHandler{
		r:      r,
		ts:     ts,
		logger: logger,
	}
	transactionRoute := delivery.r.Group("/transactions", middleware.AuthMiddleware())
	{
		transactionRoute.Handle(http.MethodPost, "", delivery.doTransaction)
		transactionRoute.Handle(http.MethodGet, "/my-transactions", delivery.viewMyTransaction)
	}
	return transactionRoute
}

func (t *transactionHandler) doTransaction(c *gin.Context) {
	transactionRequest := &dto.DoTransactionRequest{}

	errDecodeRequest := json.NewDecoder(c.Request.Body).Decode(transactionRequest)
	if errDecodeRequest != nil {
		t.logger.Sugar().Errorf("[doTransaction] failed to parse json data", zap.Error(errDecodeRequest))
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	userLoginData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userLoginData["userId"].(float64))

	transactionResponse, errTransaction := t.ts.DoTransaction(c, transactionRequest, userID)
	if errTransaction != nil {
		t.logger.Sugar().Errorf("[doTransaction] failed to do transaction", zap.Error(errTransaction))
		errResponse := utils.NewErrorResponse(c.Writer, errTransaction)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	response := utils.NewSuccessResponseWriter(c.Writer, constants.TransactionSuccess, http.StatusCreated, transactionResponse)
	c.JSON(http.StatusCreated, response)
}

func (t *transactionHandler) viewMyTransaction(c *gin.Context) {
	userLoginData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userLoginData["userId"].(float64))
	role, _ := userLoginData["userRole"].(string)

	if strings.EqualFold(constants.CustomerRole, role) {
		errCustomer := errors.ErrOnlyCustomer
		t.logger.Sugar().Errorf("[viewMyTransaction] only customer can access", zap.Error(errCustomer))
		errResponse := utils.NewErrorResponse(c.Writer, errCustomer)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	myTransactions, errMyTransaction := t.ts.ViewMyTransaction(c, userID)
	if errMyTransaction != nil {
		t.logger.Sugar().Errorf("[viewMyTransaction] failed to view my transaction", zap.Error(errMyTransaction))
		errResponse := utils.NewErrorResponse(c.Writer, errMyTransaction)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	response := utils.NewSuccessResponseWriter(c.Writer, constants.ViewMyTransaction, http.StatusOK, myTransactions)
	c.JSON(http.StatusOK, response)
}
