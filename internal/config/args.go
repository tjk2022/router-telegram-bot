package config

import (
	"errors"
	"flag"
	"strconv"
	"strings"
)

var (
	TelegramChatIds map[int]bool
	TelegramToken   string
)

func ParseArguments() error {
	flag.StringVar(&TelegramToken, "token", "", "5815872851:AAGPUcFOa9QTu_bzGrh194nL7VVYD3eh6x8")
	chatIdsInput := flag.String("chatIds", "", "124586195")

	flag.Parse()

	if TelegramToken = strings.TrimSpace(TelegramToken); TelegramToken == "" {
		return errors.New("telegram TelegramToken cannot be empty")
	}

	TelegramChatIds = make(map[int]bool)

	for _, chatId := range strings.Split(*chatIdsInput, ",") {
		value, err := strconv.Atoi(chatId)
		if err != nil {
			return errors.New("invalid chatId")
		}

		TelegramChatIds[value] = true
	}

	return nil
}
