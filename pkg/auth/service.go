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

	"github.com/bytedance/sonic"
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
	GetUserOrCreate(profile *googleResponse) (*entities.Users, error)
	SetTokenJwt(user *entities.Users) (string, error)
	GetProfile(id string) (*entities.Users, error)
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

	err = sonic.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *serivce) GetUserOrCreate(profile *googleResponse) (*entities.Users, error) {
	user, err := s.repository.GetUserByEmail(profile.Email)
	if err == nil {
		return user, err
	}
	usercreate := new(entities.Users)
	usercreate.Email = profile.Email
	usercreate.FirstName = profile.Given_name
	usercreate.LastName = profile.Family_name
	newuser, err := s.userreposistory.CreateUser(usercreate)
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

func (s *serivce) GetProfile(id string) (*entities.Users, error) {
	profile := new(entities.Users)
	return s.userreposistory.GetUser(profile)
}
