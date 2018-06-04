package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/trigun117/tproxy-bot/code"
	"os"
	"runtime"
)

func createMarkup(btn, btn1, btn2, btn3 string) tgbotapi.InlineKeyboardMarkup {
	row := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL(btn, btn1), tgbotapi.NewInlineKeyboardButtonURL(btn2, btn3))
	return tgbotapi.NewInlineKeyboardMarkup(row)
}

func startBot() {

	// Create bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	config := tgbotapi.NewUpdate(0)
	updates, _ := bot.GetUpdatesChan(config)

	for update := range updates {

		defer runtime.GC()

		if update.Message != nil && update.Message.Command() == "start" {

			// Collect data to database
			collectData(update.Message.Chat.UserName, update.Message.Chat.ID, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Select Language")

			keyb := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Russian.", "SET Russian"), tgbotapi.NewInlineKeyboardButtonData("English.", "SET English")))
			msg.ReplyMarkup = &keyb

			bot.Send(msg)

		} else if update.CallbackQuery != nil {

			switch update.CallbackQuery.Data {
			case "SET English":

				// Edit Text
				editText := tgbotapi.NewEditMessageText(int64(update.CallbackQuery.From.ID), update.CallbackQuery.Message.MessageID, "Hi, I'm a TProxy bot and with the help of me you can connect to a proxy. To get a new proxy send /start, select the language and click on the connection button. You can also go to the site where the proxy is located.")
				bot.Send(editText)

				// Edit Markup
				editMarkUp := tgbotapi.NewEditMessageReplyMarkup(int64(update.CallbackQuery.From.ID), update.CallbackQuery.Message.MessageID, createMarkup("Connect a proxy.", code.GetRandomProxy(), "Go to the site.", os.Getenv("LINK")))
				bot.Send(editMarkUp)

			case "SET Russian":

				// Edit Text
				editText := tgbotapi.NewEditMessageText(int64(update.CallbackQuery.From.ID), update.CallbackQuery.Message.MessageID, "Привет, я TProxy бот и с помощью меня ты можешь подключиться к прокси. Чтобы получить новый прокси отправь /start, выбери язык и жми на кнопку подключения. Так же ты можешь перейти на сайт, где находятся прокси.")
				bot.Send(editText)

				// Edit Markup
				editMarkUp := tgbotapi.NewEditMessageReplyMarkup(int64(update.CallbackQuery.From.ID), update.CallbackQuery.Message.MessageID, createMarkup("Подключить прокси.", code.GetRandomProxy(), "Перейти на сайт.", os.Getenv("LINK")))
				bot.Send(editMarkUp)

			}
		}
	}
}
