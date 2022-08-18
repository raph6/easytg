package easytg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	BotAPI  string
	ConvID  int64
	Message string
}

func NewBot(api string) Bot {
	return Bot{
		BotAPI: api,
	}
}

func (b Bot) Msg(message string, convId int64) Bot {
	b.Message = message
	b.ConvID = convId
	return b
}

func (b Bot) Send() error {
	bot, err := tgbotapi.NewBotAPI(b.BotAPI)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(b.ConvID, b.Message)
	bot.Send(msg)

	return nil
}
