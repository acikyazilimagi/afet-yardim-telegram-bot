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

const (
	backendURL = "https://apigo.afetharita.com/events"
)

type FeedEntry struct {
	ID              string `json:"id"`
	RawText         string `json:"raw_text"`
	Channel         string `json:"channel"`
	ExtraParameters string `json:"extra_parameters"`
	Epoch           int64  `json:"timestamp"`
}

func newFeedEntry(message *tgbotapi.Message) *FeedEntry {
	// TODO: maybe we can send username in extra parameters like: {"username": "someUsername"}
	return &FeedEntry{
		ID:      "",
		RawText: message.Text,
		Channel: "telegram",
		Epoch:   int64(message.Date),
	}
}

func sendDataToBackend(message *tgbotapi.Message) bool {
	// Checking for api key
	apiKey := os.Getenv("BACKEND_GO_API_KEY")
	if apiKey == "" {
		panic("BACKEND_GO_API_KEY is not found")
	}

	// Generating feed entry
	feedEntry := newFeedEntry(message)
	serialized, err := json.Marshal(feedEntry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred while serializing feed entry: %s \n", err.Error())
		return false
	}

	// creating request
	req, err := http.NewRequest(http.MethodPost, backendURL, bytes.NewReader(serialized))
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not prepare request data to backend: %s", err.Error())
		return false
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Api-Key", apiKey)

	// sending request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not send data to backend: %s", err.Error())
		return false
	}

	if resp.StatusCode == 201 {
		fmt.Fprintf(os.Stderr, "an error on backend response: %s", err.Error())
		return false
	}
	return true
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
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			sendDataToBackend(update.Message)
		}
	}
}
