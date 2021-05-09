package helpers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/TemurMannonov/buy_event/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(email, text string) error {
	from := mail.NewEmail("Temur Mannonov", config.Load().Mail)
	subject := "Order"
	to := mail.NewEmail("Customer", email)
	plainTextContent := text
	htmlContent := fmt.Sprintf("<strong>%s</strong>", text)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.Load().SendGridApiKey)
	response, err := client.Send(message)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("Sendgird failed to send mail")
	}

	return nil
}
