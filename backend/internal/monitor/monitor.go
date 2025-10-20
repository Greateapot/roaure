package monitor

import (
	"context"
	"log"
	"slices"
	"time"

	"github.com/Greateapot/roaure/internal/database"
	"github.com/Greateapot/roaure/internal/router"
	"github.com/Greateapot/roaure/internal/speedtest"
)

type Monitor struct {
	DownloadThreshold database.DataSize
	PollInterval      database.Time
	BadCountLimit     int32
	Schedules         []database.Schedule

	RouterClient    *router.Client
	SpeedtestClient *speedtest.Client

	DownloadSpeed  database.DataSize
	BadCount       int32
	RebootRequired bool
	Running        bool

	cancel context.CancelFunc
	ctx    context.Context
}

func NewMonitor(
	ctx context.Context,
	downloadThreshold database.DataSize,
	pollInterval database.Time,
	badCountLimit int32,
	schedules []database.Schedule,
	routerClient *router.Client,
	speedtestClient *speedtest.Client,
) *Monitor {
	m := Monitor{
		DownloadThreshold: downloadThreshold,
		PollInterval:      pollInterval,
		BadCountLimit:     badCountLimit,
		Schedules:         schedules,
		RouterClient:      routerClient,
		SpeedtestClient:   speedtestClient,
		ctx:               ctx,
	}

	return &m
}

func (m *Monitor) canRebootNow() bool {
	now := time.Now()
	weekday := now.Weekday()
	hour := int8(now.Hour())
	minute := int8(now.Minute())

	// Индекс первого расписания, удовлетворяющего условиям:
	// включен, сегодня, после времени начала и до времени окончания
	scheduleIndex := slices.IndexFunc(m.Schedules, func(schedule database.Schedule) bool {
		return schedule.Enabled &&
			slices.Contains(schedule.Weekdays, weekday) &&
			hour >= schedule.StartsAt.Hours &&
			minute >= schedule.StartsAt.Minutes &&
			hour <= schedule.EndsAt.Hours &&
			minute < schedule.EndsAt.Minutes
	})

	// Если расписание найдено - перезагружать роутер можно
	return scheduleIndex != -1
}

func (m *Monitor) canReboot() bool {
	// Замеряем скорость
	if bps, err := m.SpeedtestClient.Start(); err != nil {
		// Не удалось замерить скорость загрузки, выход
		log.Println(err)
		return false
	} else {
		m.DownloadSpeed = database.DataSize(bps)
	}

	// Скорость больше порога, все нормально
	if m.DownloadSpeed > m.DownloadThreshold {
		m.RebootRequired = false
		m.BadCount = 0
		return false
	}

	// Еще один плохой результат
	m.BadCount++

	// Недостаточно плохих замеров, чтобы действовать
	if m.BadCount < m.BadCountLimit {
		return false
	}

	// Сейчас перезагружаться нельзя
	if !m.canRebootNow() {
		m.RebootRequired = true
		return false
	}

	// Все условия выполнены, можно перезагрузить роутер
	return true
}

func (m *Monitor) reboot() {
	if err := m.RouterClient.Reboot(); err != nil {
		// Не удалось перезагрузить роутер
		log.Println(err)
	} else {
		// Успех, обнуляемся и ждем загрузки роутера
		m.RebootRequired = false
		m.BadCount = 0
		m.Stop()
		m.DelayedStart(3 * time.Minute)
	}
}

func (m *Monitor) wrapper(ctx context.Context) {
	ticker := time.NewTicker(
		0 +
			time.Duration(m.PollInterval.Hours)*time.Hour +
			time.Duration(m.PollInterval.Minutes)*time.Minute,
	)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Замеряем скорость, если нужно и можно - перезагружаем роутер
			if m.canReboot() {
				m.reboot()
			}
		}
	}
}

func (m *Monitor) Start() {
	ctx, cancel := context.WithCancel(m.ctx)
	m.cancel = cancel
	m.Running = true
	go m.wrapper(ctx)
}

func (m *Monitor) Stop() {
	if m.cancel != nil {
		m.Running = false
		m.cancel()
	}
}

func (m *Monitor) DelayedStart(delay time.Duration) {
	m.Running = true
	go func() {
		<-time.After(delay)
		m.Start()
	}()
}
