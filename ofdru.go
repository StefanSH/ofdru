package ofdru

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Client interface
type Client interface {
	GetReceipts(date time.Time) (receipts []Receipt, err error)
}

type ofdru struct {
	http.Client
	baseURL  string
	Inn      string
	Username string
	Password string
	token    *AccessToken
}

type Receipt struct {
	ID       int
	FP       string
	FD       string
	Date     string
	Products []Product
	Link     string
	Price    int
	VatPrice int
}

type Product struct {
	Name       string
	Quantity   int
	Price      int
	Vat        int
	VatPrice   int
	TotalPrice int
	FP         string
	FD         string
	FN         string
	Time       string
}

type SpecialDate struct {
	time.Time
}

func OfdRu(Inn string, Username string, Password string, baseURL string) *ofdru {
	return &ofdru{
		http.Client{
			Timeout: time.Duration(30) * time.Second,
		},
		baseURL,
		Inn,
		Username,
		Password,
		nil,
	}
}

func (o *ofdru) GetReceipts(date time.Time) (receipts []Receipt, err error) {
	o.token, err = o.auth()
	if err != nil {
		return receipts, fmt.Errorf("Ошибка авторизации %v", err)
	}

	kkts, err := o.getKkts()
	if err != nil {
		return receipts, fmt.Errorf("Ошибка получения списка ККТ %v", err)
	}
	for _, kkt := range kkts.Data {
		r, err := o.getReceipts(kkt.ID, date)
		if err != nil {
			return receipts, fmt.Errorf("Ошибка получения списка чеков по ККТ %s %v", kkt.ID, err)
		}
		receipts = append(receipts, r...)
	}

	return receipts, nil
}

func (i *SpecialDate) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	newTime, err := time.Parse("2006-01-02T15:04:05", strInput)
	if err != nil {
		return fmt.Errorf("Ошибка преобразования даты %v", err)
	}

	i.Time = newTime
	return nil
}
