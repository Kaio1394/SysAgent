package services

import (
	"SysAgent/internal/models"
	"SysAgent/internal/utils"
	"context"
	"github.com/shirou/gopsutil/v3/mem"
)

type MemoryServiceImpl struct {
}

func NewMemoryServiceImpl() *MemoryServiceImpl {
	return &MemoryServiceImpl{}
}

func (m *MemoryServiceImpl) GetMemoryInfo(ctx context.Context) (models.Memory, error) {
	var memory models.Memory
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return memory, err
	}
	memory.Used = utils.BytesToGB(vmStat.Used)
	memory.Total = utils.BytesToGB(vmStat.Total)
	memory.Free = utils.BytesToGB(vmStat.Free)
	memory.UsedPercent = vmStat.UsedPercent
	return memory, nil
}
