package services

import (
	"SysAgent/internal/models"
	"SysAgent/internal/utils"
	"context"
)

type MetricsServiceImpl struct {
}

func NewMetricsServiceImpl() *MetricsServiceImpl {
	return &MetricsServiceImpl{}
}

func (m *MetricsServiceImpl) GetMemoryInfo(ctx context.Context) (models.Memory, error) {
	memory, err := utils.GetMemoryInfo()
	if err != nil {
		return models.Memory{}, err
	}
	return memory, nil
}

func (m *MetricsServiceImpl) GetCpuInfo() (models.Cpu, error) {
	cpuInfo, err := utils.GetCpuInfo()
	if err != nil {
		return models.Cpu{}, err
	}
	return cpuInfo, nil
}
func (m *MetricsServiceImpl) GetDiskInfo() (models.Disk, error) {
	diskInfo, err := utils.GetDiskInfo()
	if err != nil {
		return models.Disk{}, err
	}
	return diskInfo, nil
}
