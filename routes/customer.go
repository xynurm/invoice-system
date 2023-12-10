package routes

import (
	"invoice-system/internal/adpaters/delivery/http"
	"invoice-system/internal/adpaters/repository"
	"invoice-system/internal/core/usecase"
	"invoice-system/pkg/mysql"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(r *gin.RouterGroup) {
	customerRepository := repository.NewCustomerRepository(mysql.DB)
	customerUsecase := usecase.NewCustomerUsecase(customerRepository)

	h := http.NewCustomerHandler(customerUsecase)

	r.POST("/customer", h.CreateCustomerHandler)
	r.GET("/customer", h.FindCustomersHandler)
	r.GET("/customer/:id", h.GetCustomerHandler)
}
