package apiimplementation

import "purchasetest/apservice"

type ApiArticleResponse struct {
	Id      string                 `json:"id"`
	Article apservice.ArticleModel `json:"article"`
}
