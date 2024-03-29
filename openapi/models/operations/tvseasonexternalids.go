// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvSeasonExternalIdsRequest struct {
	SeasonNumber int `pathParam:"style=simple,explode=false,name=season_number"`
	SeriesID     int `pathParam:"style=simple,explode=false,name=series_id"`
}

func (o *TvSeasonExternalIdsRequest) GetSeasonNumber() int {
	if o == nil {
		return 0
	}
	return o.SeasonNumber
}

func (o *TvSeasonExternalIdsRequest) GetSeriesID() int {
	if o == nil {
		return 0
	}
	return o.SeriesID
}

// TvSeasonExternalIds200ApplicationJSON - 200
type TvSeasonExternalIds200ApplicationJSON struct {
	FreebaseID  *string     `json:"freebase_id,omitempty"`
	FreebaseMid *string     `json:"freebase_mid,omitempty"`
	ID          *int64      `default:"0" json:"id"`
	TvdbID      *int64      `default:"0" json:"tvdb_id"`
	TvrageID    interface{} `json:"tvrage_id,omitempty"`
	WikidataID  *string     `json:"wikidata_id,omitempty"`
}

func (t TvSeasonExternalIds200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeasonExternalIds200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeasonExternalIds200ApplicationJSON) GetFreebaseID() *string {
	if o == nil {
		return nil
	}
	return o.FreebaseID
}

func (o *TvSeasonExternalIds200ApplicationJSON) GetFreebaseMid() *string {
	if o == nil {
		return nil
	}
	return o.FreebaseMid
}

func (o *TvSeasonExternalIds200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeasonExternalIds200ApplicationJSON) GetTvdbID() *int64 {
	if o == nil {
		return nil
	}
	return o.TvdbID
}

func (o *TvSeasonExternalIds200ApplicationJSON) GetTvrageID() interface{} {
	if o == nil {
		return nil
	}
	return o.TvrageID
}

func (o *TvSeasonExternalIds200ApplicationJSON) GetWikidataID() *string {
	if o == nil {
		return nil
	}
	return o.WikidataID
}

type TvSeasonExternalIdsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvSeasonExternalIds200ApplicationJSONObject *TvSeasonExternalIds200ApplicationJSON
}

func (o *TvSeasonExternalIdsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvSeasonExternalIdsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvSeasonExternalIdsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvSeasonExternalIdsResponse) GetTvSeasonExternalIds200ApplicationJSONObject() *TvSeasonExternalIds200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvSeasonExternalIds200ApplicationJSONObject
}
