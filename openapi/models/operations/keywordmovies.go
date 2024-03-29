// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type KeywordMoviesRequest struct {
	KeywordID    string  `pathParam:"style=simple,explode=false,name=keyword_id"`
	IncludeAdult *bool   `default:"false" queryParam:"style=form,explode=true,name=include_adult"`
	Language     *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
	Page         *int    `default:"1" queryParam:"style=form,explode=true,name=page"`
}

func (k KeywordMoviesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KeywordMoviesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *KeywordMoviesRequest) GetKeywordID() string {
	if o == nil {
		return ""
	}
	return o.KeywordID
}

func (o *KeywordMoviesRequest) GetIncludeAdult() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeAdult
}

func (o *KeywordMoviesRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *KeywordMoviesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

type KeywordMovies200ApplicationJSONResults struct {
	Adult            *bool    `default:"true" json:"adult"`
	BackdropPath     *string  `json:"backdrop_path,omitempty"`
	GenreIds         []int64  `json:"genre_ids,omitempty"`
	ID               *int64   `default:"0" json:"id"`
	OriginalLanguage *string  `json:"original_language,omitempty"`
	OriginalTitle    *string  `json:"original_title,omitempty"`
	Overview         *string  `json:"overview,omitempty"`
	Popularity       *float64 `default:"0" json:"popularity"`
	PosterPath       *string  `json:"poster_path,omitempty"`
	ReleaseDate      *string  `json:"release_date,omitempty"`
	Title            *string  `json:"title,omitempty"`
	Video            *bool    `default:"true" json:"video"`
	VoteAverage      *float64 `default:"0" json:"vote_average"`
	VoteCount        *int64   `default:"0" json:"vote_count"`
}

func (k KeywordMovies200ApplicationJSONResults) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KeywordMovies200ApplicationJSONResults) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *KeywordMovies200ApplicationJSONResults) GetAdult() *bool {
	if o == nil {
		return nil
	}
	return o.Adult
}

func (o *KeywordMovies200ApplicationJSONResults) GetBackdropPath() *string {
	if o == nil {
		return nil
	}
	return o.BackdropPath
}

func (o *KeywordMovies200ApplicationJSONResults) GetGenreIds() []int64 {
	if o == nil {
		return nil
	}
	return o.GenreIds
}

func (o *KeywordMovies200ApplicationJSONResults) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *KeywordMovies200ApplicationJSONResults) GetOriginalLanguage() *string {
	if o == nil {
		return nil
	}
	return o.OriginalLanguage
}

func (o *KeywordMovies200ApplicationJSONResults) GetOriginalTitle() *string {
	if o == nil {
		return nil
	}
	return o.OriginalTitle
}

func (o *KeywordMovies200ApplicationJSONResults) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *KeywordMovies200ApplicationJSONResults) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *KeywordMovies200ApplicationJSONResults) GetPosterPath() *string {
	if o == nil {
		return nil
	}
	return o.PosterPath
}

func (o *KeywordMovies200ApplicationJSONResults) GetReleaseDate() *string {
	if o == nil {
		return nil
	}
	return o.ReleaseDate
}

func (o *KeywordMovies200ApplicationJSONResults) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *KeywordMovies200ApplicationJSONResults) GetVideo() *bool {
	if o == nil {
		return nil
	}
	return o.Video
}

func (o *KeywordMovies200ApplicationJSONResults) GetVoteAverage() *float64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *KeywordMovies200ApplicationJSONResults) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

// KeywordMovies200ApplicationJSON - 200
type KeywordMovies200ApplicationJSON struct {
	ID           *int64                                   `default:"0" json:"id"`
	Page         *int64                                   `default:"0" json:"page"`
	Results      []KeywordMovies200ApplicationJSONResults `json:"results,omitempty"`
	TotalPages   *int64                                   `default:"0" json:"total_pages"`
	TotalResults *int64                                   `default:"0" json:"total_results"`
}

func (k KeywordMovies200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KeywordMovies200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *KeywordMovies200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *KeywordMovies200ApplicationJSON) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *KeywordMovies200ApplicationJSON) GetResults() []KeywordMovies200ApplicationJSONResults {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *KeywordMovies200ApplicationJSON) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

func (o *KeywordMovies200ApplicationJSON) GetTotalResults() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalResults
}

type KeywordMoviesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	KeywordMovies200ApplicationJSONObject *KeywordMovies200ApplicationJSON
}

func (o *KeywordMoviesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *KeywordMoviesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *KeywordMoviesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *KeywordMoviesResponse) GetKeywordMovies200ApplicationJSONObject() *KeywordMovies200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.KeywordMovies200ApplicationJSONObject
}
