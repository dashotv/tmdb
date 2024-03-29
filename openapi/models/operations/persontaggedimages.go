// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type PersonTaggedImagesRequest struct {
	PersonID int  `pathParam:"style=simple,explode=false,name=person_id"`
	Page     *int `default:"1" queryParam:"style=form,explode=true,name=page"`
}

func (p PersonTaggedImagesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonTaggedImagesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PersonTaggedImagesRequest) GetPersonID() int {
	if o == nil {
		return 0
	}
	return o.PersonID
}

func (o *PersonTaggedImagesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

type PersonTaggedImages200ApplicationJSONResultsMedia struct {
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

func (p PersonTaggedImages200ApplicationJSONResultsMedia) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonTaggedImages200ApplicationJSONResultsMedia) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetAdult() *bool {
	if o == nil {
		return nil
	}
	return o.Adult
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetBackdropPath() *string {
	if o == nil {
		return nil
	}
	return o.BackdropPath
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetGenreIds() []int64 {
	if o == nil {
		return nil
	}
	return o.GenreIds
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetMediaType() *string {
	if o == nil {
		return nil
	}
	return o.MediaType
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetOriginalLanguage() *string {
	if o == nil {
		return nil
	}
	return o.OriginalLanguage
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetOriginalTitle() *string {
	if o == nil {
		return nil
	}
	return o.OriginalTitle
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetPosterPath() *string {
	if o == nil {
		return nil
	}
	return o.PosterPath
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetReleaseDate() *string {
	if o == nil {
		return nil
	}
	return o.ReleaseDate
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetVideo() *bool {
	if o == nil {
		return nil
	}
	return o.Video
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetVoteAverage() *float64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *PersonTaggedImages200ApplicationJSONResultsMedia) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

type PersonTaggedImages200ApplicationJSONResults struct {
	AspectRatio *float64                                          `default:"0" json:"aspect_ratio"`
	FilePath    *string                                           `json:"file_path,omitempty"`
	Height      *int64                                            `default:"0" json:"height"`
	ID          *string                                           `json:"id,omitempty"`
	ImageType   *string                                           `json:"image_type,omitempty"`
	Iso6391     *string                                           `json:"iso_639_1,omitempty"`
	Media       *PersonTaggedImages200ApplicationJSONResultsMedia `json:"media,omitempty"`
	MediaType   *string                                           `json:"media_type,omitempty"`
	VoteAverage *float64                                          `default:"0" json:"vote_average"`
	VoteCount   *int64                                            `default:"0" json:"vote_count"`
	Width       *int64                                            `default:"0" json:"width"`
}

func (p PersonTaggedImages200ApplicationJSONResults) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonTaggedImages200ApplicationJSONResults) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PersonTaggedImages200ApplicationJSONResults) GetAspectRatio() *float64 {
	if o == nil {
		return nil
	}
	return o.AspectRatio
}

func (o *PersonTaggedImages200ApplicationJSONResults) GetFilePath() *string {
	if o == nil {
		return nil
	}
	return o.FilePath
}

func (o *PersonTaggedImages200ApplicationJSONResults) GetHeight() *int64 {
	if o == nil {
		return nil
	}
	return o.Height
}

func (o *PersonTaggedImages200ApplicationJSONResults) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PersonTaggedImages200ApplicationJSONResults) GetImageType() *string {
	if o == nil {
		return nil
	}
	return o.ImageType
}

func (o *PersonTaggedImages200ApplicationJSONResults) GetIso6391() *string {
	if o == nil {
		return nil
	}
	return o.Iso6391
}

func (o *PersonTaggedImages200ApplicationJSONResults) GetMedia() *PersonTaggedImages200ApplicationJSONResultsMedia {
	if o == nil {
		return nil
	}
	return o.Media
}

func (o *PersonTaggedImages200ApplicationJSONResults) GetMediaType() *string {
	if o == nil {
		return nil
	}
	return o.MediaType
}

func (o *PersonTaggedImages200ApplicationJSONResults) GetVoteAverage() *float64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *PersonTaggedImages200ApplicationJSONResults) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

func (o *PersonTaggedImages200ApplicationJSONResults) GetWidth() *int64 {
	if o == nil {
		return nil
	}
	return o.Width
}

// PersonTaggedImages200ApplicationJSON - 200
type PersonTaggedImages200ApplicationJSON struct {
	ID           *int64                                        `default:"0" json:"id"`
	Page         *int64                                        `default:"0" json:"page"`
	Results      []PersonTaggedImages200ApplicationJSONResults `json:"results,omitempty"`
	TotalPages   *int64                                        `default:"0" json:"total_pages"`
	TotalResults *int64                                        `default:"0" json:"total_results"`
}

func (p PersonTaggedImages200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonTaggedImages200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PersonTaggedImages200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PersonTaggedImages200ApplicationJSON) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *PersonTaggedImages200ApplicationJSON) GetResults() []PersonTaggedImages200ApplicationJSONResults {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *PersonTaggedImages200ApplicationJSON) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

func (o *PersonTaggedImages200ApplicationJSON) GetTotalResults() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalResults
}

type PersonTaggedImagesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	PersonTaggedImages200ApplicationJSONObject *PersonTaggedImages200ApplicationJSON
}

func (o *PersonTaggedImagesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PersonTaggedImagesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PersonTaggedImagesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PersonTaggedImagesResponse) GetPersonTaggedImages200ApplicationJSONObject() *PersonTaggedImages200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PersonTaggedImages200ApplicationJSONObject
}
