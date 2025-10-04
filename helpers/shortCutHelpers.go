package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Harikrishnan-Ashok/samBaar/state"
)

// getConfigFilePath returns the full path to the config.json file
func getConfigFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("can't get user config directory: %w", err)
	}

	appConfigDir := filepath.Join(configDir, "samBaar")
	if err := os.MkdirAll(appConfigDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("can't create config directory: %w", err)
	}

	return filepath.Join(appConfigDir, "config.json"), nil
}

// configFileExists checks whether config.json exists and is non-empty
func configFileExists() bool {
	filePath, err := getConfigFilePath()
	if err != nil {
		log.Println("Error getting config file path:", err)
		return false
	}

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) || info.Size() == 0 {
		return false
	} else if err != nil {
		log.Println("Error checking config file:", err)
		return false
	}
	return true
}

// CreateDefaultConfigFile writes an empty default config to config.json
func CreateDefaultConfigFile() {
	filePath, err := getConfigFilePath()
	if err != nil {
		log.Println("Error getting config file path:", err)
		return
	}

	data, err := json.MarshalIndent(state.DefaultConfig, "", "  ")
	if err != nil {
		log.Println("Failed to marshal default config:", err)
		return
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		log.Println("Failed to write config file:", err)
		return
	}

	fmt.Println("Default config written to:", filePath)
}

// ReadConfigFile reads and prints the contents of config.json
func ReadConfigFile() {
	filePath, err := getConfigFilePath()
	if err != nil {
		log.Println("Error getting config file path:", err)
		return
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Error reading config file:", err)
		return
	}

	fmt.Println("Config file contents:\n", string(data))
}

// loadConfig unmarshals config.json into a state.AppConfig struct
func loadConfig() (state.AppConfig, error) {
	var config state.AppConfig

	filePath, err := getConfigFilePath()
	if err != nil {
		return config, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("error parsing config JSON: %w", err)
	}

	return config, nil
}

func GetCommandForKeyWord(keyword string) (string, error) {
	config, err := loadConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	for _, shortcut := range config.Config.Shortcuts {
		if shortcut.Token == keyword {
			return shortcut.Command, nil
		}
	}

	return "", fmt.Errorf("no command found for keyword: %s", keyword)
}

// SearchForCommandKeyword ensures config exists, then searches for the command
func SearchForCommandKeyword(key string) (string, error) {
	if !configFileExists() {
		fmt.Println("Creating new config file...")
		CreateDefaultConfigFile()
	}

	fmt.Println("Reading config file...")
	ReadConfigFile()

	cmd, err := GetCommandForKeyWord(key)
	return cmd, err
}
