package ofdru

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestReceiptRaw(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/integration/v1/inn/INN1/kkt/KKT1/receipt/RawID1" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`{
  "Status": "Success",
  "Data": {
    "Tag": 3,
    "User": "ООО МКАС СПб",
    "UserInn": "7802870820 ",
    "Number": 1,
    "DateTime": "2016-07-26T12:32:00",
    "ShiftNumber": 1,
    "OperationType": 1,
    "TaxationType": 1,
    "Operator": "Администратор",
    "KKT_RegNumber": "111222333",
    "FN_FactoryNumber": "99990789388",
    "Items": [
      {
        "Name": "Услуги",
        "Price": 599000,
        "Quantity": 12,
        "Total": 599000,
        "CalculationMethod": 4,
        "SubjectType": 1,
        "NDS_Rate": 1,
        "NDS_Summ": 99833,
        "Nds00_TotalSumm": 0
      }
    ],
    "Buyer_Address": "",
    "Nds18_TotalSumm": 99833,
    "Amount_Total": 599000,
    "Amount_Cash": 0,
    "Amount_ECash": 599000,
    "Document_Number": 3,
    "FiscalSign": "MQTLUGn8",
    "DecimalFiscalSign": "3393623696",
    "KKT_MachineNumber": "1",
    "InternetSign": 1,
    "Format_Version": 2,
    "Amount_Advance": 0,
    "Amount_Loan": 0,
    "Amount_Granting": 0,
    "ExtraProperty": [
      {
        "ExtraProperty_Name": "Name",
        "ExtraProperty_Value": "Value"
      },
      {
        "ExtraProperty_Name": "Name",
        "ExtraProperty_Value": "Value"
      }
    ]
  }
}`))
			}
		}),
	)
	defer ts.Close()

	client := OfdRu("INN1", "12345", "56789", ts.URL)

	at, err := client.getReceiptRaw("RawID1", "KKT1")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, at.Status, "Success")
	assert.Equal(t, at.Data.Tag, 3)
	assert.Equal(t, at.Data.DateTime.Time, time.Date(2016, 07, 26, 12, 32, 0, 0, time.UTC))
}
