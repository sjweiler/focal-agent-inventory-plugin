package main

import (
	"encoding/json"
	"os"
)

type Request struct {
	Method string         `json:"method"`
	Args   map[string]any `json:"args,omitempty"`
}

type Response struct {
	Success bool `json:"success"`
	Data    any  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func WriteJSON(v any) {
	_ = json.NewEncoder(os.Stdout).Encode(v)
}

func WriteError(err error) {
	WriteJSON(Response{
		Success: false,
		Error:   err.Error(),
	})
}