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

func (a *AgentRepositoryImpl) CollectData(ctx context.Context, data *models.CollectMetric) error {
	return a.db.WithContext(ctx).Create(&data).Error
}

func (a *AgentRepositoryImpl) DeleteAllData(ctx context.Context) error {
	return a.db.WithContext(ctx).Where("1 = 1").Delete(&models.CollectMetric{}).Error
}

func (a *AgentRepositoryImpl) GetDataByDate(ctx context.Context, startAt string, endAt string) ([]models.CollectMetric, error) {
	var results []models.CollectMetric
	query := `
		SELECT * FROM collect_metrics
		WHERE created_at BETWEEN ? AND ?
	`
	err := a.db.Raw(query, startAt, endAt).Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
