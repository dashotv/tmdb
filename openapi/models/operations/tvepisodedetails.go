// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvEpisodeDetailsRequest struct {
	// comma separated list of endpoints within this namespace, 20 items max
	AppendToResponse *string `queryParam:"style=form,explode=true,name=append_to_response"`
	EpisodeNumber    int     `pathParam:"style=simple,explode=false,name=episode_number"`
	Language         *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
	SeasonNumber     int     `pathParam:"style=simple,explode=false,name=season_number"`
	SeriesID         int     `pathParam:"style=simple,explode=false,name=series_id"`
}

func (t TvEpisodeDetailsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvEpisodeDetailsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvEpisodeDetailsRequest) GetAppendToResponse() *string {
	if o == nil {
		return nil
	}
	return o.AppendToResponse
}

func (o *TvEpisodeDetailsRequest) GetEpisodeNumber() int {
	if o == nil {
		return 0
	}
	return o.EpisodeNumber
}

func (o *TvEpisodeDetailsRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *TvEpisodeDetailsRequest) GetSeasonNumber() int {
	if o == nil {
		return 0
	}
	return o.SeasonNumber
}

func (o *TvEpisodeDetailsRequest) GetSeriesID() int {
	if o == nil {
		return 0
	}
	return o.SeriesID
}

type TvEpisodeDetails200ApplicationJSONCrew struct {
	Adult              *bool    `default:"true" json:"adult"`
	CreditID           *string  `json:"credit_id,omitempty"`
	Department         *string  `json:"department,omitempty"`
	Gender             *int64   `default:"0" json:"gender"`
	ID                 *int64   `default:"0" json:"id"`
	Job                *string  `json:"job,omitempty"`
	KnownForDepartment *string  `json:"known_for_department,omitempty"`
	Name               *string  `json:"name,omitempty"`
	OriginalName       *string  `json:"original_name,omitempty"`
	Popularity         *float64 `default:"0" json:"popularity"`
	ProfilePath        *string  `json:"profile_path,omitempty"`
}

func (t TvEpisodeDetails200ApplicationJSONCrew) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvEpisodeDetails200ApplicationJSONCrew) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvEpisodeDetails200ApplicationJSONCrew) GetAdult() *bool {
	if o == nil {
		return nil
	}
	return o.Adult
}

func (o *TvEpisodeDetails200ApplicationJSONCrew) GetCreditID() *string {
	if o == nil {
		return nil
	}
	return o.CreditID
}

func (o *TvEpisodeDetails200ApplicationJSONCrew) GetDepartment() *string {
	if o == nil {
		return nil
	}
	return o.Department
}

func (o *TvEpisodeDetails200ApplicationJSONCrew) GetGender() *int64 {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *TvEpisodeDetails200ApplicationJSONCrew) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvEpisodeDetails200ApplicationJSONCrew) GetJob() *string {
	if o == nil {
		return nil
	}
	return o.Job
}

func (o *TvEpisodeDetails200ApplicationJSONCrew) GetKnownForDepartment() *string {
	if o == nil {
		return nil
	}
	return o.KnownForDepartment
}

func (o *TvEpisodeDetails200ApplicationJSONCrew) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TvEpisodeDetails200ApplicationJSONCrew) GetOriginalName() *string {
	if o == nil {
		return nil
	}
	return o.OriginalName
}

func (o *TvEpisodeDetails200ApplicationJSONCrew) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *TvEpisodeDetails200ApplicationJSONCrew) GetProfilePath() *string {
	if o == nil {
		return nil
	}
	return o.ProfilePath
}

type TvEpisodeDetails200ApplicationJSONGuestStars struct {
	Adult              *bool    `default:"true" json:"adult"`
	Character          *string  `json:"character,omitempty"`
	CreditID           *string  `json:"credit_id,omitempty"`
	Gender             *int64   `default:"0" json:"gender"`
	ID                 *int64   `default:"0" json:"id"`
	KnownForDepartment *string  `json:"known_for_department,omitempty"`
	Name               *string  `json:"name,omitempty"`
	Order              *int64   `default:"0" json:"order"`
	OriginalName       *string  `json:"original_name,omitempty"`
	Popularity         *float64 `default:"0" json:"popularity"`
	ProfilePath        *string  `json:"profile_path,omitempty"`
}

func (t TvEpisodeDetails200ApplicationJSONGuestStars) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvEpisodeDetails200ApplicationJSONGuestStars) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvEpisodeDetails200ApplicationJSONGuestStars) GetAdult() *bool {
	if o == nil {
		return nil
	}
	return o.Adult
}

func (o *TvEpisodeDetails200ApplicationJSONGuestStars) GetCharacter() *string {
	if o == nil {
		return nil
	}
	return o.Character
}

func (o *TvEpisodeDetails200ApplicationJSONGuestStars) GetCreditID() *string {
	if o == nil {
		return nil
	}
	return o.CreditID
}

func (o *TvEpisodeDetails200ApplicationJSONGuestStars) GetGender() *int64 {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *TvEpisodeDetails200ApplicationJSONGuestStars) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvEpisodeDetails200ApplicationJSONGuestStars) GetKnownForDepartment() *string {
	if o == nil {
		return nil
	}
	return o.KnownForDepartment
}

func (o *TvEpisodeDetails200ApplicationJSONGuestStars) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TvEpisodeDetails200ApplicationJSONGuestStars) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *TvEpisodeDetails200ApplicationJSONGuestStars) GetOriginalName() *string {
	if o == nil {
		return nil
	}
	return o.OriginalName
}

func (o *TvEpisodeDetails200ApplicationJSONGuestStars) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *TvEpisodeDetails200ApplicationJSONGuestStars) GetProfilePath() *string {
	if o == nil {
		return nil
	}
	return o.ProfilePath
}

// TvEpisodeDetails200ApplicationJSON - 200
type TvEpisodeDetails200ApplicationJSON struct {
	AirDate        *string                                        `json:"air_date,omitempty"`
	Crew           []TvEpisodeDetails200ApplicationJSONCrew       `json:"crew,omitempty"`
	EpisodeNumber  *int64                                         `default:"0" json:"episode_number"`
	GuestStars     []TvEpisodeDetails200ApplicationJSONGuestStars `json:"guest_stars,omitempty"`
	ID             *int64                                         `default:"0" json:"id"`
	Name           *string                                        `json:"name,omitempty"`
	Overview       *string                                        `json:"overview,omitempty"`
	ProductionCode *string                                        `json:"production_code,omitempty"`
	Runtime        *int64                                         `default:"0" json:"runtime"`
	SeasonNumber   *int64                                         `default:"0" json:"season_number"`
	StillPath      *string                                        `json:"still_path,omitempty"`
	VoteAverage    *float64                                       `default:"0" json:"vote_average"`
	VoteCount      *int64                                         `default:"0" json:"vote_count"`
}

func (t TvEpisodeDetails200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvEpisodeDetails200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvEpisodeDetails200ApplicationJSON) GetAirDate() *string {
	if o == nil {
		return nil
	}
	return o.AirDate
}

func (o *TvEpisodeDetails200ApplicationJSON) GetCrew() []TvEpisodeDetails200ApplicationJSONCrew {
	if o == nil {
		return nil
	}
	return o.Crew
}

func (o *TvEpisodeDetails200ApplicationJSON) GetEpisodeNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodeNumber
}

func (o *TvEpisodeDetails200ApplicationJSON) GetGuestStars() []TvEpisodeDetails200ApplicationJSONGuestStars {
	if o == nil {
		return nil
	}
	return o.GuestStars
}

func (o *TvEpisodeDetails200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvEpisodeDetails200ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TvEpisodeDetails200ApplicationJSON) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *TvEpisodeDetails200ApplicationJSON) GetProductionCode() *string {
	if o == nil {
		return nil
	}
	return o.ProductionCode
}

func (o *TvEpisodeDetails200ApplicationJSON) GetRuntime() *int64 {
	if o == nil {
		return nil
	}
	return o.Runtime
}

func (o *TvEpisodeDetails200ApplicationJSON) GetSeasonNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.SeasonNumber
}

func (o *TvEpisodeDetails200ApplicationJSON) GetStillPath() *string {
	if o == nil {
		return nil
	}
	return o.StillPath
}

func (o *TvEpisodeDetails200ApplicationJSON) GetVoteAverage() *float64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *TvEpisodeDetails200ApplicationJSON) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

type TvEpisodeDetailsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvEpisodeDetails200ApplicationJSONObject *TvEpisodeDetails200ApplicationJSON
}

func (o *TvEpisodeDetailsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvEpisodeDetailsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvEpisodeDetailsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvEpisodeDetailsResponse) GetTvEpisodeDetails200ApplicationJSONObject() *TvEpisodeDetails200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvEpisodeDetails200ApplicationJSONObject
}
