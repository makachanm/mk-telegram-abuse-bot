package telegramservice

import (
	"context"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramService struct {
	MesssageChan <-chan string
	TgAPI        *tgbotapi.BotAPI
	TargetChan   int64
}

func NewTelegramService(TgInfo Telegram, MesgChan chan string) TelegramService {
	tgAPI, err := tgbotapi.NewBotAPI(TgInfo.TgToken)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return TelegramService{TgAPI: tgAPI, MesssageChan: MesgChan, TargetChan: TgInfo.TargetChannel}
}

func (tc *TelegramService) Execute(_ context.Context) {
	for ctx := range tc.MesssageChan {
		tCtx := tgbotapi.NewMessage(tc.TargetChan, ctx)
		tCtx.ParseMode = "Markdown"

		fmt.Println("Sentting Message: ", ctx)

		_, err := tc.TgAPI.Send(tCtx)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Message Sented.")
		}

		time.Sleep(time.Second * 30)
	}
}
