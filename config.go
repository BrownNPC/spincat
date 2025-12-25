package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Size              int
	Speed             float64
	SpinSpeed         float64
	MousePassthrough  bool
	WindowDecorations bool
	Quiet             bool
}

func DefaultConfig() Config {
	return Config{
		Size:              80,
		Speed:             4.0,
		SpinSpeed:         0.75,
		Quiet:             false,
		MousePassthrough:  true,
		WindowDecorations: false,
	}
}

// Creates a default config on disk or loads an existing one if it exists.
func LoadConfig(path string) Config {
	cfgBytes, err := os.ReadFile(path)
	if err != nil {
		// File does not exist
		cfgBytes, err = json.MarshalIndent(DefaultConfig(),"","")
		if err != nil {
			panic("Failed to marshall default config: " + err.Error())
		}
		if err := os.WriteFile(path, cfgBytes, 0o666); err != nil {
			panic("failed to write default config: " + err.Error())
		}
		return DefaultConfig()
	}
	config := DefaultConfig()
	if err := json.Unmarshal(cfgBytes, &config); err != nil {
		fmt.Println("Failed to read config file:", err)
		fmt.Println("Delete the file or fix it", path)
		os.Exit(1)
	}
	return config
}
