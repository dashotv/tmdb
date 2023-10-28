// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type MovieTranslationsRequest struct {
	MovieID int `pathParam:"style=simple,explode=false,name=movie_id"`
}

func (o *MovieTranslationsRequest) GetMovieID() int {
	if o == nil {
		return 0
	}
	return o.MovieID
}

type MovieTranslations200ApplicationJSONTranslationsData struct {
	Homepage *string `json:"homepage,omitempty"`
	Overview *string `json:"overview,omitempty"`
	Runtime  *int64  `default:"0" json:"runtime"`
	Tagline  *string `json:"tagline,omitempty"`
	Title    *string `json:"title,omitempty"`
}

func (m MovieTranslations200ApplicationJSONTranslationsData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MovieTranslations200ApplicationJSONTranslationsData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MovieTranslations200ApplicationJSONTranslationsData) GetHomepage() *string {
	if o == nil {
		return nil
	}
	return o.Homepage
}

func (o *MovieTranslations200ApplicationJSONTranslationsData) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *MovieTranslations200ApplicationJSONTranslationsData) GetRuntime() *int64 {
	if o == nil {
		return nil
	}
	return o.Runtime
}

func (o *MovieTranslations200ApplicationJSONTranslationsData) GetTagline() *string {
	if o == nil {
		return nil
	}
	return o.Tagline
}

func (o *MovieTranslations200ApplicationJSONTranslationsData) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

type MovieTranslations200ApplicationJSONTranslations struct {
	Data        *MovieTranslations200ApplicationJSONTranslationsData `json:"data,omitempty"`
	EnglishName *string                                              `json:"english_name,omitempty"`
	Iso31661    *string                                              `json:"iso_3166_1,omitempty"`
	Iso6391     *string                                              `json:"iso_639_1,omitempty"`
	Name        *string                                              `json:"name,omitempty"`
}

func (o *MovieTranslations200ApplicationJSONTranslations) GetData() *MovieTranslations200ApplicationJSONTranslationsData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *MovieTranslations200ApplicationJSONTranslations) GetEnglishName() *string {
	if o == nil {
		return nil
	}
	return o.EnglishName
}

func (o *MovieTranslations200ApplicationJSONTranslations) GetIso31661() *string {
	if o == nil {
		return nil
	}
	return o.Iso31661
}

func (o *MovieTranslations200ApplicationJSONTranslations) GetIso6391() *string {
	if o == nil {
		return nil
	}
	return o.Iso6391
}

func (o *MovieTranslations200ApplicationJSONTranslations) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// MovieTranslations200ApplicationJSON - 200
type MovieTranslations200ApplicationJSON struct {
	ID           *int64                                            `default:"0" json:"id"`
	Translations []MovieTranslations200ApplicationJSONTranslations `json:"translations,omitempty"`
}

func (m MovieTranslations200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MovieTranslations200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MovieTranslations200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MovieTranslations200ApplicationJSON) GetTranslations() []MovieTranslations200ApplicationJSONTranslations {
	if o == nil {
		return nil
	}
	return o.Translations
}

type MovieTranslationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	MovieTranslations200ApplicationJSONObject *MovieTranslations200ApplicationJSON
}

func (o *MovieTranslationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MovieTranslationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MovieTranslationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MovieTranslationsResponse) GetMovieTranslations200ApplicationJSONObject() *MovieTranslations200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.MovieTranslations200ApplicationJSONObject
}