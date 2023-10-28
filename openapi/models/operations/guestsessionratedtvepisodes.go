// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type GuestSessionRatedTvEpisodesSortBy string

const (
	GuestSessionRatedTvEpisodesSortByCreatedAtAsc  GuestSessionRatedTvEpisodesSortBy = "created_at.asc"
	GuestSessionRatedTvEpisodesSortByCreatedAtDesc GuestSessionRatedTvEpisodesSortBy = "created_at.desc"
)

func (e GuestSessionRatedTvEpisodesSortBy) ToPointer() *GuestSessionRatedTvEpisodesSortBy {
	return &e
}

func (e *GuestSessionRatedTvEpisodesSortBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at.asc":
		fallthrough
	case "created_at.desc":
		*e = GuestSessionRatedTvEpisodesSortBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GuestSessionRatedTvEpisodesSortBy: %v", v)
	}
}

type GuestSessionRatedTvEpisodesRequest struct {
	GuestSessionID string                             `pathParam:"style=simple,explode=false,name=guest_session_id"`
	Language       *string                            `default:"en-US" queryParam:"style=form,explode=true,name=language"`
	Page           *int                               `default:"1" queryParam:"style=form,explode=true,name=page"`
	SortBy         *GuestSessionRatedTvEpisodesSortBy `default:"created_at.asc" queryParam:"style=form,explode=true,name=sort_by"`
}

func (g GuestSessionRatedTvEpisodesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GuestSessionRatedTvEpisodesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GuestSessionRatedTvEpisodesRequest) GetGuestSessionID() string {
	if o == nil {
		return ""
	}
	return o.GuestSessionID
}

func (o *GuestSessionRatedTvEpisodesRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *GuestSessionRatedTvEpisodesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GuestSessionRatedTvEpisodesRequest) GetSortBy() *GuestSessionRatedTvEpisodesSortBy {
	if o == nil {
		return nil
	}
	return o.SortBy
}

type GuestSessionRatedTvEpisodes200ApplicationJSONResults struct {
	AirDate        *string  `json:"air_date,omitempty"`
	EpisodeNumber  *int64   `default:"0" json:"episode_number"`
	ID             *int64   `default:"0" json:"id"`
	Name           *string  `json:"name,omitempty"`
	Overview       *string  `json:"overview,omitempty"`
	ProductionCode *string  `json:"production_code,omitempty"`
	Rating         *float64 `default:"0" json:"rating"`
	Runtime        *int64   `default:"0" json:"runtime"`
	SeasonNumber   *int64   `default:"0" json:"season_number"`
	ShowID         *int64   `default:"0" json:"show_id"`
	StillPath      *string  `json:"still_path,omitempty"`
	VoteAverage    *float64 `default:"0" json:"vote_average"`
	VoteCount      *int64   `default:"0" json:"vote_count"`
}

func (g GuestSessionRatedTvEpisodes200ApplicationJSONResults) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GuestSessionRatedTvEpisodes200ApplicationJSONResults) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSONResults) GetAirDate() *string {
	if o == nil {
		return nil
	}
	return o.AirDate
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSONResults) GetEpisodeNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodeNumber
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSONResults) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSONResults) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSONResults) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSONResults) GetProductionCode() *string {
	if o == nil {
		return nil
	}
	return o.ProductionCode
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSONResults) GetRating() *float64 {
	if o == nil {
		return nil
	}
	return o.Rating
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSONResults) GetRuntime() *int64 {
	if o == nil {
		return nil
	}
	return o.Runtime
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSONResults) GetSeasonNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.SeasonNumber
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSONResults) GetShowID() *int64 {
	if o == nil {
		return nil
	}
	return o.ShowID
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSONResults) GetStillPath() *string {
	if o == nil {
		return nil
	}
	return o.StillPath
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSONResults) GetVoteAverage() *float64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSONResults) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

// GuestSessionRatedTvEpisodes200ApplicationJSON - 200
type GuestSessionRatedTvEpisodes200ApplicationJSON struct {
	Page         *int64                                                 `default:"0" json:"page"`
	Results      []GuestSessionRatedTvEpisodes200ApplicationJSONResults `json:"results,omitempty"`
	TotalPages   *int64                                                 `default:"0" json:"total_pages"`
	TotalResults *int64                                                 `default:"0" json:"total_results"`
}

func (g GuestSessionRatedTvEpisodes200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GuestSessionRatedTvEpisodes200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSON) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSON) GetResults() []GuestSessionRatedTvEpisodes200ApplicationJSONResults {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSON) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

func (o *GuestSessionRatedTvEpisodes200ApplicationJSON) GetTotalResults() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalResults
}

type GuestSessionRatedTvEpisodesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	GuestSessionRatedTvEpisodes200ApplicationJSONObject *GuestSessionRatedTvEpisodes200ApplicationJSON
}

func (o *GuestSessionRatedTvEpisodesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GuestSessionRatedTvEpisodesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GuestSessionRatedTvEpisodesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GuestSessionRatedTvEpisodesResponse) GetGuestSessionRatedTvEpisodes200ApplicationJSONObject() *GuestSessionRatedTvEpisodes200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GuestSessionRatedTvEpisodes200ApplicationJSONObject
}