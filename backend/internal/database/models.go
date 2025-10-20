package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type DataSize float64

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
	case d >= GBit:
		return fmt.Sprintf("%.2f Gb", d/GBit)
	case d >= MBit:
		return fmt.Sprintf("%.2f Mb", d/MBit)
	case d >= KBit:
		return fmt.Sprintf("%.2f Kb", d/KBit)
	default: // bits
		return fmt.Sprintf("%.2f b", d/Bit)
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

	// Информация о iperf3-сервере для замера скорости
	Server Server `json:"server"`

	// Информация о роутере и данные для входа в админ-панель
	Router Router `json:"router"`

	// Расписание временных окон
	Schedules []Schedule `json:"schedules"`
}

// Хост/порт iperf3-сервера для замера скорости
type Server struct {
	// Хост
	Host string `json:"host"`

	// Порт
	Port int `json:"port"`
}

// Хост и данные для входа в админ-панель роутера
type Router struct {
	// Хост
	Host string `json:"host"`

	// Имя пользователя
	Username string `json:"username"`

	// Пароль пользователя
	Password string `json:"password"`
}

// Расписание временных окон, в течение которых можно перезагружать роутер
type Schedule struct {
	// ID расписания (ключ)
	ID uuid.UUID

	// Название расписания
	Title string `json:"title"`

	// Время начала окна (часы)
	StartsAt Time `json:"starts_at"`

	// Время конца окна
	EndsAt Time `json:"ends_at"`

	// Дни недели
	Weekdays []time.Weekday `json:"weekdays"`

	// Состояние (вкл/выкл)
	Enabled bool `json:"enabled"`
}

// Время (для удобного хранения часов/минут)
type Time struct {
	// Часы
	Hours int8 `json:"hours"`

	// Минуты
	Minutes int8 `json:"minutes"`
}
