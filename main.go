package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AddressExtractRequest struct {
	Data []string
}

type AddressExtractResponse struct {
	Data []string
}

type AddressDetail struct {
	Neighbourhood string `json:"neighbourhood"`
	Street        string `json:"street"`
	No            string `json:"no"`
	NameSurname   string `json:"name_surname"`
	Address       string `json:"address"`
	City          string `json:"city"`
	Distinct      string `json:"distinct"`
	Tel           string `json:"tel"`
}

type BackendAddressFormat struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	City        string `json:"city"`
	District    string `json:"district"`
	FullAddress string `json:"full_address"`
	Tel         string `json:"tel"`
}

func sendExtractAddressResponse(text string) *AddressDetail {
	req := AddressExtractRequest{
		Data: []string{text},
	}

	addressExtractApiAddress := os.Getenv("ADDRESS_EXTRACT_API")
	if addressExtractApiAddress == "" {
		return nil
	}

	serialized, _ := json.Marshal(req)
	request, _ := http.NewRequest("POST", addressExtractApiAddress, bytes.NewReader(serialized))
	request.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred while sending address request: %s \n", err.Error())
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "status returned not okay address request: %d \n", resp.StatusCode)
		return nil
	}

	var addressExtractResponse AddressExtractResponse

	readedRespBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred while reading response body: %s \n", err.Error())
		return nil
	}

	if err := json.Unmarshal(readedRespBody, &addressExtractResponse); err != nil {
		fmt.Fprintf(os.Stderr, "error occurred while deserializing response body: %s \n", err.Error())
		return nil
	}

	serializedAddressExtractResp, err := json.Marshal(addressExtractResponse.Data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred while serializing extract response: %s \n", err.Error())
		return nil
	}

	var addressDetail AddressDetail
	if err := json.Unmarshal(serializedAddressExtractResp, &addressDetail); err != nil {
		fmt.Fprintf(os.Stderr, "error occurred while deserializing addressDetail: %s \n", err.Error())
		return nil
	}

	return &addressDetail
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
		if update.ChannelPost != nil && update.ChannelPost.Text != "" { // If we got a message
			log.Printf("telegram mesaj: [%s]", update.ChannelPost.Text)
			addressResponse := sendExtractAddressResponse(update.ChannelPost.Text)
			if addressResponse == nil {
				addressResponse = &AddressDetail{}
			}

			city := UNKNOWN
			if addressResponse.City == "" {
				city = string(ExtractCity(update.ChannelPost.Text))
				addressResponse.City = city
			}

			if addressResponse.Distinct == "" {
				district := ExtractDistrict(City(city), update.ChannelPost.Text)
				addressResponse.Distinct = district
			}
		}
	}
}
