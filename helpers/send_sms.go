package helpers

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/TemurMannonov/buy_event/config"
)

type Message struct {
	Recipient string `json:"recipient"`
	MessageID string `json:"message_id"`
	SMS       struct {
		Originator string `json:"originator"`
		Content    struct {
			Text string `json:"text"`
		} `json:"content"`
	} `json:"sms"`
}

type Body struct {
	Messages []Message `json:"messages"`
}

func SendSMS(phoneNumber, text string) error {
	var (
		body    Body
		message Message
	)

	cfg := config.Load()

	client := http.Client{}

	messageID, err := generateCode(6)
	if err != nil {
		return err
	}

	message.MessageID = messageID
	message.Recipient = phoneNumber
	message.SMS.Content.Text = text
	message.SMS.Originator = cfg.PlayMobileOriginator

	body.Messages = append(body.Messages, message)
	requestBody, err := json.Marshal(body)

	request, err := http.NewRequest("POST", cfg.PlayMobileUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(cfg.PlayMobileLogin, cfg.PlayMobilePassword)
	res, err := client.Do(request)
	if err != nil {
		return errors.New("Error while sending sms code: " + err.Error())
	}

	if res.StatusCode != 200 {
		return errors.New("Playmobile failed to send sms")
	}

	return nil
}

func generateCode(max int) (string, error) {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, max)

	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return "", err
	}

	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b), nil
}
