package handlers

import (
	"SysAgent/internal/models/dto"
	"SysAgent/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AgentHandlerImpl struct {
	s *services.AgentServiceImpl
}

func NewAgentHandlerImpl(s *services.AgentServiceImpl) *AgentHandlerImpl {
	return &AgentHandlerImpl{s}
}

func (a *AgentHandlerImpl) CreateStatusAgent(c *gin.Context) {
	var dtoAgent dto.AgentCreateDto
	if err := c.ShouldBindJSON(&dtoAgent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := a.s.CreateStatusAgent(context.Background(), &dtoAgent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (a *AgentHandlerImpl) GetStatusAgent(c *gin.Context) {
	agent, err := a.s.GetStatusAgent(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": agent.Status})
}

func (a *AgentHandlerImpl) UpdateStatusAgent(c *gin.Context) {
	status := c.GetHeader("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status header is required"})
		return
	}
	err := a.s.UpdateStatus(context.Background(), status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}
