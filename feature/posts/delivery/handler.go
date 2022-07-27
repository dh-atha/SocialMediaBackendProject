package delivery

import (
	"net/http"
	"socialmediabackendproject/domain"
	"strconv"

	"github.com/labstack/echo"
)

type postHandler struct {
	postUsecase domain.PostUsecase
}

func New(e *echo.Echo, ps domain.PostUsecase) {
	handler := &postHandler{
		postUsecase: ps,
	}
	e.GET("/post/:id", handler.GetSpecificPost())
}

func (ps *postHandler) GetSpecificPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		data, err := ps.postUsecase.GetSpecificPost(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get specific user",
			"data":    ToGetSpecificPost(data),
		})
	}
}
