package invoice

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

func (fi FinanceInvoice) RecordInvoice() (InvoiceData, error) {
	return InvoiceData{},
	var invoice1 FinanceInvoice
	
	invoice1.Date = "01/01/2022"
	invoice1.Status = PAID
	invoice1.Approved = true
	invoice1.Details = "Details "
}
