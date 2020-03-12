package ofdru

import (
	"encoding/json"
	"github.com/google/logger"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type KktResult struct {
	Status string `json:"Status"`
	Data   []Kkt  `json:"Data"`
}
type Kkt struct {
	ID                      string    `json:"Id"`
	KktRegId                string    `json:"KktRegId"`
	KktName                 string    `json:"KktName"`
	SerialNumber            string    `json:"SerialNumber"`
	FnNumber                string    `json:"FnNumber"`
	CreateDate              time.Time `json:"CreateDate"` //: "2017-01-13T12:09:51",
	PaymentDate             time.Time `json:"PaymentDate"`
	CheckDate               time.Time `json:"CheckDate"`
	ActivationDate          time.Time `json:"ActivationDate"`
	FirstDocumentDate       time.Time `json:"FirstDocumentDate"`
	ContractStartDate       time.Time `json:"ContractStartDate"`
	ContractEndDate         time.Time `json:"ContractEndDate"`
	LastDocOnKktDateTime    time.Time `json:"LastDocOnKktDateTime"`
	LastDocOnOfdDateTimeUtc time.Time `json:"LastDocOnOfdDateTimeUtc"`
	FiscalAddress           string    `json:"FiscalAddress"`
	FiscalPlace             string    `json:"FiscalPlace"`
	Path                    string    `json:"Path"`
	KktModel                string    `json:"KktModel"`
	FnEndDate               time.Time `json:"FnEndDate"`
}

func (o *ofdru) getKkts() (kkts KktResult, err error) {
	req, err := http.NewRequest("GET", o.baseURL+"/api/integration/v1/inn/"+o.Inn+"/kkts?AuthToken="+o.token.AuthToken, nil)
	resp, err := o.Do(req)
	if err != nil {
		return kkts, err
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
	//var at *AccessToken
	if err := json.Unmarshal(rs, &kkts); err != nil {
		return kkts, err
	}

	return kkts, err
}

func (at *Kkt) UnmarshalJSON(b []byte) error {
	var rawStrings map[string]interface{}

	if err := json.Unmarshal(b, &rawStrings); err != nil {
		return err
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "id" {
			at.ID = v.(string)
		}
		if strings.ToLower(k) == "kktregid" {
			at.KktRegId = v.(string)
		}
		if strings.ToLower(k) == "kktname" {
			at.KktName = v.(string)
		}
		if strings.ToLower(k) == "serialnumber" {
			at.SerialNumber = v.(string)
		}
		if strings.ToLower(k) == "fnnumber" {
			at.FnNumber = v.(string)
		}
		if strings.ToLower(k) == "createdate" {
			t, err := time.Parse("2006-01-02T15:04:05", v.(string))
			if err != nil {
				return err
			}
			at.CreateDate = t
		}
		if strings.ToLower(k) == "paymentdate" {
			t, err := time.Parse("2006-01-02T15:04:05", v.(string))
			if err != nil {
				return err
			}
			at.PaymentDate = t
		}
		if strings.ToLower(k) == "checkdate" {
			t, err := time.Parse("2006-01-02T15:04:05", v.(string))
			if err != nil {
				return err
			}
			at.CheckDate = t
		}
		if strings.ToLower(k) == "activationdate" {
			t, err := time.Parse("2006-01-02T15:04:05", v.(string))
			if err != nil {
				return err
			}
			at.ActivationDate = t
		}
		if strings.ToLower(k) == "firstdocumentdate" {
			t, err := time.Parse("2006-01-02T15:04:05", v.(string))
			if err != nil {
				return err
			}
			at.FirstDocumentDate = t
		}
		if strings.ToLower(k) == "contractstartdate" {
			t, err := time.Parse("2006-01-02T15:04:05", v.(string))
			if err != nil {
				return err
			}
			at.ContractStartDate = t
		}
		if strings.ToLower(k) == "contractenddate" {
			t, err := time.Parse("2006-01-02T15:04:05", v.(string))
			if err != nil {
				return err
			}
			at.ContractEndDate = t
		}
		if strings.ToLower(k) == "lastdoconkktdatetime" {
			t, err := time.Parse("2006-01-02T15:04:05", v.(string))
			if err != nil {
				return err
			}
			at.LastDocOnKktDateTime = t
		}
		if strings.ToLower(k) == "lastdoconofddatetimeutc" {
			t, err := time.Parse("2006-01-02T15:04:05", v.(string))
			if err != nil {
				return err
			}
			at.LastDocOnOfdDateTimeUtc = t
		}
		if strings.ToLower(k) == "fiscaladdress" {
			at.FiscalAddress = v.(string)
		}
		if strings.ToLower(k) == "fiscalplace" {
			at.FiscalPlace = v.(string)
		}
		if strings.ToLower(k) == "path" {
			at.Path = v.(string)
		}
		if strings.ToLower(k) == "kktmodel" {
			at.KktModel = v.(string)
		}
		if strings.ToLower(k) == "fnenddate" {
			t, err := time.Parse("2006-01-02T15:04:05", v.(string))
			if err != nil {
				return err
			}
			at.FnEndDate = t
		}
	}

	return nil
}
