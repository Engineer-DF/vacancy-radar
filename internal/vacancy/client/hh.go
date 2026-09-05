package client

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// Возможно стоит вынести функционал браузера в отдельный слой, а в этом файле вызывать уже готовые методы.

type Scraper struct {
	logger slog.Logger
}

func New(logger slog.Logger) *Scraper {
	return &Scraper{
		logger: *logger.With("component", "scraper"), // логи теперь будут сопровождаться подписью с названием компонента
	}
}

func (s *Scraper) ScrapeData(ctx context.Context) error {
	// функция Launch() возвращает URL-адрес WebSocket (DevTools Protocol URL)

	url, err := launcher.New().Headless(true).Context(ctx).Launch()
	if err != nil {
		return fmt.Errorf("failed to launch browser: %w", err)
	}

	browser := rod.New().ControlURL(url)
	if err := browser.Connect(); err != nil {
		return fmt.Errorf("failed to connect browser via WebSocket: %w", err)
	}
	defer browser.Close() // fix 

	return nil
}
