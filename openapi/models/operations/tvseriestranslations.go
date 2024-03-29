// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvSeriesTranslationsRequest struct {
	SeriesID int `pathParam:"style=simple,explode=false,name=series_id"`
}

func (o *TvSeriesTranslationsRequest) GetSeriesID() int {
	if o == nil {
		return 0
	}
	return o.SeriesID
}

type TvSeriesTranslations200ApplicationJSONTranslationsData struct {
	Homepage *string `json:"homepage,omitempty"`
	Name     *string `json:"name,omitempty"`
	Overview *string `json:"overview,omitempty"`
	Tagline  *string `json:"tagline,omitempty"`
}

func (o *TvSeriesTranslations200ApplicationJSONTranslationsData) GetHomepage() *string {
	if o == nil {
		return nil
	}
	return o.Homepage
}

func (o *TvSeriesTranslations200ApplicationJSONTranslationsData) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TvSeriesTranslations200ApplicationJSONTranslationsData) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *TvSeriesTranslations200ApplicationJSONTranslationsData) GetTagline() *string {
	if o == nil {
		return nil
	}
	return o.Tagline
}

type TvSeriesTranslations200ApplicationJSONTranslations struct {
	Data        *TvSeriesTranslations200ApplicationJSONTranslationsData `json:"data,omitempty"`
	EnglishName *string                                                 `json:"english_name,omitempty"`
	Iso31661    *string                                                 `json:"iso_3166_1,omitempty"`
	Iso6391     *string                                                 `json:"iso_639_1,omitempty"`
	Name        *string                                                 `json:"name,omitempty"`
}

func (o *TvSeriesTranslations200ApplicationJSONTranslations) GetData() *TvSeriesTranslations200ApplicationJSONTranslationsData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *TvSeriesTranslations200ApplicationJSONTranslations) GetEnglishName() *string {
	if o == nil {
		return nil
	}
	return o.EnglishName
}

func (o *TvSeriesTranslations200ApplicationJSONTranslations) GetIso31661() *string {
	if o == nil {
		return nil
	}
	return o.Iso31661
}

func (o *TvSeriesTranslations200ApplicationJSONTranslations) GetIso6391() *string {
	if o == nil {
		return nil
	}
	return o.Iso6391
}

func (o *TvSeriesTranslations200ApplicationJSONTranslations) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// TvSeriesTranslations200ApplicationJSON - 200
type TvSeriesTranslations200ApplicationJSON struct {
	ID           *int64                                               `default:"0" json:"id"`
	Translations []TvSeriesTranslations200ApplicationJSONTranslations `json:"translations,omitempty"`
}

func (t TvSeriesTranslations200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesTranslations200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesTranslations200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeriesTranslations200ApplicationJSON) GetTranslations() []TvSeriesTranslations200ApplicationJSONTranslations {
	if o == nil {
		return nil
	}
	return o.Translations
}

type TvSeriesTranslationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvSeriesTranslations200ApplicationJSONObject *TvSeriesTranslations200ApplicationJSON
}

func (o *TvSeriesTranslationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvSeriesTranslationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvSeriesTranslationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvSeriesTranslationsResponse) GetTvSeriesTranslations200ApplicationJSONObject() *TvSeriesTranslations200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvSeriesTranslations200ApplicationJSONObject
}
