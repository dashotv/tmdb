package tmdb

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type contextKey string

const (
	contextBearerToken contextKey = "bearerToken"
)

// Tmdb wrapper client
type Tmdb struct {
	URL   string
	Key   string
	token string
	c     *ClientWithResponses
	ctx   context.Context
}

// New creates a new Tmdb wrapper client and configures Bearer authentication
//
//	url is the base URL of the Tmdb API
//
// Example: https://api.themoviedb.org
//
// https://developer.themoviedb.org/reference/intro/getting-started
//
// This wraps the generated client with a few helper functions, it also sets
// up authentication and simplifies API access.
func New(url, token string) (*Tmdb, error) {
	if token == "" {
		return nil, errors.Errorf("token cannot be empty")
	}

	c := &Tmdb{
		URL:   url,
		token: token,
	}

	c.ctx = context.WithValue(context.Background(), contextBearerToken, token)

	nc, err := newClientWithResponses(c.URL, WithRequestEditorFn(requestAuth))
	if err != nil {
		return nil, err
	}

	c.c = nc

	return c, nil
}

// requestAuth wires up the bearer token to the request
func requestAuth(ctx context.Context, req *http.Request) error {
	token, ok := ctx.Value(contextBearerToken).(string)
	if !ok {
		return errors.Errorf("no bearer token found in context")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return nil
}

func (c *Tmdb) AccountDetails(accountId int32, params *AccountDetailsParams) (*AccountDetailsResponse, error) {
	return c.c.AccountDetailsWithResponse(c.ctx, accountId, params)
}

func (c *Tmdb) AccountAddFavorite(accountId int32, params *AccountAddFavoriteParams, body AccountAddFavoriteJSONRequestBody) (*AccountAddFavoriteResponse, error) {
	return c.c.AccountAddFavoriteWithResponse(c.ctx, accountId, params, body)
}

func (c *Tmdb) AccountGetFavorites(accountId int32, params *AccountGetFavoritesParams) (*AccountGetFavoritesResponse, error) {
	return c.c.AccountGetFavoritesWithResponse(c.ctx, accountId, params)
}

func (c *Tmdb) AccountFavoriteTv(accountId int32, params *AccountFavoriteTvParams) (*AccountFavoriteTvResponse, error) {
	return c.c.AccountFavoriteTvWithResponse(c.ctx, accountId, params)
}

func (c *Tmdb) AccountLists(accountId int32, params *AccountListsParams) (*AccountListsResponse, error) {
	return c.c.AccountListsWithResponse(c.ctx, accountId, params)
}

func (c *Tmdb) AccountRatedMovies(accountId int32, params *AccountRatedMoviesParams) (*AccountRatedMoviesResponse, error) {
	return c.c.AccountRatedMoviesWithResponse(c.ctx, accountId, params)
}

func (c *Tmdb) AccountRatedTv(accountId int32, params *AccountRatedTvParams) (*AccountRatedTvResponse, error) {
	return c.c.AccountRatedTvWithResponse(c.ctx, accountId, params)
}

func (c *Tmdb) AccountRatedTvEpisodes(accountId int32, params *AccountRatedTvEpisodesParams) (*AccountRatedTvEpisodesResponse, error) {
	return c.c.AccountRatedTvEpisodesWithResponse(c.ctx, accountId, params)
}

func (c *Tmdb) AccountAddToWatchlist(accountId int32, params *AccountAddToWatchlistParams, body AccountAddToWatchlistJSONRequestBody) (*AccountAddToWatchlistResponse, error) {
	return c.c.AccountAddToWatchlistWithResponse(c.ctx, accountId, params, body)
}

func (c *Tmdb) AccountWatchlistMovies(accountId int32, params *AccountWatchlistMoviesParams) (*AccountWatchlistMoviesResponse, error) {
	return c.c.AccountWatchlistMoviesWithResponse(c.ctx, accountId, params)
}

func (c *Tmdb) AccountWatchlistTv(accountId int32, params *AccountWatchlistTvParams) (*AccountWatchlistTvResponse, error) {
	return c.c.AccountWatchlistTvWithResponse(c.ctx, accountId, params)
}

func (c *Tmdb) AuthenticationValidateKey() (*AuthenticationValidateKeyResponse, error) {
	return c.c.AuthenticationValidateKeyWithResponse(c.ctx)
}

func (c *Tmdb) AuthenticationCreateGuestSession() (*AuthenticationCreateGuestSessionResponse, error) {
	return c.c.AuthenticationCreateGuestSessionWithResponse(c.ctx)
}

func (c *Tmdb) AuthenticationDeleteSession(body AuthenticationDeleteSessionJSONRequestBody) (*AuthenticationDeleteSessionResponse, error) {
	return c.c.AuthenticationDeleteSessionWithResponse(c.ctx, body)
}

func (c *Tmdb) AuthenticationCreateSessionFromV4Token(body AuthenticationCreateSessionFromV4TokenJSONRequestBody) (*AuthenticationCreateSessionFromV4TokenResponse, error) {
	return c.c.AuthenticationCreateSessionFromV4TokenWithResponse(c.ctx, body)
}

func (c *Tmdb) AuthenticationCreateSession(body AuthenticationCreateSessionJSONRequestBody) (*AuthenticationCreateSessionResponse, error) {
	return c.c.AuthenticationCreateSessionWithResponse(c.ctx, body)
}

func (c *Tmdb) AuthenticationCreateRequestToken() (*AuthenticationCreateRequestTokenResponse, error) {
	return c.c.AuthenticationCreateRequestTokenWithResponse(c.ctx)
}

func (c *Tmdb) AuthenticationCreateSessionFromLogin(body AuthenticationCreateSessionFromLoginJSONRequestBody) (*AuthenticationCreateSessionFromLoginResponse, error) {
	return c.c.AuthenticationCreateSessionFromLoginWithResponse(c.ctx, body)
}

func (c *Tmdb) CertificationMovieList() (*CertificationMovieListResponse, error) {
	return c.c.CertificationMovieListWithResponse(c.ctx)
}

func (c *Tmdb) CertificationsTvList() (*CertificationsTvListResponse, error) {
	return c.c.CertificationsTvListWithResponse(c.ctx)
}

func (c *Tmdb) CollectionDetails(collectionId int32, params *CollectionDetailsParams) (*CollectionDetailsResponse, error) {
	return c.c.CollectionDetailsWithResponse(c.ctx, collectionId, params)
}

func (c *Tmdb) CollectionImages(collectionId int32, params *CollectionImagesParams) (*CollectionImagesResponse, error) {
	return c.c.CollectionImagesWithResponse(c.ctx, collectionId, params)
}

func (c *Tmdb) CollectionTranslations(collectionId int32) (*CollectionTranslationsResponse, error) {
	return c.c.CollectionTranslationsWithResponse(c.ctx, collectionId)
}

func (c *Tmdb) CompanyDetails(companyId int32) (*CompanyDetailsResponse, error) {
	return c.c.CompanyDetailsWithResponse(c.ctx, companyId)
}

func (c *Tmdb) CompanyAlternativeNames(companyId int32) (*CompanyAlternativeNamesResponse, error) {
	return c.c.CompanyAlternativeNamesWithResponse(c.ctx, companyId)
}

func (c *Tmdb) CompanyImages(companyId int32) (*CompanyImagesResponse, error) {
	return c.c.CompanyImagesWithResponse(c.ctx, companyId)
}

func (c *Tmdb) ConfigurationDetails() (*ConfigurationDetailsResponse, error) {
	return c.c.ConfigurationDetailsWithResponse(c.ctx)
}

func (c *Tmdb) ConfigurationCountries(params *ConfigurationCountriesParams) (*ConfigurationCountriesResponse, error) {
	return c.c.ConfigurationCountriesWithResponse(c.ctx, params)
}

func (c *Tmdb) ConfigurationJobs() (*ConfigurationJobsResponse, error) {
	return c.c.ConfigurationJobsWithResponse(c.ctx)
}

func (c *Tmdb) ConfigurationLanguages() (*ConfigurationLanguagesResponse, error) {
	return c.c.ConfigurationLanguagesWithResponse(c.ctx)
}

func (c *Tmdb) ConfigurationPrimaryTranslations() (*ConfigurationPrimaryTranslationsResponse, error) {
	return c.c.ConfigurationPrimaryTranslationsWithResponse(c.ctx)
}

func (c *Tmdb) ConfigurationTimezones() (*ConfigurationTimezonesResponse, error) {
	return c.c.ConfigurationTimezonesWithResponse(c.ctx)
}

func (c *Tmdb) CreditDetails(creditId string) (*CreditDetailsResponse, error) {
	return c.c.CreditDetailsWithResponse(c.ctx, creditId)
}

func (c *Tmdb) DiscoverMovie(params *DiscoverMovieParams) (*DiscoverMovieResponse, error) {
	return c.c.DiscoverMovieWithResponse(c.ctx, params)
}

func (c *Tmdb) DiscoverTv(params *DiscoverTvParams) (*DiscoverTvResponse, error) {
	return c.c.DiscoverTvWithResponse(c.ctx, params)
}

func (c *Tmdb) FindById(externalId string, params *FindByIdParams) (*FindByIdResponse, error) {
	return c.c.FindByIdWithResponse(c.ctx, externalId, params)
}

func (c *Tmdb) GenreMovieList(params *GenreMovieListParams) (*GenreMovieListResponse, error) {
	return c.c.GenreMovieListWithResponse(c.ctx, params)
}

func (c *Tmdb) GenreTvList(params *GenreTvListParams) (*GenreTvListResponse, error) {
	return c.c.GenreTvListWithResponse(c.ctx, params)
}

func (c *Tmdb) GuestSessionRatedMovies(guestSessionId string, params *GuestSessionRatedMoviesParams) (*GuestSessionRatedMoviesResponse, error) {
	return c.c.GuestSessionRatedMoviesWithResponse(c.ctx, guestSessionId, params)
}

func (c *Tmdb) GuestSessionRatedTv(guestSessionId string, params *GuestSessionRatedTvParams) (*GuestSessionRatedTvResponse, error) {
	return c.c.GuestSessionRatedTvWithResponse(c.ctx, guestSessionId, params)
}

func (c *Tmdb) GuestSessionRatedTvEpisodes(guestSessionId string, params *GuestSessionRatedTvEpisodesParams) (*GuestSessionRatedTvEpisodesResponse, error) {
	return c.c.GuestSessionRatedTvEpisodesWithResponse(c.ctx, guestSessionId, params)
}

func (c *Tmdb) KeywordDetails(keywordId int32) (*KeywordDetailsResponse, error) {
	return c.c.KeywordDetailsWithResponse(c.ctx, keywordId)
}

func (c *Tmdb) KeywordMovies(keywordId string, params *KeywordMoviesParams) (*KeywordMoviesResponse, error) {
	return c.c.KeywordMoviesWithResponse(c.ctx, keywordId, params)
}

func (c *Tmdb) ListCreate(params *ListCreateParams, body ListCreateJSONRequestBody) (*ListCreateResponse, error) {
	return c.c.ListCreateWithResponse(c.ctx, params, body)
}

func (c *Tmdb) ListDelete(listId int32, params *ListDeleteParams) (*ListDeleteResponse, error) {
	return c.c.ListDeleteWithResponse(c.ctx, listId, params)
}

func (c *Tmdb) ListDetails(listId int32, params *ListDetailsParams) (*ListDetailsResponse, error) {
	return c.c.ListDetailsWithResponse(c.ctx, listId, params)
}

func (c *Tmdb) ListAddMovie(listId int32, params *ListAddMovieParams, body ListAddMovieJSONRequestBody) (*ListAddMovieResponse, error) {
	return c.c.ListAddMovieWithResponse(c.ctx, listId, params, body)
}

func (c *Tmdb) ListClear(listId int32, params *ListClearParams) (*ListClearResponse, error) {
	return c.c.ListClearWithResponse(c.ctx, listId, params)
}

func (c *Tmdb) ListCheckItemStatus(listId int32, params *ListCheckItemStatusParams) (*ListCheckItemStatusResponse, error) {
	return c.c.ListCheckItemStatusWithResponse(c.ctx, listId, params)
}

func (c *Tmdb) ListRemoveMovie(listId int32, params *ListRemoveMovieParams, body ListRemoveMovieJSONRequestBody) (*ListRemoveMovieResponse, error) {
	return c.c.ListRemoveMovieWithResponse(c.ctx, listId, params, body)
}

func (c *Tmdb) ChangesMovieList(params *ChangesMovieListParams) (*ChangesMovieListResponse, error) {
	return c.c.ChangesMovieListWithResponse(c.ctx, params)
}

func (c *Tmdb) MovieLatestId() (*MovieLatestIdResponse, error) {
	return c.c.MovieLatestIdWithResponse(c.ctx)
}

func (c *Tmdb) MovieNowPlayingList(params *MovieNowPlayingListParams) (*MovieNowPlayingListResponse, error) {
	return c.c.MovieNowPlayingListWithResponse(c.ctx, params)
}

func (c *Tmdb) MoviePopularList(params *MoviePopularListParams) (*MoviePopularListResponse, error) {
	return c.c.MoviePopularListWithResponse(c.ctx, params)
}

func (c *Tmdb) MovieTopRatedList(params *MovieTopRatedListParams) (*MovieTopRatedListResponse, error) {
	return c.c.MovieTopRatedListWithResponse(c.ctx, params)
}

func (c *Tmdb) MovieUpcomingList(params *MovieUpcomingListParams) (*MovieUpcomingListResponse, error) {
	return c.c.MovieUpcomingListWithResponse(c.ctx, params)
}

func (c *Tmdb) MovieDetails(movieId int32, params *MovieDetailsParams) (*MovieDetailsResponse, error) {
	return c.c.MovieDetailsWithResponse(c.ctx, movieId, params)
}

func (c *Tmdb) MovieAccountStates(movieId int32, params *MovieAccountStatesParams) (*MovieAccountStatesResponse, error) {
	return c.c.MovieAccountStatesWithResponse(c.ctx, movieId, params)
}

func (c *Tmdb) MovieAlternativeTitles(movieId int32, params *MovieAlternativeTitlesParams) (*MovieAlternativeTitlesResponse, error) {
	return c.c.MovieAlternativeTitlesWithResponse(c.ctx, movieId, params)
}

func (c *Tmdb) MovieChanges(movieId int32, params *MovieChangesParams) (*MovieChangesResponse, error) {
	return c.c.MovieChangesWithResponse(c.ctx, movieId, params)
}

func (c *Tmdb) MovieCredits(movieId int32, params *MovieCreditsParams) (*MovieCreditsResponse, error) {
	return c.c.MovieCreditsWithResponse(c.ctx, movieId, params)
}

func (c *Tmdb) MovieExternalIds(movieId int32) (*MovieExternalIdsResponse, error) {
	return c.c.MovieExternalIdsWithResponse(c.ctx, movieId)
}

func (c *Tmdb) MovieImages(movieId int32, params *MovieImagesParams) (*MovieImagesResponse, error) {
	return c.c.MovieImagesWithResponse(c.ctx, movieId, params)
}

func (c *Tmdb) MovieKeywords(movieId string) (*MovieKeywordsResponse, error) {
	return c.c.MovieKeywordsWithResponse(c.ctx, movieId)
}

func (c *Tmdb) MovieLists(movieId int32, params *MovieListsParams) (*MovieListsResponse, error) {
	return c.c.MovieListsWithResponse(c.ctx, movieId, params)
}

func (c *Tmdb) MovieDeleteRating(movieId int32, params *MovieDeleteRatingParams) (*MovieDeleteRatingResponse, error) {
	return c.c.MovieDeleteRatingWithResponse(c.ctx, movieId, params)
}

func (c *Tmdb) MovieAddRating(movieId int32, params *MovieAddRatingParams, body MovieAddRatingJSONRequestBody) (*MovieAddRatingResponse, error) {
	return c.c.MovieAddRatingWithResponse(c.ctx, movieId, params, body)
}

func (c *Tmdb) MovieRecommendations(movieId int32, params *MovieRecommendationsParams) (*MovieRecommendationsResponse, error) {
	return c.c.MovieRecommendationsWithResponse(c.ctx, movieId, params)
}

func (c *Tmdb) MovieReleaseDates(movieId int32) (*MovieReleaseDatesResponse, error) {
	return c.c.MovieReleaseDatesWithResponse(c.ctx, movieId)
}

func (c *Tmdb) MovieReviews(movieId int32, params *MovieReviewsParams) (*MovieReviewsResponse, error) {
	return c.c.MovieReviewsWithResponse(c.ctx, movieId, params)
}

func (c *Tmdb) MovieSimilar(movieId int32, params *MovieSimilarParams) (*MovieSimilarResponse, error) {
	return c.c.MovieSimilarWithResponse(c.ctx, movieId, params)
}

func (c *Tmdb) MovieTranslations(movieId int32) (*MovieTranslationsResponse, error) {
	return c.c.MovieTranslationsWithResponse(c.ctx, movieId)
}

func (c *Tmdb) MovieVideos(movieId int32, params *MovieVideosParams) (*MovieVideosResponse, error) {
	return c.c.MovieVideosWithResponse(c.ctx, movieId, params)
}

func (c *Tmdb) MovieWatchProviders(movieId int32) (*MovieWatchProvidersResponse, error) {
	return c.c.MovieWatchProvidersWithResponse(c.ctx, movieId)
}

func (c *Tmdb) NetworkDetails(networkId int32) (*NetworkDetailsResponse, error) {
	return c.c.NetworkDetailsWithResponse(c.ctx, networkId)
}

func (c *Tmdb) DetailsCopy(networkId int32) (*DetailsCopyResponse, error) {
	return c.c.DetailsCopyWithResponse(c.ctx, networkId)
}

func (c *Tmdb) AlternativeNamesCopy(networkId int32) (*AlternativeNamesCopyResponse, error) {
	return c.c.AlternativeNamesCopyWithResponse(c.ctx, networkId)
}

func (c *Tmdb) ChangesPeopleList(params *ChangesPeopleListParams) (*ChangesPeopleListResponse, error) {
	return c.c.ChangesPeopleListWithResponse(c.ctx, params)
}

func (c *Tmdb) PersonLatestId() (*PersonLatestIdResponse, error) {
	return c.c.PersonLatestIdWithResponse(c.ctx)
}

func (c *Tmdb) PersonPopularList(params *PersonPopularListParams) (*PersonPopularListResponse, error) {
	return c.c.PersonPopularListWithResponse(c.ctx, params)
}

func (c *Tmdb) PersonDetails(personId int32, params *PersonDetailsParams) (*PersonDetailsResponse, error) {
	return c.c.PersonDetailsWithResponse(c.ctx, personId, params)
}

func (c *Tmdb) PersonChanges(personId int32, params *PersonChangesParams) (*PersonChangesResponse, error) {
	return c.c.PersonChangesWithResponse(c.ctx, personId, params)
}

func (c *Tmdb) PersonCombinedCredits(personId string, params *PersonCombinedCreditsParams) (*PersonCombinedCreditsResponse, error) {
	return c.c.PersonCombinedCreditsWithResponse(c.ctx, personId, params)
}

func (c *Tmdb) PersonExternalIds(personId int32) (*PersonExternalIdsResponse, error) {
	return c.c.PersonExternalIdsWithResponse(c.ctx, personId)
}

func (c *Tmdb) PersonImages(personId int32) (*PersonImagesResponse, error) {
	return c.c.PersonImagesWithResponse(c.ctx, personId)
}

func (c *Tmdb) PersonMovieCredits(personId int32, params *PersonMovieCreditsParams) (*PersonMovieCreditsResponse, error) {
	return c.c.PersonMovieCreditsWithResponse(c.ctx, personId, params)
}

func (c *Tmdb) PersonTaggedImages(personId int32, params *PersonTaggedImagesParams) (*PersonTaggedImagesResponse, error) {
	return c.c.PersonTaggedImagesWithResponse(c.ctx, personId, params)
}

func (c *Tmdb) Translations(personId int32) (*TranslationsResponse, error) {
	return c.c.TranslationsWithResponse(c.ctx, personId)
}

func (c *Tmdb) PersonTvCredits(personId int32, params *PersonTvCreditsParams) (*PersonTvCreditsResponse, error) {
	return c.c.PersonTvCreditsWithResponse(c.ctx, personId, params)
}

func (c *Tmdb) ReviewDetails(reviewId string) (*ReviewDetailsResponse, error) {
	return c.c.ReviewDetailsWithResponse(c.ctx, reviewId)
}

func (c *Tmdb) SearchCollection(params *SearchCollectionParams) (*SearchCollectionResponse, error) {
	return c.c.SearchCollectionWithResponse(c.ctx, params)
}

func (c *Tmdb) SearchCompany(params *SearchCompanyParams) (*SearchCompanyResponse, error) {
	return c.c.SearchCompanyWithResponse(c.ctx, params)
}

func (c *Tmdb) SearchKeyword(params *SearchKeywordParams) (*SearchKeywordResponse, error) {
	return c.c.SearchKeywordWithResponse(c.ctx, params)
}

func (c *Tmdb) SearchMovie(params *SearchMovieParams) (*SearchMovieResponse, error) {
	return c.c.SearchMovieWithResponse(c.ctx, params)
}

func (c *Tmdb) SearchMulti(params *SearchMultiParams) (*SearchMultiResponse, error) {
	return c.c.SearchMultiWithResponse(c.ctx, params)
}

func (c *Tmdb) SearchPerson(params *SearchPersonParams) (*SearchPersonResponse, error) {
	return c.c.SearchPersonWithResponse(c.ctx, params)
}

func (c *Tmdb) SearchTv(params *SearchTvParams) (*SearchTvResponse, error) {
	return c.c.SearchTvWithResponse(c.ctx, params)
}

func (c *Tmdb) TrendingAll(timeWindow TrendingAllParamsTimeWindow, params *TrendingAllParams) (*TrendingAllResponse, error) {
	return c.c.TrendingAllWithResponse(c.ctx, timeWindow, params)
}

func (c *Tmdb) TrendingMovies(timeWindow TrendingMoviesParamsTimeWindow, params *TrendingMoviesParams) (*TrendingMoviesResponse, error) {
	return c.c.TrendingMoviesWithResponse(c.ctx, timeWindow, params)
}

func (c *Tmdb) TrendingPeople(timeWindow TrendingPeopleParamsTimeWindow, params *TrendingPeopleParams) (*TrendingPeopleResponse, error) {
	return c.c.TrendingPeopleWithResponse(c.ctx, timeWindow, params)
}

func (c *Tmdb) TrendingTv(timeWindow TrendingTvParamsTimeWindow, params *TrendingTvParams) (*TrendingTvResponse, error) {
	return c.c.TrendingTvWithResponse(c.ctx, timeWindow, params)
}

func (c *Tmdb) TvSeriesAiringTodayList(params *TvSeriesAiringTodayListParams) (*TvSeriesAiringTodayListResponse, error) {
	return c.c.TvSeriesAiringTodayListWithResponse(c.ctx, params)
}

func (c *Tmdb) ChangesTvList(params *ChangesTvListParams) (*ChangesTvListResponse, error) {
	return c.c.ChangesTvListWithResponse(c.ctx, params)
}

func (c *Tmdb) TvEpisodeChangesById(episodeId int32) (*TvEpisodeChangesByIdResponse, error) {
	return c.c.TvEpisodeChangesByIdWithResponse(c.ctx, episodeId)
}

func (c *Tmdb) TvEpisodeGroupDetails(tvEpisodeGroupId string) (*TvEpisodeGroupDetailsResponse, error) {
	return c.c.TvEpisodeGroupDetailsWithResponse(c.ctx, tvEpisodeGroupId)
}

func (c *Tmdb) TvSeriesLatestId() (*TvSeriesLatestIdResponse, error) {
	return c.c.TvSeriesLatestIdWithResponse(c.ctx)
}

func (c *Tmdb) TvSeriesOnTheAirList(params *TvSeriesOnTheAirListParams) (*TvSeriesOnTheAirListResponse, error) {
	return c.c.TvSeriesOnTheAirListWithResponse(c.ctx, params)
}

func (c *Tmdb) TvSeriesPopularList(params *TvSeriesPopularListParams) (*TvSeriesPopularListResponse, error) {
	return c.c.TvSeriesPopularListWithResponse(c.ctx, params)
}

func (c *Tmdb) TvSeasonChangesById(seasonId int32, params *TvSeasonChangesByIdParams) (*TvSeasonChangesByIdResponse, error) {
	return c.c.TvSeasonChangesByIdWithResponse(c.ctx, seasonId, params)
}

func (c *Tmdb) TvSeriesTopRatedList(params *TvSeriesTopRatedListParams) (*TvSeriesTopRatedListResponse, error) {
	return c.c.TvSeriesTopRatedListWithResponse(c.ctx, params)
}

func (c *Tmdb) TvSeriesDetails(seriesId int32, params *TvSeriesDetailsParams) (*TvSeriesDetailsResponse, error) {
	return c.c.TvSeriesDetailsWithResponse(c.ctx, seriesId, params)
}

func (c *Tmdb) TvSeriesAccountStates(seriesId int32, params *TvSeriesAccountStatesParams) (*TvSeriesAccountStatesResponse, error) {
	return c.c.TvSeriesAccountStatesWithResponse(c.ctx, seriesId, params)
}

func (c *Tmdb) TvSeriesAggregateCredits(seriesId int32, params *TvSeriesAggregateCreditsParams) (*TvSeriesAggregateCreditsResponse, error) {
	return c.c.TvSeriesAggregateCreditsWithResponse(c.ctx, seriesId, params)
}

func (c *Tmdb) TvSeriesAlternativeTitles(seriesId int32) (*TvSeriesAlternativeTitlesResponse, error) {
	return c.c.TvSeriesAlternativeTitlesWithResponse(c.ctx, seriesId)
}

func (c *Tmdb) TvSeriesChanges(seriesId int32, params *TvSeriesChangesParams) (*TvSeriesChangesResponse, error) {
	return c.c.TvSeriesChangesWithResponse(c.ctx, seriesId, params)
}

func (c *Tmdb) TvSeriesContentRatings(seriesId int32) (*TvSeriesContentRatingsResponse, error) {
	return c.c.TvSeriesContentRatingsWithResponse(c.ctx, seriesId)
}

func (c *Tmdb) TvSeriesCredits(seriesId int32, params *TvSeriesCreditsParams) (*TvSeriesCreditsResponse, error) {
	return c.c.TvSeriesCreditsWithResponse(c.ctx, seriesId, params)
}

func (c *Tmdb) TvSeriesEpisodeGroups(seriesId int32) (*TvSeriesEpisodeGroupsResponse, error) {
	return c.c.TvSeriesEpisodeGroupsWithResponse(c.ctx, seriesId)
}

func (c *Tmdb) TvSeriesExternalIds(seriesId int32) (*TvSeriesExternalIdsResponse, error) {
	return c.c.TvSeriesExternalIdsWithResponse(c.ctx, seriesId)
}

func (c *Tmdb) TvSeriesImages(seriesId int32, params *TvSeriesImagesParams) (*TvSeriesImagesResponse, error) {
	return c.c.TvSeriesImagesWithResponse(c.ctx, seriesId, params)
}

func (c *Tmdb) TvSeriesKeywords(seriesId int32) (*TvSeriesKeywordsResponse, error) {
	return c.c.TvSeriesKeywordsWithResponse(c.ctx, seriesId)
}

func (c *Tmdb) TvSeriesDeleteRating(seriesId int32, params *TvSeriesDeleteRatingParams) (*TvSeriesDeleteRatingResponse, error) {
	return c.c.TvSeriesDeleteRatingWithResponse(c.ctx, seriesId, params)
}

func (c *Tmdb) TvSeriesAddRating(seriesId int32, params *TvSeriesAddRatingParams, body TvSeriesAddRatingJSONRequestBody) (*TvSeriesAddRatingResponse, error) {
	return c.c.TvSeriesAddRatingWithResponse(c.ctx, seriesId, params, body)
}

func (c *Tmdb) TvSeriesRecommendations(seriesId int32, params *TvSeriesRecommendationsParams) (*TvSeriesRecommendationsResponse, error) {
	return c.c.TvSeriesRecommendationsWithResponse(c.ctx, seriesId, params)
}

func (c *Tmdb) TvSeriesReviews(seriesId int32, params *TvSeriesReviewsParams) (*TvSeriesReviewsResponse, error) {
	return c.c.TvSeriesReviewsWithResponse(c.ctx, seriesId, params)
}

func (c *Tmdb) TvSeriesScreenedTheatrically(seriesId int32) (*TvSeriesScreenedTheatricallyResponse, error) {
	return c.c.TvSeriesScreenedTheatricallyWithResponse(c.ctx, seriesId)
}

func (c *Tmdb) TvSeasonDetails(seriesId int32, seasonNumber int32, params *TvSeasonDetailsParams) (*TvSeasonDetailsResponse, error) {
	return c.c.TvSeasonDetailsWithResponse(c.ctx, seriesId, seasonNumber, params)
}

func (c *Tmdb) TvSeasonAccountStates(seriesId int32, seasonNumber int32, params *TvSeasonAccountStatesParams) (*TvSeasonAccountStatesResponse, error) {
	return c.c.TvSeasonAccountStatesWithResponse(c.ctx, seriesId, seasonNumber, params)
}

func (c *Tmdb) TvSeasonAggregateCredits(seriesId int32, seasonNumber int32, params *TvSeasonAggregateCreditsParams) (*TvSeasonAggregateCreditsResponse, error) {
	return c.c.TvSeasonAggregateCreditsWithResponse(c.ctx, seriesId, seasonNumber, params)
}

func (c *Tmdb) TvSeasonCredits(seriesId int32, seasonNumber int32, params *TvSeasonCreditsParams) (*TvSeasonCreditsResponse, error) {
	return c.c.TvSeasonCreditsWithResponse(c.ctx, seriesId, seasonNumber, params)
}

func (c *Tmdb) TvEpisodeDetails(seriesId int32, seasonNumber int32, episodeNumber int32, params *TvEpisodeDetailsParams) (*TvEpisodeDetailsResponse, error) {
	return c.c.TvEpisodeDetailsWithResponse(c.ctx, seriesId, seasonNumber, episodeNumber, params)
}

func (c *Tmdb) TvEpisodeAccountStates(seriesId int32, seasonNumber int32, episodeNumber int32, params *TvEpisodeAccountStatesParams) (*TvEpisodeAccountStatesResponse, error) {
	return c.c.TvEpisodeAccountStatesWithResponse(c.ctx, seriesId, seasonNumber, episodeNumber, params)
}

func (c *Tmdb) TvEpisodeCredits(seriesId int32, seasonNumber int32, episodeNumber int32, params *TvEpisodeCreditsParams) (*TvEpisodeCreditsResponse, error) {
	return c.c.TvEpisodeCreditsWithResponse(c.ctx, seriesId, seasonNumber, episodeNumber, params)
}

func (c *Tmdb) TvEpisodeExternalIds(seriesId int32, seasonNumber int32, episodeNumber string) (*TvEpisodeExternalIdsResponse, error) {
	return c.c.TvEpisodeExternalIdsWithResponse(c.ctx, seriesId, seasonNumber, episodeNumber)
}

func (c *Tmdb) TvEpisodeImages(seriesId int32, seasonNumber int32, episodeNumber int32, params *TvEpisodeImagesParams) (*TvEpisodeImagesResponse, error) {
	return c.c.TvEpisodeImagesWithResponse(c.ctx, seriesId, seasonNumber, episodeNumber, params)
}

func (c *Tmdb) TvEpisodeDeleteRating(seriesId int32, seasonNumber int32, episodeNumber int32, params *TvEpisodeDeleteRatingParams) (*TvEpisodeDeleteRatingResponse, error) {
	return c.c.TvEpisodeDeleteRatingWithResponse(c.ctx, seriesId, seasonNumber, episodeNumber, params)
}

func (c *Tmdb) TvEpisodeAddRating(seriesId int32, seasonNumber int32, episodeNumber int32, params *TvEpisodeAddRatingParams, body TvEpisodeAddRatingJSONRequestBody) (*TvEpisodeAddRatingResponse, error) {
	return c.c.TvEpisodeAddRatingWithResponse(c.ctx, seriesId, seasonNumber, episodeNumber, params, body)
}

func (c *Tmdb) TvEpisodeTranslations(seriesId int32, seasonNumber int32, episodeNumber int32) (*TvEpisodeTranslationsResponse, error) {
	return c.c.TvEpisodeTranslationsWithResponse(c.ctx, seriesId, seasonNumber, episodeNumber)
}

func (c *Tmdb) TvEpisodeVideos(seriesId int32, seasonNumber int32, episodeNumber int32, params *TvEpisodeVideosParams) (*TvEpisodeVideosResponse, error) {
	return c.c.TvEpisodeVideosWithResponse(c.ctx, seriesId, seasonNumber, episodeNumber, params)
}

func (c *Tmdb) TvSeasonExternalIds(seriesId int32, seasonNumber int32) (*TvSeasonExternalIdsResponse, error) {
	return c.c.TvSeasonExternalIdsWithResponse(c.ctx, seriesId, seasonNumber)
}

func (c *Tmdb) TvSeasonImages(seriesId int32, seasonNumber int32, params *TvSeasonImagesParams) (*TvSeasonImagesResponse, error) {
	return c.c.TvSeasonImagesWithResponse(c.ctx, seriesId, seasonNumber, params)
}

func (c *Tmdb) TvSeasonTranslations(seriesId int32, seasonNumber int32) (*TvSeasonTranslationsResponse, error) {
	return c.c.TvSeasonTranslationsWithResponse(c.ctx, seriesId, seasonNumber)
}

func (c *Tmdb) TvSeasonVideos(seriesId int32, seasonNumber int32, params *TvSeasonVideosParams) (*TvSeasonVideosResponse, error) {
	return c.c.TvSeasonVideosWithResponse(c.ctx, seriesId, seasonNumber, params)
}

func (c *Tmdb) TvSeasonWatchProviders(seriesId int32, seasonNumber int32, params *TvSeasonWatchProvidersParams) (*TvSeasonWatchProvidersResponse, error) {
	return c.c.TvSeasonWatchProvidersWithResponse(c.ctx, seriesId, seasonNumber, params)
}

func (c *Tmdb) TvSeriesSimilar(seriesId string, params *TvSeriesSimilarParams) (*TvSeriesSimilarResponse, error) {
	return c.c.TvSeriesSimilarWithResponse(c.ctx, seriesId, params)
}

func (c *Tmdb) TvSeriesTranslations(seriesId int32) (*TvSeriesTranslationsResponse, error) {
	return c.c.TvSeriesTranslationsWithResponse(c.ctx, seriesId)
}

func (c *Tmdb) TvSeriesVideos(seriesId int32, params *TvSeriesVideosParams) (*TvSeriesVideosResponse, error) {
	return c.c.TvSeriesVideosWithResponse(c.ctx, seriesId, params)
}

func (c *Tmdb) TvSeriesWatchProviders(seriesId int32) (*TvSeriesWatchProvidersResponse, error) {
	return c.c.TvSeriesWatchProvidersWithResponse(c.ctx, seriesId)
}

func (c *Tmdb) WatchProvidersMovieList(params *WatchProvidersMovieListParams) (*WatchProvidersMovieListResponse, error) {
	return c.c.WatchProvidersMovieListWithResponse(c.ctx, params)
}

func (c *Tmdb) WatchProvidersAvailableRegions(params *WatchProvidersAvailableRegionsParams) (*WatchProvidersAvailableRegionsResponse, error) {
	return c.c.WatchProvidersAvailableRegionsWithResponse(c.ctx, params)
}

func (c *Tmdb) WatchProviderTvList(params *WatchProviderTvListParams) (*WatchProviderTvListResponse, error) {
	return c.c.WatchProviderTvListWithResponse(c.ctx, params)
}
