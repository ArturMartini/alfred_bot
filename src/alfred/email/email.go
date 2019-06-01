package email

import (
	"log"
	"net/smtp"
)

func Send() bool {
	from := "arturmartiniti@gmail.com"
	pass := "*******"
	to := []string{"arturmartiniti@gmail.com"}

	msg := "From: " + from + "\n" +
		"To: one@email.com,two@email.com \n" +
		"Subject: ARTUR MARTINI - HOME OFFICE[{{now.Format(02-01-2006}}]\n\n" +
		"Hello friends.\n\nIm working at home today ({{now.Format(02-01-2006}}) .\n\nTks.\n\nArtur Martini"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, to, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return false
	}
	return true
}
