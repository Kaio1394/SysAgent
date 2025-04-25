package repository

import (
	"SysAgent/internal/models"
	"context"
	"gorm.io/gorm"
)

type AgentRepositoryImpl struct {
	db *gorm.DB
}

func NewAgentRepositoryImpl(db *gorm.DB) *AgentRepositoryImpl {
	return &AgentRepositoryImpl{db: db}
}

func (a *AgentRepositoryImpl) CreateStatusAgent(ctx context.Context, model *models.Agent) error {
	return a.db.WithContext(ctx).Create(&model).Error
}

func (a *AgentRepositoryImpl) GetAllAgent(ctx context.Context) ([]models.Agent, error) {
	var agents []models.Agent
	err := a.db.WithContext(ctx).Find(&agents).Error
	return agents, err
}

func (a *AgentRepositoryImpl) UpdateAgent(ctx context.Context, agent models.Agent) error {
	return a.db.WithContext(ctx).Model(&agent).Where("uuid = ?", agent.Uuid).Updates(&agent).Error
}
