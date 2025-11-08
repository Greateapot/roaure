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
	MonitorConf *database.MonitorConf

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
	monitorConf *database.MonitorConf,
	routerClient *router.Client,
	speedtestClient *speedtest.Client,
) *Monitor {
	m := Monitor{
		MonitorConf:     monitorConf,
		RouterClient:    routerClient,
		SpeedtestClient: speedtestClient,
		ctx:             ctx,
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
	scheduleIndex := slices.IndexFunc(m.MonitorConf.Schedules, func(schedule *database.Schedule) bool {
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
	if m.DownloadSpeed > m.MonitorConf.DownloadThreshold {
		m.RebootRequired = false
		m.BadCount = 0
		return false
	}

	// Еще один плохой результат
	m.BadCount++

	// Недостаточно плохих замеров, чтобы действовать
	if m.BadCount < m.MonitorConf.BadCountLimit {
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

func (m *Monitor) Reboot() error {
	m.Stop()

	if err := m.RouterClient.Reboot(); err != nil {
		// Не удалось перезагрузить роутер
		m.RebootRequired = true
		m.Start()
		return err
	}

	// Успех, обнуляемся и ждем загрузки роутера
	m.RebootRequired = false
	m.BadCount = 0
	m.DelayedStart(3 * time.Minute)
	return nil
}

func (m *Monitor) wrapper(ctx context.Context) {
	ticker := time.NewTicker(
		0 +
			time.Duration(m.MonitorConf.PollInterval.Hours)*time.Hour +
			time.Duration(m.MonitorConf.PollInterval.Minutes)*time.Minute,
	)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Замеряем скорость, если нужно и можно - перезагружаем роутер
			if m.canReboot() {
				m.Reboot()
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
