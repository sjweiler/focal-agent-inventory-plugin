package main

import (
	"os"
	"runtime"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func CollectInventory() (any, error) {
	hostname, _ := os.Hostname()

	hostInfo, err := host.Info()
	if err != nil {
		return nil, err
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	cpuName := ""
	if len(cpuInfo) > 0 {
		cpuName = cpuInfo[0].ModelName
	}

	return map[string]any{
		"hostname": hostname,
		"os":       hostInfo.Platform,
		"os_version": hostInfo.PlatformVersion,
		"kernel":   hostInfo.KernelVersion,
		"arch":     runtime.GOARCH,
		"cpu_name": cpuName,
		"cpu_count": runtime.NumCPU(),
		"memory_total_gb": memInfo.Total / 1024 / 1024 / 1024,
		"memory_free_gb":  memInfo.Available / 1024 / 1024 / 1024,
	}, nil
}