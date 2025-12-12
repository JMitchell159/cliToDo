package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	MakeUrgent int `json:"make_urgent"`
	RemoveFromComplete int `json:"remove_from_complete"`
}

func Load() (*Config, error) {
	jsonConfig, err := os.ReadFile("config.json")
	if os.IsNotExist(err) {
		fmt.Println("Config file does not exist. Creating config.json with default values.")
		defaults := Config{
			MakeUrgent: 24,
			RemoveFromComplete: 168,
		}
		defaultFile, err := json.Marshal(defaults)
		if err != nil {
			return nil, fmt.Errorf("something went wrong while marshalling Config struct to json: %v", err)
		}
		err = os.WriteFile("config.json", defaultFile, 0660)
		if err != nil {
			return nil, fmt.Errorf("something went wrong while writing default config to file: %v", err)
		}
		return &defaults, nil
	}

	result := Config{}
	err = json.Unmarshal(jsonConfig, &result)
	if err != nil {
		return nil, fmt.Errorf("something went wrong while unmarshalling config file to struct: %v", err)
	}

	return &result, nil
}

