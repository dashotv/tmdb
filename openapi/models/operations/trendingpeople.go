// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TrendingPeopleTimeWindow string

const (
	TrendingPeopleTimeWindowDay  TrendingPeopleTimeWindow = "day"
	TrendingPeopleTimeWindowWeek TrendingPeopleTimeWindow = "week"
)

func (e TrendingPeopleTimeWindow) ToPointer() *TrendingPeopleTimeWindow {
	return &e
}

func (e *TrendingPeopleTimeWindow) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "day":
		fallthrough
	case "week":
		*e = TrendingPeopleTimeWindow(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TrendingPeopleTimeWindow: %v", v)
	}
}

type TrendingPeopleRequest struct {
	TimeWindow TrendingPeopleTimeWindow `default:"day" pathParam:"style=simple,explode=false,name=time_window"`
	// `ISO-639-1`-`ISO-3166-1` code
	Language *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
}

func (t TrendingPeopleRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TrendingPeopleRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TrendingPeopleRequest) GetTimeWindow() TrendingPeopleTimeWindow {
	if o == nil {
		return TrendingPeopleTimeWindow("")
	}
	return o.TimeWindow
}

func (o *TrendingPeopleRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

type TrendingPeople200ApplicationJSONResultsKnownFor struct {
	Adult            *bool    `default:"true" json:"adult"`
	BackdropPath     *string  `json:"backdrop_path,omitempty"`
	GenreIds         []int64  `json:"genre_ids,omitempty"`
	ID               *int64   `default:"0" json:"id"`
	MediaType        *string  `json:"media_type,omitempty"`
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

func (t TrendingPeople200ApplicationJSONResultsKnownFor) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TrendingPeople200ApplicationJSONResultsKnownFor) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetAdult() *bool {
	if o == nil {
		return nil
	}
	return o.Adult
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetBackdropPath() *string {
	if o == nil {
		return nil
	}
	return o.BackdropPath
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetGenreIds() []int64 {
	if o == nil {
		return nil
	}
	return o.GenreIds
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetMediaType() *string {
	if o == nil {
		return nil
	}
	return o.MediaType
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetOriginalLanguage() *string {
	if o == nil {
		return nil
	}
	return o.OriginalLanguage
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetOriginalTitle() *string {
	if o == nil {
		return nil
	}
	return o.OriginalTitle
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetPosterPath() *string {
	if o == nil {
		return nil
	}
	return o.PosterPath
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetReleaseDate() *string {
	if o == nil {
		return nil
	}
	return o.ReleaseDate
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetVideo() *bool {
	if o == nil {
		return nil
	}
	return o.Video
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetVoteAverage() *float64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *TrendingPeople200ApplicationJSONResultsKnownFor) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

type TrendingPeople200ApplicationJSONResults struct {
	Adult              *bool                                             `default:"true" json:"adult"`
	Gender             *int64                                            `default:"0" json:"gender"`
	ID                 *int64                                            `default:"0" json:"id"`
	KnownFor           []TrendingPeople200ApplicationJSONResultsKnownFor `json:"known_for,omitempty"`
	KnownForDepartment *string                                           `json:"known_for_department,omitempty"`
	MediaType          *string                                           `json:"media_type,omitempty"`
	Name               *string                                           `json:"name,omitempty"`
	OriginalName       *string                                           `json:"original_name,omitempty"`
	Popularity         *float64                                          `default:"0" json:"popularity"`
	ProfilePath        *string                                           `json:"profile_path,omitempty"`
}

func (t TrendingPeople200ApplicationJSONResults) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TrendingPeople200ApplicationJSONResults) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TrendingPeople200ApplicationJSONResults) GetAdult() *bool {
	if o == nil {
		return nil
	}
	return o.Adult
}

func (o *TrendingPeople200ApplicationJSONResults) GetGender() *int64 {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *TrendingPeople200ApplicationJSONResults) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TrendingPeople200ApplicationJSONResults) GetKnownFor() []TrendingPeople200ApplicationJSONResultsKnownFor {
	if o == nil {
		return nil
	}
	return o.KnownFor
}

func (o *TrendingPeople200ApplicationJSONResults) GetKnownForDepartment() *string {
	if o == nil {
		return nil
	}
	return o.KnownForDepartment
}

func (o *TrendingPeople200ApplicationJSONResults) GetMediaType() *string {
	if o == nil {
		return nil
	}
	return o.MediaType
}

func (o *TrendingPeople200ApplicationJSONResults) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TrendingPeople200ApplicationJSONResults) GetOriginalName() *string {
	if o == nil {
		return nil
	}
	return o.OriginalName
}

func (o *TrendingPeople200ApplicationJSONResults) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *TrendingPeople200ApplicationJSONResults) GetProfilePath() *string {
	if o == nil {
		return nil
	}
	return o.ProfilePath
}

// TrendingPeople200ApplicationJSON - 200
type TrendingPeople200ApplicationJSON struct {
	Page         *int64                                    `default:"0" json:"page"`
	Results      []TrendingPeople200ApplicationJSONResults `json:"results,omitempty"`
	TotalPages   *int64                                    `default:"0" json:"total_pages"`
	TotalResults *int64                                    `default:"0" json:"total_results"`
}

func (t TrendingPeople200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TrendingPeople200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TrendingPeople200ApplicationJSON) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *TrendingPeople200ApplicationJSON) GetResults() []TrendingPeople200ApplicationJSONResults {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *TrendingPeople200ApplicationJSON) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

func (o *TrendingPeople200ApplicationJSON) GetTotalResults() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalResults
}

type TrendingPeopleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TrendingPeople200ApplicationJSONObject *TrendingPeople200ApplicationJSON
}

func (o *TrendingPeopleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TrendingPeopleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TrendingPeopleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TrendingPeopleResponse) GetTrendingPeople200ApplicationJSONObject() *TrendingPeople200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TrendingPeople200ApplicationJSONObject
}
