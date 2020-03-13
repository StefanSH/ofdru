package ofdru

import (
	"encoding/json"
	"fmt"
	"github.com/google/logger"
	"io/ioutil"
	"net/http"
)

type KktResult struct {
	Status string `json:"Status"`
	Data   []Kkt  `json:"Data"`
}
type Kkt struct {
	ID                      string      `json:"Id"`
	KktRegId                string      `json:"KktRegId"`
	KktName                 string      `json:"KktName"`
	SerialNumber            string      `json:"SerialNumber"`
	FnNumber                string      `json:"FnNumber"`
	CreateDate              SpecialDate `json:"CreateDate"` //: "2017-01-13T12:09:51",
	PaymentDate             SpecialDate `json:"PaymentDate"`
	CheckDate               SpecialDate `json:"CheckDate"`
	ActivationDate          SpecialDate `json:"ActivationDate"`
	FirstDocumentDate       SpecialDate `json:"FirstDocumentDate"`
	ContractStartDate       SpecialDate `json:"ContractStartDate"`
	ContractEndDate         SpecialDate `json:"ContractEndDate"`
	LastDocOnKktDateTime    SpecialDate `json:"LastDocOnKktDateTime"`
	LastDocOnOfdDateTimeUtc SpecialDate `json:"LastDocOnOfdDateTimeUtc"`
	FiscalAddress           string      `json:"FiscalAddress"`
	FiscalPlace             string      `json:"FiscalPlace"`
	Path                    string      `json:"Path"`
	KktModel                string      `json:"KktModel"`
	FnEndDate               SpecialDate `json:"FnEndDate"`
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
		return kkts, fmt.Errorf("Ошибка чтения данных %v", err)
	}
	//var at *AccessToken
	if err := json.Unmarshal(rs, &kkts); err != nil {
		return kkts, fmt.Errorf("Ошибка преобразованя данных %v", err)
	}

	return kkts, err
}
