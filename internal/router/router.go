package router

import (
	"context"
	"database/sql"
	"mampu-wallet/internal/handler"
	"mampu-wallet/internal/middleware"
	"mampu-wallet/internal/repository"
	"mampu-wallet/internal/service"
	"mampu-wallet/internal/tools"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	DB  *sql.DB
	R   *gin.Engine
	CTX context.Context
}

func (h *Handler) Routes() {
	tools.LoadEnv()

	middleware.Add(h.R, middleware.CORSMiddleware())
	v1 := h.R.Group(os.Getenv("PREFIX_API"))

	v1.GET("/check-connection", h.CheckConnection)

	// Swagger
	v1.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Repository
	walletRepository := repository.NewWalletRepository(h.DB)
	userRepository := repository.NewUserRepository(h.DB)

	// Service
	walletService := service.NewWalletService(walletRepository, h.CTX)
	userService := service.NewUserService(userRepository, h.CTX)

	// Handler
	handler.WalletRoute(walletService, v1)
	handler.UserRoute(userService, v1)
}

// @Router /check-connection [get]
// @Accept json
// @Produce json
func (h *Handler) CheckConnection(c *gin.Context) {
	c.JSON(http.StatusOK, tools.Response{
		Status:  "success",
		Message: "Success Check Connect to API",
	})
}
