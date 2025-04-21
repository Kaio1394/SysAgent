package routes

import (
	"SysAgent/internal/handlers"
	"SysAgent/internal/services"
	"github.com/gin-gonic/gin"
)

func MemoryRoutes(r *gin.Engine) {
	s := services.NewMemoryServiceImpl()
	h := handlers.NewMemoryHandlerImpl(s)

	r.GET("/GetMemoryInfo", h.GetMemoryInfo)
}
