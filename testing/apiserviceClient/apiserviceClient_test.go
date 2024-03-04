package apiserviceclient_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"purchasetest/apiimplementation"
	apiserviceclient "purchasetest/apiimplementation/apiserviceClient"
	"purchasetest/apservice"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArticle(t *testing.T) {
	jsonResponse := `[{"article":{"name":"TV","stock":3,"unitprice":130},"id":"1"},
	{"article":{"name":"Phone","stock":5,"unitprice":210},"id":"2"}]`

	dataexpected := []apiimplementation.ApiArticleResponse{
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
	}

	t.Run("Test GetApi to retrieve a list of ApiArticleResponse", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, jsonResponse)
		}))
		defer svr.Close()
		c := svr.URL
		var data []apiimplementation.ApiArticleResponse = []apiimplementation.ApiArticleResponse{{Id: "", Article: apservice.ArticleModel{}}}

		err := apiserviceclient.GetApi(c, &data)
		if err != nil {
			t.Errorf("expected err to be nil got %v", err)
		}

		assert.Equal(t, dataexpected, data)
		if data == nil {
			t.Errorf("expected err to be nil got %v", data)
		}
	})

}
