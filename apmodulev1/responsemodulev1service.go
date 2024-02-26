package apmodulev1

import (
	"fmt"
	"time"
)

func (ap *ArticlePurchaseResponse) PerformPurchasev4(quantity int) {
	time := time.Now()

	switch {
	case ap.AT.Stock == 0:
		ap.Issuccess = false
		ap.Response = "There is no article available in stock"

	case ap.AT.Stock < quantity:
		ap.Issuccess = false
		verb := "is"
		if ap.AT.Stock > 1 {
			verb = "are"
		}
		ap.Response = fmt.Sprintf("There %v just %v(%v) available in stock", verb, ap.AT.Stock, ap.AT.Name)

	default:
		ap.AT.Date = time.Format("2006-01-02")
		ap.AT.Quantity = quantity
		ap.AT.Stock = ap.AT.Stock - quantity
		ap.AT.Total = ap.AT.Unitprice * float64(quantity)
	}

}
