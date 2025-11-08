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

func (d DataSize) Float() float64 {
	return float64(d)
}

// Конфигурация сервера
type RoaureConf struct {
	MonitorConf *MonitorConf `json:"monitor_conf"`

	// Информация о iperf3-сервере для замера скорости
	IperfServerConf *IperfServerConf `json:"iperf_server_conf"`

	// Информация о роутере и данные для входа в админ-панель
	RouterConf *RouterConf `json:"router_conf"`
}

// Конфигурация мониторинга
type MonitorConf struct {
	// Порог загрузки (бит/с)
	DownloadThreshold DataSize `json:"download_threshold"`

	// Интервал замеров
	PollInterval *Time `json:"poll_interval"`

	// Кол-во подряд плохих замеров
	BadCountLimit int32 `json:"bad_count_limit"`

	// Расписание временных окон
	Schedules []*Schedule `json:"schedules"`
}

// Хост/порт iperf3-сервера для замера скорости
type IperfServerConf struct {
	// Хост
	Host string `json:"host"`

	// Порт
	Port int32 `json:"port"`
}

// Хост и данные для входа в админ-панель роутера
type RouterConf struct {
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
	StartsAt *Time `json:"starts_at"`

	// Время конца окна
	EndsAt *Time `json:"ends_at"`

	// Дни недели
	Weekdays []time.Weekday `json:"weekdays"`

	// Состояние (вкл/выкл)
	Enabled bool `json:"enabled"`
}

// Время (для удобного хранения часов/минут)
type Time struct {
	// Часы
	Hours int32 `json:"hours"`

	// Минуты
	Minutes int32 `json:"minutes"`
}
