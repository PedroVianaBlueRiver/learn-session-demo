package main

import "fmt"

// GetItem returns the quantity of item
func GetItem(item map[int]ArticleModel, id int) (ArticleModel, bool, string) {
	art := ArticleModel{}
	if _, ok := item[id]; !ok {
		return art, false, fmt.Sprintf("article with id: %v doesn't exist", id)
	}

	return item[id], true, ""
}

func AddItem(item map[int]ArticleModel, stock, id int, name string, unitprice float64) map[int]ArticleModel {
	at := ArticleModel{name: name, unitprice: unitprice, stock: stock}
	item[id] = at

	return item
}

func DeleteItem(item map[int]ArticleModel, id int) _(map[int]ArticleModel, bool, string) {
	if _, ok := item[id]; !ok {
		return item, false, fmt.Sprintf("article with id: %v doesn't exist", id)
	}
	delete(item, id)
	return item, true, ""
}

func main() {
	articleList := NewListMapArticle()
	fmt.Println("GetItem(1) Initial Map: ", articleList)
	at, ok, msn := GetItem(articleList, 1)
	if !ok {
		fmt.Println(msn)
	} else {
		fmt.Println("Article name: ", at.name)
		fmt.Println("Article stock: ", at.stock)
		fmt.Println("Article price: ", at.unitprice)
	}

	articleList = AddItem(articleList, 400, 6, "CPU", 1300)
	fmt.Println("AddItem() Updated Map: ", articleList)

	
}
