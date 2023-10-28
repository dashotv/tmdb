package tmdb

import (
	"github.com/pkg/errors"

	"github.com/dashotv/tmdb/openapi/models/operations"
	"github.com/dashotv/tmdb/openapi/types"
)

// AccountAddFavorite wraps the generated openapi.SDK.AccountAddFavorite call
func (c *Client) AccountAddFavorite(accountID int, requestBody *operations.AccountAddFavoriteRequestBody, sessionID *string) (*AccountAddFavoriteResponse, error) {
	r, err := c.sdk.AccountAddFavorite(c.ctx, accountID, requestBody, sessionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AccountAddFavorite200ApplicationJSONObject, nil
}

// AccountAddToWatchlist wraps the generated openapi.SDK.AccountAddToWatchlist call
func (c *Client) AccountAddToWatchlist(accountID int, requestBody *operations.AccountAddToWatchlistRequestBody, sessionID *string) (*AccountAddToWatchlistResponse, error) {
	r, err := c.sdk.AccountAddToWatchlist(c.ctx, accountID, requestBody, sessionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AccountAddToWatchlist200ApplicationJSONObject, nil
}

// AccountDetails wraps the generated openapi.SDK.AccountDetails call
func (c *Client) AccountDetails(accountID int, sessionID *string) (*AccountDetailsResponse, error) {
	r, err := c.sdk.AccountDetails(c.ctx, accountID, sessionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AccountDetails200ApplicationJSONObject, nil
}

// AccountFavoriteTv wraps the generated openapi.SDK.AccountFavoriteTv call
func (c *Client) AccountFavoriteTv(request operations.AccountFavoriteTvRequest) (*AccountFavoriteTvResponse, error) {
	r, err := c.sdk.AccountFavoriteTv(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AccountFavoriteTv200ApplicationJSONObject, nil
}

// AccountGetFavorites wraps the generated openapi.SDK.AccountGetFavorites call
func (c *Client) AccountGetFavorites(request operations.AccountGetFavoritesRequest) (*AccountGetFavoritesResponse, error) {
	r, err := c.sdk.AccountGetFavorites(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AccountGetFavorites200ApplicationJSONObject, nil
}

// AccountLists wraps the generated openapi.SDK.AccountLists call
func (c *Client) AccountLists(accountID int, page *int, sessionID *string) (*AccountListsResponse, error) {
	r, err := c.sdk.AccountLists(c.ctx, accountID, page, sessionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AccountLists200ApplicationJSONObject, nil
}

// AccountRatedMovies wraps the generated openapi.SDK.AccountRatedMovies call
func (c *Client) AccountRatedMovies(request operations.AccountRatedMoviesRequest) (*AccountRatedMoviesResponse, error) {
	r, err := c.sdk.AccountRatedMovies(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AccountRatedMovies200ApplicationJSONObject, nil
}

// AccountRatedTv wraps the generated openapi.SDK.AccountRatedTv call
func (c *Client) AccountRatedTv(request operations.AccountRatedTvRequest) (*AccountRatedTvResponse, error) {
	r, err := c.sdk.AccountRatedTv(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AccountRatedTv200ApplicationJSONObject, nil
}

// AccountRatedTvEpisodes wraps the generated openapi.SDK.AccountRatedTvEpisodes call
func (c *Client) AccountRatedTvEpisodes(request operations.AccountRatedTvEpisodesRequest) (*AccountRatedTvEpisodesResponse, error) {
	r, err := c.sdk.AccountRatedTvEpisodes(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AccountRatedTvEpisodes200ApplicationJSONObject, nil
}

// AccountWatchlistMovies wraps the generated openapi.SDK.AccountWatchlistMovies call
func (c *Client) AccountWatchlistMovies(request operations.AccountWatchlistMoviesRequest) (*AccountWatchlistMoviesResponse, error) {
	r, err := c.sdk.AccountWatchlistMovies(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AccountWatchlistMovies200ApplicationJSONObject, nil
}

// AccountWatchlistTv wraps the generated openapi.SDK.AccountWatchlistTv call
func (c *Client) AccountWatchlistTv(request operations.AccountWatchlistTvRequest) (*AccountWatchlistTvResponse, error) {
	r, err := c.sdk.AccountWatchlistTv(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AccountWatchlistTv200ApplicationJSONObject, nil
}

// AlternativeNamesCopy wraps the generated openapi.SDK.AlternativeNamesCopy call
func (c *Client) AlternativeNamesCopy(networkID int) (*AlternativeNamesCopyResponse, error) {
	r, err := c.sdk.AlternativeNamesCopy(c.ctx, networkID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AlternativeNamesCopy200ApplicationJSONObject, nil
}

// AuthenticationCreateGuestSession wraps the generated openapi.SDK.AuthenticationCreateGuestSession call
func (c *Client) AuthenticationCreateGuestSession() (*AuthenticationCreateGuestSessionResponse, error) {
	r, err := c.sdk.AuthenticationCreateGuestSession(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AuthenticationCreateGuestSession200ApplicationJSONObject, nil
}

// AuthenticationCreateRequestToken wraps the generated openapi.SDK.AuthenticationCreateRequestToken call
func (c *Client) AuthenticationCreateRequestToken() (*AuthenticationCreateRequestTokenResponse, error) {
	r, err := c.sdk.AuthenticationCreateRequestToken(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AuthenticationCreateRequestToken200ApplicationJSONObject, nil
}

// AuthenticationCreateSession wraps the generated openapi.SDK.AuthenticationCreateSession call
func (c *Client) AuthenticationCreateSession(request *operations.AuthenticationCreateSessionRequestBody) (*AuthenticationCreateSessionResponse, error) {
	r, err := c.sdk.AuthenticationCreateSession(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AuthenticationCreateSession200ApplicationJSONObject, nil
}

// AuthenticationCreateSessionFromLogin wraps the generated openapi.SDK.AuthenticationCreateSessionFromLogin call
func (c *Client) AuthenticationCreateSessionFromLogin(request *operations.AuthenticationCreateSessionFromLoginRequestBody) (*AuthenticationCreateSessionFromLoginResponse, error) {
	r, err := c.sdk.AuthenticationCreateSessionFromLogin(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AuthenticationCreateSessionFromLogin200ApplicationJSONObject, nil
}

// AuthenticationCreateSessionFromV4Token wraps the generated openapi.SDK.AuthenticationCreateSessionFromV4Token call
func (c *Client) AuthenticationCreateSessionFromV4Token(request *operations.AuthenticationCreateSessionFromV4TokenRequestBody) (*AuthenticationCreateSessionFromV4TokenResponse, error) {
	r, err := c.sdk.AuthenticationCreateSessionFromV4Token(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AuthenticationCreateSessionFromV4Token200ApplicationJSONObject, nil
}

// AuthenticationDeleteSession wraps the generated openapi.SDK.AuthenticationDeleteSession call
func (c *Client) AuthenticationDeleteSession(request *operations.AuthenticationDeleteSessionRequestBody) (*AuthenticationDeleteSessionResponse, error) {
	r, err := c.sdk.AuthenticationDeleteSession(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AuthenticationDeleteSession200ApplicationJSONObject, nil
}

// AuthenticationValidateKey wraps the generated openapi.SDK.AuthenticationValidateKey call
func (c *Client) AuthenticationValidateKey() (*AuthenticationValidateKeyResponse, error) {
	r, err := c.sdk.AuthenticationValidateKey(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.AuthenticationValidateKey200ApplicationJSONObject, nil
}

// CertificationMovieList wraps the generated openapi.SDK.CertificationMovieList call
func (c *Client) CertificationMovieList() (*CertificationMovieListResponse, error) {
	r, err := c.sdk.CertificationMovieList(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.CertificationMovieList200ApplicationJSONObject, nil
}

// CertificationsTvList wraps the generated openapi.SDK.CertificationsTvList call
func (c *Client) CertificationsTvList() (*CertificationsTvListResponse, error) {
	r, err := c.sdk.CertificationsTvList(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.CertificationsTvList200ApplicationJSONObject, nil
}

// ChangesMovieList wraps the generated openapi.SDK.ChangesMovieList call
func (c *Client) ChangesMovieList(endDate *types.Date, page *int, startDate *types.Date) (*ChangesMovieListResponse, error) {
	r, err := c.sdk.ChangesMovieList(c.ctx, endDate, page, startDate)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.ChangesMovieList200ApplicationJSONObject, nil
}

// ChangesPeopleList wraps the generated openapi.SDK.ChangesPeopleList call
func (c *Client) ChangesPeopleList(endDate *types.Date, page *int, startDate *types.Date) (*ChangesPeopleListResponse, error) {
	r, err := c.sdk.ChangesPeopleList(c.ctx, endDate, page, startDate)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.ChangesPeopleList200ApplicationJSONObject, nil
}

// ChangesTvList wraps the generated openapi.SDK.ChangesTvList call
func (c *Client) ChangesTvList(endDate *types.Date, page *int, startDate *types.Date) (*ChangesTvListResponse, error) {
	r, err := c.sdk.ChangesTvList(c.ctx, endDate, page, startDate)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.ChangesTvList200ApplicationJSONObject, nil
}

// CollectionDetails wraps the generated openapi.SDK.CollectionDetails call
func (c *Client) CollectionDetails(collectionID int, language *string) (*CollectionDetailsResponse, error) {
	r, err := c.sdk.CollectionDetails(c.ctx, collectionID, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.CollectionDetails200ApplicationJSONObject, nil
}

// CollectionImages wraps the generated openapi.SDK.CollectionImages call
func (c *Client) CollectionImages(collectionID int, includeImageLanguage *string, language *string) (*CollectionImagesResponse, error) {
	r, err := c.sdk.CollectionImages(c.ctx, collectionID, includeImageLanguage, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.CollectionImages200ApplicationJSONObject, nil
}

// CollectionTranslations wraps the generated openapi.SDK.CollectionTranslations call
func (c *Client) CollectionTranslations(collectionID int) (*CollectionTranslationsResponse, error) {
	r, err := c.sdk.CollectionTranslations(c.ctx, collectionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.CollectionTranslations200ApplicationJSONObject, nil
}

// CompanyAlternativeNames wraps the generated openapi.SDK.CompanyAlternativeNames call
func (c *Client) CompanyAlternativeNames(companyID int) (*CompanyAlternativeNamesResponse, error) {
	r, err := c.sdk.CompanyAlternativeNames(c.ctx, companyID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.CompanyAlternativeNames200ApplicationJSONObject, nil
}

// CompanyDetails wraps the generated openapi.SDK.CompanyDetails call
func (c *Client) CompanyDetails(companyID int) (*CompanyDetailsResponse, error) {
	r, err := c.sdk.CompanyDetails(c.ctx, companyID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.CompanyDetails200ApplicationJSONObject, nil
}

// CompanyImages wraps the generated openapi.SDK.CompanyImages call
func (c *Client) CompanyImages(companyID int) (*CompanyImagesResponse, error) {
	r, err := c.sdk.CompanyImages(c.ctx, companyID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.CompanyImages200ApplicationJSONObject, nil
}

// ConfigurationDetails wraps the generated openapi.SDK.ConfigurationDetails call
func (c *Client) ConfigurationDetails() (*ConfigurationDetailsResponse, error) {
	r, err := c.sdk.ConfigurationDetails(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.ConfigurationDetails200ApplicationJSONObject, nil
}

// CreditDetails wraps the generated openapi.SDK.CreditDetails call
func (c *Client) CreditDetails(creditID string) (*CreditDetailsResponse, error) {
	r, err := c.sdk.CreditDetails(c.ctx, creditID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.CreditDetails200ApplicationJSONObject, nil
}

// DetailsCopy wraps the generated openapi.SDK.DetailsCopy call
func (c *Client) DetailsCopy(networkID int) (*DetailsCopyResponse, error) {
	r, err := c.sdk.DetailsCopy(c.ctx, networkID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.DetailsCopy200ApplicationJSONObject, nil
}

// DiscoverMovie wraps the generated openapi.SDK.DiscoverMovie call
func (c *Client) DiscoverMovie(request operations.DiscoverMovieRequest) (*DiscoverMovieResponse, error) {
	r, err := c.sdk.DiscoverMovie(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.DiscoverMovie200ApplicationJSONObject, nil
}

// DiscoverTv wraps the generated openapi.SDK.DiscoverTv call
func (c *Client) DiscoverTv(request operations.DiscoverTvRequest) (*DiscoverTvResponse, error) {
	r, err := c.sdk.DiscoverTv(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.DiscoverTv200ApplicationJSONObject, nil
}

// FindByID wraps the generated openapi.SDK.FindByID call
func (c *Client) FindByID(externalID string, externalSource operations.FindByIDExternalSource, language *string) (*FindByIDResponse, error) {
	r, err := c.sdk.FindByID(c.ctx, externalID, externalSource, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.FindByID200ApplicationJSONObject, nil
}

// GenreMovieList wraps the generated openapi.SDK.GenreMovieList call
func (c *Client) GenreMovieList(language *string) (*GenreMovieListResponse, error) {
	r, err := c.sdk.GenreMovieList(c.ctx, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GenreMovieList200ApplicationJSONObject, nil
}

// GenreTvList wraps the generated openapi.SDK.GenreTvList call
func (c *Client) GenreTvList(language *string) (*GenreTvListResponse, error) {
	r, err := c.sdk.GenreTvList(c.ctx, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GenreTvList200ApplicationJSONObject, nil
}

// GuestSessionRatedMovies wraps the generated openapi.SDK.GuestSessionRatedMovies call
func (c *Client) GuestSessionRatedMovies(guestSessionID string, language *string, page *int, sortBy *operations.GuestSessionRatedMoviesSortBy) (*GuestSessionRatedMoviesResponse, error) {
	r, err := c.sdk.GuestSessionRatedMovies(c.ctx, guestSessionID, language, page, sortBy)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GuestSessionRatedMovies200ApplicationJSONObject, nil
}

// GuestSessionRatedTv wraps the generated openapi.SDK.GuestSessionRatedTv call
func (c *Client) GuestSessionRatedTv(guestSessionID string, language *string, page *int, sortBy *operations.GuestSessionRatedTvSortBy) (*GuestSessionRatedTvResponse, error) {
	r, err := c.sdk.GuestSessionRatedTv(c.ctx, guestSessionID, language, page, sortBy)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GuestSessionRatedTv200ApplicationJSONObject, nil
}

// GuestSessionRatedTvEpisodes wraps the generated openapi.SDK.GuestSessionRatedTvEpisodes call
func (c *Client) GuestSessionRatedTvEpisodes(guestSessionID string, language *string, page *int, sortBy *operations.GuestSessionRatedTvEpisodesSortBy) (*GuestSessionRatedTvEpisodesResponse, error) {
	r, err := c.sdk.GuestSessionRatedTvEpisodes(c.ctx, guestSessionID, language, page, sortBy)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GuestSessionRatedTvEpisodes200ApplicationJSONObject, nil
}

// KeywordDetails wraps the generated openapi.SDK.KeywordDetails call
func (c *Client) KeywordDetails(keywordID int) (*KeywordDetailsResponse, error) {
	r, err := c.sdk.KeywordDetails(c.ctx, keywordID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.KeywordDetails200ApplicationJSONObject, nil
}

// KeywordMovies wraps the generated openapi.SDK.KeywordMovies call
func (c *Client) KeywordMovies(keywordID string, includeAdult *bool, language *string, page *int) (*KeywordMoviesResponse, error) {
	r, err := c.sdk.KeywordMovies(c.ctx, keywordID, includeAdult, language, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.KeywordMovies200ApplicationJSONObject, nil
}

// ListAddMovie wraps the generated openapi.SDK.ListAddMovie call
func (c *Client) ListAddMovie(listID int, sessionID string, requestBody *operations.ListAddMovieRequestBody) (*ListAddMovieResponse, error) {
	r, err := c.sdk.ListAddMovie(c.ctx, listID, sessionID, requestBody)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.ListAddMovie200ApplicationJSONObject, nil
}

// ListCheckItemStatus wraps the generated openapi.SDK.ListCheckItemStatus call
func (c *Client) ListCheckItemStatus(listID int, language *string, movieID *int) (*ListCheckItemStatusResponse, error) {
	r, err := c.sdk.ListCheckItemStatus(c.ctx, listID, language, movieID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.ListCheckItemStatus200ApplicationJSONObject, nil
}

// ListClear wraps the generated openapi.SDK.ListClear call
func (c *Client) ListClear(confirm bool, listID int, sessionID string) (*ListClearResponse, error) {
	r, err := c.sdk.ListClear(c.ctx, confirm, listID, sessionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.ListClear200ApplicationJSONObject, nil
}

// ListCreate wraps the generated openapi.SDK.ListCreate call
func (c *Client) ListCreate(sessionID string, requestBody *operations.ListCreateRequestBody) (*ListCreateResponse, error) {
	r, err := c.sdk.ListCreate(c.ctx, sessionID, requestBody)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.ListCreate200ApplicationJSONObject, nil
}

// ListDelete wraps the generated openapi.SDK.ListDelete call
func (c *Client) ListDelete(listID int, sessionID string) (*ListDeleteResponse, error) {
	r, err := c.sdk.ListDelete(c.ctx, listID, sessionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.ListDelete200ApplicationJSONObject, nil
}

// ListDetails wraps the generated openapi.SDK.ListDetails call
func (c *Client) ListDetails(listID int, language *string) (*ListDetailsResponse, error) {
	r, err := c.sdk.ListDetails(c.ctx, listID, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.ListDetails200ApplicationJSONObject, nil
}

// ListRemoveMovie wraps the generated openapi.SDK.ListRemoveMovie call
func (c *Client) ListRemoveMovie(listID int, sessionID string, requestBody *operations.ListRemoveMovieRequestBody) (*ListRemoveMovieResponse, error) {
	r, err := c.sdk.ListRemoveMovie(c.ctx, listID, sessionID, requestBody)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.ListRemoveMovie200ApplicationJSONObject, nil
}

// MovieAccountStates wraps the generated openapi.SDK.MovieAccountStates call
func (c *Client) MovieAccountStates(movieID int, guestSessionID *string, sessionID *string) (*MovieAccountStatesResponse, error) {
	r, err := c.sdk.MovieAccountStates(c.ctx, movieID, guestSessionID, sessionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieAccountStates200ApplicationJSONObject, nil
}

// MovieAddRating wraps the generated openapi.SDK.MovieAddRating call
func (c *Client) MovieAddRating(movieID int, requestBody *operations.MovieAddRatingRequestBody, guestSessionID *string, sessionID *string) (*MovieAddRatingResponse, error) {
	r, err := c.sdk.MovieAddRating(c.ctx, movieID, requestBody, guestSessionID, sessionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieAddRating200ApplicationJSONObject, nil
}

// MovieAlternativeTitles wraps the generated openapi.SDK.MovieAlternativeTitles call
func (c *Client) MovieAlternativeTitles(movieID int, country *string) (*MovieAlternativeTitlesResponse, error) {
	r, err := c.sdk.MovieAlternativeTitles(c.ctx, movieID, country)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieAlternativeTitles200ApplicationJSONObject, nil
}

// MovieCredits wraps the generated openapi.SDK.MovieCredits call
func (c *Client) MovieCredits(movieID int, language *string) (*MovieCreditsResponse, error) {
	r, err := c.sdk.MovieCredits(c.ctx, movieID, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieCredits200ApplicationJSONObject, nil
}

// MovieDeleteRating wraps the generated openapi.SDK.MovieDeleteRating call
func (c *Client) MovieDeleteRating(movieID int, guestSessionID *string, sessionID *string) (*MovieDeleteRatingResponse, error) {
	r, err := c.sdk.MovieDeleteRating(c.ctx, movieID, guestSessionID, sessionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieDeleteRating200ApplicationJSONObject, nil
}

// MovieDetails wraps the generated openapi.SDK.MovieDetails call
func (c *Client) MovieDetails(movieID int, appendToResponse *string, language *string) (*MovieDetailsResponse, error) {
	r, err := c.sdk.MovieDetails(c.ctx, movieID, appendToResponse, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieDetails200ApplicationJSONObject, nil
}

// MovieExternalIds wraps the generated openapi.SDK.MovieExternalIds call
func (c *Client) MovieExternalIds(movieID int) (*MovieExternalIdsResponse, error) {
	r, err := c.sdk.MovieExternalIds(c.ctx, movieID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieExternalIds200ApplicationJSONObject, nil
}

// MovieImages wraps the generated openapi.SDK.MovieImages call
func (c *Client) MovieImages(movieID int, includeImageLanguage *string, language *string) (*MovieImagesResponse, error) {
	r, err := c.sdk.MovieImages(c.ctx, movieID, includeImageLanguage, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieImages200ApplicationJSONObject, nil
}

// MovieKeywords wraps the generated openapi.SDK.MovieKeywords call
func (c *Client) MovieKeywords(movieID string) (*MovieKeywordsResponse, error) {
	r, err := c.sdk.MovieKeywords(c.ctx, movieID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieKeywords200ApplicationJSONObject, nil
}

// MovieLatestID wraps the generated openapi.SDK.MovieLatestID call
func (c *Client) MovieLatestID() (*MovieLatestIDResponse, error) {
	r, err := c.sdk.MovieLatestID(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieLatestID200ApplicationJSONObject, nil
}

// MovieLists wraps the generated openapi.SDK.MovieLists call
func (c *Client) MovieLists(movieID int, language *string, page *int) (*MovieListsResponse, error) {
	r, err := c.sdk.MovieLists(c.ctx, movieID, language, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieLists200ApplicationJSONObject, nil
}

// MovieNowPlayingList wraps the generated openapi.SDK.MovieNowPlayingList call
func (c *Client) MovieNowPlayingList(language *string, page *int, region *string) (*MovieNowPlayingListResponse, error) {
	r, err := c.sdk.MovieNowPlayingList(c.ctx, language, page, region)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieNowPlayingList200ApplicationJSONObject, nil
}

// MoviePopularList wraps the generated openapi.SDK.MoviePopularList call
func (c *Client) MoviePopularList(language *string, page *int, region *string) (*MoviePopularListResponse, error) {
	r, err := c.sdk.MoviePopularList(c.ctx, language, page, region)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MoviePopularList200ApplicationJSONObject, nil
}

// MovieRecommendations wraps the generated openapi.SDK.MovieRecommendations call
func (c *Client) MovieRecommendations(movieID int, language *string, page *int) (*MovieRecommendationsResponse, error) {
	r, err := c.sdk.MovieRecommendations(c.ctx, movieID, language, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieRecommendations200ApplicationJSONObject, nil
}

// MovieReleaseDates wraps the generated openapi.SDK.MovieReleaseDates call
func (c *Client) MovieReleaseDates(movieID int) (*MovieReleaseDatesResponse, error) {
	r, err := c.sdk.MovieReleaseDates(c.ctx, movieID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieReleaseDates200ApplicationJSONObject, nil
}

// MovieReviews wraps the generated openapi.SDK.MovieReviews call
func (c *Client) MovieReviews(movieID int, language *string, page *int) (*MovieReviewsResponse, error) {
	r, err := c.sdk.MovieReviews(c.ctx, movieID, language, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieReviews200ApplicationJSONObject, nil
}

// MovieSimilar wraps the generated openapi.SDK.MovieSimilar call
func (c *Client) MovieSimilar(movieID int, language *string, page *int) (*MovieSimilarResponse, error) {
	r, err := c.sdk.MovieSimilar(c.ctx, movieID, language, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieSimilar200ApplicationJSONObject, nil
}

// MovieTopRatedList wraps the generated openapi.SDK.MovieTopRatedList call
func (c *Client) MovieTopRatedList(language *string, page *int, region *string) (*MovieTopRatedListResponse, error) {
	r, err := c.sdk.MovieTopRatedList(c.ctx, language, page, region)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieTopRatedList200ApplicationJSONObject, nil
}

// MovieTranslations wraps the generated openapi.SDK.MovieTranslations call
func (c *Client) MovieTranslations(movieID int) (*MovieTranslationsResponse, error) {
	r, err := c.sdk.MovieTranslations(c.ctx, movieID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieTranslations200ApplicationJSONObject, nil
}

// MovieUpcomingList wraps the generated openapi.SDK.MovieUpcomingList call
func (c *Client) MovieUpcomingList(language *string, page *int, region *string) (*MovieUpcomingListResponse, error) {
	r, err := c.sdk.MovieUpcomingList(c.ctx, language, page, region)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieUpcomingList200ApplicationJSONObject, nil
}

// MovieVideos wraps the generated openapi.SDK.MovieVideos call
func (c *Client) MovieVideos(movieID int, language *string) (*MovieVideosResponse, error) {
	r, err := c.sdk.MovieVideos(c.ctx, movieID, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieVideos200ApplicationJSONObject, nil
}

// MovieWatchProviders wraps the generated openapi.SDK.MovieWatchProviders call
func (c *Client) MovieWatchProviders(movieID int) (*MovieWatchProvidersResponse, error) {
	r, err := c.sdk.MovieWatchProviders(c.ctx, movieID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.MovieWatchProviders200ApplicationJSONObject, nil
}

// NetworkDetails wraps the generated openapi.SDK.NetworkDetails call
func (c *Client) NetworkDetails(networkID int) (*NetworkDetailsResponse, error) {
	r, err := c.sdk.NetworkDetails(c.ctx, networkID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.NetworkDetails200ApplicationJSONObject, nil
}

// PersonChanges wraps the generated openapi.SDK.PersonChanges call
func (c *Client) PersonChanges(personID int, endDate *types.Date, page *int, startDate *types.Date) (*PersonChangesResponse, error) {
	r, err := c.sdk.PersonChanges(c.ctx, personID, endDate, page, startDate)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.PersonChanges200ApplicationJSONObject, nil
}

// PersonCombinedCredits wraps the generated openapi.SDK.PersonCombinedCredits call
func (c *Client) PersonCombinedCredits(personID string, language *string) (*PersonCombinedCreditsResponse, error) {
	r, err := c.sdk.PersonCombinedCredits(c.ctx, personID, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.PersonCombinedCredits200ApplicationJSONObject, nil
}

// PersonDetails wraps the generated openapi.SDK.PersonDetails call
func (c *Client) PersonDetails(personID int, appendToResponse *string, language *string) (*PersonDetailsResponse, error) {
	r, err := c.sdk.PersonDetails(c.ctx, personID, appendToResponse, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.PersonDetails200ApplicationJSONObject, nil
}

// PersonExternalIds wraps the generated openapi.SDK.PersonExternalIds call
func (c *Client) PersonExternalIds(personID int) (*PersonExternalIdsResponse, error) {
	r, err := c.sdk.PersonExternalIds(c.ctx, personID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.PersonExternalIds200ApplicationJSONObject, nil
}

// PersonImages wraps the generated openapi.SDK.PersonImages call
func (c *Client) PersonImages(personID int) (*PersonImagesResponse, error) {
	r, err := c.sdk.PersonImages(c.ctx, personID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.PersonImages200ApplicationJSONObject, nil
}

// PersonLatestID wraps the generated openapi.SDK.PersonLatestID call
func (c *Client) PersonLatestID() (*PersonLatestIDResponse, error) {
	r, err := c.sdk.PersonLatestID(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.PersonLatestID200ApplicationJSONObject, nil
}

// PersonMovieCredits wraps the generated openapi.SDK.PersonMovieCredits call
func (c *Client) PersonMovieCredits(personID int, language *string) (*PersonMovieCreditsResponse, error) {
	r, err := c.sdk.PersonMovieCredits(c.ctx, personID, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.PersonMovieCredits200ApplicationJSONObject, nil
}

// PersonPopularList wraps the generated openapi.SDK.PersonPopularList call
func (c *Client) PersonPopularList(language *string, page *int) (*PersonPopularListResponse, error) {
	r, err := c.sdk.PersonPopularList(c.ctx, language, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.PersonPopularList200ApplicationJSONObject, nil
}

// PersonTaggedImages wraps the generated openapi.SDK.PersonTaggedImages call
func (c *Client) PersonTaggedImages(personID int, page *int) (*PersonTaggedImagesResponse, error) {
	r, err := c.sdk.PersonTaggedImages(c.ctx, personID, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.PersonTaggedImages200ApplicationJSONObject, nil
}

// PersonTvCredits wraps the generated openapi.SDK.PersonTvCredits call
func (c *Client) PersonTvCredits(personID int, language *string) (*PersonTvCreditsResponse, error) {
	r, err := c.sdk.PersonTvCredits(c.ctx, personID, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.PersonTvCredits200ApplicationJSONObject, nil
}

// ReviewDetails wraps the generated openapi.SDK.ReviewDetails call
func (c *Client) ReviewDetails(reviewID string) (*ReviewDetailsResponse, error) {
	r, err := c.sdk.ReviewDetails(c.ctx, reviewID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.ReviewDetails200ApplicationJSONObject, nil
}

// SearchCollection wraps the generated openapi.SDK.SearchCollection call
func (c *Client) SearchCollection(request operations.SearchCollectionRequest) (*SearchCollectionResponse, error) {
	r, err := c.sdk.SearchCollection(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.SearchCollection200ApplicationJSONObject, nil
}

// SearchCompany wraps the generated openapi.SDK.SearchCompany call
func (c *Client) SearchCompany(query string, page *int) (*SearchCompanyResponse, error) {
	r, err := c.sdk.SearchCompany(c.ctx, query, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.SearchCompany200ApplicationJSONObject, nil
}

// SearchKeyword wraps the generated openapi.SDK.SearchKeyword call
func (c *Client) SearchKeyword(query string, page *int) (*SearchKeywordResponse, error) {
	r, err := c.sdk.SearchKeyword(c.ctx, query, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.SearchKeyword200ApplicationJSONObject, nil
}

// SearchMovie wraps the generated openapi.SDK.SearchMovie call
func (c *Client) SearchMovie(request operations.SearchMovieRequest) (*SearchMovieResponse, error) {
	r, err := c.sdk.SearchMovie(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.SearchMovie200ApplicationJSONObject, nil
}

// SearchMulti wraps the generated openapi.SDK.SearchMulti call
func (c *Client) SearchMulti(query string, includeAdult *bool, language *string, page *int) (*SearchMultiResponse, error) {
	r, err := c.sdk.SearchMulti(c.ctx, query, includeAdult, language, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.SearchMulti200ApplicationJSONObject, nil
}

// SearchPerson wraps the generated openapi.SDK.SearchPerson call
func (c *Client) SearchPerson(query string, includeAdult *bool, language *string, page *int) (*SearchPersonResponse, error) {
	r, err := c.sdk.SearchPerson(c.ctx, query, includeAdult, language, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.SearchPerson200ApplicationJSONObject, nil
}

// SearchTv wraps the generated openapi.SDK.SearchTv call
func (c *Client) SearchTv(request operations.SearchTvRequest) (*SearchTvResponse, error) {
	r, err := c.sdk.SearchTv(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.SearchTv200ApplicationJSONObject, nil
}

// Translations wraps the generated openapi.SDK.Translations call
func (c *Client) Translations(personID int) (*TranslationsResponse, error) {
	r, err := c.sdk.Translations(c.ctx, personID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.Translations200ApplicationJSONObject, nil
}

// TrendingAll wraps the generated openapi.SDK.TrendingAll call
func (c *Client) TrendingAll(timeWindow operations.TrendingAllTimeWindow, language *string) (*TrendingAllResponse, error) {
	r, err := c.sdk.TrendingAll(c.ctx, timeWindow, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TrendingAll200ApplicationJSONObject, nil
}

// TrendingMovies wraps the generated openapi.SDK.TrendingMovies call
func (c *Client) TrendingMovies(timeWindow operations.TrendingMoviesTimeWindow, language *string) (*TrendingMoviesResponse, error) {
	r, err := c.sdk.TrendingMovies(c.ctx, timeWindow, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TrendingMovies200ApplicationJSONObject, nil
}

// TrendingPeople wraps the generated openapi.SDK.TrendingPeople call
func (c *Client) TrendingPeople(timeWindow operations.TrendingPeopleTimeWindow, language *string) (*TrendingPeopleResponse, error) {
	r, err := c.sdk.TrendingPeople(c.ctx, timeWindow, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TrendingPeople200ApplicationJSONObject, nil
}

// TrendingTv wraps the generated openapi.SDK.TrendingTv call
func (c *Client) TrendingTv(timeWindow operations.TrendingTvTimeWindow, language *string) (*TrendingTvResponse, error) {
	r, err := c.sdk.TrendingTv(c.ctx, timeWindow, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TrendingTv200ApplicationJSONObject, nil
}

// TvEpisodeAccountStates wraps the generated openapi.SDK.TvEpisodeAccountStates call
func (c *Client) TvEpisodeAccountStates(request operations.TvEpisodeAccountStatesRequest) (*TvEpisodeAccountStatesResponse, error) {
	r, err := c.sdk.TvEpisodeAccountStates(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvEpisodeAccountStates200ApplicationJSONObject, nil
}

// TvEpisodeAddRating wraps the generated openapi.SDK.TvEpisodeAddRating call
func (c *Client) TvEpisodeAddRating(request operations.TvEpisodeAddRatingRequest) (*TvEpisodeAddRatingResponse, error) {
	r, err := c.sdk.TvEpisodeAddRating(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvEpisodeAddRating200ApplicationJSONObject, nil
}

// TvEpisodeChangesByID wraps the generated openapi.SDK.TvEpisodeChangesByID call
func (c *Client) TvEpisodeChangesByID(episodeID int) (*TvEpisodeChangesByIDResponse, error) {
	r, err := c.sdk.TvEpisodeChangesByID(c.ctx, episodeID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvEpisodeChangesByID200ApplicationJSONObject, nil
}

// TvEpisodeCredits wraps the generated openapi.SDK.TvEpisodeCredits call
func (c *Client) TvEpisodeCredits(episodeNumber int, seasonNumber int, seriesID int, language *string) (*TvEpisodeCreditsResponse, error) {
	r, err := c.sdk.TvEpisodeCredits(c.ctx, episodeNumber, seasonNumber, seriesID, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvEpisodeCredits200ApplicationJSONObject, nil
}

// TvEpisodeDeleteRating wraps the generated openapi.SDK.TvEpisodeDeleteRating call
func (c *Client) TvEpisodeDeleteRating(request operations.TvEpisodeDeleteRatingRequest) (*TvEpisodeDeleteRatingResponse, error) {
	r, err := c.sdk.TvEpisodeDeleteRating(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvEpisodeDeleteRating200ApplicationJSONObject, nil
}

// TvEpisodeDetails wraps the generated openapi.SDK.TvEpisodeDetails call
func (c *Client) TvEpisodeDetails(request operations.TvEpisodeDetailsRequest) (*TvEpisodeDetailsResponse, error) {
	r, err := c.sdk.TvEpisodeDetails(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvEpisodeDetails200ApplicationJSONObject, nil
}

// TvEpisodeExternalIds wraps the generated openapi.SDK.TvEpisodeExternalIds call
func (c *Client) TvEpisodeExternalIds(episodeNumber string, seasonNumber int, seriesID int) (*TvEpisodeExternalIdsResponse, error) {
	r, err := c.sdk.TvEpisodeExternalIds(c.ctx, episodeNumber, seasonNumber, seriesID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvEpisodeExternalIds200ApplicationJSONObject, nil
}

// TvEpisodeGroupDetails wraps the generated openapi.SDK.TvEpisodeGroupDetails call
func (c *Client) TvEpisodeGroupDetails(tvEpisodeGroupID string) (*TvEpisodeGroupDetailsResponse, error) {
	r, err := c.sdk.TvEpisodeGroupDetails(c.ctx, tvEpisodeGroupID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvEpisodeGroupDetails200ApplicationJSONObject, nil
}

// TvEpisodeImages wraps the generated openapi.SDK.TvEpisodeImages call
func (c *Client) TvEpisodeImages(request operations.TvEpisodeImagesRequest) (*TvEpisodeImagesResponse, error) {
	r, err := c.sdk.TvEpisodeImages(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvEpisodeImages200ApplicationJSONObject, nil
}

// TvEpisodeTranslations wraps the generated openapi.SDK.TvEpisodeTranslations call
func (c *Client) TvEpisodeTranslations(episodeNumber int, seasonNumber int, seriesID int) (*TvEpisodeTranslationsResponse, error) {
	r, err := c.sdk.TvEpisodeTranslations(c.ctx, episodeNumber, seasonNumber, seriesID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvEpisodeTranslations200ApplicationJSONObject, nil
}

// TvEpisodeVideos wraps the generated openapi.SDK.TvEpisodeVideos call
func (c *Client) TvEpisodeVideos(request operations.TvEpisodeVideosRequest) (*TvEpisodeVideosResponse, error) {
	r, err := c.sdk.TvEpisodeVideos(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvEpisodeVideos200ApplicationJSONObject, nil
}

// TvSeasonAccountStates wraps the generated openapi.SDK.TvSeasonAccountStates call
func (c *Client) TvSeasonAccountStates(seasonNumber int, seriesID int, guestSessionID *string, sessionID *string) (*TvSeasonAccountStatesResponse, error) {
	r, err := c.sdk.TvSeasonAccountStates(c.ctx, seasonNumber, seriesID, guestSessionID, sessionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeasonAccountStates200ApplicationJSONObject, nil
}

// TvSeasonAggregateCredits wraps the generated openapi.SDK.TvSeasonAggregateCredits call
func (c *Client) TvSeasonAggregateCredits(seasonNumber int, seriesID int, language *string) (*TvSeasonAggregateCreditsResponse, error) {
	r, err := c.sdk.TvSeasonAggregateCredits(c.ctx, seasonNumber, seriesID, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeasonAggregateCredits200ApplicationJSONObject, nil
}

// TvSeasonChangesByID wraps the generated openapi.SDK.TvSeasonChangesByID call
func (c *Client) TvSeasonChangesByID(seasonID int, endDate *string, page *int, startDate *string) (*TvSeasonChangesByIDResponse, error) {
	r, err := c.sdk.TvSeasonChangesByID(c.ctx, seasonID, endDate, page, startDate)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeasonChangesByID200ApplicationJSONObject, nil
}

// TvSeasonCredits wraps the generated openapi.SDK.TvSeasonCredits call
func (c *Client) TvSeasonCredits(seasonNumber int, seriesID int, language *string) (*TvSeasonCreditsResponse, error) {
	r, err := c.sdk.TvSeasonCredits(c.ctx, seasonNumber, seriesID, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeasonCredits200ApplicationJSONObject, nil
}

// TvSeasonDetails wraps the generated openapi.SDK.TvSeasonDetails call
func (c *Client) TvSeasonDetails(seasonNumber int, seriesID int, appendToResponse *string, language *string) (*TvSeasonDetailsResponse, error) {
	r, err := c.sdk.TvSeasonDetails(c.ctx, seasonNumber, seriesID, appendToResponse, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeasonDetails200ApplicationJSONObject, nil
}

// TvSeasonExternalIds wraps the generated openapi.SDK.TvSeasonExternalIds call
func (c *Client) TvSeasonExternalIds(seasonNumber int, seriesID int) (*TvSeasonExternalIdsResponse, error) {
	r, err := c.sdk.TvSeasonExternalIds(c.ctx, seasonNumber, seriesID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeasonExternalIds200ApplicationJSONObject, nil
}

// TvSeasonImages wraps the generated openapi.SDK.TvSeasonImages call
func (c *Client) TvSeasonImages(seasonNumber int, seriesID int, includeImageLanguage *string, language *string) (*TvSeasonImagesResponse, error) {
	r, err := c.sdk.TvSeasonImages(c.ctx, seasonNumber, seriesID, includeImageLanguage, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeasonImages200ApplicationJSONObject, nil
}

// TvSeasonTranslations wraps the generated openapi.SDK.TvSeasonTranslations call
func (c *Client) TvSeasonTranslations(seasonNumber int, seriesID int) (*TvSeasonTranslationsResponse, error) {
	r, err := c.sdk.TvSeasonTranslations(c.ctx, seasonNumber, seriesID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeasonTranslations200ApplicationJSONObject, nil
}

// TvSeasonVideos wraps the generated openapi.SDK.TvSeasonVideos call
func (c *Client) TvSeasonVideos(seasonNumber int, seriesID int, includeVideoLanguage *string, language *string) (*TvSeasonVideosResponse, error) {
	r, err := c.sdk.TvSeasonVideos(c.ctx, seasonNumber, seriesID, includeVideoLanguage, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeasonVideos200ApplicationJSONObject, nil
}

// TvSeasonWatchProviders wraps the generated openapi.SDK.TvSeasonWatchProviders call
func (c *Client) TvSeasonWatchProviders(seasonNumber int, seriesID int, language *string) (*TvSeasonWatchProvidersResponse, error) {
	r, err := c.sdk.TvSeasonWatchProviders(c.ctx, seasonNumber, seriesID, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeasonWatchProviders200ApplicationJSONObject, nil
}

// TvSeriesAccountStates wraps the generated openapi.SDK.TvSeriesAccountStates call
func (c *Client) TvSeriesAccountStates(seriesID int, guestSessionID *string, sessionID *string) (*TvSeriesAccountStatesResponse, error) {
	r, err := c.sdk.TvSeriesAccountStates(c.ctx, seriesID, guestSessionID, sessionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesAccountStates200ApplicationJSONObject, nil
}

// TvSeriesAddRating wraps the generated openapi.SDK.TvSeriesAddRating call
func (c *Client) TvSeriesAddRating(seriesID int, requestBody *operations.TvSeriesAddRatingRequestBody, guestSessionID *string, sessionID *string) (*TvSeriesAddRatingResponse, error) {
	r, err := c.sdk.TvSeriesAddRating(c.ctx, seriesID, requestBody, guestSessionID, sessionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesAddRating200ApplicationJSONObject, nil
}

// TvSeriesAggregateCredits wraps the generated openapi.SDK.TvSeriesAggregateCredits call
func (c *Client) TvSeriesAggregateCredits(seriesID int, language *string) (*TvSeriesAggregateCreditsResponse, error) {
	r, err := c.sdk.TvSeriesAggregateCredits(c.ctx, seriesID, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesAggregateCredits200ApplicationJSONObject, nil
}

// TvSeriesAiringTodayList wraps the generated openapi.SDK.TvSeriesAiringTodayList call
func (c *Client) TvSeriesAiringTodayList(language *string, page *int, timezone *string) (*TvSeriesAiringTodayListResponse, error) {
	r, err := c.sdk.TvSeriesAiringTodayList(c.ctx, language, page, timezone)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesAiringTodayList200ApplicationJSONObject, nil
}

// TvSeriesAlternativeTitles wraps the generated openapi.SDK.TvSeriesAlternativeTitles call
func (c *Client) TvSeriesAlternativeTitles(seriesID int) (*TvSeriesAlternativeTitlesResponse, error) {
	r, err := c.sdk.TvSeriesAlternativeTitles(c.ctx, seriesID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesAlternativeTitles200ApplicationJSONObject, nil
}

// TvSeriesChanges wraps the generated openapi.SDK.TvSeriesChanges call
func (c *Client) TvSeriesChanges(seriesID int, endDate *string, page *int, startDate *string) (*TvSeriesChangesResponse, error) {
	r, err := c.sdk.TvSeriesChanges(c.ctx, seriesID, endDate, page, startDate)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesChanges200ApplicationJSONObject, nil
}

// TvSeriesContentRatings wraps the generated openapi.SDK.TvSeriesContentRatings call
func (c *Client) TvSeriesContentRatings(seriesID int) (*TvSeriesContentRatingsResponse, error) {
	r, err := c.sdk.TvSeriesContentRatings(c.ctx, seriesID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesContentRatings200ApplicationJSONObject, nil
}

// TvSeriesCredits wraps the generated openapi.SDK.TvSeriesCredits call
func (c *Client) TvSeriesCredits(seriesID int, language *string) (*TvSeriesCreditsResponse, error) {
	r, err := c.sdk.TvSeriesCredits(c.ctx, seriesID, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesCredits200ApplicationJSONObject, nil
}

// TvSeriesDeleteRating wraps the generated openapi.SDK.TvSeriesDeleteRating call
func (c *Client) TvSeriesDeleteRating(seriesID int, guestSessionID *string, sessionID *string) (*TvSeriesDeleteRatingResponse, error) {
	r, err := c.sdk.TvSeriesDeleteRating(c.ctx, seriesID, guestSessionID, sessionID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesDeleteRating200ApplicationJSONObject, nil
}

// TvSeriesDetails wraps the generated openapi.SDK.TvSeriesDetails call
func (c *Client) TvSeriesDetails(seriesID int, appendToResponse *string, language *string) (*TvSeriesDetailsResponse, error) {
	r, err := c.sdk.TvSeriesDetails(c.ctx, seriesID, appendToResponse, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesDetails200ApplicationJSONObject, nil
}

// TvSeriesEpisodeGroups wraps the generated openapi.SDK.TvSeriesEpisodeGroups call
func (c *Client) TvSeriesEpisodeGroups(seriesID int) (*TvSeriesEpisodeGroupsResponse, error) {
	r, err := c.sdk.TvSeriesEpisodeGroups(c.ctx, seriesID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesEpisodeGroups200ApplicationJSONObject, nil
}

// TvSeriesExternalIds wraps the generated openapi.SDK.TvSeriesExternalIds call
func (c *Client) TvSeriesExternalIds(seriesID int) (*TvSeriesExternalIdsResponse, error) {
	r, err := c.sdk.TvSeriesExternalIds(c.ctx, seriesID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesExternalIds200ApplicationJSONObject, nil
}

// TvSeriesImages wraps the generated openapi.SDK.TvSeriesImages call
func (c *Client) TvSeriesImages(seriesID int, includeImageLanguage *string, language *string) (*TvSeriesImagesResponse, error) {
	r, err := c.sdk.TvSeriesImages(c.ctx, seriesID, includeImageLanguage, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesImages200ApplicationJSONObject, nil
}

// TvSeriesKeywords wraps the generated openapi.SDK.TvSeriesKeywords call
func (c *Client) TvSeriesKeywords(seriesID int) (*TvSeriesKeywordsResponse, error) {
	r, err := c.sdk.TvSeriesKeywords(c.ctx, seriesID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesKeywords200ApplicationJSONObject, nil
}

// TvSeriesLatestID wraps the generated openapi.SDK.TvSeriesLatestID call
func (c *Client) TvSeriesLatestID() (*TvSeriesLatestIDResponse, error) {
	r, err := c.sdk.TvSeriesLatestID(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesLatestID200ApplicationJSONObject, nil
}

// TvSeriesOnTheAirList wraps the generated openapi.SDK.TvSeriesOnTheAirList call
func (c *Client) TvSeriesOnTheAirList(language *string, page *int, timezone *string) (*TvSeriesOnTheAirListResponse, error) {
	r, err := c.sdk.TvSeriesOnTheAirList(c.ctx, language, page, timezone)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesOnTheAirList200ApplicationJSONObject, nil
}

// TvSeriesPopularList wraps the generated openapi.SDK.TvSeriesPopularList call
func (c *Client) TvSeriesPopularList(language *string, page *int) (*TvSeriesPopularListResponse, error) {
	r, err := c.sdk.TvSeriesPopularList(c.ctx, language, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesPopularList200ApplicationJSONObject, nil
}

// TvSeriesRecommendations wraps the generated openapi.SDK.TvSeriesRecommendations call
func (c *Client) TvSeriesRecommendations(seriesID int, language *string, page *int) (*TvSeriesRecommendationsResponse, error) {
	r, err := c.sdk.TvSeriesRecommendations(c.ctx, seriesID, language, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesRecommendations200ApplicationJSONObject, nil
}

// TvSeriesReviews wraps the generated openapi.SDK.TvSeriesReviews call
func (c *Client) TvSeriesReviews(seriesID int, language *string, page *int) (*TvSeriesReviewsResponse, error) {
	r, err := c.sdk.TvSeriesReviews(c.ctx, seriesID, language, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesReviews200ApplicationJSONObject, nil
}

// TvSeriesScreenedTheatrically wraps the generated openapi.SDK.TvSeriesScreenedTheatrically call
func (c *Client) TvSeriesScreenedTheatrically(seriesID int) (*TvSeriesScreenedTheatricallyResponse, error) {
	r, err := c.sdk.TvSeriesScreenedTheatrically(c.ctx, seriesID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesScreenedTheatrically200ApplicationJSONObject, nil
}

// TvSeriesSimilar wraps the generated openapi.SDK.TvSeriesSimilar call
func (c *Client) TvSeriesSimilar(seriesID string, language *string, page *int) (*TvSeriesSimilarResponse, error) {
	r, err := c.sdk.TvSeriesSimilar(c.ctx, seriesID, language, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesSimilar200ApplicationJSONObject, nil
}

// TvSeriesTopRatedList wraps the generated openapi.SDK.TvSeriesTopRatedList call
func (c *Client) TvSeriesTopRatedList(language *string, page *int) (*TvSeriesTopRatedListResponse, error) {
	r, err := c.sdk.TvSeriesTopRatedList(c.ctx, language, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesTopRatedList200ApplicationJSONObject, nil
}

// TvSeriesTranslations wraps the generated openapi.SDK.TvSeriesTranslations call
func (c *Client) TvSeriesTranslations(seriesID int) (*TvSeriesTranslationsResponse, error) {
	r, err := c.sdk.TvSeriesTranslations(c.ctx, seriesID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesTranslations200ApplicationJSONObject, nil
}

// TvSeriesVideos wraps the generated openapi.SDK.TvSeriesVideos call
func (c *Client) TvSeriesVideos(seriesID int, includeVideoLanguage *string, language *string) (*TvSeriesVideosResponse, error) {
	r, err := c.sdk.TvSeriesVideos(c.ctx, seriesID, includeVideoLanguage, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesVideos200ApplicationJSONObject, nil
}

// TvSeriesWatchProviders wraps the generated openapi.SDK.TvSeriesWatchProviders call
func (c *Client) TvSeriesWatchProviders(seriesID int) (*TvSeriesWatchProvidersResponse, error) {
	r, err := c.sdk.TvSeriesWatchProviders(c.ctx, seriesID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.TvSeriesWatchProviders200ApplicationJSONObject, nil
}

// WatchProviderTvList wraps the generated openapi.SDK.WatchProviderTvList call
func (c *Client) WatchProviderTvList(language *string, watchRegion *string) (*WatchProviderTvListResponse, error) {
	r, err := c.sdk.WatchProviderTvList(c.ctx, language, watchRegion)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.WatchProviderTvList200ApplicationJSONObject, nil
}

// WatchProvidersAvailableRegions wraps the generated openapi.SDK.WatchProvidersAvailableRegions call
func (c *Client) WatchProvidersAvailableRegions(language *string) (*WatchProvidersAvailableRegionsResponse, error) {
	r, err := c.sdk.WatchProvidersAvailableRegions(c.ctx, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.WatchProvidersAvailableRegions200ApplicationJSONObject, nil
}

// WatchProvidersMovieList wraps the generated openapi.SDK.WatchProvidersMovieList call
func (c *Client) WatchProvidersMovieList(language *string, watchRegion *string) (*WatchProvidersMovieListResponse, error) {
	r, err := c.sdk.WatchProvidersMovieList(c.ctx, language, watchRegion)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.WatchProvidersMovieList200ApplicationJSONObject, nil
}
