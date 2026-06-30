package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	req, err := readRequest()
	if err != nil {
		WriteError(err)
		return
	}

	data, err := handleRequest(req)
	if err != nil {
		WriteError(err)
		return
	}

	WriteJSON(Response{
		Success: true,
		Data:    data,
	})
}

func readRequest() (Request, error) {
	if len(os.Args) > 1 {
		return Request{
			Method: os.Args[1],
		}, nil
	}

	var req Request
	if err := json.NewDecoder(os.Stdin).Decode(&req); err != nil {
		return Request{}, err
	}

	return req, nil
}

func handleRequest(req Request) (any, error) {
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
		return nil, err
	}

	return data, nil
}
