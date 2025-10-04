package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type Database interface {
	DumpConfig(*Config) error
	LoadConfig() (*Config, error)
}

func NewDatabase(configPath string) Database {
	return &database{
		ConfigPath: configPath,
	}
}

type database struct {
	// Путь до .json файла конфигурации
	ConfigPath string
}

// DumpConfig implements Database.
func (d *database) DumpConfig(c *Config) error {
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	// Сохраняем во временный файл чтоб не повредить основной в случае провала
	tmpFile := d.ConfigPath + ".tmp"
	if err := os.WriteFile(tmpFile, data, 0644); err != nil {
		return fmt.Errorf("write tmp file: %w", err)
	}

	// переименовываем временный файл в основной
	if err := os.Rename(tmpFile, d.ConfigPath); err != nil {
		return fmt.Errorf("rename tmp file: %w", err)
	}
	return nil
}

// LoadConfig implements Database.
func (d *database) LoadConfig() (*Config, error) {
	data, err := os.ReadFile(d.ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &c, nil
}

var _ Database = &database{}
