# github.com/dashotv/tmdb/openapi

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README
<!-- Start SDK Installation -->
# SDK Installation

```bash
go get github.com/dashotv/tmdb/openapi
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	"github.com/dashotv/tmdb/openapi"
	"github.com/dashotv/tmdb/openapi/models/operations"
	"github.com/dashotv/tmdb/openapi/models/shared"
	"log"
)

func main() {
	s := openapi.New(
		openapi.WithSecurity(""),
	)

	var accountID int = 959345

	requestBody := &operations.AccountAddFavoriteRequestBody{
		RawBody: "{key: 16540, key1: null, key2: \"string\"}",
	}

	var sessionID *string = "string"

	ctx := context.Background()
	res, err := s.SDK.AccountAddFavorite(ctx, accountID, requestBody, sessionID)
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountAddFavorite200ApplicationJSONObject != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
# Available Resources and Operations

## [SDK](docs/sdks/sdk/README.md)

* [AccountAddFavorite](docs/sdks/sdk/README.md#accountaddfavorite) - Add Favorite
* [AccountAddToWatchlist](docs/sdks/sdk/README.md#accountaddtowatchlist) - Add To Watchlist
* [AccountDetails](docs/sdks/sdk/README.md#accountdetails) - Details
* [AccountFavoriteTv](docs/sdks/sdk/README.md#accountfavoritetv) - Favorite TV
* [AccountGetFavorites](docs/sdks/sdk/README.md#accountgetfavorites) - Favorite Movies
* [AccountLists](docs/sdks/sdk/README.md#accountlists) - Lists
* [AccountRatedMovies](docs/sdks/sdk/README.md#accountratedmovies) - Rated Movies
* [AccountRatedTv](docs/sdks/sdk/README.md#accountratedtv) - Rated TV
* [AccountRatedTvEpisodes](docs/sdks/sdk/README.md#accountratedtvepisodes) - Rated TV Episodes
* [AccountWatchlistMovies](docs/sdks/sdk/README.md#accountwatchlistmovies) - Watchlist Movies
* [AccountWatchlistTv](docs/sdks/sdk/README.md#accountwatchlisttv) - Watchlist TV
* [AlternativeNamesCopy](docs/sdks/sdk/README.md#alternativenamescopy) - Images
* [AuthenticationCreateGuestSession](docs/sdks/sdk/README.md#authenticationcreateguestsession) - Create Guest Session
* [AuthenticationCreateRequestToken](docs/sdks/sdk/README.md#authenticationcreaterequesttoken) - Create Request Token
* [AuthenticationCreateSession](docs/sdks/sdk/README.md#authenticationcreatesession) - Create Session
* [AuthenticationCreateSessionFromLogin](docs/sdks/sdk/README.md#authenticationcreatesessionfromlogin) - Create Session (with login)
* [AuthenticationCreateSessionFromV4Token](docs/sdks/sdk/README.md#authenticationcreatesessionfromv4token) - Create Session (from v4 token)
* [AuthenticationDeleteSession](docs/sdks/sdk/README.md#authenticationdeletesession) - Delete Session
* [AuthenticationValidateKey](docs/sdks/sdk/README.md#authenticationvalidatekey) - Validate Key
* [CertificationMovieList](docs/sdks/sdk/README.md#certificationmovielist) - Movie Certifications
* [CertificationsTvList](docs/sdks/sdk/README.md#certificationstvlist) - TV Certifications
* [ChangesMovieList](docs/sdks/sdk/README.md#changesmovielist) - Movie List
* [ChangesPeopleList](docs/sdks/sdk/README.md#changespeoplelist) - People List
* [ChangesTvList](docs/sdks/sdk/README.md#changestvlist) - TV List
* [CollectionDetails](docs/sdks/sdk/README.md#collectiondetails) - Details
* [CollectionImages](docs/sdks/sdk/README.md#collectionimages) - Images
* [CollectionTranslations](docs/sdks/sdk/README.md#collectiontranslations) - Translations
* [CompanyAlternativeNames](docs/sdks/sdk/README.md#companyalternativenames) - Alternative Names
* [CompanyDetails](docs/sdks/sdk/README.md#companydetails) - Details
* [CompanyImages](docs/sdks/sdk/README.md#companyimages) - Images
* [ConfigurationCountries](docs/sdks/sdk/README.md#configurationcountries) - Countries
* [ConfigurationDetails](docs/sdks/sdk/README.md#configurationdetails) - Details
* [ConfigurationJobs](docs/sdks/sdk/README.md#configurationjobs) - Jobs
* [ConfigurationLanguages](docs/sdks/sdk/README.md#configurationlanguages) - Languages
* [ConfigurationPrimaryTranslations](docs/sdks/sdk/README.md#configurationprimarytranslations) - Primary Translations
* [ConfigurationTimezones](docs/sdks/sdk/README.md#configurationtimezones) - Timezones
* [CreditDetails](docs/sdks/sdk/README.md#creditdetails) - Details
* [DetailsCopy](docs/sdks/sdk/README.md#detailscopy) - Alternative Names
* [DiscoverMovie](docs/sdks/sdk/README.md#discovermovie) - Movie
* [DiscoverTv](docs/sdks/sdk/README.md#discovertv) - TV
* [FindByID](docs/sdks/sdk/README.md#findbyid) - Find By ID
* [GenreMovieList](docs/sdks/sdk/README.md#genremovielist) - Movie List
* [GenreTvList](docs/sdks/sdk/README.md#genretvlist) - TV List
* [GuestSessionRatedMovies](docs/sdks/sdk/README.md#guestsessionratedmovies) - Rated Movies
* [GuestSessionRatedTv](docs/sdks/sdk/README.md#guestsessionratedtv) - Rated TV
* [GuestSessionRatedTvEpisodes](docs/sdks/sdk/README.md#guestsessionratedtvepisodes) - Rated TV Episodes
* [KeywordDetails](docs/sdks/sdk/README.md#keyworddetails) - Details
* [KeywordMovies](docs/sdks/sdk/README.md#keywordmovies) - Movies
* [ListAddMovie](docs/sdks/sdk/README.md#listaddmovie) - Add Movie
* [ListCheckItemStatus](docs/sdks/sdk/README.md#listcheckitemstatus) - Check Item Status
* [ListClear](docs/sdks/sdk/README.md#listclear) - Clear
* [ListCreate](docs/sdks/sdk/README.md#listcreate) - Create
* [ListDelete](docs/sdks/sdk/README.md#listdelete) - Delete
* [ListDetails](docs/sdks/sdk/README.md#listdetails) - Details
* [ListRemoveMovie](docs/sdks/sdk/README.md#listremovemovie) - Remove Movie
* [MovieAccountStates](docs/sdks/sdk/README.md#movieaccountstates) - Account States
* [MovieAddRating](docs/sdks/sdk/README.md#movieaddrating) - Add Rating
* [MovieAlternativeTitles](docs/sdks/sdk/README.md#moviealternativetitles) - Alternative Titles
* [MovieChanges](docs/sdks/sdk/README.md#moviechanges) - Changes
* [MovieCredits](docs/sdks/sdk/README.md#moviecredits) - Credits
* [MovieDeleteRating](docs/sdks/sdk/README.md#moviedeleterating) - Delete Rating
* [MovieDetails](docs/sdks/sdk/README.md#moviedetails) - Details
* [MovieExternalIds](docs/sdks/sdk/README.md#movieexternalids) - External IDs
* [MovieImages](docs/sdks/sdk/README.md#movieimages) - Images
* [MovieKeywords](docs/sdks/sdk/README.md#moviekeywords) - Keywords
* [MovieLatestID](docs/sdks/sdk/README.md#movielatestid) - Latest
* [MovieLists](docs/sdks/sdk/README.md#movielists) - Lists
* [MovieNowPlayingList](docs/sdks/sdk/README.md#movienowplayinglist) - Now Playing
* [MoviePopularList](docs/sdks/sdk/README.md#moviepopularlist) - Popular
* [MovieRecommendations](docs/sdks/sdk/README.md#movierecommendations) - Recommendations
* [MovieReleaseDates](docs/sdks/sdk/README.md#moviereleasedates) - Release Dates
* [MovieReviews](docs/sdks/sdk/README.md#moviereviews) - Reviews
* [MovieSimilar](docs/sdks/sdk/README.md#moviesimilar) - Similar
* [MovieTopRatedList](docs/sdks/sdk/README.md#movietopratedlist) - Top Rated
* [MovieTranslations](docs/sdks/sdk/README.md#movietranslations) - Translations
* [MovieUpcomingList](docs/sdks/sdk/README.md#movieupcominglist) - Upcoming
* [MovieVideos](docs/sdks/sdk/README.md#movievideos) - Videos
* [MovieWatchProviders](docs/sdks/sdk/README.md#moviewatchproviders) - Watch Providers
* [NetworkDetails](docs/sdks/sdk/README.md#networkdetails) - Details
* [PersonChanges](docs/sdks/sdk/README.md#personchanges) - Changes
* [PersonCombinedCredits](docs/sdks/sdk/README.md#personcombinedcredits) - Combined Credits
* [PersonDetails](docs/sdks/sdk/README.md#persondetails) - Details
* [PersonExternalIds](docs/sdks/sdk/README.md#personexternalids) - External IDs
* [PersonImages](docs/sdks/sdk/README.md#personimages) - Images
* [PersonLatestID](docs/sdks/sdk/README.md#personlatestid) - Latest
* [PersonMovieCredits](docs/sdks/sdk/README.md#personmoviecredits) - Movie Credits
* [PersonPopularList](docs/sdks/sdk/README.md#personpopularlist) - Popular
* [PersonTaggedImages](docs/sdks/sdk/README.md#persontaggedimages) - Tagged Images
* [PersonTvCredits](docs/sdks/sdk/README.md#persontvcredits) - TV Credits
* [ReviewDetails](docs/sdks/sdk/README.md#reviewdetails) - Details
* [SearchCollection](docs/sdks/sdk/README.md#searchcollection) - Collection
* [SearchCompany](docs/sdks/sdk/README.md#searchcompany) - Company
* [SearchKeyword](docs/sdks/sdk/README.md#searchkeyword) - Keyword
* [SearchMovie](docs/sdks/sdk/README.md#searchmovie) - Movie
* [SearchMulti](docs/sdks/sdk/README.md#searchmulti) - Multi
* [SearchPerson](docs/sdks/sdk/README.md#searchperson) - Person
* [SearchTv](docs/sdks/sdk/README.md#searchtv) - TV
* [Translations](docs/sdks/sdk/README.md#translations) - Translations
* [TrendingAll](docs/sdks/sdk/README.md#trendingall) - All
* [TrendingMovies](docs/sdks/sdk/README.md#trendingmovies) - Movies
* [TrendingPeople](docs/sdks/sdk/README.md#trendingpeople) - People
* [TrendingTv](docs/sdks/sdk/README.md#trendingtv) - TV
* [TvEpisodeAccountStates](docs/sdks/sdk/README.md#tvepisodeaccountstates) - Account States
* [TvEpisodeAddRating](docs/sdks/sdk/README.md#tvepisodeaddrating) - Add Rating
* [TvEpisodeChangesByID](docs/sdks/sdk/README.md#tvepisodechangesbyid) - Changes
* [TvEpisodeCredits](docs/sdks/sdk/README.md#tvepisodecredits) - Credits
* [TvEpisodeDeleteRating](docs/sdks/sdk/README.md#tvepisodedeleterating) - Delete Rating
* [TvEpisodeDetails](docs/sdks/sdk/README.md#tvepisodedetails) - Details
* [TvEpisodeExternalIds](docs/sdks/sdk/README.md#tvepisodeexternalids) - External IDs
* [TvEpisodeGroupDetails](docs/sdks/sdk/README.md#tvepisodegroupdetails) - Details
* [TvEpisodeImages](docs/sdks/sdk/README.md#tvepisodeimages) - Images
* [TvEpisodeTranslations](docs/sdks/sdk/README.md#tvepisodetranslations) - Translations
* [TvEpisodeVideos](docs/sdks/sdk/README.md#tvepisodevideos) - Videos
* [TvSeasonAccountStates](docs/sdks/sdk/README.md#tvseasonaccountstates) - Account States
* [TvSeasonAggregateCredits](docs/sdks/sdk/README.md#tvseasonaggregatecredits) - Aggregate Credits
* [TvSeasonChangesByID](docs/sdks/sdk/README.md#tvseasonchangesbyid) - Changes
* [TvSeasonCredits](docs/sdks/sdk/README.md#tvseasoncredits) - Credits
* [TvSeasonDetails](docs/sdks/sdk/README.md#tvseasondetails) - Details
* [TvSeasonExternalIds](docs/sdks/sdk/README.md#tvseasonexternalids) - External IDs
* [TvSeasonImages](docs/sdks/sdk/README.md#tvseasonimages) - Images
* [TvSeasonTranslations](docs/sdks/sdk/README.md#tvseasontranslations) - Translations
* [TvSeasonVideos](docs/sdks/sdk/README.md#tvseasonvideos) - Videos
* [TvSeasonWatchProviders](docs/sdks/sdk/README.md#tvseasonwatchproviders) - Watch Providers
* [TvSeriesAccountStates](docs/sdks/sdk/README.md#tvseriesaccountstates) - Account States
* [TvSeriesAddRating](docs/sdks/sdk/README.md#tvseriesaddrating) - Add Rating
* [TvSeriesAggregateCredits](docs/sdks/sdk/README.md#tvseriesaggregatecredits) - Aggregate Credits
* [TvSeriesAiringTodayList](docs/sdks/sdk/README.md#tvseriesairingtodaylist) - Airing Today
* [TvSeriesAlternativeTitles](docs/sdks/sdk/README.md#tvseriesalternativetitles) - Alternative Titles
* [TvSeriesChanges](docs/sdks/sdk/README.md#tvserieschanges) - Changes
* [TvSeriesContentRatings](docs/sdks/sdk/README.md#tvseriescontentratings) - Content Ratings
* [TvSeriesCredits](docs/sdks/sdk/README.md#tvseriescredits) - Credits
* [TvSeriesDeleteRating](docs/sdks/sdk/README.md#tvseriesdeleterating) - Delete Rating
* [TvSeriesDetails](docs/sdks/sdk/README.md#tvseriesdetails) - Details
* [TvSeriesEpisodeGroups](docs/sdks/sdk/README.md#tvseriesepisodegroups) - Episode Groups
* [TvSeriesExternalIds](docs/sdks/sdk/README.md#tvseriesexternalids) - External IDs
* [TvSeriesImages](docs/sdks/sdk/README.md#tvseriesimages) - Images
* [TvSeriesKeywords](docs/sdks/sdk/README.md#tvserieskeywords) - Keywords
* [TvSeriesLatestID](docs/sdks/sdk/README.md#tvserieslatestid) - Latest
* [TvSeriesOnTheAirList](docs/sdks/sdk/README.md#tvseriesontheairlist) - On The Air
* [TvSeriesPopularList](docs/sdks/sdk/README.md#tvseriespopularlist) - Popular
* [TvSeriesRecommendations](docs/sdks/sdk/README.md#tvseriesrecommendations) - Recommendations
* [TvSeriesReviews](docs/sdks/sdk/README.md#tvseriesreviews) - Reviews
* [TvSeriesScreenedTheatrically](docs/sdks/sdk/README.md#tvseriesscreenedtheatrically) - Screened Theatrically
* [TvSeriesSimilar](docs/sdks/sdk/README.md#tvseriessimilar) - Similar
* [TvSeriesTopRatedList](docs/sdks/sdk/README.md#tvseriestopratedlist) - Top Rated
* [TvSeriesTranslations](docs/sdks/sdk/README.md#tvseriestranslations) - Translations
* [TvSeriesVideos](docs/sdks/sdk/README.md#tvseriesvideos) - Videos
* [TvSeriesWatchProviders](docs/sdks/sdk/README.md#tvserieswatchproviders) - Watch Providers
* [WatchProviderTvList](docs/sdks/sdk/README.md#watchprovidertvlist) - TV Providers
* [WatchProvidersAvailableRegions](docs/sdks/sdk/README.md#watchprovidersavailableregions) - Available Regions
* [WatchProvidersMovieList](docs/sdks/sdk/README.md#watchprovidersmovielist) - Movie Providers
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->



<!-- End Dev Containers -->

<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


<!-- End Error Handling -->

<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.themoviedb.org` | None |

For example:


```go
package main

import (
	"context"
	"github.com/dashotv/tmdb/openapi"
	"github.com/dashotv/tmdb/openapi/models/operations"
	"github.com/dashotv/tmdb/openapi/models/shared"
	"log"
)

func main() {
	s := openapi.New(
		openapi.WithSecurity(""),
		openapi.WithServerIndex(0),
	)

	var accountID int = 959345

	requestBody := &operations.AccountAddFavoriteRequestBody{
		RawBody: "{key: 16540, key1: null, key2: \"string\"}",
	}

	var sessionID *string = "string"

	ctx := context.Background()
	res, err := s.SDK.AccountAddFavorite(ctx, accountID, requestBody, sessionID)
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountAddFavorite200ApplicationJSONObject != nil {
		// handle response
	}
}

```


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:


```go
package main

import (
	"context"
	"github.com/dashotv/tmdb/openapi"
	"github.com/dashotv/tmdb/openapi/models/operations"
	"github.com/dashotv/tmdb/openapi/models/shared"
	"log"
)

func main() {
	s := openapi.New(
		openapi.WithSecurity(""),
		openapi.WithServerURL("https://api.themoviedb.org"),
	)

	var accountID int = 959345

	requestBody := &operations.AccountAddFavoriteRequestBody{
		RawBody: "{key: 16540, key1: null, key2: \"string\"}",
	}

	var sessionID *string = "string"

	ctx := context.Background()
	res, err := s.SDK.AccountAddFavorite(ctx, accountID, requestBody, sessionID)
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountAddFavorite200ApplicationJSONObject != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->

<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->

<!-- Start Go Types -->
# Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

## Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
