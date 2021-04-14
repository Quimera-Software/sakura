// Copyright (c) 2020. Quimera Software S.p.A.

package sakura

import (
	"errors"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI = nil

func getTelegramBot() (error, *tgbotapi.BotAPI) {
	if bot == nil {
		token := cfg.Telegram.Token
		if token == "" {
			return errors.New("no telegram token set"), nil
		}

		var err error
		bot, err = tgbotapi.NewBotAPI(token)
		if err != nil {
			return err, nil
		}

		if bot == nil {
			return errors.New("unable to start telegram bot"), nil
		}
	}

	return nil, bot
}

func telegramBroadcast(message string) error {
	err, bot := getTelegramBot()
	if err != nil {
		return err
	}

	for _, chatId := range cfg.Telegram.Channels {
		msg := tgbotapi.NewMessage(chatId, message)

		_, err := bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
