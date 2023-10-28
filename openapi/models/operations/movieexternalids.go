// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type MovieExternalIdsRequest struct {
	MovieID int `pathParam:"style=simple,explode=false,name=movie_id"`
}

func (o *MovieExternalIdsRequest) GetMovieID() int {
	if o == nil {
		return 0
	}
	return o.MovieID
}

// MovieExternalIds200ApplicationJSON - 200
type MovieExternalIds200ApplicationJSON struct {
	FacebookID  *string     `json:"facebook_id,omitempty"`
	ID          *int64      `default:"0" json:"id"`
	ImdbID      *string     `json:"imdb_id,omitempty"`
	InstagramID interface{} `json:"instagram_id,omitempty"`
	TwitterID   interface{} `json:"twitter_id,omitempty"`
	WikidataID  interface{} `json:"wikidata_id,omitempty"`
}

func (m MovieExternalIds200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MovieExternalIds200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MovieExternalIds200ApplicationJSON) GetFacebookID() *string {
	if o == nil {
		return nil
	}
	return o.FacebookID
}

func (o *MovieExternalIds200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MovieExternalIds200ApplicationJSON) GetImdbID() *string {
	if o == nil {
		return nil
	}
	return o.ImdbID
}

func (o *MovieExternalIds200ApplicationJSON) GetInstagramID() interface{} {
	if o == nil {
		return nil
	}
	return o.InstagramID
}

func (o *MovieExternalIds200ApplicationJSON) GetTwitterID() interface{} {
	if o == nil {
		return nil
	}
	return o.TwitterID
}

func (o *MovieExternalIds200ApplicationJSON) GetWikidataID() interface{} {
	if o == nil {
		return nil
	}
	return o.WikidataID
}

type MovieExternalIdsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	MovieExternalIds200ApplicationJSONObject *MovieExternalIds200ApplicationJSON
}

func (o *MovieExternalIdsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MovieExternalIdsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MovieExternalIdsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MovieExternalIdsResponse) GetMovieExternalIds200ApplicationJSONObject() *MovieExternalIds200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.MovieExternalIds200ApplicationJSONObject
}