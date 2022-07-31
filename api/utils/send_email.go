package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// 出品者へメール送信
func SendEmail(user_name, email, subject, plainTextContent, htmlContent string) {
	from := mail.NewEmail("UruApi", os.Getenv("AdminEmail"))
	to := mail.NewEmail(user_name, email)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SendGridApiKey"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}