package apiimplementation

import (
	"io/ioutil"
	"net/http"
	"os"
	apiserviceclient "purchasetest/apiimplementation/apiserviceClient"
	"purchasetest/utilities"
)

func GetApi(endpoint string, v any) error {
	url := apiserviceclient.NewApi().Url + endpoint
	response, err := http.Get(url)
	if err != nil {

		os.Exit(1)
		return err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var amm = v
	err = utilities.DeserializeJson(responseData, &amm)
	if err != nil {
		return err
	}
	return nil
}

func GetArticleApi() ([]ApiArticleResponse, error) {
	url := "/articlemap"
	var data []ApiArticleResponse
	error := GetApi(url, &data)

	if error != nil {
		return nil, error
	}

	return data, nil
}

func GetArticleByIdApi(id string) (ApiArticleResponse, error) {
	url := "/articlemap/" + id
	var data ApiArticleResponse
	error := GetApi(url, &data)

	if error != nil {
		return ApiArticleResponse{}, error
	}

	return data, nil
}
