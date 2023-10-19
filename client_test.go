package tmdb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tmdbURL = os.Getenv("TMDB_API_URL")
var tmdbToken = os.Getenv("TMDB_API_TOKEN")

func testTmdb(t *testing.T) *Tmdb {
	c, err := New(tmdbURL, tmdbToken)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	return c
}

func TestTmdb_AuthAndSearch(t *testing.T) {
	if tmdbURL == "" {
		t.Skip("TMDB_API_URL not set")
	}
	if tmdbToken == "" {
		t.Skip("TMDB_API_TOKEN not set")
	}

	c, err := New(tmdbURL, tmdbToken)
	assert.NoError(t, err)
	assert.NotNil(t, c)

	r, err := c.SearchMovie(&SearchMovieParams{Query: "The Matrix"})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AccountDetails(t *testing.T) {
	c := testTmdb(t)

	r, err := c.AccountDetails(1, &AccountDetailsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AccountAddFavorite(t *testing.T) {
	t.Skip("400")
	c := testTmdb(t)

	r, err := c.AccountAddFavorite(1, &AccountAddFavoriteParams{}, AccountAddFavoriteJSONRequestBody{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AccountGetFavorites(t *testing.T) {
	c := testTmdb(t)

	r, err := c.AccountGetFavorites(1, &AccountGetFavoritesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AccountFavoriteTv(t *testing.T) {
	c := testTmdb(t)

	r, err := c.AccountFavoriteTv(1, &AccountFavoriteTvParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AccountLists(t *testing.T) {
	c := testTmdb(t)

	r, err := c.AccountLists(1, &AccountListsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AccountRatedMovies(t *testing.T) {
	c := testTmdb(t)

	r, err := c.AccountRatedMovies(1, &AccountRatedMoviesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AccountRatedTv(t *testing.T) {
	c := testTmdb(t)

	r, err := c.AccountRatedTv(1, &AccountRatedTvParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AccountRatedTvEpisodes(t *testing.T) {
	c := testTmdb(t)

	r, err := c.AccountRatedTvEpisodes(1, &AccountRatedTvEpisodesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AccountAddToWatchlist(t *testing.T) {
	t.Skip("400")
	c := testTmdb(t)

	r, err := c.AccountAddToWatchlist(1, &AccountAddToWatchlistParams{}, AccountAddToWatchlistJSONRequestBody{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AccountWatchlistMovies(t *testing.T) {
	c := testTmdb(t)

	r, err := c.AccountWatchlistMovies(1, &AccountWatchlistMoviesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AccountWatchlistTv(t *testing.T) {
	c := testTmdb(t)

	r, err := c.AccountWatchlistTv(1, &AccountWatchlistTvParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AuthenticationValidateKey(t *testing.T) {
	c := testTmdb(t)

	r, err := c.AuthenticationValidateKey()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AuthenticationCreateGuestSession(t *testing.T) {
	// t.Skip("TODO")
	c := testTmdb(t)

	r, err := c.AuthenticationCreateGuestSession()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AuthenticationDeleteSession(t *testing.T) {
	t.Skip("400")
	c := testTmdb(t)

	r, err := c.AuthenticationDeleteSession(AuthenticationDeleteSessionJSONRequestBody{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AuthenticationCreateSessionFromV4Token(t *testing.T) {
	t.Skip("400")
	c := testTmdb(t)

	r, err := c.AuthenticationCreateSessionFromV4Token(AuthenticationCreateSessionFromV4TokenJSONRequestBody{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AuthenticationCreateSession(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.AuthenticationCreateSession(AuthenticationCreateSessionJSONRequestBody{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AuthenticationCreateRequestToken(t *testing.T) {
	// t.Skip("TODO")
	c := testTmdb(t)

	r, err := c.AuthenticationCreateRequestToken()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AuthenticationCreateSessionFromLogin(t *testing.T) {
	t.Skip("400")
	c := testTmdb(t)

	r, err := c.AuthenticationCreateSessionFromLogin(AuthenticationCreateSessionFromLoginJSONRequestBody{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_CertificationMovieList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.CertificationMovieList()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_CertificationsTvList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.CertificationsTvList()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_CollectionDetails(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.CollectionDetails(1, &CollectionDetailsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_CollectionImages(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.CollectionImages(1, &CollectionImagesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_CollectionTranslations(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.CollectionTranslations(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_CompanyDetails(t *testing.T) {
	c := testTmdb(t)

	r, err := c.CompanyDetails(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_CompanyAlternativeNames(t *testing.T) {
	c := testTmdb(t)

	r, err := c.CompanyAlternativeNames(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_CompanyImages(t *testing.T) {
	c := testTmdb(t)

	r, err := c.CompanyImages(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ConfigurationDetails(t *testing.T) {
	c := testTmdb(t)

	r, err := c.ConfigurationDetails()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ConfigurationCountries(t *testing.T) {
	c := testTmdb(t)

	r, err := c.ConfigurationCountries(&ConfigurationCountriesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ConfigurationJobs(t *testing.T) {
	c := testTmdb(t)

	r, err := c.ConfigurationJobs()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ConfigurationLanguages(t *testing.T) {
	c := testTmdb(t)

	r, err := c.ConfigurationLanguages()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ConfigurationPrimaryTranslations(t *testing.T) {
	c := testTmdb(t)

	r, err := c.ConfigurationPrimaryTranslations()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ConfigurationTimezones(t *testing.T) {
	c := testTmdb(t)

	r, err := c.ConfigurationTimezones()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_CreditDetails(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.CreditDetails("1")
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_DiscoverMovie(t *testing.T) {
	c := testTmdb(t)

	r, err := c.DiscoverMovie(&DiscoverMovieParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_DiscoverTv(t *testing.T) {
	c := testTmdb(t)

	r, err := c.DiscoverTv(&DiscoverTvParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_FindById(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.FindById("tt0436992", &FindByIdParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_GenreMovieList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.GenreMovieList(&GenreMovieListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_GenreTvList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.GenreTvList(&GenreTvListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_GuestSessionRatedMovies(t *testing.T) {
	c := testTmdb(t)

	r, err := c.GuestSessionRatedMovies("id", &GuestSessionRatedMoviesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_GuestSessionRatedTv(t *testing.T) {
	c := testTmdb(t)

	r, err := c.GuestSessionRatedTv("id", &GuestSessionRatedTvParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_GuestSessionRatedTvEpisodes(t *testing.T) {
	c := testTmdb(t)

	r, err := c.GuestSessionRatedTvEpisodes("id", &GuestSessionRatedTvEpisodesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_KeywordDetails(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.KeywordDetails(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_KeywordMovies(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.KeywordMovies("scifi", &KeywordMoviesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ListCreate(t *testing.T) {
	t.Skip("400")
	c := testTmdb(t)

	r, err := c.ListCreate(&ListCreateParams{}, ListCreateJSONRequestBody{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ListDelete(t *testing.T) {
	t.Skip("401")
	c := testTmdb(t)

	r, err := c.ListDelete(1, &ListDeleteParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ListDetails(t *testing.T) {
	c := testTmdb(t)

	r, err := c.ListDetails(1, &ListDetailsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ListAddMovie(t *testing.T) {
	t.Skip("401")
	c := testTmdb(t)

	r, err := c.ListAddMovie(1, &ListAddMovieParams{}, ListAddMovieJSONRequestBody{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ListClear(t *testing.T) {
	t.Skip("401")
	c := testTmdb(t)

	r, err := c.ListClear(1, &ListClearParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ListCheckItemStatus(t *testing.T) {
	c := testTmdb(t)

	r, err := c.ListCheckItemStatus(1, &ListCheckItemStatusParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ListRemoveMovie(t *testing.T) {
	t.Skip("401")
	c := testTmdb(t)

	r, err := c.ListRemoveMovie(1, &ListRemoveMovieParams{}, ListRemoveMovieJSONRequestBody{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ChangesMovieList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.ChangesMovieList(&ChangesMovieListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieLatestId(t *testing.T) {
	c := testTmdb(t)

	r, err := c.MovieLatestId()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieNowPlayingList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.MovieNowPlayingList(&MovieNowPlayingListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MoviePopularList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.MoviePopularList(&MoviePopularListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieTopRatedList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.MovieTopRatedList(&MovieTopRatedListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieUpcomingList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.MovieUpcomingList(&MovieUpcomingListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieDetails(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieDetails(1, &MovieDetailsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieAccountStates(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieAccountStates(1, &MovieAccountStatesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieAlternativeTitles(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieAlternativeTitles(1, &MovieAlternativeTitlesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieChanges(t *testing.T) {
	c := testTmdb(t)

	r, err := c.MovieChanges(1, &MovieChangesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieCredits(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieCredits(1, &MovieCreditsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieExternalIds(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieExternalIds(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieImages(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieImages(1, &MovieImagesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieKeywords(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieKeywords("1")
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieLists(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieLists(1, &MovieListsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieDeleteRating(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieDeleteRating(1, &MovieDeleteRatingParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieAddRating(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieAddRating(1, &MovieAddRatingParams{}, MovieAddRatingJSONRequestBody{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieRecommendations(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieRecommendations(1, &MovieRecommendationsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieReleaseDates(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieReleaseDates(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieReviews(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieReviews(1, &MovieReviewsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieSimilar(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieSimilar(1, &MovieSimilarParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieTranslations(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieTranslations(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieVideos(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieVideos(1, &MovieVideosParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_MovieWatchProviders(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.MovieWatchProviders(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_NetworkDetails(t *testing.T) {
	c := testTmdb(t)

	r, err := c.NetworkDetails(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_DetailsCopy(t *testing.T) {
	c := testTmdb(t)

	r, err := c.DetailsCopy(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_AlternativeNamesCopy(t *testing.T) {
	c := testTmdb(t)

	r, err := c.AlternativeNamesCopy(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ChangesPeopleList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.ChangesPeopleList(&ChangesPeopleListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_PersonLatestId(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.PersonLatestId()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_PersonPopularList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.PersonPopularList(&PersonPopularListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_PersonDetails(t *testing.T) {
	c := testTmdb(t)

	r, err := c.PersonDetails(1, &PersonDetailsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_PersonChanges(t *testing.T) {
	c := testTmdb(t)

	r, err := c.PersonChanges(1, &PersonChangesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_PersonCombinedCredits(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.PersonCombinedCredits("1", &PersonCombinedCreditsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_PersonExternalIds(t *testing.T) {
	c := testTmdb(t)

	r, err := c.PersonExternalIds(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_PersonImages(t *testing.T) {
	c := testTmdb(t)

	r, err := c.PersonImages(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_PersonMovieCredits(t *testing.T) {
	c := testTmdb(t)

	r, err := c.PersonMovieCredits(1, &PersonMovieCreditsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_PersonTaggedImages(t *testing.T) {
	c := testTmdb(t)

	r, err := c.PersonTaggedImages(1, &PersonTaggedImagesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_Translations(t *testing.T) {
	c := testTmdb(t)

	r, err := c.Translations(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_PersonTvCredits(t *testing.T) {
	c := testTmdb(t)

	r, err := c.PersonTvCredits(1, &PersonTvCreditsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ReviewDetails(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.ReviewDetails("1")
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_SearchCollection(t *testing.T) {
	c := testTmdb(t)

	r, err := c.SearchCollection(&SearchCollectionParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_SearchCompany(t *testing.T) {
	c := testTmdb(t)

	r, err := c.SearchCompany(&SearchCompanyParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_SearchKeyword(t *testing.T) {
	c := testTmdb(t)

	r, err := c.SearchKeyword(&SearchKeywordParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_SearchMovie(t *testing.T) {
	c := testTmdb(t)

	r, err := c.SearchMovie(&SearchMovieParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_SearchMulti(t *testing.T) {
	c := testTmdb(t)

	r, err := c.SearchMulti(&SearchMultiParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_SearchPerson(t *testing.T) {
	c := testTmdb(t)

	r, err := c.SearchPerson(&SearchPersonParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_SearchTv(t *testing.T) {
	c := testTmdb(t)

	r, err := c.SearchTv(&SearchTvParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TrendingAll(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TrendingAll(TrendingAllParamsTimeWindowDay, &TrendingAllParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TrendingMovies(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TrendingMovies(TrendingMoviesParamsTimeWindowDay, &TrendingMoviesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TrendingPeople(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TrendingPeople(TrendingPeopleParamsTimeWindowDay, &TrendingPeopleParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TrendingTv(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TrendingTv(TrendingTvParamsTimeWindowDay, &TrendingTvParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesAiringTodayList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesAiringTodayList(&TvSeriesAiringTodayListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_ChangesTvList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.ChangesTvList(&ChangesTvListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvEpisodeChangesById(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvEpisodeChangesById(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvEpisodeGroupDetails(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.TvEpisodeGroupDetails("1")
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesLatestId(t *testing.T) {
	t.Skip("404")
	c := testTmdb(t)

	r, err := c.TvSeriesLatestId()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesOnTheAirList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesOnTheAirList(&TvSeriesOnTheAirListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesPopularList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesPopularList(&TvSeriesPopularListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeasonChangesById(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeasonChangesById(1, &TvSeasonChangesByIdParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesTopRatedList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesTopRatedList(&TvSeriesTopRatedListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesDetails(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesDetails(1, &TvSeriesDetailsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesAccountStates(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesAccountStates(1, &TvSeriesAccountStatesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesAggregateCredits(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesAggregateCredits(1, &TvSeriesAggregateCreditsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesAlternativeTitles(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesAlternativeTitles(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesChanges(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesChanges(1, &TvSeriesChangesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesContentRatings(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesContentRatings(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesCredits(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesCredits(1, &TvSeriesCreditsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesEpisodeGroups(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesEpisodeGroups(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesExternalIds(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesExternalIds(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesImages(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesImages(1, &TvSeriesImagesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesKeywords(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesKeywords(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesDeleteRating(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesDeleteRating(1, &TvSeriesDeleteRatingParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesAddRating(t *testing.T) {
	t.Skip("400")
	c := testTmdb(t)

	r, err := c.TvSeriesAddRating(1, &TvSeriesAddRatingParams{}, TvSeriesAddRatingJSONRequestBody{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesRecommendations(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesRecommendations(1, &TvSeriesRecommendationsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesReviews(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesReviews(1, &TvSeriesReviewsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesScreenedTheatrically(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesScreenedTheatrically(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeasonDetails(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeasonDetails(1, 1, &TvSeasonDetailsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeasonAccountStates(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeasonAccountStates(1, 1, &TvSeasonAccountStatesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeasonAggregateCredits(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeasonAggregateCredits(1, 1, &TvSeasonAggregateCreditsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeasonCredits(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeasonCredits(1, 1, &TvSeasonCreditsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvEpisodeDetails(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvEpisodeDetails(1, 1, 1, &TvEpisodeDetailsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvEpisodeAccountStates(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvEpisodeAccountStates(1, 1, 1, &TvEpisodeAccountStatesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvEpisodeCredits(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvEpisodeCredits(1, 1, 1, &TvEpisodeCreditsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvEpisodeExternalIds(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvEpisodeExternalIds(1, 1, "1")
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvEpisodeImages(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvEpisodeImages(1, 1, 1, &TvEpisodeImagesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvEpisodeDeleteRating(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvEpisodeDeleteRating(1, 1, 1, &TvEpisodeDeleteRatingParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvEpisodeAddRating(t *testing.T) {
	t.Skip("400")
	c := testTmdb(t)

	r, err := c.TvEpisodeAddRating(1, 1, 1, &TvEpisodeAddRatingParams{}, TvEpisodeAddRatingJSONRequestBody{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvEpisodeTranslations(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvEpisodeTranslations(1, 1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvEpisodeVideos(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvEpisodeVideos(1, 1, 1, &TvEpisodeVideosParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeasonExternalIds(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeasonExternalIds(1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeasonImages(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeasonImages(1, 1, &TvSeasonImagesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeasonTranslations(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeasonTranslations(1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeasonVideos(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeasonVideos(1, 1, &TvSeasonVideosParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeasonWatchProviders(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeasonWatchProviders(1, 1, &TvSeasonWatchProvidersParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesSimilar(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesSimilar("1", &TvSeriesSimilarParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesTranslations(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesTranslations(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesVideos(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesVideos(1, &TvSeriesVideosParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_TvSeriesWatchProviders(t *testing.T) {
	c := testTmdb(t)

	r, err := c.TvSeriesWatchProviders(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_WatchProvidersMovieList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.WatchProvidersMovieList(&WatchProvidersMovieListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_WatchProvidersAvailableRegions(t *testing.T) {
	c := testTmdb(t)

	r, err := c.WatchProvidersAvailableRegions(&WatchProvidersAvailableRegionsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}

func TestTmdb_WatchProviderTvList(t *testing.T) {
	c := testTmdb(t)

	r, err := c.WatchProviderTvList(&WatchProviderTvListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode())
}
