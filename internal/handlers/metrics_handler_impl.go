package handlers

import (
	"SysAgent/internal/logger"
	"SysAgent/internal/models"
	"SysAgent/internal/models/dto"
	"SysAgent/internal/services"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type MetricsHandlerImpl struct {
	s  *services.MetricsServiceImpl
	sa *services.AgentServiceImpl
}

func NewMetricsHandlerImpl(s *services.MetricsServiceImpl, sa *services.AgentServiceImpl) *MetricsHandlerImpl {
	return &MetricsHandlerImpl{s, sa}
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

func (h *MetricsHandlerImpl) GetDataCollection(c *gin.Context) {
	memory, err := h.s.GetMemoryInfo(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cpuInfo, err := h.s.GetCpuInfo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	diskInfo, err := h.s.GetDiskInfo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	infos := models.Metric{Memory: memory, Cpu: cpuInfo, Disk: diskInfo}
	c.JSON(http.StatusOK, gin.H{"infos": infos})
}
func (h *MetricsHandlerImpl) StartCollect(c *gin.Context) {
	var timeNumber int
	timeStr := c.GetHeader("time")
	if timeStr == "" {
		timeStr = "5"
	}
	timeNumber, err := strconv.Atoi(timeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.sa.UpdateStatus(context.Background(), "RUNNING")
	if err != nil {
		logger.Log.Errorln(err)
		return
	}
	go func() {
		for {
			logger.Log.Infoln("collect started")
			agent, err := h.sa.GetStatusAgent(context.Background())
			if err != nil {
				logger.Log.Errorln(err)
				return
			}
			if agent.Status == "STOPPED" {
				logger.Log.Errorln("stopped agent")
				return
			}
			err = h.sa.UpdateStatus(context.Background(), "RUNNING")
			if err != nil {
				logger.Log.Errorln(err)
				return
			}
			memory, err := h.s.GetMemoryInfo(context.Background())
			if err != nil {
				logger.Log.Errorln(err)
				return
			}
			cpuInfo, err := h.s.GetCpuInfo()
			if err != nil {
				logger.Log.Errorln(err)
				return
			}
			diskInfo, err := h.s.GetDiskInfo()
			if err != nil {
				logger.Log.Errorln(err)
				return
			}

			infos := models.Metric{Memory: memory, Cpu: cpuInfo, Disk: diskInfo}
			bytes, err := json.Marshal(infos)
			if err != nil {
				logger.Log.Errorln(err)
				return
			}
			jsonResult := string(bytes)
			var dataDto dto.CollectMetricCreateDto
			dataDto.JsonResult = jsonResult

			err = h.sa.CollectData(context.Background(), &dataDto)
			if err != nil {
				logger.Log.Errorln(err)
				return
			}
			logger.Log.Infoln("collect started")
			time.Sleep(time.Duration(timeNumber) * time.Second)
		}
	}()
	c.JSON(http.StatusOK, gin.H{"info": "start collect success"})
}
func (h *MetricsHandlerImpl) StopCollect(c *gin.Context) {
	agent, err := h.sa.GetStatusAgent(context.Background())
	if err != nil {
		logger.Log.Errorln(err)
		return
	}
	if agent.Status == "STOPPED" {
		logger.Log.Errorln("agent already stopped")
		c.JSON(http.StatusOK, gin.H{"info": "agent already stopped"})
		return
	}
	err = h.sa.UpdateStatus(context.Background(), "STOPPED")
	if err != nil {
		logger.Log.Errorln(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "stop collect success"})
}

func (h *MetricsHandlerImpl) DeleteDataCollection(c *gin.Context) {
	err := h.sa.DeleteAllData(context.Background())
	if err != nil {
		logger.Log.Errorln(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "delete data collection success"})
}

func (h *MetricsHandlerImpl) GetDataByDate(c *gin.Context) {
	startTimeStr := c.GetHeader("startTime")
	if startTimeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startTime required"})
		return
	}
	enbTime := c.GetHeader("endTime")
	if enbTime == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endTime required"})
		return
	}
	listData, err := h.sa.GetDataByDate(context.Background(), startTimeStr, enbTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": listData})
}
