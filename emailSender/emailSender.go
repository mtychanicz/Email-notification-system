package emailSender

import (
	"fmt"
	dbModels "golang-projects/dbModels"
	"io/ioutil"
	"log"
	"net/smtp"
	"strings"
)

func Send(email string, message *dbModels.Message) {
	data, err := ioutil.ReadFile("config.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	credentials := string(data)
	strSlice := strings.Split(credentials, ";")
	login := strSlice[0]
	password := strSlice[1]

	auth := smtp.PlainAuth("", login, password, "smtp.wp.pl")

	to := []string{email}

	data, err = ioutil.ReadFile("preparedEmails/greetMessage.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	emailMessage := []byte("From: " + login + "\r\n" +

		"To: " + email  + "\r\n" +

		"Subject: Test\r\n" +

		"\r\n" +

		message.Content) // Wyslij wiadomość

	err = smtp.SendMail("smtp.wp.pl:587", auth, login, to, emailMessage)

	if err != nil {
		log.Fatal(err)
	}
}
