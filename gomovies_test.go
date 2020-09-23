package gomovies

import "testing"

func TestSearchByTitle(t *testing.T) {
	service := NewService("API_KEY")

	_, err := service.SearchByTitle("Lord of the rings")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSearchByID(t *testing.T) {
	service := NewService("API_KEY")

	_, err := service.SearchByID("tt0120737")
	if err != nil {
		t.Error(err)
		return
	}
}
