// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type MovieAlternativeTitlesRequest struct {
	MovieID int `pathParam:"style=simple,explode=false,name=movie_id"`
	// specify a ISO-3166-1 value to filter the results
	Country *string `queryParam:"style=form,explode=true,name=country"`
}

func (o *MovieAlternativeTitlesRequest) GetMovieID() int {
	if o == nil {
		return 0
	}
	return o.MovieID
}

func (o *MovieAlternativeTitlesRequest) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

type MovieAlternativeTitles200ApplicationJSONTitles struct {
	Iso31661 *string `json:"iso_3166_1,omitempty"`
	Title    *string `json:"title,omitempty"`
	Type     *string `json:"type,omitempty"`
}

func (o *MovieAlternativeTitles200ApplicationJSONTitles) GetIso31661() *string {
	if o == nil {
		return nil
	}
	return o.Iso31661
}

func (o *MovieAlternativeTitles200ApplicationJSONTitles) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *MovieAlternativeTitles200ApplicationJSONTitles) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

// MovieAlternativeTitles200ApplicationJSON - 200
type MovieAlternativeTitles200ApplicationJSON struct {
	ID     *int64                                           `default:"0" json:"id"`
	Titles []MovieAlternativeTitles200ApplicationJSONTitles `json:"titles,omitempty"`
}

func (m MovieAlternativeTitles200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MovieAlternativeTitles200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MovieAlternativeTitles200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MovieAlternativeTitles200ApplicationJSON) GetTitles() []MovieAlternativeTitles200ApplicationJSONTitles {
	if o == nil {
		return nil
	}
	return o.Titles
}

type MovieAlternativeTitlesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	MovieAlternativeTitles200ApplicationJSONObject *MovieAlternativeTitles200ApplicationJSON
}

func (o *MovieAlternativeTitlesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MovieAlternativeTitlesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MovieAlternativeTitlesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MovieAlternativeTitlesResponse) GetMovieAlternativeTitles200ApplicationJSONObject() *MovieAlternativeTitles200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.MovieAlternativeTitles200ApplicationJSONObject
}
