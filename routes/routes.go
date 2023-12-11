package routes

import "github.com/gin-gonic/gin"

func RouteInit(r *gin.RouterGroup) {
	ItemRoutes(r)
	CustomerRoutes(r)
	InvoiceRoutes(r)
}
