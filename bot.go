package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/trigun117/tproxy-bot/code"
	"net/http"
	"os"
	"strings"
)

func start() {

	//Create bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	//Set webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://"+os.Getenv("IP")+":8443/"+bot.Token, "cert.pem"))
	if err != nil {
		panic(err)
	}

	//Listening for updates
	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)

	for update := range updates {

		//Collect data to database
		collectData(update.Message.Chat.UserName, update.Message.Chat.ID, update.Message.Text)

		code.GetProxies()

		//Get random proxy and put it to link
		randomNumber := code.Random(0, len(code.P.Proxies))
		randomProxy := strings.Split(code.P.Proxies[randomNumber], ":")
		link := `http://t.me/socks?server=` + randomProxy[0] + `&port=` + randomProxy[1]

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, я TProxy бот и с помощью меня ты можешь подключиться к прокси. Чтобы получить новый прокси просто напиши мне что-либо и жми на кнопку подключения.")

		butt := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Нажми чтобы подключить прокси.", link))
		keyb := tgbotapi.NewInlineKeyboardMarkup(butt)
		msg.ReplyMarkup = &keyb

		msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, "Так же ты можешь перейти на сайт, где находятся прокси.")

		butt1 := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Нажми чтобы перейти на сайт.", "http://telegram-socks.tk/"))
		keyb1 := tgbotapi.NewInlineKeyboardMarkup(butt1)
		msg1.ReplyMarkup = &keyb1

		//Send message
		bot.Send(msg)
		bot.Send(msg1)
	}
}
