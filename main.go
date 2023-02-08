package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"os"
)

type FeedEntry struct {
	IsResolved bool   `json:"is_resolved"`
	FullText   string `json:"full_text"`
	Channel    string `json:"channel"`
}

func newFeedEntry(fullText string) *FeedEntry {
	return &FeedEntry{
		IsResolved: false,
		FullText:   fullText,
		Channel:    "telegram",
	}
}

func sendDataToBackend(text string) {
	feedEntry := newFeedEntry(text)
	serialized, err := json.Marshal(feedEntry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred while serializing feed entry: %s \n", err.Error())
		return
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.afetharita.com/feeds/entries", bytes.NewReader(serialized))
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not prepare request data to backend: %s", err.Error())
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "AfetHarita "+os.Getenv("BACKEND_TOKEN"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not send data to backend: %s", err.Error())
		return
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 { // 2xx
		fmt.Fprintf(os.Stderr, "an error on backend response: %s", err.Error())
		return
	}
}

func main() {
	botKey := os.Getenv("BOT_KEY")
	if botKey == "" {
		panic("BOT_KEY is not found")
	}

	bot, err := tgbotapi.NewBotAPI(botKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			// TODO send username: username, time_sent: timestamp in the future as extraParams
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			sendDataToBackend(update.Message.Text)
		}
	}
}
