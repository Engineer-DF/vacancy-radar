package main

import (
	"log"
	"log/slog"
	"os"

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
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var cfg Config
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Configuration initialization error: %v", err)
	}

	logger.Info("Configuration loaded successfully.")

	telegram.StartBot(cfg.TelegramToken)

	// TODO: добавить дальнейший процесс запуска
	// Убрать говнокод
}
