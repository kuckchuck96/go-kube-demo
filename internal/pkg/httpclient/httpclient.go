package httpclient

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

type (
	HttpClient interface {
		Make(HttpRequest) ([]byte, error)
	}

	httpClient struct {
		apiClient *http.Client
	}

	HttpRequest struct {
		Method  string
		URL     string
		Header  http.Header
		Params  map[string]string
		Body    any
		Timeout time.Duration
	}
)

func New() HttpClient {
	return &httpClient{
		apiClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: true,
			},
		},
	}
}

func (client *httpClient) Make(httpRequest HttpRequest) ([]byte, error) {
	req, err := buildHttpRequest(httpRequest)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), httpRequest.Timeout)
	defer cancel()

	res, err := client.apiClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, errors.WithMessage(err, "Error doing http request")
	}
	defer res.Body.Close()

	return buildHttpResponse(res)
}

func buildHttpRequest(httpRequest HttpRequest) (*http.Request, error) {
	url, err := url.Parse(httpRequest.URL)
	if err != nil {
		return nil, errors.WithMessage(err, "Error building http url")
	}

	if httpRequest.Params != nil || len(httpRequest.Params) > 0 {
		query := url.Query()
		for k, v := range httpRequest.Params {
			query.Set(k, v)
		}
		url.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(httpRequest.Method, url.String(), nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Error creating request with context")
	}
	return req, nil
}

func buildHttpResponse(res *http.Response) ([]byte, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "Error reading response body")
	}
	return body, nil
}
