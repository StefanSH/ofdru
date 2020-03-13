package ofdru

import (
	"encoding/json"
	"fmt"
	"github.com/google/logger"
	"io/ioutil"
	"net/http"
)

type ReceiptRawResult struct {
	Status string     `json:"Status"`
	Data   ReceiptRaw `json:"Data"`
}

type Item struct {
	Name              string `json:"Name"`
	Price             int    `json:"Price"`
	Quantity          int    `json:"Quantity"`
	Total             int    `json:"Total"`
	CalculationMethod int    `json:"CalculationMethod"`
	SubjectType       int    `json:"SubjectType"`
	NDSRate           int    `json:"NDS_Rate"`
	NDSSumm           int    `json:"NDS_Summ"`
	Nds00TotalSumm    int    `json:"Nds00_TotalSumm"`
}

type ExtraProperty struct {
	ExtraPropertyName  string `json:"ExtraProperty_Name"`
	ExtraPropertyValue string `json:"ExtraProperty_Value"`
}
type ReceiptRaw struct {
	Tag               int             `json:"Tag"`
	User              string          `json:"User"`
	UserInn           string          `json:"UserInn"`
	Number            int             `json:"Number"`
	DateTime          SpecialDate     `json:"DateTime"`
	ShiftNumber       int             `json:"ShiftNumber"`
	OperationType     int             `json:"OperationType"`
	TaxationType      int             `json:"TaxationType"`
	Operator          string          `json:"Operator"`
	KKTRegNumber      string          `json:"KKT_RegNumber"`
	FNFactoryNumber   string          `json:"FN_FactoryNumber"`
	Items             []Item          `json:"Items"`
	BuyerAddress      string          `json:"Buyer_Address"`
	Nds18TotalSumm    int             `json:"Nds18_TotalSumm"`
	AmountTotal       int             `json:"Amount_Total"`
	AmountCash        int             `json:"Amount_Cash"`
	AmountECash       int             `json:"Amount_ECash"`
	DocumentNumber    int             `json:"Document_Number"`
	FiscalSign        string          `json:"FiscalSign"`
	DecimalFiscalSign string          `json:"DecimalFiscalSign"`
	KKTMachineNumber  string          `json:"KKT_MachineNumber"`
	InternetSign      int             `json:"InternetSign"`
	FormatVersion     int             `json:"Format_Version"`
	AmountAdvance     int             `json:"Amount_Advance"`
	AmountLoan        int             `json:"Amount_Loan"`
	AmountGranting    int             `json:"Amount_Granting"`
	ExtraProperty     []ExtraProperty `json:"ExtraProperty"`
}

func (o *ofdru) getReceiptRaw(rawID string, kkt string) (receipts ReceiptRawResult, err error) {
	req, err := http.NewRequest("GET", o.baseURL+"/api/integration/v1/inn/"+o.Inn+"/kkt/"+kkt+"/receipt/"+rawID+"?AuthToken="+o.token.AuthToken, nil)
	resp, err := o.Do(req)
	if err != nil {
		return receipts, err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			logger.Error("error closing response body")
		}
	}()
	rs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return receipts, fmt.Errorf("Ошибка чтения данных %v", err)
	}
	if err := json.Unmarshal(rs, &receipts); err != nil {
		return receipts, fmt.Errorf("Ошибка преобразованя данных %v", err)
	}

	return receipts, err
}
