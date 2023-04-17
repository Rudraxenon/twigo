package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Message struct {
	Body           string `xml:",chardata"`
	To             string `xml:"to,attr,omitempty"`
	From           string `xml:"from,attr,omitempty"`
	StatusCallback string `xml:"statusCallback,attr,omitempty"`
}

func main() {

	godotenv.Load()
	accountSid := os.Getenv("SID")
	authToken := os.Getenv("TOKEN")

	// fmt.Println(accountSid)
	// fmt.Println(authToken)

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	r := mux.NewRouter()

	params := &twilioApi.CreateMessageParams{}
	params.SetTo("+918159030925")
	params.SetFrom("+15075640946")
	params.SetBody("Hello from Go!")
	params.SetStatusCallback("https://demo.twilio.com/welcome/sms/reply")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	r.HandleFunc("/reply", func(w http.ResponseWriter, r *http.Request) {
		// smsSid := r.FormValue("SmsSid")
		// from := r.FormValue("From")
		body := r.FormValue("Body")

		// fmt.Println(body)

		if body == "1" || strings.Contains(body, "something in punjabi") {
			message1 := Message{
				Body:           "hola",
				To:             "+918159030925",
				From:           "+15075640946",
				StatusCallback: "https://demo.twilio.com/welcome/sms/reply",
			}
			// params.SetBody(message1)
			// resp, err := client.Api.CreateMessage(params)
			xmlResponse, err := xml.Marshal(message1)
			if err != nil {
				panic(err)
			}
			w.Write(xmlResponse)
		} else if body == "2" || strings.Contains(body, "something else in punjabi") {
			message2 := Message{
				Body:           "mushi mushi",
				To:             "+918159030925",
				From:           "+15075640946",
				StatusCallback: "https://demo.twilio.com/welcome/sms/reply",
			}
			// params.SetBody(message2)
			// resp, err := client.Api.CreateMessage(params)
			xmlResponse, err := xml.Marshal(message2)
			if err != nil {
				panic(err)
			}
			w.Write(xmlResponse)
		} else {
			message3 := Message{
				Body:           "namaste",
				To:             "+918159030925",
				From:           "+15075640946",
				StatusCallback: "https://demo.twilio.com/welcome/sms/reply",
			}
			// params.SetBody(message3)
			// resp, err := client.Api.CreateMessage(params)
			xmlResponse, err := xml.Marshal(message3)
			if err != nil {
				panic(err)
			}
			w.Write(xmlResponse)
		}
	}).Methods("POST")

	e := http.ListenAndServe(":8888", r)
	if e != nil {
		fmt.Println(e)
	}

}
