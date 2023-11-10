package tmdb

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dashotv/tmdb/openapi/models/operations"
	"github.com/dashotv/tmdb/openapi/types"
)

func TestClient_CertificationMovieList(t *testing.T) {
	c := testClient(t)

	r, err := c.CertificationMovieList()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_CertificationsTvList(t *testing.T) {
	c := testClient(t)

	r, err := c.CertificationsTvList()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_ChangesMovieList(t *testing.T) {
	c := testClient(t)
	var endDate *types.Date = nil
	var page *int = nil
	var startDate *types.Date = nil
	r, err := c.ChangesMovieList(endDate, page, startDate)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_ChangesPeopleList(t *testing.T) {
	c := testClient(t)
	var endDate *types.Date = nil
	var page *int = nil
	var startDate *types.Date = nil
	r, err := c.ChangesPeopleList(endDate, page, startDate)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_ChangesTvList(t *testing.T) {
	c := testClient(t)
	var endDate *types.Date = nil
	var page *int = nil
	var startDate *types.Date = nil
	r, err := c.ChangesTvList(endDate, page, startDate)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_CollectionDetails(t *testing.T) {
	c := testClient(t)
	var collectionID int = collectionID_int
	var language *string = nil
	r, err := c.CollectionDetails(collectionID, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_CollectionImages(t *testing.T) {
	t.Skip("404: valid id's return 404")
	// c := testClient(t)
	// var collectionID int = collectionID_int
	// var includeImageLanguage *string = nil
	// var language *string = nil
	// r, err := c.CollectionImages(collectionID, includeImageLanguage, language)
	// assert.NoError(t, err)
	// assert.NotNil(t, r)
}

func TestClient_CollectionTranslations(t *testing.T) {
	c := testClient(t)
	var collectionID int = collectionID_int
	r, err := c.CollectionTranslations(collectionID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_CompanyAlternativeNames(t *testing.T) {
	c := testClient(t)
	var companyID int = companyID_int
	r, err := c.CompanyAlternativeNames(companyID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_CompanyDetails(t *testing.T) {
	c := testClient(t)
	var companyID int = companyID_int
	r, err := c.CompanyDetails(companyID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_CompanyImages(t *testing.T) {
	c := testClient(t)
	var companyID int = companyID_int
	r, err := c.CompanyImages(companyID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_ConfigurationDetails(t *testing.T) {
	c := testClient(t)

	r, err := c.ConfigurationDetails()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_CreditDetails(t *testing.T) {
	c := testClient(t)
	var creditID string = creditID_string
	r, err := c.CreditDetails(creditID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_DiscoverMovie(t *testing.T) {
	c := testClient(t)
	var request operations.DiscoverMovieRequest = operations_DiscoverMovieRequest
	r, err := c.DiscoverMovie(request)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_DiscoverTv(t *testing.T) {
	c := testClient(t)
	var request operations.DiscoverTvRequest = operations_DiscoverTvRequest
	r, err := c.DiscoverTv(request)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_FindByID(t *testing.T) {
	c := testClient(t)
	var externalID string = externalID_string
	var externalSource operations.FindByIDExternalSource = operations_FindByIDExternalSource
	var language *string = nil
	r, err := c.FindByID(externalID, externalSource, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GenreMovieList(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	r, err := c.GenreMovieList(language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GenreTvList(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	r, err := c.GenreTvList(language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_KeywordDetails(t *testing.T) {
	c := testClient(t)
	var keywordID int = keywordID_int
	r, err := c.KeywordDetails(keywordID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_KeywordMovies(t *testing.T) {
	c := testClient(t)
	var keywordID string = keywordID_string
	var includeAdult *bool = nil
	var language *string = nil
	var page *int = nil
	r, err := c.KeywordMovies(keywordID, includeAdult, language, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieAlternativeTitles(t *testing.T) {
	c := testClient(t)
	var movieID int = movieID_int
	var country *string = nil
	r, err := c.MovieAlternativeTitles(movieID, country)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieCredits(t *testing.T) {
	c := testClient(t)
	var movieID int = movieID_int
	var language *string = nil
	r, err := c.MovieCredits(movieID, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieDetails(t *testing.T) {
	c := testClient(t)
	var movieID int = movieID_int
	var appendToResponse *string = nil
	var language *string = nil
	r, err := c.MovieDetails(movieID, appendToResponse, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieExternalIds(t *testing.T) {
	c := testClient(t)
	var movieID int = movieID_int
	r, err := c.MovieExternalIds(movieID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieImages(t *testing.T) {
	c := testClient(t)
	var movieID int = movieID_int
	var includeImageLanguage *string = nil
	var language *string = nil
	r, err := c.MovieImages(movieID, includeImageLanguage, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieKeywords(t *testing.T) {
	c := testClient(t)
	var movieID string = movieID_string
	r, err := c.MovieKeywords(movieID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieLatestID(t *testing.T) {
	c := testClient(t)

	r, err := c.MovieLatestID()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieNowPlayingList(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	var page *int = nil
	var region *string = nil
	r, err := c.MovieNowPlayingList(language, page, region)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MoviePopularList(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	var page *int = nil
	var region *string = nil
	r, err := c.MoviePopularList(language, page, region)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieRecommendations(t *testing.T) {
	c := testClient(t)
	var movieID int = movieID_int
	var language *string = nil
	var page *int = nil
	r, err := c.MovieRecommendations(movieID, language, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieReleaseDates(t *testing.T) {
	c := testClient(t)
	var movieID int = movieID_int
	r, err := c.MovieReleaseDates(movieID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieReviews(t *testing.T) {
	c := testClient(t)
	var movieID int = movieID_int
	var language *string = nil
	var page *int = nil
	r, err := c.MovieReviews(movieID, language, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieSimilar(t *testing.T) {
	c := testClient(t)
	var movieID int = movieID_int
	var language *string = nil
	var page *int = nil
	r, err := c.MovieSimilar(movieID, language, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieTopRatedList(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	var page *int = nil
	var region *string = nil
	r, err := c.MovieTopRatedList(language, page, region)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieTranslations(t *testing.T) {
	c := testClient(t)
	var movieID int = movieID_int
	r, err := c.MovieTranslations(movieID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieUpcomingList(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	var page *int = nil
	var region *string = nil
	r, err := c.MovieUpcomingList(language, page, region)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieVideos(t *testing.T) {
	c := testClient(t)
	var movieID int = movieID_int
	var language *string = nil
	r, err := c.MovieVideos(movieID, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_MovieWatchProviders(t *testing.T) {
	c := testClient(t)
	var movieID int = movieID_int
	r, err := c.MovieWatchProviders(movieID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_NetworkDetails(t *testing.T) {
	c := testClient(t)
	var networkID int = networkID_int
	r, err := c.NetworkDetails(networkID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_PersonChanges(t *testing.T) {
	c := testClient(t)
	var personID int = personID_int
	var endDate *types.Date = nil
	var page *int = nil
	var startDate *types.Date = nil
	r, err := c.PersonChanges(personID, endDate, page, startDate)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_PersonCombinedCredits(t *testing.T) {
	c := testClient(t)
	var personID string = personID_string
	var language *string = nil
	r, err := c.PersonCombinedCredits(personID, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_PersonDetails(t *testing.T) {
	c := testClient(t)
	var personID int = personID_int
	var appendToResponse *string = nil
	var language *string = nil
	r, err := c.PersonDetails(personID, appendToResponse, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_PersonExternalIds(t *testing.T) {
	c := testClient(t)
	var personID int = personID_int
	r, err := c.PersonExternalIds(personID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_PersonImages(t *testing.T) {
	c := testClient(t)
	var personID int = personID_int
	r, err := c.PersonImages(personID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_PersonLatestID(t *testing.T) {
	c := testClient(t)

	r, err := c.PersonLatestID()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_PersonMovieCredits(t *testing.T) {
	c := testClient(t)
	var personID int = personID_int
	var language *string = nil
	r, err := c.PersonMovieCredits(personID, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_PersonPopularList(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	var page *int = nil
	r, err := c.PersonPopularList(language, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_PersonTaggedImages(t *testing.T) {
	c := testClient(t)
	var personID int = personID_int
	var page *int = nil
	r, err := c.PersonTaggedImages(personID, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_PersonTvCredits(t *testing.T) {
	c := testClient(t)
	var personID int = personID_int
	var language *string = nil
	r, err := c.PersonTvCredits(personID, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_ReviewDetails(t *testing.T) {
	c := testClient(t)
	var reviewID string = reviewID_string
	r, err := c.ReviewDetails(reviewID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_SearchCollection(t *testing.T) {
	c := testClient(t)
	var request operations.SearchCollectionRequest = operations_SearchCollectionRequest
	r, err := c.SearchCollection(request)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_SearchCompany(t *testing.T) {
	c := testClient(t)
	var query string = query_string
	var page *int = nil
	r, err := c.SearchCompany(query, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_SearchKeyword(t *testing.T) {
	c := testClient(t)
	var query string = query_string
	var page *int = nil
	r, err := c.SearchKeyword(query, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_SearchMovie(t *testing.T) {
	c := testClient(t)
	var request operations.SearchMovieRequest = operations_SearchMovieRequest
	r, err := c.SearchMovie(request)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_SearchMulti(t *testing.T) {
	c := testClient(t)
	var query string = query_string
	var includeAdult *bool = nil
	var language *string = nil
	var page *int = nil
	r, err := c.SearchMulti(query, includeAdult, language, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_SearchPerson(t *testing.T) {
	c := testClient(t)
	var query string = query_string
	var includeAdult *bool = nil
	var language *string = nil
	var page *int = nil
	r, err := c.SearchPerson(query, includeAdult, language, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_SearchTv(t *testing.T) {
	c := testClient(t)
	var request operations.SearchTvRequest = operations_SearchTvRequest
	r, err := c.SearchTv(request)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_Translations(t *testing.T) {
	c := testClient(t)
	var personID int = personID_int
	r, err := c.Translations(personID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TrendingAll(t *testing.T) {
	c := testClient(t)
	var timeWindow operations.TrendingAllTimeWindow = operations_TrendingAllTimeWindow
	var language *string = nil
	r, err := c.TrendingAll(timeWindow, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TrendingMovies(t *testing.T) {
	c := testClient(t)
	var timeWindow operations.TrendingMoviesTimeWindow = operations_TrendingMoviesTimeWindow
	var language *string = nil
	r, err := c.TrendingMovies(timeWindow, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TrendingPeople(t *testing.T) {
	c := testClient(t)
	var timeWindow operations.TrendingPeopleTimeWindow = operations_TrendingPeopleTimeWindow
	var language *string = nil
	r, err := c.TrendingPeople(timeWindow, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TrendingTv(t *testing.T) {
	c := testClient(t)
	var timeWindow operations.TrendingTvTimeWindow = operations_TrendingTvTimeWindow
	var language *string = nil
	r, err := c.TrendingTv(timeWindow, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvEpisodeChangesByID(t *testing.T) {
	c := testClient(t)
	var episodeID int = episodeID_int
	r, err := c.TvEpisodeChangesByID(episodeID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvEpisodeCredits(t *testing.T) {
	c := testClient(t)
	var episodeNumber int = episodeNumber_int
	var seasonNumber int = seasonNumber_int
	var seriesID int = seriesID_int
	var language *string = nil
	r, err := c.TvEpisodeCredits(episodeNumber, seasonNumber, seriesID, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvEpisodeDetails(t *testing.T) {
	c := testClient(t)
	var request operations.TvEpisodeDetailsRequest = operations_TvEpisodeDetailsRequest
	r, err := c.TvEpisodeDetails(request)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvEpisodeExternalIds(t *testing.T) {
	c := testClient(t)
	var episodeNumber string = episodeNumber_string
	var seasonNumber int = seasonNumber_int
	var seriesID int = seriesID_int
	r, err := c.TvEpisodeExternalIds(episodeNumber, seasonNumber, seriesID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvEpisodeGroupDetails(t *testing.T) {
	c := testClient(t)
	var tvEpisodeGroupID string = tvEpisodeGroupID_string
	r, err := c.TvEpisodeGroupDetails(tvEpisodeGroupID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvEpisodeImages(t *testing.T) {
	c := testClient(t)
	var request operations.TvEpisodeImagesRequest = operations_TvEpisodeImagesRequest
	r, err := c.TvEpisodeImages(request)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvEpisodeTranslations(t *testing.T) {
	c := testClient(t)
	var episodeNumber int = episodeNumber_int
	var seasonNumber int = seasonNumber_int
	var seriesID int = seriesID_int
	r, err := c.TvEpisodeTranslations(episodeNumber, seasonNumber, seriesID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvEpisodeVideos(t *testing.T) {
	c := testClient(t)
	var request operations.TvEpisodeVideosRequest = operations_TvEpisodeVideosRequest
	r, err := c.TvEpisodeVideos(request)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeasonAggregateCredits(t *testing.T) {
	c := testClient(t)
	var seasonNumber int = seasonNumber_int
	var seriesID int = seriesID_int
	var language *string = nil
	r, err := c.TvSeasonAggregateCredits(seasonNumber, seriesID, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeasonChangesByID(t *testing.T) {
	c := testClient(t)
	var seasonID int = seasonID_int
	var endDate *string = nil
	var page *int = nil
	var startDate *string = nil
	r, err := c.TvSeasonChangesByID(seasonID, endDate, page, startDate)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeasonCredits(t *testing.T) {
	c := testClient(t)
	var seasonNumber int = seasonNumber_int
	var seriesID int = seriesID_int
	var language *string = nil
	r, err := c.TvSeasonCredits(seasonNumber, seriesID, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeasonDetails(t *testing.T) {
	c := testClient(t)
	var seasonNumber int = seasonNumber_int
	var seriesID int = seriesID_int
	var appendToResponse *string = nil
	var language *string = nil
	r, err := c.TvSeasonDetails(seasonNumber, seriesID, appendToResponse, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeasonExternalIds(t *testing.T) {
	c := testClient(t)
	var seasonNumber int = seasonNumber_int
	var seriesID int = seriesID_int
	r, err := c.TvSeasonExternalIds(seasonNumber, seriesID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeasonImages(t *testing.T) {
	c := testClient(t)
	var seasonNumber int = seasonNumber_int
	var seriesID int = seriesID_int
	var includeImageLanguage *string = nil
	var language *string = nil
	r, err := c.TvSeasonImages(seasonNumber, seriesID, includeImageLanguage, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeasonTranslations(t *testing.T) {
	c := testClient(t)
	var seasonNumber int = seasonNumber_int
	var seriesID int = seriesID_int
	r, err := c.TvSeasonTranslations(seasonNumber, seriesID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeasonVideos(t *testing.T) {
	c := testClient(t)
	var seasonNumber int = seasonNumber_int
	var seriesID int = seriesID_int
	var includeVideoLanguage *string = nil
	var language *string = nil
	r, err := c.TvSeasonVideos(seasonNumber, seriesID, includeVideoLanguage, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeasonWatchProviders(t *testing.T) {
	c := testClient(t)
	var seasonNumber int = seasonNumber_int
	var seriesID int = seriesID_int
	var language *string = nil
	r, err := c.TvSeasonWatchProviders(seasonNumber, seriesID, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesAggregateCredits(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	var language *string = nil
	r, err := c.TvSeriesAggregateCredits(seriesID, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesAiringTodayList(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	var page *int = nil
	var timezone *string = nil
	r, err := c.TvSeriesAiringTodayList(language, page, timezone)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesAlternativeTitles(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	r, err := c.TvSeriesAlternativeTitles(seriesID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesContentRatings(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	r, err := c.TvSeriesContentRatings(seriesID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesCredits(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	var language *string = nil
	r, err := c.TvSeriesCredits(seriesID, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesDetails(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	var appendToResponse *string = nil
	var language *string = nil
	r, err := c.TvSeriesDetails(seriesID, appendToResponse, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesEpisodeGroups(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	r, err := c.TvSeriesEpisodeGroups(seriesID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesExternalIds(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	r, err := c.TvSeriesExternalIds(seriesID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesImages(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	var includeImageLanguage *string = nil
	var language *string = nil
	r, err := c.TvSeriesImages(seriesID, includeImageLanguage, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesKeywords(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	r, err := c.TvSeriesKeywords(seriesID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesLatestID(t *testing.T) {
	c := testClient(t)

	r, err := c.TvSeriesLatestID()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesOnTheAirList(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	var page *int = nil
	var timezone *string = nil
	r, err := c.TvSeriesOnTheAirList(language, page, timezone)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesPopularList(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	var page *int = nil
	r, err := c.TvSeriesPopularList(language, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesRecommendations(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	var language *string = nil
	var page *int = nil
	r, err := c.TvSeriesRecommendations(seriesID, language, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesReviews(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	var language *string = nil
	var page *int = nil
	r, err := c.TvSeriesReviews(seriesID, language, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesScreenedTheatrically(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	r, err := c.TvSeriesScreenedTheatrically(seriesID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesSimilar(t *testing.T) {
	c := testClient(t)
	var seriesID string = seriesID_string
	var language *string = nil
	var page *int = nil
	r, err := c.TvSeriesSimilar(seriesID, language, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesTopRatedList(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	var page *int = nil
	r, err := c.TvSeriesTopRatedList(language, page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesTranslations(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	r, err := c.TvSeriesTranslations(seriesID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesVideos(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	var includeVideoLanguage *string = nil
	var language *string = nil
	r, err := c.TvSeriesVideos(seriesID, includeVideoLanguage, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_TvSeriesWatchProviders(t *testing.T) {
	c := testClient(t)
	var seriesID int = seriesID_int
	r, err := c.TvSeriesWatchProviders(seriesID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_WatchProviderTvList(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	var watchRegion *string = nil
	r, err := c.WatchProviderTvList(language, watchRegion)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_WatchProvidersAvailableRegions(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	r, err := c.WatchProvidersAvailableRegions(language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_WatchProvidersMovieList(t *testing.T) {
	c := testClient(t)
	var language *string = nil
	var watchRegion *string = nil
	r, err := c.WatchProvidersMovieList(language, watchRegion)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}
