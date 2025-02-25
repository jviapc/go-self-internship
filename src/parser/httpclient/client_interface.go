package httpclient

type ClientInterface interface {
	Get(url string) ([]byte, error)
}
