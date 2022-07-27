package factory

import (
	postsData "socialmediabackendproject/feature/posts/data"
	postDelivery "socialmediabackendproject/feature/posts/delivery"
	postUsecase "socialmediabackendproject/feature/posts/usecase"
	usersData "socialmediabackendproject/feature/users/data"
	userDelivery "socialmediabackendproject/feature/users/delivery"
	userUsecase "socialmediabackendproject/feature/users/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	userData := usersData.New(db)
	useCase := userUsecase.New(userData)
	userDelivery.New(e, useCase)

	postData := postsData.New(db)
	postUsecase := postUsecase.New(postData)
	postDelivery.New(e, postUsecase)
}
