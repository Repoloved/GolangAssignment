package invoice

import (
	"errors"
	"strings"
)

// Finance invoice
type FinanceInvoice struct {
	Date     string
	Status   InvoiceStatus // status: "paid", "unpaid"
	Approved bool
	Details  []Detail
}

type InvoiceStatus string

const (
	PAID   InvoiceStatus = "paid"
	UNPAID InvoiceStatus = "unpaid"
)

type Detail struct {
	Description string
	Total       int
}

func (fi FinanceInvoice) TotalInvoice() float64 {
	var total float64
	for _, detail := range fi.Details {
		total += float64(detail.Total)
	}
	return total
}

func (fi FinanceInvoice) RecordInvoice() (InvoiceData, error) {

	if strings.TrimSpace(fi.Date) == "" {
		return InvoiceData{}, errors.New("invoice date is empty")
	}

	if len(fi.Details) == 0 {
		return InvoiceData{}, errors.New("invoice details is empty")
	}

	if fi.Status == "" || (fi.Status != PAID && fi.Status != UNPAID) {
		return InvoiceData{}, errors.New("invoice status is not valid")
	}

	for d := range fi.Details {
		if fi.Details[d].Total <= 0 {
			return InvoiceData{}, errors.New("total price is not valid")
		}
	}

	var invoiceData InvoiceData

	invoiceData.Date = formatDate(fi.Date)
	invoiceData.TotalInvoice = fi.TotalInvoice()
	invoiceData.Departemen = Finance

	return invoiceData, nil
}
