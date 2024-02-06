package main

type ArticleModel struct {
	name      string
	stock     int
	unitprice float64
}

type PurchaseModels struct {
	articleId int
	date      string
	quantity  int
	total     float64
}

func NewListMapArticle() map[int]ArticleModel {
	return map[int]ArticleModel{
		1: {
			name:      "TV",
			stock:     4,
			unitprice: 130,
		},
		2: {
			name:      "Phone",
			stock:     5,
			unitprice: 210,
		},
		4: {
			name:      "AirFryer",
			stock:     30,
			unitprice: 190,
		},
	}
}

func NewListMapPruchase(articleId, quantity int, date string, total float64) *PurchaseModels {
	return &PurchaseModels{
		articleId: articleId,
		date:      date,
		quantity:  quantity,
		total:     total,
	}
}
