package service

import (
	"encoding/json"
	"go-kube-demo/internal/pkg/httpclient"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type (
	User interface {
		GetUsers() ([]UserDto, error)
	}

	user struct {
		httpClient httpclient.HttpClient
	}

	UserDto struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Address  any    `json:"address"`
		Phone    string `json:"phone"`
		Website  string `json:"website"`
		Company  any    `json:"company"`
	}
)

func NewUserService(httpClient httpclient.HttpClient) User {
	return &user{
		httpClient,
	}
}

func (user *user) GetUsers() ([]UserDto, error) {
	req := httpclient.HttpRequest{
		Method:  http.MethodGet,
		URL:     "https://jsonplaceholder.typicode.com/users",
		Timeout: 5 * time.Second,
	}

	res, err := user.httpClient.Make(req)
	if err != nil {
		return nil, errors.WithMessage(err, "Error getting users")
	}

	var users []UserDto
	if err := json.Unmarshal(res, &users); err != nil {
		return nil, errors.WithMessage(err, "Error unmarshalling users data")
	}
	return users, nil
}
