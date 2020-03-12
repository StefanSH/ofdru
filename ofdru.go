package ofdru

import (
	"net/http"
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
	token    AccessToken
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
	TotalPrice string
	FP         string
	FD         string
	FN         string
	Time       string
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
		AccessToken{},
	}
}

func (o *ofdru) GetReceipts(date time.Time) (receipts []Receipt, err error) {

	return receipts, nil
}
