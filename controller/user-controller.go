package controller

import (
	"fmt"
	"learn-go-api/dto"
	"learn-go-api/helper"
	"learn-go-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserController interface {
	Update(context *gin.Context)
	Profile(contex *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService service.JWTService
}

func NewUSerCOntroller(userService service.UserService, jwtService service.JWTService) UserController{
	return &userController{
		userService: userService,
		jwtService: jwtService,
	}
}

func (c *userController) Update(contex *gin.Context) {
	var UserUpdateDTO dto.UserUpdateDTO
	errDTO := contex.ShouldBind(&UserUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		contex.AbortWithStatusJSON(http.StatusBadRequest, res)
		
		return
	}

	authHeader := contex.GetHeader("Authorization")
	token, errorToken := c.jwtService.ValidateToken((authHeader))
	if errorToken != nil {
		panic(errorToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	UserUpdateDTO.ID = id
	u := c.userService.Update(UserUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	contex.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	user := c.userService.Profile((fmt.Sprintf("%v", claims["user_id"])))
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}