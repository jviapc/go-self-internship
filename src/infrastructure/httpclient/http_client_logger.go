package httpclient

import (
	"log"
	"wiki/parser/parser/httpclient"
)

type logger struct {
	httpClient httpclient.ClientInterface
}

func NewClientLogger(httpClient httpclient.ClientInterface) httpclient.ClientInterface {
	if httpClient == nil {
		panic("httpClient cannot be nil")
	}

	return &logger{httpClient: httpClient}
}

func (that *logger) Get(url string) ([]byte, error) {
	log.Printf("Loading page %s", url)

	return that.httpClient.Get(url)
}
