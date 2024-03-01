package apiserviceclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"purchasetest/utilities"
)

type apiUrl struct {
	Url string
}

type httpResponseCustom struct {
	MsnError      error
	StatusCode    int
	StatusMessage string
}

func NewApi() *apiUrl {
	return &apiUrl{Url: "https://65e0b4ced3db23f76249e825.mockapi.io/"}
}

func NewHttpResponseCustom(msnerror error, statusCode int, statusMessage string) *httpResponseCustom {
	return &httpResponseCustom{MsnError: msnerror,
		StatusCode:    statusCode,
		StatusMessage: statusMessage}
}

func GetApi(endpoint string, v any) error {
	url := NewApi().Url + endpoint
	response, err := http.Get(url)
	if err != nil {

		os.Exit(1)
		return err
	}

	defer response.Body.Close()

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

func PostApi(endpoint string, v any) httpResponseCustom {
	json_data, errorJson := utilities.SerializeJson(v)

	if errorJson != nil {
		return *NewHttpResponseCustom(errorJson, 0, "")
	}

	resp, err := http.Post(NewApi().Url+endpoint, "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		return *NewHttpResponseCustom(err, resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()

	return *NewHttpResponseCustom(err, resp.StatusCode, resp.Status)

}

func PutApi(endpoint string, Model any) httpResponseCustom {

	json_data, errorJson := utilities.SerializeJson(Model)
	if errorJson != nil {
		return *NewHttpResponseCustom(errorJson, 0, "")
	}

	resp, err := http.NewRequest(http.MethodPut, NewApi().Url+endpoint, bytes.NewBuffer(json_data))
	if err != nil {
		return *NewHttpResponseCustom(err, resp.Response.StatusCode, resp.Response.Status)
	}

	resp.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp1, err1 := client.Do(resp)
	if err1 != nil {
		return *NewHttpResponseCustom(err1, resp.Response.StatusCode, resp.Response.Status)
	}
	defer resp1.Body.Close()
	return *NewHttpResponseCustom(err1, resp1.StatusCode, resp1.Status)
}
