package main

import (
	"fmt"
	"purchasetest/apmodulev1"
	"time"
)

func main() {

	//********************************************************************************************************************
	//  articlepurchaseservice.go
	fmt.Println("*************************************************articlepurchaseservice.go*******************************************************************")
	time1 := time.Now()
	fmt.Println(time1.Format("2006-01-01 "))
	ap := apmodulev1.NewArticlePurchase("TV", 3, 300)
	fmt.Printf("ArticlePurchase struct %v", ap)
	fmt.Println()
	p1 := apmodulev1.PerformPurchase(3, ap)
	fmt.Printf("PerformPurchase() result = %v", p1)
	fmt.Println()

	//method with a receiver argument
	p2 := apmodulev1.NewArticlePurchase("TV", 3, 300)
	fmt.Printf("ArticlePurchase struct %v", p2)
	fmt.Println()
	p2.PerformPurchasev2(2)
	fmt.Printf("PerformPurchasev2() result = %v", p2)
	fmt.Println()

	//********************************************************************************************************************
	//  arrayarticlepurchase.go
	fmt.Println("*************************************************arrayarticlepurchase.go*******************************************************************")
	arrayp := apmodulev1.NewMultipleArticlePurchase()
	arrayp.PerformPurchasev3(3, "Phone")
	fmt.Println(*arrayp)

	//**********************************************************************************************************************
	//  articlepurchaseResponse.go
	fmt.Println("*************************************************articlepurchaseResponse.go*******************************************************************")
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

		apresponse1 := apmodulev1.NewArticlePurchaseResponse(listPurchase[i].name, listPurchase[i].stock, listPurchase[i].unitprice)
		apresponse1.PerformPurchasev4(listPurchase[i].quantity)
		if apresponse1.Issuccess {
			fmt.Println(apresponse1.Response)
			fmt.Printf("Article: %v\n", apresponse1.AT.Name)
			fmt.Printf("Quantity: %v\n", apresponse1.AT.Quantity)
			fmt.Printf("Total : $ %2f\n", apresponse1.AT.Total)
		} else {
			fmt.Println(apresponse1.Response)
		}
	}
}
