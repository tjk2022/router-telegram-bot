package telegram

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"router-telegram/internal/config"
)

const API = "https://api.telegram.org/bot"

type Response struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text *string `json:"text"`
	Chat Chat    `json:"chat"`
}

type Chat struct {
	Id int `json:"id"`
}

func ParseResponse(r *http.Response) (*Response, error) {
	var response Response

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		log.Printf("Could not decode incoming update %s", err.Error())
		return nil, err
	}

	return &response, nil
}

func SendRequest(method string, queries map[string]string) *http.Response {
	var query strings.Builder

	for q, b := range queries {
		query.WriteString(q + "=" + b + "&")
	}

	fmt.Println(API + config.TelegramToken + "/" + method + "?" + strings.TrimRight(query.String(), "&"))
	resp, err := http.Get(API + config.TelegramToken + "/" + method + "?" + strings.TrimRight(query.String(), "&"))

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return resp
}
