// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type MovieDetailsRequest struct {
	MovieID int `pathParam:"style=simple,explode=false,name=movie_id"`
	// comma separated list of endpoints within this namespace, 20 items max
	AppendToResponse *string `queryParam:"style=form,explode=true,name=append_to_response"`
	Language         *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
}

func (m MovieDetailsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MovieDetailsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MovieDetailsRequest) GetMovieID() int {
	if o == nil {
		return 0
	}
	return o.MovieID
}

func (o *MovieDetailsRequest) GetAppendToResponse() *string {
	if o == nil {
		return nil
	}
	return o.AppendToResponse
}

func (o *MovieDetailsRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

type MovieDetails200ApplicationJSONGenres struct {
	ID   *int64  `default:"0" json:"id"`
	Name *string `json:"name,omitempty"`
}

func (m MovieDetails200ApplicationJSONGenres) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MovieDetails200ApplicationJSONGenres) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MovieDetails200ApplicationJSONGenres) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MovieDetails200ApplicationJSONGenres) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type MovieDetails200ApplicationJSONProductionCompanies struct {
	ID            *int64  `default:"0" json:"id"`
	LogoPath      *string `json:"logo_path,omitempty"`
	Name          *string `json:"name,omitempty"`
	OriginCountry *string `json:"origin_country,omitempty"`
}

func (m MovieDetails200ApplicationJSONProductionCompanies) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MovieDetails200ApplicationJSONProductionCompanies) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MovieDetails200ApplicationJSONProductionCompanies) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MovieDetails200ApplicationJSONProductionCompanies) GetLogoPath() *string {
	if o == nil {
		return nil
	}
	return o.LogoPath
}

func (o *MovieDetails200ApplicationJSONProductionCompanies) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MovieDetails200ApplicationJSONProductionCompanies) GetOriginCountry() *string {
	if o == nil {
		return nil
	}
	return o.OriginCountry
}

type MovieDetails200ApplicationJSONProductionCountries struct {
	Iso31661 *string `json:"iso_3166_1,omitempty"`
	Name     *string `json:"name,omitempty"`
}

func (o *MovieDetails200ApplicationJSONProductionCountries) GetIso31661() *string {
	if o == nil {
		return nil
	}
	return o.Iso31661
}

func (o *MovieDetails200ApplicationJSONProductionCountries) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type MovieDetails200ApplicationJSONSpokenLanguages struct {
	EnglishName *string `json:"english_name,omitempty"`
	Iso6391     *string `json:"iso_639_1,omitempty"`
	Name        *string `json:"name,omitempty"`
}

func (o *MovieDetails200ApplicationJSONSpokenLanguages) GetEnglishName() *string {
	if o == nil {
		return nil
	}
	return o.EnglishName
}

func (o *MovieDetails200ApplicationJSONSpokenLanguages) GetIso6391() *string {
	if o == nil {
		return nil
	}
	return o.Iso6391
}

func (o *MovieDetails200ApplicationJSONSpokenLanguages) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// MovieDetails200ApplicationJSON - 200
type MovieDetails200ApplicationJSON struct {
	Adult               *bool                                               `default:"true" json:"adult"`
	BackdropPath        *string                                             `json:"backdrop_path,omitempty"`
	BelongsToCollection interface{}                                         `json:"belongs_to_collection,omitempty"`
	Budget              *int64                                              `default:"0" json:"budget"`
	Genres              []MovieDetails200ApplicationJSONGenres              `json:"genres,omitempty"`
	Homepage            *string                                             `json:"homepage,omitempty"`
	ID                  *int64                                              `default:"0" json:"id"`
	ImdbID              *string                                             `json:"imdb_id,omitempty"`
	OriginalLanguage    *string                                             `json:"original_language,omitempty"`
	OriginalTitle       *string                                             `json:"original_title,omitempty"`
	Overview            *string                                             `json:"overview,omitempty"`
	Popularity          *float64                                            `default:"0" json:"popularity"`
	PosterPath          *string                                             `json:"poster_path,omitempty"`
	ProductionCompanies []MovieDetails200ApplicationJSONProductionCompanies `json:"production_companies,omitempty"`
	ProductionCountries []MovieDetails200ApplicationJSONProductionCountries `json:"production_countries,omitempty"`
	ReleaseDate         *string                                             `json:"release_date,omitempty"`
	Revenue             *int64                                              `default:"0" json:"revenue"`
	Runtime             *int64                                              `default:"0" json:"runtime"`
	SpokenLanguages     []MovieDetails200ApplicationJSONSpokenLanguages     `json:"spoken_languages,omitempty"`
	Status              *string                                             `json:"status,omitempty"`
	Tagline             *string                                             `json:"tagline,omitempty"`
	Title               *string                                             `json:"title,omitempty"`
	Video               *bool                                               `default:"true" json:"video"`
	VoteAverage         *float64                                            `default:"0" json:"vote_average"`
	VoteCount           *int64                                              `default:"0" json:"vote_count"`
}

func (m MovieDetails200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MovieDetails200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MovieDetails200ApplicationJSON) GetAdult() *bool {
	if o == nil {
		return nil
	}
	return o.Adult
}

func (o *MovieDetails200ApplicationJSON) GetBackdropPath() *string {
	if o == nil {
		return nil
	}
	return o.BackdropPath
}

func (o *MovieDetails200ApplicationJSON) GetBelongsToCollection() interface{} {
	if o == nil {
		return nil
	}
	return o.BelongsToCollection
}

func (o *MovieDetails200ApplicationJSON) GetBudget() *int64 {
	if o == nil {
		return nil
	}
	return o.Budget
}

func (o *MovieDetails200ApplicationJSON) GetGenres() []MovieDetails200ApplicationJSONGenres {
	if o == nil {
		return nil
	}
	return o.Genres
}

func (o *MovieDetails200ApplicationJSON) GetHomepage() *string {
	if o == nil {
		return nil
	}
	return o.Homepage
}

func (o *MovieDetails200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MovieDetails200ApplicationJSON) GetImdbID() *string {
	if o == nil {
		return nil
	}
	return o.ImdbID
}

func (o *MovieDetails200ApplicationJSON) GetOriginalLanguage() *string {
	if o == nil {
		return nil
	}
	return o.OriginalLanguage
}

func (o *MovieDetails200ApplicationJSON) GetOriginalTitle() *string {
	if o == nil {
		return nil
	}
	return o.OriginalTitle
}

func (o *MovieDetails200ApplicationJSON) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *MovieDetails200ApplicationJSON) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *MovieDetails200ApplicationJSON) GetPosterPath() *string {
	if o == nil {
		return nil
	}
	return o.PosterPath
}

func (o *MovieDetails200ApplicationJSON) GetProductionCompanies() []MovieDetails200ApplicationJSONProductionCompanies {
	if o == nil {
		return nil
	}
	return o.ProductionCompanies
}

func (o *MovieDetails200ApplicationJSON) GetProductionCountries() []MovieDetails200ApplicationJSONProductionCountries {
	if o == nil {
		return nil
	}
	return o.ProductionCountries
}

func (o *MovieDetails200ApplicationJSON) GetReleaseDate() *string {
	if o == nil {
		return nil
	}
	return o.ReleaseDate
}

func (o *MovieDetails200ApplicationJSON) GetRevenue() *int64 {
	if o == nil {
		return nil
	}
	return o.Revenue
}

func (o *MovieDetails200ApplicationJSON) GetRuntime() *int64 {
	if o == nil {
		return nil
	}
	return o.Runtime
}

func (o *MovieDetails200ApplicationJSON) GetSpokenLanguages() []MovieDetails200ApplicationJSONSpokenLanguages {
	if o == nil {
		return nil
	}
	return o.SpokenLanguages
}

func (o *MovieDetails200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *MovieDetails200ApplicationJSON) GetTagline() *string {
	if o == nil {
		return nil
	}
	return o.Tagline
}

func (o *MovieDetails200ApplicationJSON) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *MovieDetails200ApplicationJSON) GetVideo() *bool {
	if o == nil {
		return nil
	}
	return o.Video
}

func (o *MovieDetails200ApplicationJSON) GetVoteAverage() *float64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *MovieDetails200ApplicationJSON) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

type MovieDetailsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	MovieDetails200ApplicationJSONObject *MovieDetails200ApplicationJSON
}

func (o *MovieDetailsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MovieDetailsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MovieDetailsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MovieDetailsResponse) GetMovieDetails200ApplicationJSONObject() *MovieDetails200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.MovieDetails200ApplicationJSONObject
}