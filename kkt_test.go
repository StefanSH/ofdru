package ofdru

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKkt(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/integration/v1/inn/INN1/kkts" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`{
  "Status": "Success",
  "Data": [
    {
      "Id": "00000000-0000-0000-0000-000000000000",
      "KktRegId": "9304171212297195",
      "KktName": "Касса 1",
      "SerialNumber": "44444444444443421132",
      "FnNumber": "0666666666666660",
      "CreateDate": "2017-01-13T12:09:51",
      "PaymentDate": "2017-01-13T12:15:43",
      "CheckDate": "2017-01-13T12:12:47",
      "ActivationDate": "2017-01-13T12:15:48",
      "FirstDocumentDate": "2017-01-13T14:15:48",
      "ContractStartDate": "2017-01-13T12:12:47",
      "ContractEndDate": "2018-02-12T12:12:47",
      "LastDocOnKktDateTime": "2017-02-12T10:12:00",
      "LastDocOnOfdDateTimeUtc": "2017-02-12T07:13:10",
      "FiscalAddress": "https://ofd.ru/",
      "FiscalPlace": "https://ofd.ru/",
      "Path": "/Мои кассы/Список касс 1/",
      "KktModel": "АТОЛ 42ФС",
      "FnEndDate": "2019-10-17T12:47:57"
    }
  ]
}`))
			}
		}),
	)
	defer ts.Close()

	client := OfdRu("INN1", "12345", "56789", ts.URL)

	at, err := client.getKkts()
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, at.Status, "Success")
}
