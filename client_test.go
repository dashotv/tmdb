package tmdb

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dashotv/tmdb/openapi/models/operations"
)

var tmdbToken = os.Getenv("TMDB_API_TOKEN")
var query = "The Simpsons"
var collectionID = 131295
var companyID = 7505
var creditID = "52fe4250c3a36847f80149f3"
var externalID = "tt0111161"
var networkID = 49
var reviewID = "640b2aeecaaca20079decdcc"
var episodeID = 63056
var seasonID = 3572
var tvEpisodeGroupID = "5acf93e60e0a26346d0000ce"

var operations_TvEpisodeVideosRequest = operations.TvEpisodeVideosRequest{SeriesID: 1396, SeasonNumber: 1, EpisodeNumber: 1}
var operations_TrendingAllTimeWindow = operations.TrendingAllTimeWindow(operations.TrendingAllTimeWindowDay)
var operations_TrendingMoviesTimeWindow = operations.TrendingMoviesTimeWindow(operations.TrendingMoviesTimeWindowDay)
var operations_TrendingPeopleTimeWindow = operations.TrendingPeopleTimeWindow(operations.TrendingPeopleTimeWindowDay)
var operations_TrendingTvTimeWindow = operations.TrendingTvTimeWindow(operations.TrendingPeopleTimeWindowDay)
var operations_FindByIDExternalSource = operations.FindByIDExternalSource("tt0111161")
var operations_DiscoverMovieRequest = operations.DiscoverMovieRequest{}
var operations_DiscoverTvRequest = operations.DiscoverTvRequest{}
var operations_SearchCollectionRequest = operations.SearchCollectionRequest{}
var operations_SearchMovieRequest = operations.SearchMovieRequest{}
var operations_SearchTvRequest = operations.SearchTvRequest{}
var operations_TvEpisodeDetailsRequest = operations.TvEpisodeDetailsRequest{SeriesID: 1396, SeasonNumber: 1, EpisodeNumber: 1}
var operations_TvEpisodeImagesRequest = operations.TvEpisodeImagesRequest{SeriesID: 1396, SeasonNumber: 1, EpisodeNumber: 1}

func testClient(t *testing.T) *Client {
	c := New(tmdbToken)
	assert.Equal(t, tmdbToken, c.Token)
	return c
}

func TestNew(t *testing.T) {
	c := New(tmdbToken)
	assert.Equal(t, tmdbToken, c.Token)
}

func stringValue(p *string) string {
	if p == nil {
		return ""
	}
	return *p
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
		fmt.Printf("%d: %s\n", *v.ID, stringValue(v.Title))
	}
}
