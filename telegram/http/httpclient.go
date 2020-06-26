package http

import (
	"net/http"
	"net/url"
)

type HttpPoster interface {
	PostForm(url string, values url.Values) (*http.Response, error)
}

type HttpClient struct{}

func (h *HttpClient) PostForm(url string, values url.Values) (*http.Response, error) {
	return http.PostForm(url, values)
}

func GetHttpClient() *HttpClient {
	return &HttpClient{}
}
