package main

import "fmt"

type ArticleModel struct {
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	Unitprice float64 `json:"unitprice"`
}

type PurchaseModels struct {
	ArticleId int     `json:"articleId"`
	Date      string  `json:"date"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}

type APModelRespose struct {
	Nat map[int]ArticleModel `json:"articlemodel"`
	Npm *PurchaseModels      `json:"purchasemodel"`
}

func NewListMapArticle() map[int]ArticleModel {

	listMap := make(map[int]ArticleModel)
	listMap[1] = ArticleModel{
		Name:      "TV",
		Stock:     4,
		Unitprice: 130,
	}
	listMap[2] = ArticleModel{
		Name:      "Phone",
		Stock:     5,
		Unitprice: 210,
	}
	listMap[4] = ArticleModel{
		Name:      "AirFryer",
		Stock:     30,
		Unitprice: 190,
	}

	return listMap
	// return map[int]ArticleModel{
	// 	1: {
	// 		name:      "TV",
	// 		stock:     4,
	// 		unitprice: 130,
	// 	},
	// 	2: {
	// 		name:      "Phone",
	// 		stock:     5,
	// 		unitprice: 210,
	// 	},
	// 	4: {
	// 		name:      "AirFryer",
	// 		stock:     30,
	// 		unitprice: 190,
	// 	},
	// }
}

func NewAPModelRespose(at map[int]ArticleModel, pm PurchaseModels) *APModelRespose {
	return &APModelRespose{
		Nat: at,
		Npm: &pm,
	}
}

func NewListMapPruchase(articleId, quantity int, date string, total float64) *PurchaseModels {
	return &PurchaseModels{
		ArticleId: articleId,
		Date:      date,
		Quantity:  quantity,
		Total:     total,
	}
}

func (am ArticleModel) validateStock(quantity int) (bool, string) {
	var issuccess bool
	var response string
	switch {
	case am.Stock == 0:
		issuccess = false
		response = "There is no article available in stock"

	case am.Stock < quantity:
		issuccess = false
		verb := "is"
		if am.Stock > 1 {
			verb = "are"
		}
		response = fmt.Sprintf("There %v just %v(%v) available in stock", verb, am.Stock, am.Name)

	default:
		issuccess = true
		response = ""
	}

	return issuccess, response
}
