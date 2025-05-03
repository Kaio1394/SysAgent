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
	r.GET("/agent/metric/cpu", h.GetCpuInfo)
	r.GET("/agent/metric/memory", h.GetMemoryInfo)
	r.GET("/agent/metric/disk", h.GetDiskInfo)
	r.GET("/agent/metric/infos", h.GetDataCollection)

	r.DELETE("/agent/data-collection", h.DeleteDataCollection)
	r.GET("/agent/data-collection", h.GetDataByDate)

	r.POST("/agent/data-collection/start", h.StartCollect)
	r.POST("/agent/data-collection/stop", h.StopCollect)
}
