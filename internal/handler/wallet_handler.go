package handler

import (
	"mampu-wallet/internal/domain"
	"mampu-wallet/internal/tools"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	Service domain.WalletService
	r       *gin.RouterGroup
}

func WalletRoute(service domain.WalletService, r *gin.RouterGroup) {
	h := &WalletHandler{
		Service: service,
		r:       r,
	}
	v2 := h.r.Group("wallets")

	v2.GET("/:user_id", h.GetBalance)
	v2.POST("", h.Withdraw)
}

// @Tags Wallet
// @Summary Get Balance Wallet
// @Description Get Balance Wallet
// @Router /wallets/{user_id} [get]
// @Accept json
// @Produce json
// @param user_id path string true "User ID"
func (h *WalletHandler) GetBalance(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	balance, code, err := h.Service.GetBalance(userID)
	if err != nil {
		c.JSON(code, tools.Response{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Balance : " + tools.FormatRupiah(int(balance)),
	})
}

// @Tags Wallet
// @Summary Withdraw Wallet
// @Description Withdraw Wallet
// @Router /wallets [post]
// @Accept json
// @Produce json
// @Param request body domain.Withdraw true "Payload Body for Withdraw [RAW]"
func (h *WalletHandler) Withdraw(c *gin.Context) {
	var form *domain.Withdraw
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, tools.Response{
			Status:  "failed",
			Message: "Error : " + err.Error(),
		})
		return
	}

	balance, code, err := h.Service.Withdraw(form)
	if err != nil {
		c.JSON(code, tools.Response{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Withdraw Success, Balance Now : " + tools.FormatRupiah(int(balance)),
	})
}
