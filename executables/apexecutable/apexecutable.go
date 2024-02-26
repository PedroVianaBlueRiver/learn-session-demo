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
