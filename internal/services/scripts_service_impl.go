package services

import (
	"SysAgent/internal/models"
	"SysAgent/internal/repository"
	"context"
)

type ScriptsServiceImpl struct {
	r *repository.ScriptRepositoryImpl
}

func NewScriptsServiceImpl(r *repository.ScriptRepositoryImpl) *ScriptsServiceImpl {
	return &ScriptsServiceImpl{r: r}
}

func (s *ScriptsServiceImpl) GetAllScripts(ctx context.Context) ([]models.Script, error) {
	scripts, err := s.r.GetAllScripts(ctx)
	if err != nil {
		return nil, err
	}
	return scripts, nil
}
