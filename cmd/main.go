package main

import (
	"context"
	"log"
	"os"

	Bot "github.com/fishkaoff/telegram-client/internal/bot"
	"github.com/fishkaoff/telegram-client/internal/middlewares"
	"github.com/fishkaoff/telegram-client/internal/transport/grpc"
	transport "github.com/fishkaoff/telegram-client/internal/transport/rest"
	"github.com/fishkaoff/telegram-client/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("test.env")
	if err != nil {
		log.Fatal("Cannot load env file")
	}
}

func main() {
	TGTOKEN := os.Getenv("TGTOKEN")
	CHECKSERVERURL := os.Getenv("CHECKSERVERURL")
	DBADDR := os.Getenv("DBMICRO")

	// logger
	logger := logger.Logger{}
	log := logger.Start()
	defer log.Sync() // flushes buffer, if any
	logger.LogMessage("Logger Started")


	// init depends
	middlewares := middlewares.NewMiddlewares()
	api := transport.NewApi(CHECKSERVERURL)


	grpcClient, err := grpc.NewGRPCClient(DBADDR)
	if err != nil {
		log.Fatal(err.Error())
	}
	grpcStruct := grpc.NewGRPCStruct(grpcClient, context.Background())


	// start bot
	bot, err := tgbotapi.NewBotAPI(TGTOKEN)
	if err != nil {
		logger.LogErrorAndQuit(err.Error())
	}

	tg := Bot.NewBot(bot, &logger,middlewares, api, grpcStruct)
	tg.Start()

}
