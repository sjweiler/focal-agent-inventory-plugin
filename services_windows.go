package main

import (
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

func ListServices() (any, error) {
	m, err := mgr.Connect()
	if err != nil {
		return nil, err
	}
	defer m.Disconnect()

	names, err := m.ListServices()
	if err != nil {
		return nil, err
	}

	var results []map[string]any

	for _, name := range names {
		s, err := m.OpenService(name)
		if err != nil {
			continue
		}

		status, err := s.Query()
		s.Close()

		if err != nil {
			continue
		}

		results = append(results, map[string]any{
			"name":  name,
			"state": serviceStateString(status.State),
		})
	}

	return results, nil
}

func serviceStateString(state svc.State) string {
	switch state {
	case svc.Stopped:
		return "stopped"
	case svc.StartPending:
		return "start_pending"
	case svc.StopPending:
		return "stop_pending"
	case svc.Running:
		return "running"
	case svc.ContinuePending:
		return "continue_pending"
	case svc.PausePending:
		return "pause_pending"
	case svc.Paused:
		return "paused"
	default:
		return "unknown"
	}
}
