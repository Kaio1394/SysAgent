package services

import (
	"SysAgent/internal/models"
	"SysAgent/internal/models/dto"
	"SysAgent/internal/repository"
	"context"
	"github.com/jinzhu/copier"
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
	var scriptModel models.Script
	_ = copier.CopyWithOption(&scriptModel, script, copier.Option{IgnoreEmpty: true})
	err := s.r.CreateNewScript(ctx, &scriptModel)
	if err != nil {
		return err
	}
	return nil
}
