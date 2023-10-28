// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type CollectionDetailsRequest struct {
	CollectionID int     `pathParam:"style=simple,explode=false,name=collection_id"`
	Language     *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
}

func (c CollectionDetailsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CollectionDetailsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CollectionDetailsRequest) GetCollectionID() int {
	if o == nil {
		return 0
	}
	return o.CollectionID
}

func (o *CollectionDetailsRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

type CollectionDetails200ApplicationJSONParts struct {
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

func (c CollectionDetails200ApplicationJSONParts) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CollectionDetails200ApplicationJSONParts) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CollectionDetails200ApplicationJSONParts) GetAdult() *bool {
	if o == nil {
		return nil
	}
	return o.Adult
}

func (o *CollectionDetails200ApplicationJSONParts) GetBackdropPath() *string {
	if o == nil {
		return nil
	}
	return o.BackdropPath
}

func (o *CollectionDetails200ApplicationJSONParts) GetGenreIds() []int64 {
	if o == nil {
		return nil
	}
	return o.GenreIds
}

func (o *CollectionDetails200ApplicationJSONParts) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CollectionDetails200ApplicationJSONParts) GetMediaType() *string {
	if o == nil {
		return nil
	}
	return o.MediaType
}

func (o *CollectionDetails200ApplicationJSONParts) GetOriginalLanguage() *string {
	if o == nil {
		return nil
	}
	return o.OriginalLanguage
}

func (o *CollectionDetails200ApplicationJSONParts) GetOriginalTitle() *string {
	if o == nil {
		return nil
	}
	return o.OriginalTitle
}

func (o *CollectionDetails200ApplicationJSONParts) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *CollectionDetails200ApplicationJSONParts) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *CollectionDetails200ApplicationJSONParts) GetPosterPath() *string {
	if o == nil {
		return nil
	}
	return o.PosterPath
}

func (o *CollectionDetails200ApplicationJSONParts) GetReleaseDate() *string {
	if o == nil {
		return nil
	}
	return o.ReleaseDate
}

func (o *CollectionDetails200ApplicationJSONParts) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *CollectionDetails200ApplicationJSONParts) GetVideo() *bool {
	if o == nil {
		return nil
	}
	return o.Video
}

func (o *CollectionDetails200ApplicationJSONParts) GetVoteAverage() *float64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *CollectionDetails200ApplicationJSONParts) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

// CollectionDetails200ApplicationJSON - 200
type CollectionDetails200ApplicationJSON struct {
	BackdropPath *string                                    `json:"backdrop_path,omitempty"`
	ID           *int64                                     `default:"0" json:"id"`
	Name         *string                                    `json:"name,omitempty"`
	Overview     *string                                    `json:"overview,omitempty"`
	Parts        []CollectionDetails200ApplicationJSONParts `json:"parts,omitempty"`
	PosterPath   *string                                    `json:"poster_path,omitempty"`
}

func (c CollectionDetails200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CollectionDetails200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CollectionDetails200ApplicationJSON) GetBackdropPath() *string {
	if o == nil {
		return nil
	}
	return o.BackdropPath
}

func (o *CollectionDetails200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CollectionDetails200ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CollectionDetails200ApplicationJSON) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *CollectionDetails200ApplicationJSON) GetParts() []CollectionDetails200ApplicationJSONParts {
	if o == nil {
		return nil
	}
	return o.Parts
}

func (o *CollectionDetails200ApplicationJSON) GetPosterPath() *string {
	if o == nil {
		return nil
	}
	return o.PosterPath
}

type CollectionDetailsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	CollectionDetails200ApplicationJSONObject *CollectionDetails200ApplicationJSON
}

func (o *CollectionDetailsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CollectionDetailsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CollectionDetailsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CollectionDetailsResponse) GetCollectionDetails200ApplicationJSONObject() *CollectionDetails200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CollectionDetails200ApplicationJSONObject
}