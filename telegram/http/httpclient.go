package http

import (
	"fmt"
	"image/png"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/image/webp"
)

type HttpPoster interface {
	PostForm(url string, values url.Values) (*http.Response, error)
	Get(url string) (*http.Response, error)
	DownloadFile(url string, name string) (string, error)
}

type HttpClient struct{}

func (h *HttpClient) PostForm(url string, values url.Values) (*http.Response, error) {
	fmt.Println("Entering HttpClient.PostForm")
	fmt.Println("Using URL " + url)
	fmt.Println("Using values...")
	fmt.Println(values)
	resp, err := http.PostForm(url, values)
	if err != nil {
		fmt.Println(err)
	}
	return resp, err
}

func (h *HttpClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

func (h *HttpClient) DownloadFile(url string, name string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	img, err := webp.Decode(resp.Body)
	if err != nil {
		fmt.Println("Error decoding WEBP", err)
		return "", err
	}

	defer resp.Body.Close()

	os.Mkdir("downloads", 0700)
	path := filepath.Join("downloads", filepath.Base(name))
	fmt.Println("Saving to filepath", path)
	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		fmt.Println("Error saving PNG", err)
		return "", err
	}
	return path, err
}

func GetHttpClient() *HttpClient {
	return &HttpClient{}
}
