package main

import (
	"fmt"
	"purchasetest/apservice"
	"purchasetest/mapstructmodel"
)

func main() {
	articleList := mapstructmodel.NewListMapArticle()
	fmt.Println("************************************************* GetItem() *******************************************************************")
	fmt.Println("GetItem(1) Initial Map: ", articleList)
	at, ok, msn := apservice.GetItem(articleList, 1)
	if !ok {
		fmt.Println(msn)
	} else {
		fmt.Println("Article name: ", at.Name)
		fmt.Println("Article stock: ", at.Stock)
		fmt.Println("Article price: ", at.Unitprice)
	}
	fmt.Println("************************************************* AddItem() *******************************************************************")
	articleList, ok, msn = apservice.AddItem(articleList, 400, 6, "CPU", 1300)
	if !ok {
		fmt.Println(msn)
	} else {
		fmt.Println("AddItem() Updated Map: ", articleList)
	}

	fmt.Println("************************************************* updateItem() *******************************************************************")
	articleList, ok, msn = apservice.UpdateItem(articleList, 2, 20, "Phone", 310)

	if !ok {
		fmt.Println(msn)
	} else {
		fmt.Println("updateItem() Updated Map: ", articleList)
	}

	fmt.Println("************************************************* DeleteItem() *******************************************************************")
	articleList, ok, msn = apservice.DeleteItem(articleList, 1)
	if !ok {
		fmt.Println(msn)
	} else {
		fmt.Println("DeleteItem() Updated Map: ", articleList)
	}

	fmt.Println("************************************************* createPurchase() *******************************************************************")
	response := apservice.NewAPModelRespose(mapstructmodel.NewListMapArticle(), apservice.PurchaseModels{})
	ok2, msn2 := apservice.CreatePurchase(1, 1, response)
	if !ok2 {
		fmt.Println(msn2)
	} else {
		fmt.Println("Updated Map: ", response.Nat)
		fmt.Println("purchase: ", response.Npm)
	}

	fmt.Println("************************************************* Json() *******************************************************************")

	// 	datosJson, err := serializeJson(response.Nat)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(string(datosJson))

	// 	jsonString := `{"1":{"name":"TV","stock":3,"unitprice":130},"2":{"name":"Phone","stock":5,"unitprice":210},"4":{"name":"AirFryer","stock":30,"unitprice":190}}`
	// 	var amm map[int]ArticleModel
	// 	err2 := deserializeJson([]byte(jsonString), &amm)
	// 	if err2 != nil {
	// 		fmt.Println(err2)
	// 	}
	// 	fmt.Printf("%+v\n", amm)

	// }
}

// func main() {

// 	//********************************************************************************************************************
// 	//  articlepurchaseservice.go
// 	fmt.Println("*************************************************articlepurchaseservice.go*******************************************************************")
// 	time1 := time.Now()
// 	fmt.Println(time1.Format("2006-01-01 "))
// 	ap := NewArticlePurchase("TV", 3, 300)
// 	fmt.Printf("ArticlePurchase struct %v", ap)
// 	fmt.Println()
// 	p1 := PerformPurchase(3, ap)
// 	fmt.Printf("PerformPurchase() result = %v", p1)
// 	fmt.Println()

// 	//method with a receiver argument
// 	p2 := NewArticlePurchase("TV", 3, 300)
// 	fmt.Printf("ArticlePurchase struct %v", p2)
// 	fmt.Println()
// 	p2.PerformPurchasev2(2)
// 	fmt.Printf("PerformPurchasev2() result = %v", p2)
// 	fmt.Println()

// 	//********************************************************************************************************************
// 	//  arrayarticlepurchase.go
// 	fmt.Println("*************************************************arrayarticlepurchase.go*******************************************************************")
// 	arrayp := NewMultipleArticlePurchase()
// 	arrayp.PerformPurchasev3(3, "Phone")
// 	fmt.Println(*arrayp)

// 	//**********************************************************************************************************************
// 	//  articlepurchaseResponse.go
// 	fmt.Println("*************************************************articlepurchaseResponse.go*******************************************************************")
// 	cont := 0
// 	listPurchase := []struct {
// 		name      string
// 		stock     int
// 		unitprice float64
// 		quantity  int
// 	}{
// 		{
// 			name:      "TV",
// 			stock:     3,
// 			unitprice: 300.50,
// 			quantity:  4,
// 		},
// 		{
// 			name:      "Phone",
// 			stock:     3,
// 			unitprice: 623.50,
// 			quantity:  2,
// 		},
// 	}

// 	for i := 0; i < len(listPurchase); i++ {
// 		cont = cont + 1
// 		fmt.Printf("Purchase #%v ---------------------------------------------------------------------------------------\n", cont)

// 		apresponse1 := NewArticlePurchaseResponse(listPurchase[i].name, listPurchase[i].stock, listPurchase[i].unitprice)
// 		apresponse1.PerformPurchasev4(listPurchase[i].quantity)
// 		if apresponse1.issuccess {
// 			fmt.Println(apresponse1.response)
// 			fmt.Printf("Article: %v\n", apresponse1.at.name)
// 			fmt.Printf("Quantity: %v\n", apresponse1.at.quantity)
// 			fmt.Printf("Total : $ %2f\n", apresponse1.at.total)
// 		} else {
// 			fmt.Println(apresponse1.response)
// 		}
// 	}
// }
