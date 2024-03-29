// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvSeriesDeleteRatingRequest struct {
	SeriesID       int     `pathParam:"style=simple,explode=false,name=series_id"`
	GuestSessionID *string `queryParam:"style=form,explode=true,name=guest_session_id"`
	SessionID      *string `queryParam:"style=form,explode=true,name=session_id"`
}

func (o *TvSeriesDeleteRatingRequest) GetSeriesID() int {
	if o == nil {
		return 0
	}
	return o.SeriesID
}

func (o *TvSeriesDeleteRatingRequest) GetGuestSessionID() *string {
	if o == nil {
		return nil
	}
	return o.GuestSessionID
}

func (o *TvSeriesDeleteRatingRequest) GetSessionID() *string {
	if o == nil {
		return nil
	}
	return o.SessionID
}

// TvSeriesDeleteRating200ApplicationJSON - 200
type TvSeriesDeleteRating200ApplicationJSON struct {
	StatusCode    *int64  `default:"0" json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
}

func (t TvSeriesDeleteRating200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesDeleteRating200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesDeleteRating200ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

func (o *TvSeriesDeleteRating200ApplicationJSON) GetStatusMessage() *string {
	if o == nil {
		return nil
	}
	return o.StatusMessage
}

type TvSeriesDeleteRatingResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvSeriesDeleteRating200ApplicationJSONObject *TvSeriesDeleteRating200ApplicationJSON
}

func (o *TvSeriesDeleteRatingResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvSeriesDeleteRatingResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvSeriesDeleteRatingResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvSeriesDeleteRatingResponse) GetTvSeriesDeleteRating200ApplicationJSONObject() *TvSeriesDeleteRating200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvSeriesDeleteRating200ApplicationJSONObject
}
