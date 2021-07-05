package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// Config has the configuration keys required for the application setup ans running
type Config struct {
	AppURL  string `json:"app_url"`
	AppEnv  string `json:"app_env"`
	AppPort uint   `json:"app_port"`
}

// GetAbsolutePath returns the project absolute path
func GetAbsolutePath() string {
	_, b, _, _ := runtime.Caller(0)

	return fmt.Sprintf("%s/", filepath.Join(filepath.Dir(b), "../.."))
}

// Read attempts to open and read the contents of the config file provided
func Read(filename string) (*Config, error) {
	// Attempt to open the provided file
	file, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("config.Read.OpenConfigFile:: %v", err)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	config := &Config{}

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("config.Read.DecodeJSONFile:: %v", err)
	}

	return config, nil
}
