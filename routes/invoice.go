package routes

import (
	"invoice-system/internal/adpaters/delivery/http"

	"invoice-system/internal/adpaters/repository"
	"invoice-system/internal/core/usecase"
	"invoice-system/pkg/mysql"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(r *gin.RouterGroup) {
	invoiceRepository := repository.NewInvoiceRepository(mysql.DB)
	invoiceUsecase := usecase.NewInvoiceUsecase(invoiceRepository)

	h := http.NewInvoiceHandler(invoiceUsecase)

	r.POST("/invoice", h.CreateInvoiceHandler)
	r.GET("/invoice", h.FindInvoicesHandler)
	r.GET("/invoice/:id", h.GetInvoiceHandler)
}
