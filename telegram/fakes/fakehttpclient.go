package fakes

import (
	"net/http"
	"net/url"
)

type FakeHttpClient struct {
	PostFormCounter     int
	GetCounter          int
	DownloadFileCounter int
}

func (h *FakeHttpClient) PostForm(url string, values url.Values) (*http.Response, error) {
	h.PostFormCounter++
	return nil, nil
}

func (h *FakeHttpClient) Get(url string) (*http.Response, error) {
	h.GetCounter++
	return nil, nil
}

func (h *FakeHttpClient) DownloadFile(url string, name string) (string, error) {
	h.DownloadFileCounter++
	h.GetCounter++
	return name, nil
}

func GetFakeHttpClient() *FakeHttpClient {
	return &FakeHttpClient{}
}
