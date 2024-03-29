package telegrambotnotificationgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TelegramService struct {
	Token  string
	ChatId string
}

func NewTelegramService(token, chatId string) *TelegramService {
	return &TelegramService{
		Token:  token,
		ChatId: chatId,
	}
}

func (t *TelegramService) keyword(keyword string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", t.Token, keyword)
}

func (t *TelegramService) SendMessage(message string) error {
	client := http.Client{}

	data := map[string]string{
		"chat_id": t.ChatId,
		"text":    message,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", t.keyword("sendMessage"), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message: %s", resp.Status)
	}

	return nil
}
