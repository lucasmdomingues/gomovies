package gomovies

import "testing"

const apikey = ""

func TestSearchByTitle(t *testing.T) {
	_, err := SearchByTitle(apikey, "Lord of the rings")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSearchByID(t *testing.T) {
	_, err := SearchByID(apikey, "tt0120737")
	if err != nil {
		t.Error(err)
		return
	}
}
