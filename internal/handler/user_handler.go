package handler

import (
	"mampu-wallet/internal/domain"
	"mampu-wallet/internal/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service domain.UserService
	r       *gin.RouterGroup
}

func UserRoute(service domain.UserService, r *gin.RouterGroup) {
	h := &UserHandler{
		Service: service,
		r:       r,
	}
	v2 := h.r.Group("users")

	v2.GET("", h.GetAll)
}

// @Tags User
// @Summary Get All User
// @Description Get All User
// @Router /users [get]
// @Accept json
// @Produce json
// @param limit query integer false "Limit of pagination"
// @param page query integer false "Page of pagination"
func (h *UserHandler) GetAll(c *gin.Context) {
	pagination, err := tools.Paginate(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.Response{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	result, pagination, code, err := h.Service.GetAllUser(pagination)
	if err != nil {
		c.JSON(code, tools.Response{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Get All Users",
		Data:    result,
		Meta:    pagination,
	})
}
