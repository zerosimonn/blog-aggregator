package config

import (
	"os"
	"encoding/json"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(homeDir, configFileName)
	return dir, nil
}

func Read() (Config, error)	{
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	} 
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}
	result := Config{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return Config{}, err
	}
	return result, nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(configFilePath, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}