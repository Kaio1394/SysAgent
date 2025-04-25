package repository

import (
	"context"
	"gorm.io/gorm"
)

type MetricRepositoryImpl struct {
	Db *gorm.DB
}

func NewMetricRepositoryImpl(db *gorm.DB) *MetricRepositoryImpl {
	return &MetricRepositoryImpl{Db: db}
}

func (m *MetricRepositoryImpl) InsertDataCollection(ctx context.Context) {

}
