package auth

import (
	"fmt"
	"gobackend/pkg/entities"
	"gobackend/pkg/user"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v5"
)

type googleResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Verified    bool   `json:"verified_email"`
	Picture     string `json:"picture"`
	Name        string `json:"name"`
	Given_name  string `json:"given_name"`
	Family_name string `json:"family_name"`
}

type Service interface {
	GetUserInfoGoogle(token string) (*googleResponse, error)
	GetUserOrCreate(email string) (*entities.Users, error)
	SetTokenJwt(user *entities.Users) (string, error)
}

type serivce struct {
	repository      Repository
	userreposistory user.Repository
}

func NewService(r Repository, userrepo user.Repository) Service {
	return &serivce{
		repository:      r,
		userreposistory: userrepo,
	}
}
func (service *serivce) GetUserInfoGoogle(token string) (*googleResponse, error) {
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
		return nil, err

	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var data googleResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *serivce) GetUserOrCreate(email string) (*entities.Users, error) {
	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	newuser, err := s.userreposistory.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return newuser, err

}

func (s *serivce) SetTokenJwt(user *entities.Users) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["issuperadmin"] = user.Issuperadmin
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix()
	t, err := token.SignedString([]byte(os.Getenv("SIGNKEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}
