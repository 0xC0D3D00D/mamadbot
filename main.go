package main

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func main() {
	bot_key := os.Getenv("TELEGRAM_BOT_API_KEY")
	if bot_key == "" {
		log.Panic("Set TELEGRAM_BOT_API_KEY in your environment")
	}
	bot, err := tgbotapi.NewBotAPI(bot_key)
	if err != nil {
		log.Panic(err)
	}

	//	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		lg, _ := json.Marshal(update)
		log.Printf("%s\n", lg)
		var userID int
		var chatID int64
		if update.Message != nil {
			userID = update.Message.From.ID
			chatID = update.Message.Chat.ID
		} else if update.CallbackQuery != nil {
			userID = update.CallbackQuery.From.ID
			chatID = update.CallbackQuery.Message.Chat.ID

			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
		} else {
			continue
		}

		_ = userID

		/*
			if HasBadwords(update.Message.Text) {
				msg := tgbotapi.NewMessage(chatID, MessageBadWords)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		*/

		if reply := PrepareReply(update.Message.Text); reply != "" {
			msg := tgbotapi.NewMessage(chatID, reply)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}

		if faal := GetFaal(update.Message.Text); faal != "" {
			msg := tgbotapi.NewMessage(chatID, faal)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
