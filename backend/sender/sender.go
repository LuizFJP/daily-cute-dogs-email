package sender

import (
	"bytes"
	"crypto/tls"
	"daily-cute-dogs-email/backend/db"
	"daily-cute-dogs-email/backend/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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
)

var fetchResponse *models.Response

type message struct {
	*gomail.Message
}

func Start() {
	for {
		sleepUntilSixAM()
		fetchDog()
		sendEmail()
	}
}

func sleepUntilSixAM() {
	zone, err := time.LoadLocation("America/Belem")
	if err != nil {
		fmt.Println(err)
	}
	
	now := time.Now()
	sixAM := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, zone)
	hoursLeftAbsolute := math.Abs(time.Until(sixAM).Hours())
	hoursUntilSixAM := (24 - time.Duration(hoursLeftAbsolute)) * time.Hour

	time.Sleep(hoursUntilSixAM)
}

func fetchDog() {
	response := &models.Response{}
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, response)

	fetchResponse = response
}

func sendEmail() {
	messager := &message{gomail.NewMessage()}

	mdb := db.Start()
	defer mdb.Finish()

	emails, err := mdb.GetEmails()
	if err != nil {
		fmt.Println(err)
	}

	dialer := messager.login()

	for _, to := range emails {
		messager.loadHeader(to.Email)
		messager.createBody()
		messager.dialer(dialer)
	}
}

func (m *message) login() *gomail.Dialer {
	dialer := gomail.NewDialer("smtp.gmail.com", 587, from, password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return dialer
}

func (m *message) loadHeader(to string) {
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Seu doguinho fofo do dia chegou!")
}

func (m *message) createBody() {
	t, _ := template.ParseFiles("./backend/body.html")
	var body bytes.Buffer

	t.Execute(&body, struct {
		Link string
	}{
		Link: fetchResponse.Message,
	})

	m.SetBody("text/html", body.String())
}

func (m *message) dialer(d *gomail.Dialer) {
	if err := d.DialAndSend(m.Message); err != nil {
		fmt.Println(err)
	}
}
