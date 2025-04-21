package handlers

import (
	"SysAgent/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemoryHandlerImpl struct {
	s *services.MemoryServiceImpl
}

func NewMemoryHandlerImpl(s *services.MemoryServiceImpl) *MemoryHandlerImpl {
	return &MemoryHandlerImpl{s}
}

func (h *MemoryHandlerImpl) GetMemoryInfo(c *gin.Context) {
	memory, err := h.s.GetMemoryInfo(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"memory": memory})
}
