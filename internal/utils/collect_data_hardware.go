package utils

import (
	"SysAgent/internal/models"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"strconv"
)

func bytesToGB(b uint64) float64 {
	return TruncateFloat(float64(b) / (1024 * 1024 * 1024))
}

func TruncateFloat(value float64) float64 {
	formatted := fmt.Sprintf("%.2f", value)
	valueFinal, _ := strconv.ParseFloat(formatted, 64)
	return valueFinal
}

func GetMemoryInfo() (models.Memory, error) {
	var memory models.Memory
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return memory, err
	}
	memory.Used = bytesToGB(vmStat.Used)
	memory.Total = bytesToGB(vmStat.Total)
	memory.Free = bytesToGB(vmStat.Free)
	memory.UsedPercent = vmStat.UsedPercent
	return memory, nil
}

func GetCpuInfo() (models.Cpu, error) {
	var cpuInfo models.Cpu
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return models.Cpu{}, err
	}
	info, err := cpu.Info()
	if err != nil {
		return models.Cpu{}, err
	}
	cpuInfo.UsagePercent = TruncateFloat(percent[0])
	cpuInfo.Frequency = info[0].Mhz
	cpuInfo.Core = info[0].Cores
	cpuInfo.Model = info[0].ModelName
	return cpuInfo, nil
}

func GetDiskInfo() (models.Disk, error) {
	var diskInfo models.Disk
	diskStat, err := disk.Usage("/")
	if err != nil {
		return models.Disk{}, err
	}

	diskInfo.TotalSpace = TruncateFloat(float64(diskStat.Total) / 1e9)
	diskInfo.UsedSpace = TruncateFloat(float64(diskStat.Used) / 1e9)
	diskInfo.FreeSpace = TruncateFloat(float64(diskStat.Free) / 1e9)

	return diskInfo, nil
}
