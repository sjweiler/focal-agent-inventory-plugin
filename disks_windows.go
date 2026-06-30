package main

import "github.com/shirou/gopsutil/v4/disk"

func ListDisks() (any, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	var results []map[string]any

	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}

		results = append(results, map[string]any{
			"device":     p.Device,
			"mountpoint": p.Mountpoint,
			"filesystem": p.Fstype,
			"total_gb":   usage.Total / 1024 / 1024 / 1024,
			"free_gb":    usage.Free / 1024 / 1024 / 1024,
			"used_percent": usage.UsedPercent,
		})
	}

	return results, nil
}