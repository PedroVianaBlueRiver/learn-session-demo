package mapstructmodel

import "purchasetest/structmodel"

func NewListMapArticle() map[int]structmodel.ArticleModel {

	listMap := make(map[int]structmodel.ArticleModel)
	listMap[1] = structmodel.ArticleModel{
		Name:      "TV",
		Stock:     4,
		Unitprice: 130,
	}
	listMap[2] = structmodel.ArticleModel{
		Name:      "Phone",
		Stock:     5,
		Unitprice: 210,
	}
	listMap[4] = structmodel.ArticleModel{
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
