// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvSeriesPopularListRequest struct {
	Language *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
	Page     *int    `default:"1" queryParam:"style=form,explode=true,name=page"`
}

func (t TvSeriesPopularListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesPopularListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesPopularListRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *TvSeriesPopularListRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

type TvSeriesPopularList200ApplicationJSONResults struct {
	BackdropPath     *string  `json:"backdrop_path,omitempty"`
	FirstAirDate     *string  `json:"first_air_date,omitempty"`
	GenreIds         []int64  `json:"genre_ids,omitempty"`
	ID               *int64   `default:"0" json:"id"`
	Name             *string  `json:"name,omitempty"`
	OriginCountry    []string `json:"origin_country,omitempty"`
	OriginalLanguage *string  `json:"original_language,omitempty"`
	OriginalName     *string  `json:"original_name,omitempty"`
	Overview         *string  `json:"overview,omitempty"`
	Popularity       *float64 `default:"0" json:"popularity"`
	PosterPath       *string  `json:"poster_path,omitempty"`
	VoteAverage      *int64   `default:"0" json:"vote_average"`
	VoteCount        *int64   `default:"0" json:"vote_count"`
}

func (t TvSeriesPopularList200ApplicationJSONResults) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesPopularList200ApplicationJSONResults) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesPopularList200ApplicationJSONResults) GetBackdropPath() *string {
	if o == nil {
		return nil
	}
	return o.BackdropPath
}

func (o *TvSeriesPopularList200ApplicationJSONResults) GetFirstAirDate() *string {
	if o == nil {
		return nil
	}
	return o.FirstAirDate
}

func (o *TvSeriesPopularList200ApplicationJSONResults) GetGenreIds() []int64 {
	if o == nil {
		return nil
	}
	return o.GenreIds
}

func (o *TvSeriesPopularList200ApplicationJSONResults) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeriesPopularList200ApplicationJSONResults) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TvSeriesPopularList200ApplicationJSONResults) GetOriginCountry() []string {
	if o == nil {
		return nil
	}
	return o.OriginCountry
}

func (o *TvSeriesPopularList200ApplicationJSONResults) GetOriginalLanguage() *string {
	if o == nil {
		return nil
	}
	return o.OriginalLanguage
}

func (o *TvSeriesPopularList200ApplicationJSONResults) GetOriginalName() *string {
	if o == nil {
		return nil
	}
	return o.OriginalName
}

func (o *TvSeriesPopularList200ApplicationJSONResults) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *TvSeriesPopularList200ApplicationJSONResults) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *TvSeriesPopularList200ApplicationJSONResults) GetPosterPath() *string {
	if o == nil {
		return nil
	}
	return o.PosterPath
}

func (o *TvSeriesPopularList200ApplicationJSONResults) GetVoteAverage() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *TvSeriesPopularList200ApplicationJSONResults) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

// TvSeriesPopularList200ApplicationJSON - 200
type TvSeriesPopularList200ApplicationJSON struct {
	Page         *int64                                         `default:"0" json:"page"`
	Results      []TvSeriesPopularList200ApplicationJSONResults `json:"results,omitempty"`
	TotalPages   *int64                                         `default:"0" json:"total_pages"`
	TotalResults *int64                                         `default:"0" json:"total_results"`
}

func (t TvSeriesPopularList200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesPopularList200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesPopularList200ApplicationJSON) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *TvSeriesPopularList200ApplicationJSON) GetResults() []TvSeriesPopularList200ApplicationJSONResults {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *TvSeriesPopularList200ApplicationJSON) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

func (o *TvSeriesPopularList200ApplicationJSON) GetTotalResults() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalResults
}

type TvSeriesPopularListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvSeriesPopularList200ApplicationJSONObject *TvSeriesPopularList200ApplicationJSON
}

func (o *TvSeriesPopularListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvSeriesPopularListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvSeriesPopularListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvSeriesPopularListResponse) GetTvSeriesPopularList200ApplicationJSONObject() *TvSeriesPopularList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvSeriesPopularList200ApplicationJSONObject
}