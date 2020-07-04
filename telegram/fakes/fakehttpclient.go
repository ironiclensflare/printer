package fakes

import (
	"net/http"
	"net/url"
)

type FakeHttpClient struct {
	PostFormCounter int
	GetCounter      int
}

func (h *FakeHttpClient) PostForm(url string, values url.Values) (*http.Response, error) {
	h.PostFormCounter++
	return nil, nil
}

func (h *FakeHttpClient) Get(url string) (*http.Response, error) {
	h.GetCounter++
	return nil, nil
}

func GetFakeHttpClient() *FakeHttpClient {
	return &FakeHttpClient{}
}
