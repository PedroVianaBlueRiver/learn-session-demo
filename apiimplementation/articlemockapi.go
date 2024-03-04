package apiimplementation

import apiserviceclient "purchasetest/apiimplementation/apiserviceClient"

const sUrl string = "https://65e0b4ced3db23f76249e825.mockapi.io/"

func GetArticleApi() ([]ApiArticleResponse, error) {
	url := apiserviceclient.NewApi(sUrl).Url + "/articlemap"
	var data []ApiArticleResponse
	error := apiserviceclient.GetApi(url, &data)

	if error != nil {
		return nil, error
	}

	return data, nil
}

func GetArticleByIdApi(id string) (ApiArticleResponse, error) {
	url := apiserviceclient.NewApi(sUrl).Url + "/articlemap/" + id
	var data ApiArticleResponse
	error := apiserviceclient.GetApi(url, &data)

	if error != nil {
		return ApiArticleResponse{}, error
	}

	return data, nil
}

func PostCreateArticleApi(article *ApiArticleResponse) any {
	url := apiserviceclient.NewApi(sUrl).Url + "/articlemap/"
	response := apiserviceclient.PostApi(url, article)

	return response
}

func PutUpdateArticleApi(article *ApiArticleResponse) any {
	url := apiserviceclient.NewApi(sUrl).Url + "/articlemap/" + article.Id
	response := apiserviceclient.PutApi(url, article)

	return response
}

func DeleteArticleApi(id string) any {
	url := apiserviceclient.NewApi(sUrl).Url + "/articlemap/" + id
	response := apiserviceclient.DeleteApi(url)

	return response
}
