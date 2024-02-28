package main

import (
	"fmt"
	"purchasetest/apservice"
	"purchasetest/mapstructmodel"
	"purchasetest/utilities"
)

func main() {

	fmt.Println("************************************************* SerializeJson() *******************************************************************")
	maparticle := mapstructmodel.NewListMapArticle()
	datosJson, err := utilities.SerializeJson(maparticle)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(datosJson))

	fmt.Println("************************************************* DeserializeJson() *******************************************************************")
	jsonString := `{"1":{"name":"TV","stock":3,"unitprice":130},"2":{"name":"Phone","stock":5,"unitprice":210},"4":{"name":"AirFryer","stock":30,"unitprice":190}}`
	var amm map[int]apservice.ArticleModel
	err2 := utilities.DeserializeJson([]byte(jsonString), &amm)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("%+v\n", amm)

}
