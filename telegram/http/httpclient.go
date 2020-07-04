package http

import (
	"net/http"
	"net/url"
)

type HttpPoster interface {
	PostForm(url string, values url.Values) (*http.Response, error)
	Get(url string) (*http.Response, error)
}

type HttpClient struct{}

func (h *HttpClient) PostForm(url string, values url.Values) (*http.Response, error) {
	return http.PostForm(url, values)
}

func (h *HttpClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

func GetHttpClient() *HttpClient {
	return &HttpClient{}
}
