package vksdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// https://vk.com/dev/users.get
func (c *vkApiClient) UsersGet(userIDs []interface{}, fields []string, nameCase NameCase) (*UsersGetResponse, error) {
	userIDsList := make([]string, 0, len(userIDs))
	for _, ui := range userIDs {
		if str, ok := ui.(string); ok {
			userIDsList = append(userIDsList, str)
			continue
		}
		if num, ok := ui.(int); ok {
			userIDsList = append(userIDsList, fmt.Sprintf("%d", num))
			continue
		}
		return nil, errors.New("unexpected ID type, want string or int")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", baseApiUrl, "users.get"), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("user_ids", strings.Join(userIDsList, ","))
	q.Add("fields", strings.Join(fields, ","))
	q.Add("name_case", string(nameCase))

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

	result := UsersGetResponse{}
	err = json.Unmarshal(bytes, &result)
	return &result, err
}
