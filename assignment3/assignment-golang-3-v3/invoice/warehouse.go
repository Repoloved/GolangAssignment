package invoice

import (
	"errors"
	"strings"
)

// Warehouse invoice

type WarehouseInvoice struct {
	Date        string
	InvoiceType InvoiceTypeName
	Approved    bool
	Products    []Product
}

type InvoiceTypeName string

const (
	PURCHASE InvoiceTypeName = "purchase"
	SALES    InvoiceTypeName = "sales"
)

type Product struct {
	Name     string
	Unit     int
	Price    float64
	Discount float64
}

func (p Product) TotalPrice() float64 {
	return float64(p.Unit) * p.Price * (1 - p.Discount)
}

func (wi WarehouseInvoice) TotalInvoice() float64 {
	var total float64
	for _, product := range wi.Products {
		total += product.TotalPrice()
	}
	return total
}

func (wi WarehouseInvoice) RecordInvoice() (InvoiceData, error) {
	if strings.TrimSpace(wi.Date) == "" {
		return InvoiceData{}, errors.New("invoice date is empty")
	}

	if wi.InvoiceType == "" || (wi.InvoiceType != PURCHASE && wi.InvoiceType != SALES) {
		return InvoiceData{}, errors.New("invoice type is not valid")
	}

	if len(wi.Products) == 0 {
		return InvoiceData{}, errors.New("invoice products is empty")
	}

	for p := range wi.Products {
		if wi.Products[p].Unit <= 0 {
			return InvoiceData{}, errors.New("unit product is not valid")
		}

		if wi.Products[p].Price <= 0 {
			return InvoiceData{}, errors.New("price product is not valid")
		}
	}

	var invoiceData InvoiceData

	invoiceData.Date = wi.Date
	invoiceData.TotalInvoice = wi.TotalInvoice()
	invoiceData.Departemen = Warehouse

	return invoiceData, nil
}
