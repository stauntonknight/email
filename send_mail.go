package  main

import (
	"bytes"
	"log"
	"net/smtp"
	"html/template"
	"fmt"
)

func main() {
	t, err := template.ParseFiles("email.html")
	if err != nil {
	    log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "Naveen",
		URL:  "http://google.com",
	}
	if err = t.Execute(buf, templateData); err != nil {
	    log.Fatal(err)
	}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + "Subject" + "!\n"
	msg := subject + mime + "\n" + buf.String()
	client, err := smtp.Dial("localhost:25")
	if err !=  nil {
		log.Fatal(err)
	}
	defer client.Close()
	client.Mail("knite@gmail.com")
	client.Rcpt("ydv.naveen@yahoo.com")
	wc, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, msg)
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = client.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
