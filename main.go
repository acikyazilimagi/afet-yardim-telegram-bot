package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const envBotKey = "BOT_KEY"

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	registerSignals(ctx, cancel)

	botKey := os.Getenv(envBotKey)
	if botKey == "" {
		log.Fatalf("%s is not found\n", envBotKey)
	}

	bot, err := tgbotapi.NewBotAPI(botKey)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			log.Println("shutting down...")
			return

		case update := <-updates:
			if update.Message != nil {
				addressResponse, err := sendExtractAddressResponse(ctx, update.Message.Text)
				if err != nil {
					addressResponse = &AddressDetail{}
					log.Println(err)
				}

				city := UNKNOWN
				if addressResponse.City == "" {
					city = ExtractCity(update.Message.Text)
					addressResponse.City = string(city)
				}

				if addressResponse.District == "" {
					district := ExtractDistrict(City(city), update.Message.Text)
					addressResponse.District = district
				}

				// TODO: waiting contract from backend
				//sendDataToBackend()
			}
		}
	}
}

func registerSignals(ctx context.Context, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Recovered in registerSignals", r)
			}
		}()

		select {
		case <-ctx.Done():
		case <-ch:
		}

		signal.Stop(ch)
		cancel()
	}()
}
