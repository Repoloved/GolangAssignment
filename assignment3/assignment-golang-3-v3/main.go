package main

import (
	"a21hc3NpZ25tZW50/invoice"
	"fmt"
	"log"
)

func RecapDataInvoice(data []invoice.Invoice) ([]invoice.InvoiceData, error) {
	var result []invoice.InvoiceData

	for _, d := range data {
		inv, err := d.RecordInvoice()
		isDupe := false

		if err != nil {
			return nil, err
		}

		for idx, i := range result {
			if i.Date == inv.Date && i.Departemen == inv.Departemen {
				result[idx].TotalInvoice += inv.TotalInvoice
				isDupe = true
				break
			}
		}

		if d.GetApproved() && !isDupe {
			if inv.Departemen == invoice.Finance {
				if d.(invoice.FinanceInvoice).GetStatus() == invoice.PAID {
					result = append(result, inv)
				}
			} else if inv.Departemen == invoice.Warehouse {
				if d.(invoice.WarehouseInvoice).GetInvoiceType() == invoice.PURCHASE {
					result = append(result, inv)
				}
			} else {
				result = append(result, inv)
			}
		}

	}
	return result, nil
}

func main() {
	listInvoice := []invoice.Invoice{
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{Description: "pembelian nota", Total: 4000}, {Description: "Pembelian alat tulis", Total: 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{Description: "pembelian nota", Total: 4000}, {Description: "Pembelian alat tulis", Total: 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.WarehouseInvoice{
			Date: "01-February-2020",
			Products: []invoice.Product{
				{Name: "product A", Unit: 10, Price: 10000, Discount: 0.1},
				{Name: "product C", Unit: 5, Price: 15000, Discount: 0.2},
			},
			InvoiceType: invoice.PURCHASE,
			Approved:    true,
		},
		invoice.MarketingInvoice{
			Date:        "01/02/2020",
			StartDate:   "20/01/2020",
			EndDate:     "25/01/2020",
			Approved:    true,
			PricePerDay: 10000,
			AnotherFee:  5000,
		},
	}

	result, err := RecapDataInvoice(listInvoice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
