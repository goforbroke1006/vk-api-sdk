package vk_api_sdk

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

	response, err := http.Get(fmt.Sprintf("%s/%s?%s&access_token=%s&v=%s",
		baseApiUrl,
		"users.get",
		fmt.Sprintf("user_ids=%s&fields=%s&name_case=%s",
			strings.Join(userIDsList, ","),
			strings.Join(fields, ","),
			nameCase,
		),
		c.accessToken,
		currentVersion,
	))
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	result := UsersGetResponse{}
	err = json.Unmarshal(bytes, &result)
	return &result, err
}
