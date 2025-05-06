package handlers

import (
	"SysAgent/internal/models/dto"
	"SysAgent/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ScriptHandlerImpl struct {
	s *services.ScriptsServiceImpl
}

func NewScriptHandlerImpl(s *services.ScriptsServiceImpl) *ScriptHandlerImpl {
	return &ScriptHandlerImpl{s}
}

func (h *ScriptHandlerImpl) CreateScriptHandler(c *gin.Context) {
	var dtoScript dto.ScriptCreateDto
	err := c.ShouldBindJSON(&dtoScript)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.s.CreateNewScript(context.Background(), &dtoScript)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"script": dtoScript})
}

func (h *ScriptHandlerImpl) GetAllScriptHandler(c *gin.Context) {
	model, err := h.s.GetAllScripts(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"scripts": model})
}

func (h *ScriptHandlerImpl) GetScriptHandlerByUuid(c *gin.Context) {
	uuid := c.GetHeader("Uuid")
	script, err := h.s.GetScriptByUuid(context.Background(), uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"script": script})
}

func (h *ScriptHandlerImpl) ExecuteScriptHandlerByUuid(c *gin.Context) {
	uuid := c.GetHeader("Uuid")
	result, err := h.s.ExecuteScript(context.Background(), uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"output": result})
}
