package auth

import (
	"fmt"
	"gobackend/pkg/models"
	"io"
	"net/http"
	"net/url"

	"github.com/goccy/go-json"
)

type Service interface {
	GetUserInfoGoogle(token string) models.GoogleResponse
}

type serivce struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &serivce{
		repository: r,
	}
}
func (service *serivce) GetUserInfoGoogle(token string) models.GoogleResponse {
	reqURL, err := url.Parse("https://www.googleapis.com/oauth2/v1/userinfo")
	if err != nil {
		panic(err)
	}
	ptoken := fmt.Sprintf("Bearer %s", token)
	res := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Authorization": {ptoken}},
	}
	req, err := http.DefaultClient.Do(res)
	if err != nil {
		panic(err)

	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var data models.GoogleResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return data
}
