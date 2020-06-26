package fakes

import (
	"net/http"
	"net/url"
)

type FakeHttpClient struct {
	PostFormCounter int
}

func (h *FakeHttpClient) PostForm(url string, values url.Values) (*http.Response, error) {
	h.PostFormCounter++
	return nil, nil
}

func GetFakeHttpClient() *FakeHttpClient {
	return &FakeHttpClient{}
}
