package main

import (
	"fmt"
	"purchasetest/apiimplementation"
	"purchasetest/apservice"
)

func main() {
	// fmt.Println(apiimplementation.GetArticleApi())
	// fmt.Println(apiimplementation.GetArticleByIdApi("4"))

	at := apiimplementation.NewApiArticleResponse("", *apservice.NewArticleModel("Bed Srping air", 20, 400))
	postresponse := apiimplementation.PostCreateArticleApi(at)
	fmt.Println(postresponse)

	// Putat := apiimplementation.NewApiArticleResponse("9", *apservice.NewArticleModel("Nike Jersey England", 24, 802))
	// putresponse := apiimplementation.PutUpdateArticleApi(Putat)
	// fmt.Println(putresponse)

	// deleteAt := apiimplementation.DeleteArticleApi("12")
	// fmt.Println(deleteAt)
}
