package sender

import (
	"bytes"
	"crypto/tls"
	"daily-cute-dogs-email/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	_ "github.com/joho/godotenv/autoload"
	gomail "gopkg.in/mail.v2"
)

var (
	from     string = os.Getenv("FROM_EMAIL")
	password string = os.Getenv("FROM_PASSWORD")
	to       string = os.Getenv("SEND_EMAIL")
)

type message struct {
	*gomail.Message
}

func Start() {
	for {
		sendEmail()
		time.Sleep(1 * time.Hour)

	}
}

func sendEmail() {
	messager := &message{gomail.NewMessage()}

	messager.loadHeader()
	messager.createBody()
	messager.dialer()

}

func (m *message) loadHeader() {
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Seu doguinho fofo do dia chegou!")
}

func (m *message) createBody() {
	t, _ := template.ParseFiles("body.html")
	var body bytes.Buffer

	fetch := m.fetchDog()

	t.Execute(&body, struct {
		Link string
	}{
		Link: fetch.Message,
	})

	m.SetBody("text/html", body.String())
}

func (m *message) fetchDog() *models.Response {
	response := &models.Response{}
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, response)
	return response
}

func (m *message) dialer() {
	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m.Message); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
