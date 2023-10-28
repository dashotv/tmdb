// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type PersonMovieCreditsRequest struct {
	PersonID int     `pathParam:"style=simple,explode=false,name=person_id"`
	Language *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
}

func (p PersonMovieCreditsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonMovieCreditsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PersonMovieCreditsRequest) GetPersonID() int {
	if o == nil {
		return 0
	}
	return o.PersonID
}

func (o *PersonMovieCreditsRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

type PersonMovieCredits200ApplicationJSONCast struct {
	Adult            *bool    `default:"true" json:"adult"`
	BackdropPath     *string  `json:"backdrop_path,omitempty"`
	Character        *string  `json:"character,omitempty"`
	CreditID         *string  `json:"credit_id,omitempty"`
	GenreIds         []int64  `json:"genre_ids,omitempty"`
	ID               *int64   `default:"0" json:"id"`
	Order            *int64   `default:"0" json:"order"`
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

func (p PersonMovieCredits200ApplicationJSONCast) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonMovieCredits200ApplicationJSONCast) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetAdult() *bool {
	if o == nil {
		return nil
	}
	return o.Adult
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetBackdropPath() *string {
	if o == nil {
		return nil
	}
	return o.BackdropPath
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetCharacter() *string {
	if o == nil {
		return nil
	}
	return o.Character
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetCreditID() *string {
	if o == nil {
		return nil
	}
	return o.CreditID
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetGenreIds() []int64 {
	if o == nil {
		return nil
	}
	return o.GenreIds
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetOriginalLanguage() *string {
	if o == nil {
		return nil
	}
	return o.OriginalLanguage
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetOriginalTitle() *string {
	if o == nil {
		return nil
	}
	return o.OriginalTitle
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetPosterPath() *string {
	if o == nil {
		return nil
	}
	return o.PosterPath
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetReleaseDate() *string {
	if o == nil {
		return nil
	}
	return o.ReleaseDate
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetVideo() *bool {
	if o == nil {
		return nil
	}
	return o.Video
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetVoteAverage() *float64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *PersonMovieCredits200ApplicationJSONCast) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

type PersonMovieCredits200ApplicationJSONCrew struct {
	Adult            *bool    `default:"true" json:"adult"`
	BackdropPath     *string  `json:"backdrop_path,omitempty"`
	CreditID         *string  `json:"credit_id,omitempty"`
	Department       *string  `json:"department,omitempty"`
	GenreIds         []int64  `json:"genre_ids,omitempty"`
	ID               *int64   `default:"0" json:"id"`
	Job              *string  `json:"job,omitempty"`
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

func (p PersonMovieCredits200ApplicationJSONCrew) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonMovieCredits200ApplicationJSONCrew) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetAdult() *bool {
	if o == nil {
		return nil
	}
	return o.Adult
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetBackdropPath() *string {
	if o == nil {
		return nil
	}
	return o.BackdropPath
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetCreditID() *string {
	if o == nil {
		return nil
	}
	return o.CreditID
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetDepartment() *string {
	if o == nil {
		return nil
	}
	return o.Department
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetGenreIds() []int64 {
	if o == nil {
		return nil
	}
	return o.GenreIds
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetJob() *string {
	if o == nil {
		return nil
	}
	return o.Job
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetOriginalLanguage() *string {
	if o == nil {
		return nil
	}
	return o.OriginalLanguage
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetOriginalTitle() *string {
	if o == nil {
		return nil
	}
	return o.OriginalTitle
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetPosterPath() *string {
	if o == nil {
		return nil
	}
	return o.PosterPath
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetReleaseDate() *string {
	if o == nil {
		return nil
	}
	return o.ReleaseDate
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetVideo() *bool {
	if o == nil {
		return nil
	}
	return o.Video
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetVoteAverage() *float64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *PersonMovieCredits200ApplicationJSONCrew) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

// PersonMovieCredits200ApplicationJSON - 200
type PersonMovieCredits200ApplicationJSON struct {
	Cast []PersonMovieCredits200ApplicationJSONCast `json:"cast,omitempty"`
	Crew []PersonMovieCredits200ApplicationJSONCrew `json:"crew,omitempty"`
	ID   *int64                                     `default:"0" json:"id"`
}

func (p PersonMovieCredits200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonMovieCredits200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PersonMovieCredits200ApplicationJSON) GetCast() []PersonMovieCredits200ApplicationJSONCast {
	if o == nil {
		return nil
	}
	return o.Cast
}

func (o *PersonMovieCredits200ApplicationJSON) GetCrew() []PersonMovieCredits200ApplicationJSONCrew {
	if o == nil {
		return nil
	}
	return o.Crew
}

func (o *PersonMovieCredits200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

type PersonMovieCreditsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	PersonMovieCredits200ApplicationJSONObject *PersonMovieCredits200ApplicationJSON
}

func (o *PersonMovieCreditsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PersonMovieCreditsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PersonMovieCreditsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PersonMovieCreditsResponse) GetPersonMovieCredits200ApplicationJSONObject() *PersonMovieCredits200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PersonMovieCredits200ApplicationJSONObject
}
