package apiimplementation

import "purchasetest/apservice"

type ApiArticleResponse struct {
	Id      string                 `json:"id"`
	Article apservice.ArticleModel `json:"article"`
}

func NewApiArticleResponse(id string, article apservice.ArticleModel) *ApiArticleResponse {
	return &ApiArticleResponse{
		Id:      id,
		Article: article,
	}
}
