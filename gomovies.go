package gomovies

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const URL = "http://www.omdbapi.com"

type service struct {
	ApiKey string
}

type Service interface {
	SearchByTitle(title string) (*Movie, error)
	SearchByID(id string) (*Movie, error)
}

func NewService(apiKey string) Service {
	return &service{
		ApiKey: apiKey,
	}
}

func (s *service) SearchByTitle(title string) (*Movie, error) {
	if title == "" {
		return nil, errors.New("Oops, title cannot be empty")
	}

	title = strings.ToLower(title)
	title = strings.Replace(title, " ", "+", len(title))

	url := fmt.Sprintf("%s/?apikey=%s&t=%s", URL, s.ApiKey, title)

	return httpGet(url)
}

func (s *service) SearchByID(id string) (*Movie, error) {
	if id == "" {
		return nil, errors.New("Oops, id cannot be empty")
	}

	url := fmt.Sprintf("%s/?apikey=%s&i=%s", URL, s.ApiKey, id)

	return httpGet(url)
}

func httpGet(url string) (*Movie, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var movie *Movie

	err = json.Unmarshal(body, &movie)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Oops, %s", movie.Error)
	}

	return movie, nil
}
