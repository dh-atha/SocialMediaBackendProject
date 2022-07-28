package delivery

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"socialmediabackendproject/config"
	"socialmediabackendproject/domain"
	"socialmediabackendproject/feature/common"
	"socialmediabackendproject/feature/middlewares"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
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
	e.GET("/profile", handler.MyProfile(), useJWT)
	e.PUT("/profile", handler.UpdateProfile(), useJWT)
	e.DELETE("/profile", handler.DeleteProfile(), useJWT)
	e.PUT("/profilepic", handler.UpdateProfilePic(), useJWT)
}

func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newUser Register
		err := c.Bind(&newUser)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "error parsing data")
		}

		err = validator.New().Struct(newUser)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		data, err := uh.userUsecase.Register(newUser.ToDomain())
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

func (uh *userHandler) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var updateData domain.User
		err := c.Bind(&updateData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "error parsing data")
		}

		userID := common.ExtractData(c)
		data, err := uh.userUsecase.UpdateUser(updateData, uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update data",
			"data":    data,
		})
	}
}

func (uh *userHandler) DeleteProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := uh.userUsecase.DeleteUser(uint(common.ExtractData(c)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, "success delete user")
	}
}

func (uh *userHandler) UpdateProfilePic() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
		data, err := uh.userUsecase.GetSpecificUser(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		file, err := c.FormFile("profilepic")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error()+"error parsing data")
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		getExt := strings.Split(file.Filename, ".")
		dst, err := os.Create(fmt.Sprint("uploads/profilepic/", strconv.Itoa(int(data.ID)), "-", data.Name, ".", getExt[len(getExt)-1]))
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, "success update profile picture")
	}
}
