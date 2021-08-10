package gomovies

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type service struct {
	ApiKey  string
	BaseURL url.URL
}

type Service interface {
	SearchByTitle(title string) (*Movie, error)
	SearchByID(id string) (*Movie, error)
}

func NewService(apiKey string) Service {
	return &service{
		ApiKey: apiKey,
		BaseURL: url.URL{
			Scheme: "https",
			Host:   "www.omdbapi.com",
		},
	}
}

func (s *service) SearchByTitle(title string) (*Movie, error) {
	url, err := s.BaseURL.Parse("")
	if err != nil {
		return nil, err
	}

	q := url.Query()

	q.Add("apikey", s.ApiKey)
	q.Add("t", title)

	url.RawQuery = q.Encode()

	return httpGet(url.String())
}

func (s *service) SearchByID(id string) (*Movie, error) {
	url, err := s.BaseURL.Parse("")
	if err != nil {
		return nil, err
	}

	q := url.Query()

	q.Add("apikey", s.ApiKey)
	q.Add("i", id)

	url.RawQuery = q.Encode()

	return httpGet(url.String())
}

func httpGet(url string) (*Movie, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var movieError *MovieError

	err = json.Unmarshal(body, &movieError)
	if err != nil {
		return nil, err
	}

	if movieError.Error != "" {
		return nil, errors.New(movieError.Error)
	}

	var movie *Movie

	err = json.Unmarshal(body, &movie)
	if err != nil {
		return nil, err
	}

	return movie, nil
}
