package main

import "github.com/shirou/gopsutil/v4/process"

func ListProcesses() (any, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var results []map[string]any

	for _, p := range processes {
		name, _ := p.Name()
		username, _ := p.Username()
		cpuPercent, _ := p.CPUPercent()
		memPercent, _ := p.MemoryPercent()

		results = append(results, map[string]any{
			"pid":           p.Pid,
			"name":          name,
			"user":          username,
			"cpu_percent":   cpuPercent,
			"memory_percent": memPercent,
		})
	}

	return results, nil
}