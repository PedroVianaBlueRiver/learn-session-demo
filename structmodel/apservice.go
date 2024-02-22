package structmodel

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
func UpdateItem(item map[int]ArticleModel, id, stock int, name string, unitprice float64) (map[int]ArticleModel, bool, string) {
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
func (apr *APModelRespose) CreatePurchase(id, quantity int) (bool, string) {
	time := time.Now()
	lstAM := apr.Nat
	fmt.Println("Initial Map Article: ", lstAM)

	//Get item by Id
	itemAM, ok, msn := GetItem(lstAM, id)

	if !ok {
		return false, msn
	}

	//Validate existing stock
	okvalidstock, msnresponse := itemAM.ValidateStock(quantity)
	if !okvalidstock {
		return false, msnresponse
	}
	itemAM.Stock = itemAM.Stock - quantity

	//update the item after the purchase
	totalpurchase := itemAM.Unitprice * float64(quantity)
	lstAM, ok1, msn1 := UpdateItem(lstAM, id, itemAM.Stock, itemAM.Name, itemAM.Unitprice)

	if !ok1 {
		return false, msn1
	}

	//Assign new values
	apr.Npm = NewListMapPruchase(id, quantity, time.Format("2006-01-02"), totalpurchase)
	apr.Nat = lstAM
	return true, ""

}

func CreatePurchase(id, quantity int, ai APInterface) (bool, string) {
	return ai.CreatePurchase(id, quantity)
}
