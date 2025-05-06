package services

import (
	"SysAgent/internal/models"
	"SysAgent/internal/models/dto"
	"SysAgent/internal/repository"
	"SysAgent/internal/utils"
	"context"
	"errors"
	"time"
)

const (
	LAYOUT string = "2006-01-02"
)

type ScriptsServiceImpl struct {
	r *repository.ScriptRepositoryImpl
}

func NewScriptServiceImpl(r *repository.ScriptRepositoryImpl) *ScriptsServiceImpl {
	return &ScriptsServiceImpl{r: r}
}

func (s *ScriptsServiceImpl) GetAllScripts(ctx context.Context) ([]models.Script, error) {
	scripts, err := s.r.GetAllScripts(ctx)
	if err != nil {
		return nil, err
	}
	return scripts, nil
}

func (s *ScriptsServiceImpl) CreateNewScript(ctx context.Context, script *dto.ScriptCreateDto) error {
	start, err := time.Parse(LAYOUT, script.StartDate)
	if err != nil {
		return err
	}
	end, err := time.Parse(LAYOUT, script.EndDate)
	if err != nil {
		return err
	}
	var scriptModel models.Script
	scriptModel.Activated = script.Activated
	scriptModel.Command = script.Command
	scriptModel.StartDate = &start
	scriptModel.EndDate = &end
	scriptModel.ExecuteHour = script.ExecuteHour
	scriptModel.Weekdays = script.Weekdays
	err = s.r.CreateNewScript(ctx, &scriptModel)
	if err != nil {
		return err
	}
	return nil
}

func (s *ScriptsServiceImpl) GetScriptByUuid(ctx context.Context, uuid string) (models.Script, error) {
	if uuid == "" {
		return models.Script{}, errors.New("uuid is empty")
	}
	script, err := s.r.GetScriptByUuid(ctx, uuid)
	if err != nil {
		return models.Script{}, err
	}
	return script, nil
}

func (s *ScriptsServiceImpl) ExecuteScript(ctx context.Context, uuid string) (string, error) {
	if uuid == "" {
		return "", errors.New("uuid is empty")
	}
	script, err := s.r.GetScriptByUuid(ctx, uuid)
	if err != nil {
		return "", err
	}
	result, err := utils.ExecuteScript(script.Command)
	if err != nil {
		return "", err
	}
	return result, nil
}
