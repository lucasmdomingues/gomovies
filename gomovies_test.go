package gomovies

import "testing"

const key = "45ace70e"

func TestGetMovie(t *testing.T) {

	movie, err := GetMovie(key, "Lord of the rings", "tt0120737")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("%+v", movie)
}
