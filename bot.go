package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
	"os"
	"strings"
)

func start() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://"+os.Getenv("IP")+":8443/"+bot.Token, "cert.pem"))
	if err != nil {
		panic(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)

	for update := range updates {

		collectData(update.Message.Chat.UserName, update.Message.Chat.ID, update.Message.Text)

		getProxies()

		randomNumber := random(0, len(p.Proxies))
		randomProxy := strings.Split(p.Proxies[randomNumber], ":")
		link := `http://t.me/socks?server=` + randomProxy[0] + `&port=` + randomProxy[1]

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, я TProxy бот и с помощью меня ты можешь подключиться к прокси. Чтобы получить новый прокси просто напиши мне что-либо и жми на кнопку подключения.")

		butt := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Нажми чтобы подключить прокси.", link))
		keyb := tgbotapi.NewInlineKeyboardMarkup(butt)
		msg.ReplyMarkup = &keyb

		msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, "Так же ты можешь перейти на сайт, где находятся прокси.")

		butt1 := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Нажми чтобы перейти на сайт.", "http://telegram-socks.tk/"))
		keyb1 := tgbotapi.NewInlineKeyboardMarkup(butt1)
		msg1.ReplyMarkup = &keyb1

		bot.Send(msg)
		bot.Send(msg1)
	}
}
