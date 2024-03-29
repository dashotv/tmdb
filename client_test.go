package tmdb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New(tmdbToken)
	assert.Equal(t, tmdbToken, c.Token)
}

func TestClient_Search(t *testing.T) {
	assert.NotEmpty(t, tmdbToken, "TMDB_API_TOKEN is empty")

	c := New(tmdbToken)
	assert.Equal(t, tmdbToken, c.Token)

	p := SearchMovieRequest{
		Query: "The Simpsons",
	}
	r, err := c.SearchMovie(p)
	assert.NoError(t, err)
	assert.NotNil(t, r)

	for _, v := range r.Results {
		assert.NotEmpty(t, *v.ID)
		fmt.Printf("%d: %s\n", *v.ID, StringValue(v.Title))
	}
}
