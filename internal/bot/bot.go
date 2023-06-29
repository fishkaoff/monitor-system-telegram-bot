package bot

import (
	"os"

	"github.com/fishkaoff/telegram-client/internal/transport"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Middlwares interface {
	CheckUrl(URL string) bool
	CheckMatches(webSites []string, site string) bool
}

type Loggerer interface {
	LogMessage(message string)
	LogError(err string)
	LogErrorAndQuit(err string)
}

type Api interface {
	SendUrls(request []byte) []byte 
}

type StorageMicroservice interface {
	Save(chatID int64, site string) string
	Delete(chatID int64, site string) string
	Get(chatID int64) []string
}

type Bot struct {
	// for work with telegram
	bot *tgbotapi.BotAPI

	// logging
	sugar Loggerer

	// middlewares
	mw Middlwares

	// api
	checkUrlsMicro transport.Transport

	// storage
	storageMicro StorageMicroservice

	// app req
	userStatus map[int64]int
}

func NewBot(bot *tgbotapi.BotAPI, sugar Loggerer,mw Middlwares,checkUrlsMicro Api, storageMicro StorageMicroservice) *Bot {
	return &Bot{bot: bot, sugar: sugar,mw: mw, checkUrlsMicro: checkUrlsMicro, storageMicro: storageMicro, userStatus: make(map[int64]int, 0)}
}

func (b *Bot) Start() {
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := b.bot.GetUpdatesChan(updateConfig)

	b.sugar.LogMessage("Bot Started")
	b.handleUpdates(updates)

}

func (b *Bot) Stop() {
	os.Exit(1)
}


