package routes

import (
	"SysAgent/internal/handlers"
	"SysAgent/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterMetricsRoute(r *gin.Engine) {
	s := services.NewMetricsServiceImpl()
	h := handlers.NewMetricsHandlerImpl(s)
	r.GET("/cpu", h.GetCpuInfo)
	r.GET("/memory", h.GetMemoryInfo)
	r.GET("/disk", h.GetDiskInfo)
	r.GET("/infos", h.GetDataCollection)
}
