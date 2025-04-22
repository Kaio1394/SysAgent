package handlers

import (
	"SysAgent/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MetricsHandlerImpl struct {
	s *services.MetricsServiceImpl
}

func NewMetricsHandlerImpl(s *services.MetricsServiceImpl) *MetricsHandlerImpl {
	return &MetricsHandlerImpl{s}
}

func (h *MetricsHandlerImpl) GetMemoryInfo(c *gin.Context) {
	memory, err := h.s.GetMemoryInfo(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"memory": memory})
}

func (h *MetricsHandlerImpl) GetCpuInfo(c *gin.Context) {
	cpuInfo, err := h.s.GetCpuInfo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"cpuInfo": cpuInfo})
}

func (h *MetricsHandlerImpl) GetDiskInfo(c *gin.Context) {
	diskInfo, err := h.s.GetDiskInfo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"diskInfo": diskInfo})
}

func (h *MetricsHandlerImpl) StartCollect(c *gin.Context) {}
func (h *MetricsHandlerImpl) StopCollect(c *gin.Context)  {}
