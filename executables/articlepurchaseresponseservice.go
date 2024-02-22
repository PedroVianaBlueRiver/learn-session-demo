package main

import (
	"fmt"
	"time"
)

func (ap *ArticlePurchaseResponse) PerformPurchasev4(quantity int) {
	time := time.Now()

	switch {
	case ap.at.stock == 0:
		ap.issuccess = false
		ap.response = "There is no article available in stock"

	case ap.at.stock < quantity:
		ap.issuccess = false
		verb := "is"
		if ap.at.stock > 1 {
			verb = "are"
		}
		ap.response = fmt.Sprintf("There %v just %v(%v) available in stock", verb, ap.at.stock, ap.at.name)

	default:
		ap.at.date = time.Format("2006-01-02")
		ap.at.quantity = quantity
		ap.at.stock = ap.at.stock - quantity
		ap.at.total = ap.at.unitprice * float64(quantity)
	}

}
