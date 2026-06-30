package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var req Request

	if err := json.NewDecoder(os.Stdin).Decode(&req); err != nil {
		WriteError(err)
		return
	}

	var data any
	var err error

	switch req.Method {
	case "inventory.collect":
		data, err = CollectInventory()
	case "disks.list":
		data, err = ListDisks()
	case "network.list":
		data, err = ListNetwork()
	case "processes.list":
		data, err = ListProcesses()
	case "services.list":
		data, err = ListServices()
	default:
		err = fmt.Errorf("unknown method: %s", req.Method)
	}

	if err != nil {
		WriteError(err)
		return
	}

	WriteJSON(Response{
		Success: true,
		Data:    data,
	})
}