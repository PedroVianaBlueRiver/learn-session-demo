package apiimplementation

import apiserviceclient "purchasetest/apiimplementation/apiserviceClient"

const sUrl string = "https://65e0b4ced3db23f76249e825.mockapi.io/"

func GetArticleApi() ([]ApiArticleResponse, any) {
	url := apiserviceclient.NewApi(sUrl).Url + "/articlemap"
	var data []ApiArticleResponse
	response := apiserviceclient.GetApi(url, &data)

	return data, response
}

func GetArticleByIdApi(id string) (ApiArticleResponse, any) {
	url := apiserviceclient.NewApi(sUrl).Url + "/articlemap/" + id
	var data ApiArticleResponse
	response := apiserviceclient.GetApi(url, &data)

	return data, response
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
