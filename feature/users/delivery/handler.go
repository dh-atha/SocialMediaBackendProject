package delivery

import (
	"net/http"
	"socialmediabackendproject/domain"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase domain.UserUsecase
}

func New(e *echo.Echo, us domain.UserUsecase) {
	handler := &userHandler{
		userUsecase: us,
	}
	e.POST("/login", handler.Login())
	e.POST("/users", handler.Register())
	e.GET("/users", handler.GetAllUser())
	e.GET("/users/:id", handler.GetSpecificUser())
}

func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newUser domain.User
		err := c.Bind(&newUser)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "error parsing data")
		}

		data, err := uh.userUsecase.Register(newUser)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "user created",
			"data":    data,
		})
	}
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginData domain.User
		err := c.Bind(&loginData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "error parsing data")
		}

		data, token, err := uh.userUsecase.Login(loginData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success login",
			"data":    ToGetUser(data),
			"token":   token,
		})
	}
}

func (uh *userHandler) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := uh.userUsecase.GetAllUser()
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}

		var convertToGetUser []GetUser
		for i := 0; i < len(data); i++ {
			convertToGetUser = append(convertToGetUser, ToGetUser(data[i]))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all users data",
			"data":    convertToGetUser,
		})
	}
}

func (uh *userHandler) GetSpecificUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
