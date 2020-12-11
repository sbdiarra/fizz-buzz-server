// integration

package handler_test

import (
	"fiz-buzz-server/http/handler"
	"fiz-buzz-server/http/router"
	"fiz-buzz-server/repository/in_memory"
	"fiz-buzz-server/service/fizzbuzz"
	"fiz-buzz-server/service/stats"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFizzBuzzRouteInvalidParam(t *testing.T) {
	path := "/api/compute/fizzbuzz"
	r := getMockHandler()
	server := httptest.NewServer(r)
	defer server.Close()

	tests := []struct {
		name           string
		queryParam     string
		wantStatusCode int
	}{
		{name: "1rst queryParam missing",
			queryParam:     path + "?secM=2&limit=10&label1=label1&label2=label2",
			wantStatusCode: 400,
		},
		{name: "2rst queryParam missing",
			queryParam:     path + "?fstM=3&limit=10&label1=label1&label2=label2",
			wantStatusCode: 400,
		},
		{name: "limit queryParam missing",
			queryParam:     path + "?secM=2&fstM=3&label1=label1&label2=label2",
			wantStatusCode: 400,
		},

		{name: "label1 queryParam missing",
			queryParam:     path + "?secM=2&fstM=3&limit=10&label2=label2",
			wantStatusCode: 400,
		},

		{name: "label2 queryParam missing",
			queryParam:     path + "?secM=2&fstM=3&limit=10&secM=2&fstM=3&label1=label1",
			wantStatusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.Get(server.URL + tt.queryParam)
			assert.Equal(t, tt.wantStatusCode, r.StatusCode)
		})
	}

}

func TestGetFizzBuzzRouteValidParam(t *testing.T) {
	path := "/api/compute/fizzbuzz"
	r := getMockHandler()
	server := httptest.NewServer(r)
	defer server.Close()

	tests := []struct {
		name           string
		queryParam     string
		wantStatusCode int
		wantResponse   string
	}{
		{name: "should return the correct sequence",
			queryParam:     path + "?fstM=2&secM=5&limit=10&label1=fizz&label2=buzz",
			wantStatusCode: 200,
			wantResponse:   "{\"result\":\"1,fizz,3,fizz,buzz,fizz,7,fizz,9,fizzbuzz\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.Get(server.URL + tt.queryParam)
			bodyByte, err := ioutil.ReadAll(r.Body)
			body := string(bodyByte)
			defer r.Body.Close()
			if err != nil {
				t.Fatalf("GET /err = %s; want nil", err)
			}
			t.Log(body)
			assert.Equal(t, tt.wantStatusCode, r.StatusCode)
			assert.Equal(t, tt.wantResponse, body)

		})
	}

}

func TestGetHitsRoute(t *testing.T) {
	hitPath := "/api/metrics/besthits"
	fizbuzzPath := "/api/compute/fizzbuzz?fstM=2&secM=5&limit=10&label1=fizz&label2=buzz"
	r := getMockHandler()
	server := httptest.NewServer(r)
	defer server.Close()

	tests := []struct {
		name           string
		callTofbPath   func()
		wantStatusCode int
		wantResponse   string
	}{
		{name: "return_204_when_no_hit",
			callTofbPath:   func() {},
			wantStatusCode: 204,
			wantResponse:   "",
		},
		{name: "return_200_and_the_most_call_query",
			callTofbPath: func() {
				for i := 1; i <= 4; i++ {
					_, _ = http.Get(server.URL + fizbuzzPath)
				}
			},
			wantStatusCode: 200,
			wantResponse:   "{\"hits\":4,\"associateParams\":\"2, 5, 10, fizz, buzz\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.callTofbPath()
			r, _ := http.Get(server.URL + hitPath)
			bodyByte, err := ioutil.ReadAll(r.Body)
			body := string(bodyByte)
			defer r.Body.Close()
			if err != nil {
				t.Fatalf("GET /err = %s; want nil", err)
			}
			t.Log(body)
			assert.Equal(t, tt.wantStatusCode, r.StatusCode)
			assert.Equal(t, tt.wantResponse, body)

		})
	}
}

func getMockHandler() *gin.Engine {

	fizBuzzService := fizzbuzz.NewFizzBuzzService()
	inMemStore := in_memory.NewInMemoryStore()
	statService := stats.NewStatService(inMemStore)
	fbHandler := handler.NewFizzBuzzHandler(fizBuzzService, statService)
	mockRouter := router.InitRouter(fbHandler)

	return mockRouter
}
