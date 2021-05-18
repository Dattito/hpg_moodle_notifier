package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type checkTokenResponse struct {
	Exception string
	Errorcode string
	Message   string
	Token     string
}

func CheckToken(token string) (bool, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("MOODLE_HOST")+"/webservice/rest/server.php", nil)
	if err != nil {
		return false, err
	}
	q := req.URL.Query()
	q.Add("wstoken", token)
	q.Add("moodlewsrestformat", "json")

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)

	var tokenResp checkTokenResponse

	json.Unmarshal(resp_body, &tokenResp)

	valid := tokenResp.Errorcode != "invalidtoken"

	return valid, nil
}

func GetMoodleToken(username string, password string) (string, error) {

	data := url.Values{
		"username": {username},
		"password": {password},
		"service":  {"moodle_mobile_app"},
	}

	resp, err := http.PostForm(os.Getenv("MOODLE_HOST")+"/login/token.php", data)
	if err != nil {
		return "", err
	}

	resp_body, _ := ioutil.ReadAll(resp.Body)

	var tokenResp checkTokenResponse

	json.Unmarshal(resp_body, &tokenResp)

	return tokenResp.Token, nil
}
