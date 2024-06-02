package tmdb

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		panic(err)
	}
}

func TestNew(t *testing.T) {
	c := New(tmdbToken)
	assert.Equal(t, tmdbToken, c.Token)
}

func TestClient_Search(t *testing.T) {
	require.NotEmpty(t, tmdbToken, "TMDB_API_TOKEN is empty")

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
