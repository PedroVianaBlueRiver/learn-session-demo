package main

import (
	"fmt"
	"time"
)

// GetItem returns the quantity of item
func GetItem(item map[int]ArticleModel, id int) (ArticleModel, bool, string) {
	art := ArticleModel{}
	if _, ok := item[id]; !ok {
		return art, false, fmt.Sprintf("article with id: %v doesn't exist", id)
	}

	return item[id], true, ""
}

// Add Item of the map
func AddItem(item map[int]ArticleModel, stock, id int, name string, unitprice float64) (map[int]ArticleModel, bool, string) {
	if _, ok := item[id]; ok {
		return item, false, fmt.Sprintf("article with id: %v already exists", id)
	}
	at := ArticleModel{Name: name, Unitprice: unitprice, Stock: stock}
	item[id] = at

	return item, true, ""
}

// Delete item of the map
func DeleteItem(item map[int]ArticleModel, id int) (map[int]ArticleModel, bool, string) {
	if _, ok := item[id]; !ok {
		return item, false, fmt.Sprintf("article with id: %v doesn't exist", id)
	}
	delete(item, id)
	return item, true, ""
}

// update item of the map
func updateItem(item map[int]ArticleModel, id, stock int, name string, unitprice float64) (map[int]ArticleModel, bool, string) {
	if _, ok := item[id]; !ok {
		return item, false, fmt.Sprintf("article with id: %v doesn't exist", id)
	}

	at := item[id]

	at.Stock = stock
	at.Name = name
	at.Unitprice = unitprice
	item[id] = at

	return item, true, ""
}

// create a purchaseModel and update the map[int]ArticleModel after performing a purchase
func (apr *APModelRespose) createPurchase(id, quantity int) (bool, string) {
	time := time.Now()
	lstAM := apr.Nat
	fmt.Println("Initial Map Article: ", lstAM)

	//Get item by Id
	itemAM, ok, msn := GetItem(lstAM, id)

	if !ok {
		return false, msn
	}

	//Validate existing stock
	okvalidstock, msnresponse := itemAM.validateStock(quantity)
	if !okvalidstock {
		return false, msnresponse
	}
	itemAM.Stock = itemAM.Stock - quantity

	//update the item after the purchase
	totalpurchase := itemAM.Unitprice * float64(quantity)
	lstAM, ok1, msn1 := updateItem(lstAM, id, itemAM.Stock, itemAM.Name, itemAM.Unitprice)

	if !ok1 {
		return false, msn1
	}

	//Assign new values
	apr.Npm = NewListMapPruchase(id, quantity, time.Format("2006-01-02"), totalpurchase)
	apr.Nat = lstAM
	return true, ""

}

func createPurchase(id, quantity int, ai APInterface) (bool, string) {
	return ai.createPurchase(id, quantity)
}

func main() {
	articleList := NewListMapArticle()
	fmt.Println("************************************************* GetItem() *******************************************************************")
	fmt.Println("GetItem(1) Initial Map: ", articleList)
	at, ok, msn := GetItem(articleList, 1)
	if !ok {
		fmt.Println(msn)
	} else {
		fmt.Println("Article name: ", at.Name)
		fmt.Println("Article stock: ", at.Stock)
		fmt.Println("Article price: ", at.Unitprice)
	}
	fmt.Println("************************************************* AddItem() *******************************************************************")
	articleList, ok, msn = AddItem(articleList, 400, 6, "CPU", 1300)
	if !ok {
		fmt.Println(msn)
	} else {
		fmt.Println("AddItem() Updated Map: ", articleList)
	}

	fmt.Println("************************************************* updateItem() *******************************************************************")
	articleList, ok, msn = updateItem(articleList, 2, 20, "Phone", 310)

	if !ok {
		fmt.Println(msn)
	} else {
		fmt.Println("updateItem() Updated Map: ", articleList)
	}

	fmt.Println("************************************************* DeleteItem() *******************************************************************")
	articleList, ok, msn = DeleteItem(articleList, 1)
	if !ok {
		fmt.Println(msn)
	} else {
		fmt.Println("DeleteItem() Updated Map: ", articleList)
	}

	fmt.Println("************************************************* createPurchase() *******************************************************************")
	response := NewAPModelRespose(NewListMapArticle(), PurchaseModels{})
	ok2, msn2 := createPurchase(1, 1, response)
	if !ok2 {
		fmt.Println(msn2)
	} else {
		fmt.Println("Updated Map: ", response.Nat)
		fmt.Println("purchase: ", response.Npm)
	}

	fmt.Println("************************************************* Json() *******************************************************************")

	datosJson, err := serializeJson(response.Nat)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(datosJson))

	jsonString := `{"1":{"name":"TV","stock":3,"unitprice":130},"2":{"name":"Phone","stock":5,"unitprice":210},"4":{"name":"AirFryer","stock":30,"unitprice":190}}`
	var amm map[int]ArticleModel
	err2 := deserializeJson([]byte(jsonString), &amm)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("%+v\n", amm)

}
