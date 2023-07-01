package core

import (
	"abusebot/core/db"
	dbmodels "abusebot/core/db/models"
	botservices "abusebot/core/services/bot"
	telegramservice "abusebot/core/services/telegram"
	"abusebot/utils"
	"context"
	"fmt"
	"time"

	"github.com/reugn/go-quartz/quartz"
)

type MainService struct {
	MainScheduler  quartz.Scheduler
	DBCore         *db.DatabaseCore
	MessageChannel chan string
	Config         utils.Configure
}

func NewService(dbc *db.DatabaseCore, config utils.Configure) MainService {
	return MainService{
		MainScheduler:  quartz.NewStdScheduler(),
		DBCore:         dbc,
		MessageChannel: make(chan string, 512),
		Config:         config,
	}
}

func (c *MainService) InitMainService() {
	ctx := context.Background()

	botserv := c.makeBotService()
	tgserv := c.makeTelegramService()

	c.MainScheduler.Start(ctx)

	c.MainScheduler.ScheduleJob(ctx, botserv, quartz.NewSimpleTrigger(time.Second*30))
	fmt.Println("BotService Attached.")

	go tgserv.Execute(ctx)
	fmt.Println("TelegramService Attached.")

	c.MainScheduler.Wait(ctx)
}

func (c *MainService) makeBotService() *botservices.BotService {
	AbuM := dbmodels.NewAbuseModel(c.DBCore)

	InsData := botservices.Misskey{
		MisskeyToken: c.Config.Token.MisskeyToken,
		InstanceURL:  c.Config.InsURL,
	}

	botService := botservices.NewBotService(InsData, &AbuM, c.MessageChannel)

	return &botService
}

func (c *MainService) makeTelegramService() *telegramservice.TelegramService {
	TgData := telegramservice.Telegram{
		TgToken:       c.Config.Token.TelegramToken,
		TargetChannel: c.Config.TargetChan,
	}

	tgService := telegramservice.NewTelegramService(TgData, c.MessageChannel)
	return &tgService
}
