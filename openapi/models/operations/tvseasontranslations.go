// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvSeasonTranslationsRequest struct {
	SeasonNumber int `pathParam:"style=simple,explode=false,name=season_number"`
	SeriesID     int `pathParam:"style=simple,explode=false,name=series_id"`
}

func (o *TvSeasonTranslationsRequest) GetSeasonNumber() int {
	if o == nil {
		return 0
	}
	return o.SeasonNumber
}

func (o *TvSeasonTranslationsRequest) GetSeriesID() int {
	if o == nil {
		return 0
	}
	return o.SeriesID
}

type TvSeasonTranslations200ApplicationJSONTranslationsData struct {
	Name     *string `json:"name,omitempty"`
	Overview *string `json:"overview,omitempty"`
}

func (o *TvSeasonTranslations200ApplicationJSONTranslationsData) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TvSeasonTranslations200ApplicationJSONTranslationsData) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

type TvSeasonTranslations200ApplicationJSONTranslations struct {
	Data        *TvSeasonTranslations200ApplicationJSONTranslationsData `json:"data,omitempty"`
	EnglishName *string                                                 `json:"english_name,omitempty"`
	Iso31661    *string                                                 `json:"iso_3166_1,omitempty"`
	Iso6391     *string                                                 `json:"iso_639_1,omitempty"`
	Name        *string                                                 `json:"name,omitempty"`
}

func (o *TvSeasonTranslations200ApplicationJSONTranslations) GetData() *TvSeasonTranslations200ApplicationJSONTranslationsData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *TvSeasonTranslations200ApplicationJSONTranslations) GetEnglishName() *string {
	if o == nil {
		return nil
	}
	return o.EnglishName
}

func (o *TvSeasonTranslations200ApplicationJSONTranslations) GetIso31661() *string {
	if o == nil {
		return nil
	}
	return o.Iso31661
}

func (o *TvSeasonTranslations200ApplicationJSONTranslations) GetIso6391() *string {
	if o == nil {
		return nil
	}
	return o.Iso6391
}

func (o *TvSeasonTranslations200ApplicationJSONTranslations) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// TvSeasonTranslations200ApplicationJSON - 200
type TvSeasonTranslations200ApplicationJSON struct {
	ID           *int64                                               `default:"0" json:"id"`
	Translations []TvSeasonTranslations200ApplicationJSONTranslations `json:"translations,omitempty"`
}

func (t TvSeasonTranslations200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeasonTranslations200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeasonTranslations200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeasonTranslations200ApplicationJSON) GetTranslations() []TvSeasonTranslations200ApplicationJSONTranslations {
	if o == nil {
		return nil
	}
	return o.Translations
}

type TvSeasonTranslationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvSeasonTranslations200ApplicationJSONObject *TvSeasonTranslations200ApplicationJSON
}

func (o *TvSeasonTranslationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvSeasonTranslationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvSeasonTranslationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvSeasonTranslationsResponse) GetTvSeasonTranslations200ApplicationJSONObject() *TvSeasonTranslations200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvSeasonTranslations200ApplicationJSONObject
}
