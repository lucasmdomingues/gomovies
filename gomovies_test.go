package gomovies

import "testing"

const key = ""

func TestGetMovieByTitle(t *testing.T) {

	movie, err := GetMovieByTitle(key, "Lord of the rings")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("%#v", movie)
}
