package apiserviceclient_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	constants "purchasetest/Constants"
	"purchasetest/apiimplementation"
	apiserviceclient "purchasetest/apiimplementation/apiserviceClient"
	"purchasetest/apservice"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArticle(t *testing.T) {
	jsonResponse := `[{"article":{"name":"TV","stock":3,"unitprice":130},"id":"1"},
	{"article":{"name":"Phone","stock":5,"unitprice":210},"id":"2"}]`

	type testData struct {
		statusexpected     string
		statusCodeexpected int
		modelexpected      []apiimplementation.ApiArticleResponse
	}

	expectedData := testData{
		statusCodeexpected: int(constants.StatusOK),
		statusexpected:     "200 OK",
		modelexpected: []apiimplementation.ApiArticleResponse{
			{
				Id: "1",
				Article: apservice.ArticleModel{
					Name:      "TV",
					Stock:     3,
					Unitprice: 130,
				},
			},
			{
				Id: "2",
				Article: apservice.ArticleModel{
					Name:      "Phone",
					Stock:     5,
					Unitprice: 210,
				},
			},
		},
	}

	t.Run("Test GetApi to retrieve a list of ApiArticleResponse", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, jsonResponse)
		}))
		defer svr.Close()
		c := svr.URL
		var data []apiimplementation.ApiArticleResponse

		resp := apiserviceclient.GetApi(c, &data)
		if resp.StatusCode != expectedData.statusCodeexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusCodeexpected, resp.StatusCode)
		}
		if resp.StatusMessage != expectedData.statusexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusexpected, resp.StatusMessage)
		}
		if data == nil {
			t.Errorf("data is nil")
		}

		assert.Equal(t, expectedData.modelexpected, data)

	})

}

func TestGetArticleById(t *testing.T) {
	jsonResponse := `{
		"article": {
			"name": "Phone",
			"stock": 5,
			"unitprice": 210
		},
		"id": "2"
	}`

	type testData struct {
		statusexpected     string
		statusCodeexpected int
		modelexpected      apiimplementation.ApiArticleResponse
	}

	expectedData := testData{
		statusCodeexpected: int(constants.StatusOK),
		statusexpected:     "200 OK",
		modelexpected:      *apiimplementation.NewApiArticleResponse("2", *apservice.NewArticleModel("Phone", 5, 210)),
	}

	t.Run("Test GetApi By Id to retrieve a single result", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, jsonResponse)
		}))
		defer svr.Close()
		c := svr.URL
		var data apiimplementation.ApiArticleResponse

		resp := apiserviceclient.GetApi(c, &data)
		if resp.StatusCode != expectedData.statusCodeexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusCodeexpected, resp.StatusCode)
		}
		if resp.StatusMessage != expectedData.statusexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusexpected, resp.StatusMessage)
		}

		assert.Equal(t, expectedData.modelexpected, data)

	})

}

func TestGetArticle404NotFound(t *testing.T) {
	jsonResponse := ``

	type testData struct {
		statusexpected     string
		statusCodeexpected int
		modelexpected      apiimplementation.ApiArticleResponse
	}

	expectedData := testData{
		statusCodeexpected: int(constants.StatusNotFound),
		statusexpected:     "404 Not Found",
		modelexpected:      apiimplementation.ApiArticleResponse{},
	}

	t.Run("Test GetApi to obtain 404 not found", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, jsonResponse)
		}))
		defer svr.Close()
		c := svr.URL
		var data apiimplementation.ApiArticleResponse

		resp := apiserviceclient.GetApi(c, &data)
		if resp.StatusCode != expectedData.statusCodeexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusCodeexpected, resp.StatusCode)
		}
		if resp.StatusMessage != expectedData.statusexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusexpected, resp.StatusMessage)
		}

		assert.Equal(t, expectedData.modelexpected, data)

	})

}

func TestGetArticle400BadRequest(t *testing.T) {
	jsonResponse := ``

	type testData struct {
		statusexpected     string
		statusCodeexpected int
		modelexpected      apiimplementation.ApiArticleResponse
	}

	expectedData := testData{
		statusCodeexpected: int(constants.StatusBadRequest),
		statusexpected:     "400 Bad Request",
		modelexpected:      apiimplementation.ApiArticleResponse{},
	}

	t.Run("Test GetApi to obtain 400 Bad Request", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, jsonResponse)
		}))
		defer svr.Close()
		c := svr.URL
		var data apiimplementation.ApiArticleResponse

		resp := apiserviceclient.GetApi(c, &data)
		if resp.StatusCode != expectedData.statusCodeexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusCodeexpected, resp.StatusCode)
		}
		if resp.StatusMessage != expectedData.statusexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusexpected, resp.StatusMessage)
		}

		assert.Equal(t, expectedData.modelexpected, data)

	})

}
