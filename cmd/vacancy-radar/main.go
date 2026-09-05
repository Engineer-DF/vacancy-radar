package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Engineer-DF/vacancy-radar/internal/telegram"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken       string `env:"TELEGRAM_BOT_TOKEN,required"`
	DBURI               string `env:"DB_URI"`
	RabbitURL           string `env:"RABBIT_URL"`
	AdminChannelChatID  int64  `env:"ADMIN_CHANNEL_CHAT_ID"`
	PublicChannelChatID int64  `env:"PUBLIC_CHANNEL_CHAT_ID"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger.Info("Starting Vacancy Radar application...")

	if err := run(ctx, *logger); err != nil {
		logger.Error("fatal startup error", slog.Any("err", err))
		os.Exit(1)
	}
}

func run(ctx context.Context, logger slog.Logger) error {
	// !os.IsNotExist(err) проверяет что ошибка не вызвана отсутствием .env файла

	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env file")
	}

	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Configuration initialization error: %v", err)
	}
	logger.Info("Configuration loaded successfully.")

	err := telegram.StartBot(ctx, logger, cfg.TelegramToken)
	if err != nil {
		logger.Error("Starting bot failed: %v", slog.Any("err", err))
		return err
	}
	return nil
}
