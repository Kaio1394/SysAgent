package routes

import (
	"SysAgent/internal/handlers"
	"SysAgent/internal/repository"
	"SysAgent/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterScriptRoute(db *gorm.DB, r *gin.Engine) {
	repo := repository.NewScriptRepositoryImpl(db)
	serv := services.NewScriptServiceImpl(repo)
	h := handlers.NewScriptHandlerImpl(serv)

	r.POST("/agent/script", h.CreateScriptHandler)
	r.GET("/agent/script/all", h.GetAllScriptHandler)
	r.POST("/agent/script/execute", h.ExecuteScriptHandlerByUuid)

	r.POST("/agent/script/execute/sync", h.ExecuteScryptSync)
	r.GET("/agent/script/execute/sync", h.GetResult)
}
