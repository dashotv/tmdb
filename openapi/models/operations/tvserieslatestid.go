// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvSeriesLatestID200ApplicationJSONLastEpisodeToAir struct {
	AirDate        *string     `json:"air_date,omitempty"`
	EpisodeNumber  *int64      `default:"0" json:"episode_number"`
	ID             *int64      `default:"0" json:"id"`
	Name           *string     `json:"name,omitempty"`
	Overview       *string     `json:"overview,omitempty"`
	ProductionCode *string     `json:"production_code,omitempty"`
	Runtime        interface{} `json:"runtime,omitempty"`
	SeasonNumber   *int64      `default:"0" json:"season_number"`
	ShowID         *int64      `default:"0" json:"show_id"`
	StillPath      interface{} `json:"still_path,omitempty"`
	VoteAverage    *float64    `default:"0" json:"vote_average"`
	VoteCount      *int64      `default:"0" json:"vote_count"`
}

func (t TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) GetAirDate() *string {
	if o == nil {
		return nil
	}
	return o.AirDate
}

func (o *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) GetEpisodeNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodeNumber
}

func (o *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) GetProductionCode() *string {
	if o == nil {
		return nil
	}
	return o.ProductionCode
}

func (o *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) GetRuntime() interface{} {
	if o == nil {
		return nil
	}
	return o.Runtime
}

func (o *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) GetSeasonNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.SeasonNumber
}

func (o *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) GetShowID() *int64 {
	if o == nil {
		return nil
	}
	return o.ShowID
}

func (o *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) GetStillPath() interface{} {
	if o == nil {
		return nil
	}
	return o.StillPath
}

func (o *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) GetVoteAverage() *float64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

type TvSeriesLatestID200ApplicationJSONSeasons struct {
	AirDate      interface{} `json:"air_date,omitempty"`
	EpisodeCount *int64      `default:"0" json:"episode_count"`
	ID           *int64      `default:"0" json:"id"`
	Name         *string     `json:"name,omitempty"`
	Overview     *string     `json:"overview,omitempty"`
	PosterPath   interface{} `json:"poster_path,omitempty"`
	SeasonNumber *int64      `default:"0" json:"season_number"`
}

func (t TvSeriesLatestID200ApplicationJSONSeasons) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesLatestID200ApplicationJSONSeasons) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesLatestID200ApplicationJSONSeasons) GetAirDate() interface{} {
	if o == nil {
		return nil
	}
	return o.AirDate
}

func (o *TvSeriesLatestID200ApplicationJSONSeasons) GetEpisodeCount() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodeCount
}

func (o *TvSeriesLatestID200ApplicationJSONSeasons) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeriesLatestID200ApplicationJSONSeasons) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TvSeriesLatestID200ApplicationJSONSeasons) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *TvSeriesLatestID200ApplicationJSONSeasons) GetPosterPath() interface{} {
	if o == nil {
		return nil
	}
	return o.PosterPath
}

func (o *TvSeriesLatestID200ApplicationJSONSeasons) GetSeasonNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.SeasonNumber
}

// TvSeriesLatestID200ApplicationJSON - 200
type TvSeriesLatestID200ApplicationJSON struct {
	Adult               *bool                                               `default:"true" json:"adult"`
	BackdropPath        interface{}                                         `json:"backdrop_path,omitempty"`
	CreatedBy           []interface{}                                       `json:"created_by,omitempty"`
	EpisodeRunTime      []interface{}                                       `json:"episode_run_time,omitempty"`
	FirstAirDate        *string                                             `json:"first_air_date,omitempty"`
	Genres              []interface{}                                       `json:"genres,omitempty"`
	Homepage            *string                                             `json:"homepage,omitempty"`
	ID                  *int64                                              `default:"0" json:"id"`
	InProduction        *bool                                               `default:"true" json:"in_production"`
	Languages           []interface{}                                       `json:"languages,omitempty"`
	LastAirDate         *string                                             `json:"last_air_date,omitempty"`
	LastEpisodeToAir    *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir `json:"last_episode_to_air,omitempty"`
	Name                *string                                             `json:"name,omitempty"`
	Networks            []interface{}                                       `json:"networks,omitempty"`
	NextEpisodeToAir    interface{}                                         `json:"next_episode_to_air,omitempty"`
	NumberOfEpisodes    *int64                                              `default:"0" json:"number_of_episodes"`
	NumberOfSeasons     *int64                                              `default:"0" json:"number_of_seasons"`
	OriginCountry       []string                                            `json:"origin_country,omitempty"`
	OriginalLanguage    *string                                             `json:"original_language,omitempty"`
	OriginalName        *string                                             `json:"original_name,omitempty"`
	Overview            *string                                             `json:"overview,omitempty"`
	Popularity          *float64                                            `default:"0" json:"popularity"`
	PosterPath          interface{}                                         `json:"poster_path,omitempty"`
	ProductionCompanies []interface{}                                       `json:"production_companies,omitempty"`
	ProductionCountries []interface{}                                       `json:"production_countries,omitempty"`
	Seasons             []TvSeriesLatestID200ApplicationJSONSeasons         `json:"seasons,omitempty"`
	SpokenLanguages     []interface{}                                       `json:"spoken_languages,omitempty"`
	Status              *string                                             `json:"status,omitempty"`
	Tagline             *string                                             `json:"tagline,omitempty"`
	Type                *string                                             `json:"type,omitempty"`
	VoteAverage         *float64                                            `default:"0" json:"vote_average"`
	VoteCount           *int64                                              `default:"0" json:"vote_count"`
}

func (t TvSeriesLatestID200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesLatestID200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesLatestID200ApplicationJSON) GetAdult() *bool {
	if o == nil {
		return nil
	}
	return o.Adult
}

func (o *TvSeriesLatestID200ApplicationJSON) GetBackdropPath() interface{} {
	if o == nil {
		return nil
	}
	return o.BackdropPath
}

func (o *TvSeriesLatestID200ApplicationJSON) GetCreatedBy() []interface{} {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *TvSeriesLatestID200ApplicationJSON) GetEpisodeRunTime() []interface{} {
	if o == nil {
		return nil
	}
	return o.EpisodeRunTime
}

func (o *TvSeriesLatestID200ApplicationJSON) GetFirstAirDate() *string {
	if o == nil {
		return nil
	}
	return o.FirstAirDate
}

func (o *TvSeriesLatestID200ApplicationJSON) GetGenres() []interface{} {
	if o == nil {
		return nil
	}
	return o.Genres
}

func (o *TvSeriesLatestID200ApplicationJSON) GetHomepage() *string {
	if o == nil {
		return nil
	}
	return o.Homepage
}

func (o *TvSeriesLatestID200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeriesLatestID200ApplicationJSON) GetInProduction() *bool {
	if o == nil {
		return nil
	}
	return o.InProduction
}

func (o *TvSeriesLatestID200ApplicationJSON) GetLanguages() []interface{} {
	if o == nil {
		return nil
	}
	return o.Languages
}

func (o *TvSeriesLatestID200ApplicationJSON) GetLastAirDate() *string {
	if o == nil {
		return nil
	}
	return o.LastAirDate
}

func (o *TvSeriesLatestID200ApplicationJSON) GetLastEpisodeToAir() *TvSeriesLatestID200ApplicationJSONLastEpisodeToAir {
	if o == nil {
		return nil
	}
	return o.LastEpisodeToAir
}

func (o *TvSeriesLatestID200ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TvSeriesLatestID200ApplicationJSON) GetNetworks() []interface{} {
	if o == nil {
		return nil
	}
	return o.Networks
}

func (o *TvSeriesLatestID200ApplicationJSON) GetNextEpisodeToAir() interface{} {
	if o == nil {
		return nil
	}
	return o.NextEpisodeToAir
}

func (o *TvSeriesLatestID200ApplicationJSON) GetNumberOfEpisodes() *int64 {
	if o == nil {
		return nil
	}
	return o.NumberOfEpisodes
}

func (o *TvSeriesLatestID200ApplicationJSON) GetNumberOfSeasons() *int64 {
	if o == nil {
		return nil
	}
	return o.NumberOfSeasons
}

func (o *TvSeriesLatestID200ApplicationJSON) GetOriginCountry() []string {
	if o == nil {
		return nil
	}
	return o.OriginCountry
}

func (o *TvSeriesLatestID200ApplicationJSON) GetOriginalLanguage() *string {
	if o == nil {
		return nil
	}
	return o.OriginalLanguage
}

func (o *TvSeriesLatestID200ApplicationJSON) GetOriginalName() *string {
	if o == nil {
		return nil
	}
	return o.OriginalName
}

func (o *TvSeriesLatestID200ApplicationJSON) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *TvSeriesLatestID200ApplicationJSON) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *TvSeriesLatestID200ApplicationJSON) GetPosterPath() interface{} {
	if o == nil {
		return nil
	}
	return o.PosterPath
}

func (o *TvSeriesLatestID200ApplicationJSON) GetProductionCompanies() []interface{} {
	if o == nil {
		return nil
	}
	return o.ProductionCompanies
}

func (o *TvSeriesLatestID200ApplicationJSON) GetProductionCountries() []interface{} {
	if o == nil {
		return nil
	}
	return o.ProductionCountries
}

func (o *TvSeriesLatestID200ApplicationJSON) GetSeasons() []TvSeriesLatestID200ApplicationJSONSeasons {
	if o == nil {
		return nil
	}
	return o.Seasons
}

func (o *TvSeriesLatestID200ApplicationJSON) GetSpokenLanguages() []interface{} {
	if o == nil {
		return nil
	}
	return o.SpokenLanguages
}

func (o *TvSeriesLatestID200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TvSeriesLatestID200ApplicationJSON) GetTagline() *string {
	if o == nil {
		return nil
	}
	return o.Tagline
}

func (o *TvSeriesLatestID200ApplicationJSON) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *TvSeriesLatestID200ApplicationJSON) GetVoteAverage() *float64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *TvSeriesLatestID200ApplicationJSON) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

type TvSeriesLatestIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvSeriesLatestID200ApplicationJSONObject *TvSeriesLatestID200ApplicationJSON
}

func (o *TvSeriesLatestIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvSeriesLatestIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvSeriesLatestIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvSeriesLatestIDResponse) GetTvSeriesLatestID200ApplicationJSONObject() *TvSeriesLatestID200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvSeriesLatestID200ApplicationJSONObject
}
