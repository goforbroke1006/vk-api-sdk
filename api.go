package vk_api_sdk

func New(accessToken string) *vkApiClient {
	return &vkApiClient{
		accessToken: accessToken,
	}
}

type vkApiClient struct {
	accessToken string
}
