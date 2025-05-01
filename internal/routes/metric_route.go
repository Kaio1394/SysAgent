package routes

import (
	"SysAgent/internal/handlers"
	"SysAgent/internal/repository"
	"SysAgent/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterMetricsRoute(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewAgentRepositoryImpl(db)
	sa := services.NewAgentServiceImpl(repo)
	s := services.NewMetricsServiceImpl()
	h := handlers.NewMetricsHandlerImpl(s, sa)
	r.GET("/cpu", h.GetCpuInfo)
	r.GET("/memory", h.GetMemoryInfo)
	r.GET("/disk", h.GetDiskInfo)
	r.GET("/infos", h.GetDataCollection)

	r.DELETE("/data/collection", h.DeleteDataCollection)
	r.GET("/data/collection", h.GetDataByDate)

	r.POST("/start", h.StartCollect)
	r.POST("/stop", h.StopCollect)
}
