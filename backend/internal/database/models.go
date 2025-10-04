package database

import (
	"fmt"
	"time"
)

type DataSize uint64

const (
	Bit   DataSize = 1
	KBit           = 1024 * Bit
	MBit           = 1024 * KBit
	GBit           = 1024 * MBit
	Byte           = 8 * Bit
	KByte          = 1024 * Byte
	MByte          = 1024 * KByte
	GByte          = 1024 * MByte
)

func (d DataSize) String() string {
	switch {
	case d >= GByte:
		return fmt.Sprintf("%.2f GB", float64(d)/float64(GByte))
	case d >= MByte:
		return fmt.Sprintf("%.2f MB", float64(d)/float64(MByte))
	case d >= KByte:
		return fmt.Sprintf("%.2f KB", float64(d)/float64(KByte))
	case d >= Byte:
		return fmt.Sprintf("%.2f B", float64(d)/float64(Byte))
	default: // bits
		return fmt.Sprintf("%d b", d)
	}
}

// Конфигурация сервера
type Config struct {
	// Порог загрузки (бит/с)
	DownloadThreshold DataSize `json:"downloadThreshold"`

	// Интервал замеров
	PollInterval Time `json:"pollInterval"`

	// Кол-во подряд плохих замеров
	BadCountLimit int32 `json:"badCountLimit"`

	// Расписание временных окон
	Schedules []Schedule `json:"schedules"`
}

// Расписание временных окон, в течение которых можно перезагружать роутер
type Schedule struct {
	// название расписания (удобнее, чем ID)
	Title string `json:"title"`

	// время начала окна (часы)
	StartsAt Time `json:"starts_at"`

	// время конца окна
	EndsAt Time `json:"ends_at"`

	// дни недели
	Weekdays []time.Weekday `json:"weekdays"`

	// состояние (вкл/выкл)
	Enabled bool `json:"enabled"`
}

// Время (для удобного хранения часов/минут)
type Time struct {
	// Часы
	Hours int8 `json:"hours"`

	// Минуты
	Minutes int8 `json:"minutes"`
}
