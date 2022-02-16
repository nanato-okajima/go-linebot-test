package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"

	"go-linebot-test/weather"
)

func main() {
	// LINE bot クライアントの作成
	_ = godotenv.Load()
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	result, err := weather.GetWeather()
	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage(result)
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
