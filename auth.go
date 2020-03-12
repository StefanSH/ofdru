package ofdru

import (
	"bytes"
	"encoding/json"
	"github.com/google/logger"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AccessToken struct {
	AuthToken         string    `json:"AuthToken"`
	ExpirationDateUtc time.Time `json:"ExpirationDateUtc"`
}

func (o *ofdru) auth() (authToken *AccessToken, err error) {
	data := url.Values{}
	data.Set("Login", o.Username)
	data.Add("Password", o.Password)
	req, err := http.NewRequest("POST", o.baseURL+"/api/Authorization/CreateAuthToken", bytes.NewBufferString(data.Encode()))

	resp, err := o.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			logger.Error("error closing response body")
		}
	}()
	rs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var at *AccessToken
	if err := json.Unmarshal(rs, &at); err != nil {
		return nil, err
	}

	return at, nil
}

func (at *AccessToken) UnmarshalJSON(b []byte) error {
	var rawStrings map[string]interface{}

	if err := json.Unmarshal(b, &rawStrings); err != nil {
		return err
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "authtoken" {
			at.AuthToken = v.(string)
		}

		if strings.ToLower(k) == "expirationdateutc" {
			t, err := time.Parse("2006-01-02T15:04:05", v.(string))
			if err != nil {
				return err
			}
			at.ExpirationDateUtc = t
		}
	}

	return nil
}
