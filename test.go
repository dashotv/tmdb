package tmdb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dashotv/tmdb/openapi/models/operations"
)

var tmdbToken = os.Getenv("TMDB_TOKEN")

var collectionID_int = 131295
var companyID_int = 7505
var creditID_string = "52fe4250c3a36847f80149f3"
var episodeID_int = 63056
var episodeNumber_int = 1
var episodeNumber_string = "1"
var externalID_string = "tt0111161"
var keywordID_int = 818
var keywordID_string = "818"
var movieID_int = 278
var movieID_string = "278"
var networkID_int = 49
var personID_int = 31
var personID_string = "31"
var query_string = "The Simpsons"
var reviewID_string = "640b2aeecaaca20079decdcc"
var seasonID_int = 3572
var seasonNumber_int = 1
var seriesID_int = 1396
var seriesID_string = "1396"
var tvEpisodeGroupID_string = "5acf93e60e0a26346d0000ce"

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
