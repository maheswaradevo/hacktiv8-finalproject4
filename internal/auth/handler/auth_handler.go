package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/auth"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
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
