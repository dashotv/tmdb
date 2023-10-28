// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvEpisodeAddRatingRequestBody struct {
	RawBody string `json:"RAW_BODY"`
}

func (o *TvEpisodeAddRatingRequestBody) GetRawBody() string {
	if o == nil {
		return ""
	}
	return o.RawBody
}

type TvEpisodeAddRatingRequest struct {
	RequestBody    *TvEpisodeAddRatingRequestBody `request:"mediaType=application/json"`
	EpisodeNumber  int                            `pathParam:"style=simple,explode=false,name=episode_number"`
	GuestSessionID *string                        `queryParam:"style=form,explode=true,name=guest_session_id"`
	SeasonNumber   int                            `pathParam:"style=simple,explode=false,name=season_number"`
	SeriesID       int                            `pathParam:"style=simple,explode=false,name=series_id"`
	SessionID      *string                        `queryParam:"style=form,explode=true,name=session_id"`
}

func (o *TvEpisodeAddRatingRequest) GetRequestBody() *TvEpisodeAddRatingRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *TvEpisodeAddRatingRequest) GetEpisodeNumber() int {
	if o == nil {
		return 0
	}
	return o.EpisodeNumber
}

func (o *TvEpisodeAddRatingRequest) GetGuestSessionID() *string {
	if o == nil {
		return nil
	}
	return o.GuestSessionID
}

func (o *TvEpisodeAddRatingRequest) GetSeasonNumber() int {
	if o == nil {
		return 0
	}
	return o.SeasonNumber
}

func (o *TvEpisodeAddRatingRequest) GetSeriesID() int {
	if o == nil {
		return 0
	}
	return o.SeriesID
}

func (o *TvEpisodeAddRatingRequest) GetSessionID() *string {
	if o == nil {
		return nil
	}
	return o.SessionID
}

// TvEpisodeAddRating200ApplicationJSON - 200
type TvEpisodeAddRating200ApplicationJSON struct {
	StatusCode    *int64  `default:"0" json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
}

func (t TvEpisodeAddRating200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvEpisodeAddRating200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvEpisodeAddRating200ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

func (o *TvEpisodeAddRating200ApplicationJSON) GetStatusMessage() *string {
	if o == nil {
		return nil
	}
	return o.StatusMessage
}

type TvEpisodeAddRatingResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvEpisodeAddRating200ApplicationJSONObject *TvEpisodeAddRating200ApplicationJSON
}

func (o *TvEpisodeAddRatingResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvEpisodeAddRatingResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvEpisodeAddRatingResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvEpisodeAddRatingResponse) GetTvEpisodeAddRating200ApplicationJSONObject() *TvEpisodeAddRating200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvEpisodeAddRating200ApplicationJSONObject
}
