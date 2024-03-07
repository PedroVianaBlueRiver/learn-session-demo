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
			w.WriteHeader(http.StatusOK)
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
			w.WriteHeader(http.StatusOK)
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

func TestGetArticle500StatusInternalServerError(t *testing.T) {
	jsonResponse := ``

	type testData struct {
		statusexpected     string
		statusCodeexpected int
		modelexpected      apiimplementation.ApiArticleResponse
	}

	expectedData := testData{
		statusCodeexpected: int(constants.StatusInternalServerError),
		statusexpected:     "500 Internal Server Error",
		modelexpected:      apiimplementation.ApiArticleResponse{},
	}

	t.Run("Test GetApi to obtain 500 Internal Server Error", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
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

func TestPost201Created(t *testing.T) {

	modelPayload := *apiimplementation.NewApiArticleResponse("", *apservice.NewArticleModel("Phone", 5, 210))

	type testData struct {
		statusexpected     string
		statusCodeexpected int
	}

	expectedData := testData{
		statusCodeexpected: int(constants.StatusCreated),
		statusexpected:     "201 Created",
	}

	t.Run("Test PostApi for creating an ApiArticleResponse ", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
		}))
		defer svr.Close()
		c := svr.URL

		resp := apiserviceclient.PostApi(c, modelPayload)
		if resp.StatusCode != expectedData.statusCodeexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusCodeexpected, resp.StatusCode)
		}
		if resp.StatusMessage != expectedData.statusexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusexpected, resp.StatusMessage)
		}

	})
}

func TestPost400BadRequest(t *testing.T) {

	modelPayload := *apiimplementation.NewApiArticleResponse("", *apservice.NewArticleModel("Phone", 5, 210))

	type testData struct {
		statusexpected     string
		statusCodeexpected int
	}

	expectedData := testData{
		statusCodeexpected: int(constants.StatusBadRequest),
		statusexpected:     "400 Bad Request",
	}

	t.Run("Test PostApi to obtain 400 Bad Request", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		}))
		defer svr.Close()
		c := svr.URL

		resp := apiserviceclient.PostApi(c, modelPayload)
		if resp.StatusCode != expectedData.statusCodeexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusCodeexpected, resp.StatusCode)
		}
		if resp.StatusMessage != expectedData.statusexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusexpected, resp.StatusMessage)
		}

	})
}

func TestPost500StatusInternalServerError(t *testing.T) {

	modelPayload := *apiimplementation.NewApiArticleResponse("", *apservice.NewArticleModel("Phone", 5, 210))

	type testData struct {
		statusexpected     string
		statusCodeexpected int
	}

	expectedData := testData{
		statusCodeexpected: int(constants.StatusInternalServerError),
		statusexpected:     "500 Internal Server Error",
	}

	t.Run("Test PostApi to obtain 500 Internal Server Error", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer svr.Close()
		c := svr.URL

		resp := apiserviceclient.PostApi(c, modelPayload)
		if resp.StatusCode != expectedData.statusCodeexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusCodeexpected, resp.StatusCode)
		}
		if resp.StatusMessage != expectedData.statusexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusexpected, resp.StatusMessage)
		}

	})
}

func TestPutApi200OK(t *testing.T) {

	modelPayload := *apiimplementation.NewApiArticleResponse("2", *apservice.NewArticleModel("Phone", 5, 210))

	type testData struct {
		statusexpected     string
		statusCodeexpected int
	}

	expectedData := testData{
		statusCodeexpected: int(constants.StatusOK),
		statusexpected:     "200 OK",
	}

	t.Run("Test PutApi for updating an ApiArticleResponse ", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer svr.Close()
		c := svr.URL

		resp := apiserviceclient.PutApi(c, modelPayload)
		if resp.StatusCode != expectedData.statusCodeexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusCodeexpected, resp.StatusCode)
		}
		if resp.StatusMessage != expectedData.statusexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusexpected, resp.StatusMessage)
		}

	})
}

func TestPut400BadRequest(t *testing.T) {

	modelPayload := *apiimplementation.NewApiArticleResponse("2", *apservice.NewArticleModel("Phone", 5, 210))

	type testData struct {
		statusexpected     string
		statusCodeexpected int
	}

	expectedData := testData{
		statusCodeexpected: int(constants.StatusBadRequest),
		statusexpected:     "400 Bad Request",
	}

	t.Run("Test PutApi to obtain 400 Bad Request", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		}))
		defer svr.Close()
		c := svr.URL

		resp := apiserviceclient.PutApi(c, modelPayload)
		if resp.StatusCode != expectedData.statusCodeexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusCodeexpected, resp.StatusCode)
		}
		if resp.StatusMessage != expectedData.statusexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusexpected, resp.StatusMessage)
		}

	})
}

func TestPut500StatusInternalServerError(t *testing.T) {

	modelPayload := *apiimplementation.NewApiArticleResponse("", *apservice.NewArticleModel("Phone", 5, 210))

	type testData struct {
		statusexpected     string
		statusCodeexpected int
	}

	expectedData := testData{
		statusCodeexpected: int(constants.StatusInternalServerError),
		statusexpected:     "500 Internal Server Error",
	}

	t.Run("Test PutApi to obtain 500 Internal Server Error", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer svr.Close()
		c := svr.URL

		resp := apiserviceclient.PutApi(c, modelPayload)
		if resp.StatusCode != expectedData.statusCodeexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusCodeexpected, resp.StatusCode)
		}
		if resp.StatusMessage != expectedData.statusexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusexpected, resp.StatusMessage)
		}

	})
}

func TestDeleteApi400BadRequest(t *testing.T) {

	idEndpint := "2"

	type testData struct {
		statusexpected     string
		statusCodeexpected int
	}

	expectedData := testData{
		statusCodeexpected: int(constants.StatusBadRequest),
		statusexpected:     "400 Bad Request",
	}

	t.Run("Test Delete Api to obtain 400 Bad Request", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		}))
		defer svr.Close()
		c := svr.URL + "/" + idEndpint

		resp := apiserviceclient.DeleteApi(c)
		if resp.StatusCode != expectedData.statusCodeexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusCodeexpected, resp.StatusCode)
		}
		if resp.StatusMessage != expectedData.statusexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusexpected, resp.StatusMessage)
		}

	})
}

func TestDeleteApi500StatusInternalServerError(t *testing.T) {
	idEndpint := "2"

	type testData struct {
		statusexpected     string
		statusCodeexpected int
	}

	expectedData := testData{
		statusCodeexpected: int(constants.StatusInternalServerError),
		statusexpected:     "500 Internal Server Error",
	}

	t.Run("Test DeleteApi to obtain 500 Internal Server Error", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer svr.Close()
		c := svr.URL + "/" + idEndpint

		resp := apiserviceclient.DeleteApi(c)
		if resp.StatusCode != expectedData.statusCodeexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusCodeexpected, resp.StatusCode)
		}
		if resp.StatusMessage != expectedData.statusexpected {
			t.Errorf("Expected result doesn't match with Actual result ..... Expected =  %v ..... Actual %v", expectedData.statusexpected, resp.StatusMessage)
		}

	})
}
