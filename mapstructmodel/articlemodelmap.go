package mapstructmodel

import "purchasetest/apservice"

func NewListMapArticle() map[int]apservice.ArticleModel {

	listMap := make(map[int]apservice.ArticleModel)
	listMap[1] = apservice.ArticleModel{
		Name:      "TV",
		Stock:     4,
		Unitprice: 130,
	}
	listMap[2] = apservice.ArticleModel{
		Name:      "Phone",
		Stock:     5,
		Unitprice: 210,
	}
	listMap[4] = apservice.ArticleModel{
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
