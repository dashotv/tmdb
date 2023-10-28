// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvSeriesAccountStatesRequest struct {
	SeriesID       int     `pathParam:"style=simple,explode=false,name=series_id"`
	GuestSessionID *string `queryParam:"style=form,explode=true,name=guest_session_id"`
	SessionID      *string `queryParam:"style=form,explode=true,name=session_id"`
}

func (o *TvSeriesAccountStatesRequest) GetSeriesID() int {
	if o == nil {
		return 0
	}
	return o.SeriesID
}

func (o *TvSeriesAccountStatesRequest) GetGuestSessionID() *string {
	if o == nil {
		return nil
	}
	return o.GuestSessionID
}

func (o *TvSeriesAccountStatesRequest) GetSessionID() *string {
	if o == nil {
		return nil
	}
	return o.SessionID
}

type TvSeriesAccountStates200ApplicationJSONRated struct {
	Value *int64 `default:"0" json:"value"`
}

func (t TvSeriesAccountStates200ApplicationJSONRated) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesAccountStates200ApplicationJSONRated) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesAccountStates200ApplicationJSONRated) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

// TvSeriesAccountStates200ApplicationJSON - 200
type TvSeriesAccountStates200ApplicationJSON struct {
	Favorite  *bool                                         `default:"true" json:"favorite"`
	ID        *int64                                        `default:"0" json:"id"`
	Rated     *TvSeriesAccountStates200ApplicationJSONRated `json:"rated,omitempty"`
	Watchlist *bool                                         `default:"true" json:"watchlist"`
}

func (t TvSeriesAccountStates200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesAccountStates200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesAccountStates200ApplicationJSON) GetFavorite() *bool {
	if o == nil {
		return nil
	}
	return o.Favorite
}

func (o *TvSeriesAccountStates200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeriesAccountStates200ApplicationJSON) GetRated() *TvSeriesAccountStates200ApplicationJSONRated {
	if o == nil {
		return nil
	}
	return o.Rated
}

func (o *TvSeriesAccountStates200ApplicationJSON) GetWatchlist() *bool {
	if o == nil {
		return nil
	}
	return o.Watchlist
}

type TvSeriesAccountStatesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvSeriesAccountStates200ApplicationJSONObject *TvSeriesAccountStates200ApplicationJSON
}

func (o *TvSeriesAccountStatesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvSeriesAccountStatesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvSeriesAccountStatesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvSeriesAccountStatesResponse) GetTvSeriesAccountStates200ApplicationJSONObject() *TvSeriesAccountStates200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvSeriesAccountStates200ApplicationJSONObject
}
