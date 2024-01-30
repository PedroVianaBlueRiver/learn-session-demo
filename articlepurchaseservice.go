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
		ap.response = fmt.Sprintln("There is no article available in stock")

	case ap.at.quantity < quantity:
		ap.issuccess = false
		ap.response = fmt.Sprintf("There is just %v available in stock", ap.at.stock)

	default:
		ap.at.date = time.Format("2006-01-02")
		ap.at.quantity = quantity
		ap.at.stock = ap.at.stock - quantity
		ap.at.total = ap.at.unitprice * float64(quantity)
	}

}

func main() {
	apresponse := NewArticlePurchaseResponse("TV", 3, 300)
	apresponse.PerformPurchasev4(4)
	fmt.Print(apresponse.response)
}
