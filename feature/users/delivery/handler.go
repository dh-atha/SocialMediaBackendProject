package delivery

import (
	"net/http"
	"socialmediabackendproject/config"
	"socialmediabackendproject/domain"
	"socialmediabackendproject/feature/common"
	"socialmediabackendproject/feature/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	userUsecase domain.UserUsecase
}

func New(e *echo.Echo, us domain.UserUsecase) {
	handler := &userHandler{
		userUsecase: us,
	}
	useJWT := middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
	e.POST("/login", handler.Login())
	e.POST("/users", handler.Register())
	e.GET("/users", handler.GetAllUser())
	e.GET("/users/:id", handler.GetSpecificUser())
	e.PUT("/profile",handler.UpdateUser(), useJWT)
	e.GET("/profile", handler.MyProfile(), useJWT)
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
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		data, err := uh.userUsecase.GetSpecificUser(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get specific user",
			"data":    ToGetSpecificUser(data),
		})
	}
}

func (uh *userHandler) MyProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
		data, err := uh.userUsecase.GetSpecificUser(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success showing my profile",
			"data":    ToGetSpecificUser(data),
		})
	}
}


func (uh *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var updateUser domain.User
		err := c.Bind(&updateUser)
		id := common.ExtractData(c)
		
		data, err := uh.userUsecase.UpdateUser(uint(id), updateUser)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "user updated",
			"data":    data,
		})
	}
}