package apiimplementation

import apiserviceclient "purchasetest/apiimplementation/apiserviceClient"

func GetArticleApi() ([]ApiArticleResponse, error) {
	url := "/articlemap"
	var data []ApiArticleResponse
	error := apiserviceclient.GetApi(url, &data)

	if error != nil {
		return nil, error
	}

	return data, nil
}

func GetArticleByIdApi(id string) (ApiArticleResponse, error) {
	url := "/articlemap/" + id
	var data ApiArticleResponse
	error := apiserviceclient.GetApi(url, &data)

	if error != nil {
		return ApiArticleResponse{}, error
	}

	return data, nil
}

func PostCreateArticleApi(article *ApiArticleResponse) any {
	url := "/articlemap/"
	response := apiserviceclient.PostApi(url, article)

	return response
}

func PutUpdateArticleApi(article *ApiArticleResponse) any {
	url := "/articlemap/" + article.Id
	response := apiserviceclient.PutApi(url, article)

	return response
}

func DeleteArticleApi(id string) any {
	url := "/articlemap/" + id
	response := apiserviceclient.DeleteApi(url)

	return response
}
