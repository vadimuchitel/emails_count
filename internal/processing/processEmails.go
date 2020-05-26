package processing

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func logMessage(message interface{}, err error) {
	fmt.Println(message, err)
}

func sendResponse(w http.ResponseWriter, responseCode int, payload []byte) {
	w.WriteHeader(responseCode)
	if payload == nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(payload); err != nil {
		logMessage("PayloadErr", err)
	}
}

type Emails struct {
	Emails []string `json:"emails"`
}

type Response struct {
	NumEmails int `json:"num_emails"`
}

func ProcessEmails(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, nil)
		return
	}
	emails := Emails{}
	if err := json.Unmarshal(body, &emails); err != nil {
		sendResponse(w, http.StatusBadRequest, nil)
		return
	}

	numEmails := getNumEmails(emails.Emails)
	response := Response{
		NumEmails: numEmails,
	}
	payload, err := json.Marshal(response)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, nil)
		return
	}
	sendResponse(w, http.StatusOK, payload)
}

func getNumEmails(emails []string) int {
	var a struct{}
	cleanEmails := make(map[string]struct{})

	for _, email := range emails {
		cleanEmail, err := cleanEmail(email)
		logMessage(cleanEmail, err)
		if err != nil {
			continue // non-gmail accounts are ignored
		}
		if _, ok := cleanEmails[cleanEmail]; ok {
			continue
		}
		cleanEmails[cleanEmail] = a
	}
	return len(cleanEmails)
}

func cleanEmail(email string) (string, error) {
	name := strings.Split(strings.TrimSpace(email), "@gmail.com")
	l := len(name)
	if l == 1 {
		return "", errors.New("Not gmail.com email")
	}
	if l > 2 {
		return "", errors.New("@gmail.com@gmail.com is not a valid format")
	}
	username := strings.Split(name[0], "+")[0]
	username = strings.Replace(username, ".", "", -1)
	if username == "" {
		return "", errors.New("Empty Username")
	}
	return username, nil
}
