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

type commentHandler struct {
	commentUsecase domain.CommentUsecase
}

func New(e *echo.Echo, cu domain.CommentUsecase) {
	handler := &commentHandler{
		commentUsecase: cu,
	}
	useJWT := middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
	e.POST("comments/:id", handler.AddComment(), useJWT)
	e.DELETE("comments/:id", handler.DeleteComment(), useJWT)
}

func (ch *commentHandler) AddComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "error get id")
		}

		var commentData domain.Comment = domain.Comment{
			Post_ID: uint(id),
			User_ID: uint(common.ExtractData(c)),
			Caption: c.FormValue("caption"),
		}

		data, err := ch.commentUsecase.AddComment(commentData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success add comment",
			"data":    data,
		})
	}
}

func (ch *commentHandler) DeleteComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "error get id")
		}

		var commentData = domain.Comment{
			ID:      uint(id),
			User_ID: uint(common.ExtractData(c)),
		}

		err = ch.commentUsecase.DeleteComment(commentData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, "success delete comment")
	}
}
