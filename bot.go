package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/trigun117/tproxy-bot/code"
	//"net/http"
	"os"
)

func start() {

	//Create bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}
	/*
		//Set webhook
		_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://"+os.Getenv("IP")+":8443/"+bot.Token, "cert.pem"))
		if err != nil {
			panic(err)
		}

		//Listening for updates
		updates := bot.ListenForWebhook("/" + bot.Token)
		go http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	*/
	config := tgbotapi.NewUpdate(0)
	updates, _ := bot.GetUpdatesChan(config)

	for update := range updates {

		if update.Message != nil && update.Message.Command() == "start" {

			//Collect data to database
			collectData(update.Message.Chat.UserName, update.Message.Chat.ID, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Select Language")

			butt := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Russian.", "SET Russian"), tgbotapi.NewInlineKeyboardButtonData("English.", "SET English"))
			keyb := tgbotapi.NewInlineKeyboardMarkup(butt)
			msg.ReplyMarkup = &keyb

			bot.Send(msg)
		} else if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "SET English":

				//Edit Text
				editText := tgbotapi.NewEditMessageText(int64(int64(update.CallbackQuery.From.ID)), update.CallbackQuery.Message.MessageID, "Hi, I'm a TProxy bot and with the help of me you can connect to a proxy. To get a new proxy send /start, select the language and click on the connection button. You can also go to the site where the proxy is located.")
				bot.Send(editText)

				//Edit buttons
				butt := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Connect a proxy.", code.GetRandomProxy()), tgbotapi.NewInlineKeyboardButtonURL("Go to the site.", "http://telegram-socks.tk/"))
				keyb := tgbotapi.NewInlineKeyboardMarkup(butt)

				//Edit Markup
				editMarkUp := tgbotapi.NewEditMessageReplyMarkup(int64(update.CallbackQuery.From.ID), update.CallbackQuery.Message.MessageID, keyb)
				bot.Send(editMarkUp)

			case "SET Russian":

				//Edit Text
				editText := tgbotapi.NewEditMessageText(int64(update.CallbackQuery.From.ID), update.CallbackQuery.Message.MessageID, "Привет, я TProxy бот и с помощью меня ты можешь подключиться к прокси. Чтобы получить новый прокси отправь /start, выбери язык и жми на кнопку подключения. Так же ты можешь перейти на сайт, где находятся прокси.")
				bot.Send(editText)

				//Edit buttons
				butt := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Подключить прокси.", code.GetRandomProxy()), tgbotapi.NewInlineKeyboardButtonURL("Перейти на сайт.", "http://telegram-socks.tk/"))
				keyb := tgbotapi.NewInlineKeyboardMarkup(butt)

				//Edit Markup
				editMarkUp := tgbotapi.NewEditMessageReplyMarkup(int64(update.CallbackQuery.From.ID), update.CallbackQuery.Message.MessageID, keyb)
				bot.Send(editMarkUp)

			}
		}
	}
}
