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
	"github.com/google/uuid"
)

type Service interface {
	GetUserInfoGoogle(token string) (*GoogleResponse, error)
	GetUserOrCreate(profile *GoogleResponse) (*entities.User, error)
	SetTokenJwt(user *entities.User) (string, error)
	GetProfile(id string) (*entities.User, error)
	UpdateProfile(user *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
	GenerateToken(user *entities.User) (string, error)
}

type serivce struct {
	authrepository  Repository
	userreposistory user.Repository
}

func NewService(authrepository Repository, userrepo user.Repository) Service {
	return &serivce{
		authrepository:  authrepository,
		userreposistory: userrepo,
	}
}
func (service *serivce) GetUserInfoGoogle(token string) (*GoogleResponse, error) {
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
	var data GoogleResponse

	err = sonic.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *serivce) GetUserOrCreate(profile *GoogleResponse) (*entities.User, error) {
	user, err := s.authrepository.GetUserByEmail(profile.Email)
	if err == nil {
		return user, err
	}
	usercreate := new(entities.User)
	usercreate.Email = profile.Email
	usercreate.FirstName = profile.Given_name
	usercreate.LastName = profile.Family_name
	usercreate.Image = profile.Picture
	newuser, err := s.userreposistory.CreateUserWithOutValidate(usercreate)
	if err != nil {
		return nil, err
	}
	return newuser, err

}

func (s *serivce) SetTokenJwt(user *entities.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["issuperadmin"] = user.Issuperadmin
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()
	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (s *serivce) GetProfile(id string) (*entities.User, error) {
	profile := new(entities.User)
	id_uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	profile.ID = id_uuid
	return s.userreposistory.GetUser(profile)
}

func (s *serivce) UpdateProfile(user *entities.User) error {
	return s.userreposistory.UpdateUser(user)
}

func (s *serivce) GetUserByEmail(email string) (*entities.User, error) {
	return s.userreposistory.GetUserByEmail(email)
}

func (service *serivce) GenerateToken(user *entities.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["issuperadmin"] = user.Issuperadmin
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	return t, err
}
