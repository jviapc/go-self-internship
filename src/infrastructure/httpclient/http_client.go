package httpclient

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
}

func (h *Client) Get(url string) ([]byte, error) {
	if len(url) <= 0 {
		return []byte{}, fmt.Errorf("httpclient client error: url must not be empty")
	}

	res, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("httpclient client error: (net/httpclient.Get) %v", err)
	}
	defer res.Body.Close() //todo: handle

	if res.StatusCode != 200 {
		return []byte{}, fmt.Errorf("httpclient client error: unexpected status code %d", res.StatusCode)
	}

	return io.ReadAll(res.Body)
}
