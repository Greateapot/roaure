package led

import (
	"sync"

	"github.com/warthog618/go-gpiocdev"
)

type LED struct {
	line    *gpiocdev.Line
	Enabled bool
	mu      sync.Mutex
}

func NewLED(chip string, lineOffset int) (*LED, error) {
	line, err := gpiocdev.RequestLine(
		chip,
		lineOffset,
		gpiocdev.AsOutput(0),
	)
	if err != nil {
		return nil, err
	}

	return &LED{
		line:    line,
		Enabled: false,
	}, nil
}

func (l *LED) On() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.Enabled {
		return nil
	}

	if err := l.line.SetValue(1); err != nil {
		return err
	}

	l.Enabled = true
	return nil
}

func (l *LED) Off() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if !l.Enabled {
		return nil
	}

	if err := l.line.SetValue(0); err != nil {
		return err
	}

	l.Enabled = false
	return nil
}

func (l *LED) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Enabled = false
	return l.line.Close()
}
