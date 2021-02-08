package vksdk

import "net/http"

func New(accessToken string) *vkApiClient {
	return &vkApiClient{
		accessToken: accessToken,
		httpClient:  &http.Client{},
	}
}
func NewWithHTTPClient(accessToken string, httpClient *http.Client) *vkApiClient {
	return &vkApiClient{
		accessToken: accessToken,
		httpClient:  httpClient,
	}
}

type vkApiClient struct {
	accessToken string
	httpClient  *http.Client
}
