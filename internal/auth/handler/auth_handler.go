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

type UserHandler struct {
	r  *gin.RouterGroup
	us auth.UserService
}

func NewUserHandler(r *gin.RouterGroup, us auth.UserService) *gin.RouterGroup {
	delivery := UserHandler{
		r:  r,
		us: us,
	}
	userRoute := delivery.r.Group("/users")
	{
		userRoute.Handle(http.MethodPost, "/register", delivery.register)
	}
	return userRoute
}

func (u *UserHandler) register(c *gin.Context) {
	registerRequest := &dto.UserRegisterRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(registerRequest)
	if err != nil {
		zap.S().Errorf("[register] failed to parse json data, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	res, err := u.us.Register(c, registerRequest)
	if err != nil {
		zap.S().Errorf("[register] failed to register user, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, constants.RegisterSuccess, http.StatusCreated, res)
	c.JSON(http.StatusCreated, response)
}
