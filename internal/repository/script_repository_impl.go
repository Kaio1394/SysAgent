package repository

import (
	"SysAgent/internal/models"
	"context"
	"gorm.io/gorm"
)

type ScriptRepositoryImpl struct {
	db *gorm.DB
}

func NewScriptRepositoryImpl(db *gorm.DB) *ScriptRepositoryImpl {
	return &ScriptRepositoryImpl{db: db}
}

func (r *ScriptRepositoryImpl) CreateNewScript(ctx context.Context, script *models.Script) error {
	return r.db.WithContext(ctx).Create(script).Error
}

func (r *ScriptRepositoryImpl) GetAllScripts(ctx context.Context) ([]models.Script, error) {
	var scripts []models.Script
	err := r.db.WithContext(ctx).Find(&scripts)
	if err != nil {
		return scripts, err.Error
	}
	return scripts, nil
}

func (r *ScriptRepositoryImpl) GetScriptByUuid(ctx context.Context, uuid string) (models.Script, error) {
	var script models.Script
	err := r.db.WithContext(ctx).First(&script, "uuid = ?", uuid).Error
	if err != nil {
		return script, err
	}
	return script, nil
}
