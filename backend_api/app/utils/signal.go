package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type signalRequestData struct {
	Message    string   `json:"message"`
	Number     string   `json:"number"`
	Recipients []string `json:"recipients"`
}

func SendSignalMessage(phoneNumber string, message string) error {
	srd := signalRequestData{
		Message:    message,
		Number:     os.Getenv("SIGNAL_NUMBER"),
		Recipients: []string{phoneNumber},
	}
	payloadBuffer := new(bytes.Buffer)

	json.NewEncoder(payloadBuffer).Encode(srd)

	req, _ := http.NewRequest("POST", os.Getenv("SIGNAL_HOST")+"/v2/send", payloadBuffer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		if resp.StatusCode == 400 {
			return errors.New("phone_number not valid")
		}
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
		return errors.New("something went wrong")
	}
	return nil
}
