package ofdru

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestReceipt(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/integration/v1/inn/INN1/kkt/KKT1/receipts" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`{
  "Status": "Success",
  "Data": [
    {
      "Id": "3a6e3b83-a0b0-4587-bfb3-1b7539b05cf3",
      "CDateUtc": "2016-07-26T09:32:41",
      "Tag": 0,
      "IsBso": false,
      "IsCorrection": false,
      "OperationType": "Income",
      "UserInn": "7802870820",
      "KktRegNumber": "111222333",
      "FnNumber": "99990789388",
      "DocNumber": 3,
      "DocDateTime": "2016-07-26T12:32:00",
      "DocShiftNumber": 1,
      "ReceiptNumber": 1,
      "DocRawId": "3a6e3b83-a0b0-4587-bfb3-1b7539b05cf3",
      "TotalSumm": 0,
      "CashSumm": 0,
      "ECashSumm": 0,
      "PrepaidSumm": 0,
      "CreditSumm": 0,
      "ProvisionSumm": 0,
      "TaxTotalSumm": 0,
      "Tax10Summ": 0,
      "Tax18Summ": 0,
      "Tax110Summ": 0,
      "Tax118Summ": 0,
      "Tax0Summ": 0,
      "TaxNaSumm": 0,
      "Depth": 3
    }
  ]
}`))
			}
		}),
	)
	defer ts.Close()

	client := OfdRu("INN1", "12345", "56789", ts.URL)

	at, err := client.getReceipts("KKT1", time.Now())
	if err != nil {
		assert.Error(t, err)
	}
	for _, receipt := range at {
		assert.Equal(t, receipt.ID, 0)
		assert.Equal(t, receipt.Link, "/rec/INN1/111222333/99990789388/3/")
	}
}
