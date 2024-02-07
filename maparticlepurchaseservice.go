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
	at := ArticleModel{name: name, unitprice: unitprice, stock: stock}
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

	at.stock = stock
	at.name = name
	at.unitprice = unitprice
	item[id] = at

	return item, true, ""
}

// create a purchaseModel and update the map[int]ArticleModel after performing a purchase
func createPurchase(id, quantity int) (map[int]ArticleModel, PurchaseModels, bool, string) {
	time := time.Now()
	lstAM := NewListMapArticle()
	fmt.Println("Initial Map Article: ", lstAM)

	//Get item by Id
	itemAM, ok, msn := GetItem(lstAM, id)

	if !ok {
		return map[int]ArticleModel{}, PurchaseModels{}, false, msn
	}

	//Validate existing stock
	okvalidstock, msnresponse := itemAM.validateStock(quantity)
	itemAM.stock = itemAM.stock - quantity
	if !okvalidstock {
		return map[int]ArticleModel{}, PurchaseModels{}, false, msnresponse
	}

	//update the item after the purchase
	totalpurchase := itemAM.unitprice * float64(quantity)
	lstAM, ok1, msn1 := updateItem(lstAM, id, itemAM.stock, itemAM.name, itemAM.unitprice)

	if !ok1 {
		return map[int]ArticleModel{}, PurchaseModels{}, false, msn1
	}

	purchasemodel := NewListMapPruchase(id, quantity, time.Format("2006-01-02"), totalpurchase)

	return lstAM, *purchasemodel, true, ""

}

func main() {
	articleList := NewListMapArticle()
	fmt.Println("************************************************* GetItem() *******************************************************************")
	fmt.Println("GetItem(1) Initial Map: ", articleList)
	at, ok, msn := GetItem(articleList, 1)
	if !ok {
		fmt.Println(msn)
	} else {
		fmt.Println("Article name: ", at.name)
		fmt.Println("Article stock: ", at.stock)
		fmt.Println("Article price: ", at.unitprice)
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
	mapresponse, purchaseresponse, ok2, msn2 := createPurchase(1, 5)
	if !ok2 {
		fmt.Println(msn2)
	} else {
		fmt.Println("Updated Map: ", mapresponse)
		fmt.Println("purchase: ", purchaseresponse)
	}
}
