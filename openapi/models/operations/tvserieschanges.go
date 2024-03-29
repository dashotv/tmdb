// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvSeriesChangesRequest struct {
	SeriesID  int     `pathParam:"style=simple,explode=false,name=series_id"`
	EndDate   *string `queryParam:"style=form,explode=true,name=end_date"`
	Page      *int    `default:"1" queryParam:"style=form,explode=true,name=page"`
	StartDate *string `queryParam:"style=form,explode=true,name=start_date"`
}

func (t TvSeriesChangesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesChangesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesChangesRequest) GetSeriesID() int {
	if o == nil {
		return 0
	}
	return o.SeriesID
}

func (o *TvSeriesChangesRequest) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *TvSeriesChangesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *TvSeriesChangesRequest) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

type TvSeriesChanges200ApplicationJSONChangesItemsOriginalValuePoster struct {
	FilePath *string `json:"file_path,omitempty"`
	Iso6391  *string `json:"iso_639_1,omitempty"`
}

func (o *TvSeriesChanges200ApplicationJSONChangesItemsOriginalValuePoster) GetFilePath() *string {
	if o == nil {
		return nil
	}
	return o.FilePath
}

func (o *TvSeriesChanges200ApplicationJSONChangesItemsOriginalValuePoster) GetIso6391() *string {
	if o == nil {
		return nil
	}
	return o.Iso6391
}

type TvSeriesChanges200ApplicationJSONChangesItemsOriginalValue struct {
	Poster *TvSeriesChanges200ApplicationJSONChangesItemsOriginalValuePoster `json:"poster,omitempty"`
}

func (o *TvSeriesChanges200ApplicationJSONChangesItemsOriginalValue) GetPoster() *TvSeriesChanges200ApplicationJSONChangesItemsOriginalValuePoster {
	if o == nil {
		return nil
	}
	return o.Poster
}

type TvSeriesChanges200ApplicationJSONChangesItemsValuePoster struct {
	FilePath *string `json:"file_path,omitempty"`
	Iso6391  *string `json:"iso_639_1,omitempty"`
}

func (o *TvSeriesChanges200ApplicationJSONChangesItemsValuePoster) GetFilePath() *string {
	if o == nil {
		return nil
	}
	return o.FilePath
}

func (o *TvSeriesChanges200ApplicationJSONChangesItemsValuePoster) GetIso6391() *string {
	if o == nil {
		return nil
	}
	return o.Iso6391
}

type TvSeriesChanges200ApplicationJSONChangesItemsValue struct {
	Poster *TvSeriesChanges200ApplicationJSONChangesItemsValuePoster `json:"poster,omitempty"`
}

func (o *TvSeriesChanges200ApplicationJSONChangesItemsValue) GetPoster() *TvSeriesChanges200ApplicationJSONChangesItemsValuePoster {
	if o == nil {
		return nil
	}
	return o.Poster
}

type TvSeriesChanges200ApplicationJSONChangesItems struct {
	Action        *string                                                     `json:"action,omitempty"`
	ID            *string                                                     `json:"id,omitempty"`
	Iso31661      *string                                                     `json:"iso_3166_1,omitempty"`
	Iso6391       *string                                                     `json:"iso_639_1,omitempty"`
	OriginalValue *TvSeriesChanges200ApplicationJSONChangesItemsOriginalValue `json:"original_value,omitempty"`
	Time          *string                                                     `json:"time,omitempty"`
	Value         *TvSeriesChanges200ApplicationJSONChangesItemsValue         `json:"value,omitempty"`
}

func (o *TvSeriesChanges200ApplicationJSONChangesItems) GetAction() *string {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *TvSeriesChanges200ApplicationJSONChangesItems) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeriesChanges200ApplicationJSONChangesItems) GetIso31661() *string {
	if o == nil {
		return nil
	}
	return o.Iso31661
}

func (o *TvSeriesChanges200ApplicationJSONChangesItems) GetIso6391() *string {
	if o == nil {
		return nil
	}
	return o.Iso6391
}

func (o *TvSeriesChanges200ApplicationJSONChangesItems) GetOriginalValue() *TvSeriesChanges200ApplicationJSONChangesItemsOriginalValue {
	if o == nil {
		return nil
	}
	return o.OriginalValue
}

func (o *TvSeriesChanges200ApplicationJSONChangesItems) GetTime() *string {
	if o == nil {
		return nil
	}
	return o.Time
}

func (o *TvSeriesChanges200ApplicationJSONChangesItems) GetValue() *TvSeriesChanges200ApplicationJSONChangesItemsValue {
	if o == nil {
		return nil
	}
	return o.Value
}

type TvSeriesChanges200ApplicationJSONChanges struct {
	Items []TvSeriesChanges200ApplicationJSONChangesItems `json:"items,omitempty"`
	Key   *string                                         `json:"key,omitempty"`
}

func (o *TvSeriesChanges200ApplicationJSONChanges) GetItems() []TvSeriesChanges200ApplicationJSONChangesItems {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *TvSeriesChanges200ApplicationJSONChanges) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

// TvSeriesChanges200ApplicationJSON - 200
type TvSeriesChanges200ApplicationJSON struct {
	Changes []TvSeriesChanges200ApplicationJSONChanges `json:"changes,omitempty"`
}

func (o *TvSeriesChanges200ApplicationJSON) GetChanges() []TvSeriesChanges200ApplicationJSONChanges {
	if o == nil {
		return nil
	}
	return o.Changes
}

type TvSeriesChangesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvSeriesChanges200ApplicationJSONObject *TvSeriesChanges200ApplicationJSON
}

func (o *TvSeriesChangesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvSeriesChangesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvSeriesChangesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvSeriesChangesResponse) GetTvSeriesChanges200ApplicationJSONObject() *TvSeriesChanges200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvSeriesChanges200ApplicationJSONObject
}
