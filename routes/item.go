package routes

import (
	"invoice-system/internal/adpaters/delivery/http"
	"invoice-system/internal/adpaters/repository"
	"invoice-system/internal/core/usecase"
	"invoice-system/pkg/mysql"

	"github.com/gin-gonic/gin"
)

func ItemRoutes(r *gin.RouterGroup) {
	itemRepository := repository.NewItemRepository(mysql.DB)
	itemUsecase := usecase.NewItemUsecase(itemRepository)

	h := http.NewItemHandler(itemUsecase)

	r.POST("/item", h.CreateItemHandler)
	r.GET("/item", h.FindItemsHandler)
}
