package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gojek/heimdall/httpclient"
)

type Requester interface {
	GetUserInfo(token string) (*UserInfo, error)
}

type TwitchRequest struct {
	client httpclient.Client
}

func NewTwitchRequest(client httpclient.Client) *TwitchRequest {
	return &TwitchRequest{client: client}
}

func (t *TwitchRequest) validateToken(token string) (*ValidResponse, error) {
	token = "OAuth " + token
	header := http.Header{
		"Authorization": {token},
	}
	response, err := t.client.Get("https://id.twitch.tv/oauth2/validate", header)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var result ValidResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *TwitchRequest) GetUserInfo(token string) (*UserInfo, error) {
	validResponse, err := t.validateToken(token)
	if err != nil {
		return nil, err
	}
	token = "Bearer " + token
	header := http.Header{
		"Authorization": {token},
		"Client-Id":     {validResponse.ClientId},
	}
	response, err := t.client.Get(fmt.Sprintf("https://api.twitch.tv/helix/users?id=%s", validResponse.UserId), header)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var result UsersInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data[0], nil
}
