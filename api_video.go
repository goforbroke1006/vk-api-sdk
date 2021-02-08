package vksdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *vkApiClient) VideoGet(ownerID int64, albumID int64, ownedIdVideId ...int64) (*VideoGetResponse, error) {
	if len(ownedIdVideId)%2 != 0 {
		return nil, errors.New("ownedIdVideId has unpaired quantity")
	}

	videos := make([]string, 0, len(ownedIdVideId)/2)
	for i := 0; i < len(ownedIdVideId); i += 2 {
		fullVideoID := fmt.Sprintf("%d_%d", ownedIdVideId[i], ownedIdVideId[i+1])
		videos = append(videos, fullVideoID)
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", baseApiUrl, "video.get"), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("owner_id", fmt.Sprintf("%d", ownerID))
	q.Add("album_id", fmt.Sprintf("%d", albumID))
	q.Add("videos", strings.Join(videos, ","))
	q.Add("extended", "1")

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

	result := VideoGetResponse{}
	err = json.Unmarshal(bytes, &result)
	return &result, err
}

func (c *vkApiClient) VideoGetAlbums(ownerID int64) (*VideoGetAlbumsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", baseApiUrl, "video.getAlbums"), nil)
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

	result := VideoGetAlbumsResponse{}
	err = json.Unmarshal(bytes, &result)
	return &result, err
}
