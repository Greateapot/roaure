package database

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Database interface {
	// Сохранить конфиг
	DumpConfig(*RoaureConf) error

	// Загрузить конфиг
	LoadConfig() (*RoaureConf, error)

	// Создать конфиг по умолчанию и сохранить
	NewConfig() (*RoaureConf, error)
}

type database struct {
	// Путь до .json файла конфигурации
	ConfigPath string
}

func NewDatabase(configPath string) Database {
	d := database{ConfigPath: configPath}

	return &d
}

// DumpConfig implements Database.
func (d *database) DumpConfig(c *RoaureConf) error {
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
func (d *database) LoadConfig() (*RoaureConf, error) {
	data, err := os.ReadFile(d.ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var c RoaureConf
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &c, nil
}

// NewConfig implements Database.
func (d *database) NewConfig() (*RoaureConf, error) {
	c := &RoaureConf{
		MonitorConf: &MonitorConf{
			DownloadThreshold: 100 * KBit,
			PollInterval:      &Time{Hours: 00, Minutes: 30},
			BadCountLimit:     3,
			Schedules: []*Schedule{
				{
					Title:    "New Schedule",
					StartsAt: &Time{Hours: 10, Minutes: 00},
					EndsAt:   &Time{Hours: 17, Minutes: 00},
					Weekdays: []time.Weekday{
						time.Monday,
						time.Tuesday,
						time.Wednesday,
						time.Thursday,
						time.Friday,
					},
					Enabled: false,
				},
			},
		},
		IperfServerConf: &IperfServerConf{
			Host: "localhost",
			Port: 5201,
		},
		RouterConf: &RouterConf{
			Host:     "192.168.1.1",
			Username: "admin",
			Password: "admin",
		},
	}

	data, err := json.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("marshal config: %w", err)
	}

	// Сохраняем во временный файл чтоб не повредить основной в случае провала
	tmpFile := d.ConfigPath + ".tmp"
	if err := os.WriteFile(tmpFile, data, 0644); err != nil {
		return nil, fmt.Errorf("write tmp file: %w", err)
	}

	// переименовываем временный файл в основной
	if err := os.Rename(tmpFile, d.ConfigPath); err != nil {
		return nil, fmt.Errorf("rename tmp file: %w", err)
	}

	return c, nil
}

var _ Database = &database{}
