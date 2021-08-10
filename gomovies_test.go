package gomovies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchByTitle(t *testing.T) {
	service := NewService("45ace70e")

	t.Run("should be able return a movie by title", func(t *testing.T) {
		_, err := service.SearchByTitle("Lord of the rings")
		assert.Nil(t, err)
	})

	t.Run("should by able return error passing a invalid movie title", func(t *testing.T) {
		_, err := service.SearchByTitle("this movie not exists")
		assert.NotNil(t, err)
	})
}

func TestSearchByID(t *testing.T) {
	service := NewService("45ace70e")

	t.Run("should be able return a movie by id", func(t *testing.T) {
		_, err := service.SearchByID("tt0092675")
		assert.Nil(t, err)
	})

	t.Run("should by able return error passing a invalid movie id", func(t *testing.T) {
		_, err := service.SearchByID("tt")
		assert.NotNil(t, err)
	})
}
