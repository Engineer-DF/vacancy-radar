package telegram

import (
	"log"

	"github.com/Engineer-DF/vacancy-radar/internal/vacancy/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(token string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = true

	log.Printf("Logged in to account: %v", &bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			command := update.Message.Command()

			var replyText string

			switch command {
			case "start":
				replyText = "Привет! Я бот, предназначенный для удобного поиска вакансий по заданным фильтрам! Чтобы получить инструкцию " +
					"введите команду /help"
			case "help":
				replyText = "Список доступных команд:\n/start\n/test\n\nа большего и не жди, сынку"
			case "test":
				go func() {
					err := client.PrototypeRodRequest()
					if err != nil {
						log.Printf("Error request to API hh.ru: %v", err)
						return
					}
				}()
				replyText = "Смотри что тебе в терминале накидало, а не сюда."
			default:
				replyText = "Не знаю такого. Воспользуйтесь /help чтобы вывести список доступных команд."
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)

			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				log.Printf("Ошибка отправки сообщения: %v", err)
			}
		}

	}
	return nil
}
