package ofdru

import (
	"encoding/json"
	"fmt"
	"github.com/google/logger"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ReceiptResult struct {
	Status string               `json:"Status"`
	Data   []ReceiptInformation `json:"Data"`
}

type ReceiptInformation struct {
	ID             string      `json:"Id"`
	CDateUtc       SpecialDate `json:"CDateUtc"` //"2016-07-26T09:32:41"
	Tag            int         `json:"Tag"`
	IsBso          bool        `json:"IsBso"`
	IsCorrection   bool        `json:"IsCorrection"`
	OperationType  string      `json:"OperationType"`
	UserInn        string      `json:"UserInn"`
	KktRegNumber   string      `json:"KktRegNumber"`
	FnNumber       string      `json:"FnNumber"`
	DocNumber      int         `json:"DocNumber"`
	DocDateTime    SpecialDate `json:"DocDateTime"`
	DocShiftNumber int         `json:"DocShiftNumber"`
	ReceiptNumber  int         `json:"ReceiptNumber"`
	DocRawId       string      `json:"DocRawId"`
	TotalSumm      int         `json:"TotalSumm"`
	CashSumm       int         `json:"CashSumm"`
	ECashSumm      int         `json:"ECashSumm"`
	PrepaidSumm    int         `json:"PrepaidSumm"`
	CreditSumm     int         `json:"CreditSumm"`
	ProvisionSumm  int         `json:"ProvisionSumm"`
	TaxTotalSumm   int         `json:"TaxTotalSumm"`
	Tax10Summ      int         `json:"Tax10Summ"`
	Tax18Summ      int         `json:"Tax18Summ"`
	Tax110Summ     int         `json:"Tax110Summ"`
	Tax118Summ     int         `json:"Tax118Summ"`
	Tax0Summ       int         `json:"Tax0Summ"`
	TaxNaSumm      int         `json:"TaxNaSumm"`
	Depth          int         `json:"Depth"`
}

func (o *ofdru) getReceipts(kkt string, date time.Time) ([]Receipt, error) {
	wg := &sync.WaitGroup{}
	ch := make(chan []Receipt, 24)
	errCh := make(chan error, 24)

	wg.Add(24)
	for i:= 0; i < 24; i++ {
		go o.getReceiptsInHour(wg, ch, errCh, i, kkt, date)
	}
	wg.Wait()
	close(errCh)
	close(ch)
	errs := make([]string, 0)
	for errPart := range errCh {
		errs = append(errs, errPart.Error())
	}
	err := fmt.Errorf(strings.Join(errs, "\n"))

	receipts := make([]Receipt, 0)
	for receiptsPart := range ch {
		receipts = append(receipts, receiptsPart...)
	}

	return receipts, err
}

func (o *ofdru) getReceiptsInHour(wg *sync.WaitGroup, receiptCh chan []Receipt, errCh chan error, hour int, kkt string, date time.Time) {
	defer wg.Done()
	startDate := time.Date(date.Year(), date.Month(), date.Day(), hour, 0, 0, 0, date.Location())
	endDate := time.Date(date.Year(), date.Month(), date.Day(), hour, 59, 59, 59, date.Location())
	req, err := http.NewRequest("GET", o.baseURL+"/api/integration/v1/inn/"+o.Inn+"/kkt/"+kkt+"/receipts?AuthToken="+o.token.AuthToken+"&dateFrom="+startDate.Format("2006-01-02T15:04:05")+"&dateTo="+endDate.Format("2006-01-02T15:04:05"), nil)
	resp, err := o.Do(req)
	if err != nil {
		errCh <- err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			logger.Error("error closing response body")
		}
	}()
	rs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		formatErr := fmt.Errorf("ошибка чтения данных %v с %s по %s", err, startDate, endDate)
		errCh <- formatErr
	}
	var r ReceiptResult
	if err := json.Unmarshal(rs, &r); err != nil {
		formatErr := fmt.Errorf("ошибка преобразованя данных %v", err)
		errCh <- formatErr
	}

	receipts := make([]Receipt, 0, len(r.Data))
	for _, receipt := range r.Data {
		rec, err := o.getReceiptRaw(receipt.ID, kkt)
		if err != nil {
			errCh <- err
		}
		var products []Product
		for _, item := range rec.Data.Items {
			products = append(products, Product{
				Name:       item.Name,
				Quantity:   item.Quantity,
				Price:      item.Price,
				Vat:        item.NDSRate,
				VatPrice:   item.NDSSumm,
				TotalPrice: item.Total,
				FP:         rec.Data.FiscalSign,
				FD:         strconv.Itoa(rec.Data.DocumentNumber),
				FN:         receipt.FnNumber,
				Time:       rec.Data.DateTime.Time.Format(time.RFC3339Nano),
			})
		}

		receipts = append(receipts, Receipt{
			KktRegId: rec.Data.KKTRegNumber,
			ID:       rec.Data.DocumentNumber,
			FP:       rec.Data.FiscalSign,
			FD:       strconv.Itoa(rec.Data.DocumentNumber),
			Date:     rec.Data.DateTime.Time.Format(time.RFC3339Nano),
			Products: products,
			Link:     fmt.Sprintf("/rec/%s/%s/%s/%d/%s", o.Inn, receipt.KktRegNumber, receipt.FnNumber, receipt.DocNumber, rec.Data.DecimalFiscalSign),
			Price:    rec.Data.AmountTotal,
			VatPrice: rec.Data.Nds18TotalSumm,
		})
	}
	receiptCh <- receipts
}
