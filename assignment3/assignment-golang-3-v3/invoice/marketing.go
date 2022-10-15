package invoice

import (
	"errors"
	"strings"
	"time"
)

// Marketing invoice
type MarketingInvoice struct {
	Date        string
	StartDate   string
	EndDate     string
	PricePerDay int
	AnotherFee  int
	Approved    bool
}

func parseToDate(dateInString string) time.Time {
	tmp := strings.Split(dateInString, "/")
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	date, _ := time.Parse("2006-01-02", strings.Join(tmp, "-"))
	return date
}

func (mi MarketingInvoice) TotalDay() float64 {
	startDate := parseToDate(mi.StartDate)
	endDate := parseToDate(mi.EndDate)
	return (endDate.Sub(startDate).Hours() / 24) + 1
}

func (mi MarketingInvoice) TotalInvoice() float64 {
	return mi.TotalDay()*float64(mi.PricePerDay) + float64(mi.AnotherFee)
}

func (mi MarketingInvoice) GetApproved() bool {
	return mi.Approved
}

func (mi MarketingInvoice) RecordInvoice() (InvoiceData, error) {

	if mi.Date == "" {
		return InvoiceData{}, errors.New("invoice date is empty")
	}

	if mi.StartDate == "" || mi.EndDate == "" {
		return InvoiceData{}, errors.New("travel date is empty")
	}

	if mi.PricePerDay <= 0 {
		return InvoiceData{}, errors.New("price per day is not valid")
	}

	var invoiceData InvoiceData

	invoiceData.Date = formatDate(mi.Date)
	invoiceData.TotalInvoice = mi.TotalInvoice()
	invoiceData.Departemen = Marketing

	return invoiceData, nil
}
