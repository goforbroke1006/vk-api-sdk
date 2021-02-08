package vksdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *vkApiClient) PhotosGet(ownerID int64, albumID int64) (*PhotosGetResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", baseApiUrl, "photos.get"), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("owner_id", fmt.Sprintf("%d", ownerID))
	q.Add("album_id", fmt.Sprintf("%d", albumID))
	//q.Add("album_id", "saved")

	q.Add("access_token", c.accessToken)
	q.Add("v", currentVersion)

	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := PhotosGetResponse{}
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

func (c *vkApiClient) PhotosGetAlbums(ownerID int64) (*PhotosGetAlbumsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", baseApiUrl, "photos.getAlbums"), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("owner_id", fmt.Sprintf("%d", ownerID))
	q.Add("need_system", "1")

	q.Add("access_token", c.accessToken)
	q.Add("v", currentVersion)

	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := PhotosGetAlbumsResponse{}
	err = json.Unmarshal(bytes, &result)
	return &result, err
}
