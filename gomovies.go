package gomovies

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const apiOmdbPrefix = "http://www.omdbapi.com"

func GetMovie(key, title, idMovie string) (*Movie, error) {

	url := fmt.Sprintf("%s/?apikey=%s&t=%s&i=%s", apiOmdbPrefix, key, strings.Replace(title, " ", "+", len(title)), idMovie)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
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

	if movie.Response != "True" {
		return nil, fmt.Errorf("%v", movie.Error)
	}

	return movie, nil
}
