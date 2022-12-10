package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/auth"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/middleware"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/utils"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/constants"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
	"go.uber.org/zap"
)

type userHandler struct {
	r      *gin.RouterGroup
	us     auth.UserService
	logger *zap.Logger
}

func NewUserHandler(r *gin.RouterGroup, us auth.UserService, logger *zap.Logger) *gin.RouterGroup {
	delivery := userHandler{
		r:      r,
		us:     us,
		logger: logger,
	}
	userRoute := delivery.r.Group("/users")
	{
		userRoute.Handle(http.MethodPost, "/register", delivery.register)
		userRoute.Handle(http.MethodPost, "/login", delivery.login)
	}
	userProtectedRoute := delivery.r.Group("/users", middleware.AuthMiddleware())
	{
		userProtectedRoute.Handle(http.MethodPatch, "/topup", delivery.topupBalance)
	}
	return userRoute
}

func (u *userHandler) register(c *gin.Context) {
	registerRequest := &dto.UserRegisterRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(registerRequest)
	if err != nil {
		u.logger.Sugar().Errorf("[register] failed to parse json data, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	validate := validator.New()
	validateError := validate.Struct(registerRequest)
	if validateError != nil {
		validateError = errors.ErrInvalidRequestBody
		u.logger.Sugar().Errorf("[register] there's data that not through the validate process")
		errResponse := utils.NewErrorResponse(c.Writer, validateError)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	isValid := utils.IsValidEmail(registerRequest.Email)
	if !isValid {
		errEmail := errors.ErrEmailFormat
		u.logger.Sugar().Errorf("wrong email format")
		errResponse := utils.NewErrorResponse(c.Writer, errEmail)
		c.JSON(http.StatusBadRequest, errResponse)
		return

	}
	res, err := u.us.Register(c, registerRequest)
	if err != nil {
		u.logger.Sugar().Errorf("[register] failed to register user, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, constants.RegisterSuccess, http.StatusCreated, res)
	c.JSON(http.StatusCreated, response)
}

func (u *userHandler) login(c *gin.Context) {
	loginRequest := &dto.UserSignInRequest{}
	errDecodeRequest := json.NewDecoder(c.Request.Body).Decode(loginRequest)
	if errDecodeRequest != nil {
		u.logger.Sugar().Errorf("[login] failed to parse json data, err: %v", errDecodeRequest)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	response, errLogin := u.us.Login(c, loginRequest)
	if errLogin != nil {
		u.logger.Sugar().Errorf("[login] failed to login, err: %v", errLogin)
		errResponse := utils.NewErrorResponse(c.Writer, errLogin)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	loginResponse := utils.NewSuccessResponseWriter(c.Writer, constants.LoginSuccess, http.StatusOK, response)
	c.JSON(http.StatusOK, loginResponse)
}

func (u userHandler) topupBalance(c *gin.Context) {
	balanceRequest := &dto.UserTopupBalanceRequest{}
	errDecodeRequest := json.NewDecoder(c.Request.Body).Decode(balanceRequest)
	if errDecodeRequest != nil {
		u.logger.Sugar().Errorf("[topupBalance] failed to parse json data, err: %v", errDecodeRequest)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	userLoginData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userLoginData["userId"].(float64))

	topupBalanceResponse, errTopUpBalance := u.us.TopupBalance(c, balanceRequest, userID)
	if errTopUpBalance != nil {
		u.logger.Sugar().Errorf("[topupBalance] failed to top up user's balance, err: %v", errTopUpBalance)
		errResponse := utils.NewErrorResponse(c.Writer, errTopUpBalance)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	topUpResponse := utils.NewSuccessResponseWriter(c.Writer, constants.TopUpSuccess, http.StatusOK, topupBalanceResponse)
	c.JSON(http.StatusOK, topUpResponse)
}
