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
		ap.response = fmt.Sprintf("There is no article available in stock")

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

func main() {
	cont := 0
	listPurchase := []struct {
		name      string
		stock     int
		unitprice float64
		quantity  int
	}{
		{
			name:      "TV",
			stock:     3,
			unitprice: 300.50,
			quantity:  4,
		},
		{
			name:      "Phone",
			stock:     3,
			unitprice: 623.50,
			quantity:  2,
		},
	}

	for i := 0; i < len(listPurchase); i++ {
		cont = cont + 1
		fmt.Printf("Purchase #%v ---------------------------------------------------------------------------------------\n", cont)

		apresponse1 := NewArticlePurchaseResponse(listPurchase[i].name, listPurchase[i].stock, listPurchase[i].unitprice)
		apresponse1.PerformPurchasev4(listPurchase[i].quantity)
		if apresponse1.issuccess {
			fmt.Println(apresponse1.response)
			fmt.Printf("Article: %v\n", apresponse1.at.name)
			fmt.Printf("Quantity: %v\n", apresponse1.at.quantity)
			fmt.Printf("Total : $ %2f\n", apresponse1.at.total)
		} else {
			fmt.Println(apresponse1.response)
		}
	}
}
