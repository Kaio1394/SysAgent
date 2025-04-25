package routes

import (
	"SysAgent/internal/handlers"
	"SysAgent/internal/repository"
	"SysAgent/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAgentRoutes(db *gorm.DB, r *gin.Engine) {
	repo := repository.NewAgentRepositoryImpl(db)
	serv := services.NewAgentServiceImpl(repo)
	h := handlers.NewAgentHandlerImpl(serv)

	r.POST("/agent", h.CreateStatusAgent)
	r.GET("/agent/status", h.GetStatusAgent)
	r.PUT("/agent/status", h.UpdateStatusAgent)
}
