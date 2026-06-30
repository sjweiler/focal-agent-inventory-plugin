package main

import "github.com/shirou/gopsutil/v4/net"

func ListNetwork() (any, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var results []map[string]any

	for _, iface := range interfaces {
		var addresses []string

		for _, addr := range iface.Addrs {
			addresses = append(addresses, addr.Addr)
		}

		results = append(results, map[string]any{
			"name":      iface.Name,
			"mac":       iface.HardwareAddr,
			"addresses": addresses,
			"flags":     iface.Flags,
		})
	}

	return results, nil
}