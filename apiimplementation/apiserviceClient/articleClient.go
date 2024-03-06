package apiserviceclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
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

func NewApi(url string) *apiUrl {
	return &apiUrl{Url: url}
}

func NewHttpResponseCustom(msnerror error, statusCode int, statusMessage string) *httpResponseCustom {
	return &httpResponseCustom{
		MsnError:      msnerror,
		StatusCode:    statusCode,
		StatusMessage: statusMessage,
	}
}

func GetApi(endpoint string, v any) httpResponseCustom {
	url := endpoint
	var amm = v
	response, err := http.Get(url)
	if err != nil {

		return *NewHttpResponseCustom(err, response.StatusCode, response.Status)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return *NewHttpResponseCustom(err, 0, "")
	}

	err = utilities.DeserializeJson(responseData, &amm)
	if err != nil {
		return *NewHttpResponseCustom(err, response.StatusCode, response.Status)
	}

	return *NewHttpResponseCustom(err, response.StatusCode, response.Status)
}

func PostApi(endpoint string, v any) httpResponseCustom {
	json_data, errorJson := utilities.SerializeJson(v)

	if errorJson != nil {
		return *NewHttpResponseCustom(errorJson, 0, "")
	}

	resp, err := http.Post(endpoint, "application/json",
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

	resp, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(json_data))
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

func DeleteApi(endpoint string) httpResponseCustom {
	client := &http.Client{}

	// create a new DELETE request
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return *NewHttpResponseCustom(err, 0, "")
	}

	// send the request
	resp, err := client.Do(req)
	return *NewHttpResponseCustom(err, resp.StatusCode, resp.Status)

}
