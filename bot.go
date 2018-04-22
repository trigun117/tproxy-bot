package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"os"
	"strings"
)

func start() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if err := collectData(update.Message.Chat.UserName, update.Message.Chat.ID, update.Message.Text); err != nil {
			continue
		}

		getProxies()

		randomNumber := random(0, len(p.Proxies))
		randomProxy := strings.Split(p.Proxies[randomNumber], ":")
		link := `http://t.me/socks?server=` + randomProxy[0] + `&port=` + randomProxy[1]

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет я TProxy бот и с помощью меня ты можешь подключится к прокси. Чтобы получить новый прокси просто напиши мне что либо и жми на кнопку подключения.")

		butt := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Нажми чтобы подключить прокси.", link))
		keyb := tgbotapi.NewInlineKeyboardMarkup(butt)
		msg.ReplyMarkup = &keyb

		msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, "Так же ты можешь перейти на сайт где находятся прокси.")

		butt1 := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Нажми чтобы перейти на сайт.", "http://telegram-socks.tk/"))
		keyb1 := tgbotapi.NewInlineKeyboardMarkup(butt1)
		msg1.ReplyMarkup = &keyb1

		bot.Send(msg)
		bot.Send(msg1)
	}
}
