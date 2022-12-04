package main

import (
	"bytes"
	"net/url"
	"os/exec"
	"strconv"

	"router-telegram/internal/config"
	"router-telegram/internal/telegram"
)

func main() {
	if err := config.ParseArguments(); err != nil {
		panic(err)
	}

	resp := telegram.SendRequest("getUpdates", map[string]string{
		"allowed_updates": "[message]",
	})
	if resp == nil {
		return
	}

	data, err := telegram.ParseResponse(resp)
	if err != nil {
		return
	}

	var updateId int
	for _, message := range data.Result {
		updateId = message.UpdateId

		if message.Message.Text != nil && config.TelegramChatIds[message.Message.Chat.Id] {
			if *message.Message.Text == "/start" {
				telegram.SendRequest("sendMessage", map[string]string{
					"chat_id":      strconv.Itoa(message.Message.Chat.Id),
					"text":         "Welcome",
					"reply_markup": "{\"resize_keyboard\":true,\"keyboard\":[[{\"text\": \"/ping\"}]]}",
				})
			}
			if *message.Message.Text == "/ping" {
				telegram.SendRequest("sendMessage", map[string]string{
					"chat_id": strconv.Itoa(message.Message.Chat.Id),
					"text":    url.QueryEscape(execUptime()),
				})
			}
		}
	}

	if updateId != 0 {
		telegram.SendRequest("getUpdates", map[string]string{
			"offset": strconv.Itoa(updateId + 1),
		})
	}
}

func execUptime() string {
	cmd := exec.Command("uptime")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return err.Error()
	}

	return out.String()
}
