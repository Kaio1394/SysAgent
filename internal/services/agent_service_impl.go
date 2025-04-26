package services

import (
	"SysAgent/internal/models"
	"SysAgent/internal/models/dto"
	"SysAgent/internal/repository"
	"context"
	"errors"
	"github.com/jinzhu/copier"
)

type AgentServiceImpl struct {
	r *repository.AgentRepositoryImpl
}

func NewAgentServiceImpl(r *repository.AgentRepositoryImpl) *AgentServiceImpl {
	return &AgentServiceImpl{r: r}
}

func (s *AgentServiceImpl) CreateStatusAgent(ctx context.Context, dto *dto.AgentCreateDto) error {
	var agent models.Agent
	agents, err := s.r.GetAllAgent(ctx)
	if err != nil {
		return err
	}
	if len(agents) > 0 {
		return errors.New("status agent already created")
	}
	_ = copier.CopyWithOption(&agent, &dto, copier.Option{IgnoreEmpty: true})

	return s.r.CreateStatusAgent(ctx, &agent)
}

func (s *AgentServiceImpl) GetStatusAgent(ctx context.Context) (*models.Agent, error) {
	var agents []models.Agent
	agents, err := s.r.GetAllAgent(ctx)
	if err != nil {
		return nil, err
	}
	if len(agents) == 0 {
		return nil, errors.New("status agent not found")
	}
	agent := agents[0]
	return &agent, nil
}

func (s *AgentServiceImpl) UpdateStatus(ctx context.Context, status string) error {
	agents, err := s.r.GetAllAgent(ctx)
	if err != nil {
		return err
	}
	if len(agents) == 0 {
		return errors.New("status agent not found")
	}
	agent := agents[0]
	agent.Status = status
	err = s.r.UpdateAgent(ctx, agent)
	if err != nil {
		return err
	}
	return nil
}

func (s *AgentServiceImpl) CollectData(ctx context.Context, data *dto.CollectMetricCreateDto) error {
	var dataModel models.CollectMetric
	_ = copier.CopyWithOption(&dataModel, data, copier.Option{IgnoreEmpty: true})
	err := s.r.CollectData(ctx, &dataModel)
	if err != nil {
		return err
	}
	return nil
}
