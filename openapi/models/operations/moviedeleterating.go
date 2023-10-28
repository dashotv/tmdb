// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type MovieDeleteRatingRequest struct {
	MovieID        int     `pathParam:"style=simple,explode=false,name=movie_id"`
	GuestSessionID *string `queryParam:"style=form,explode=true,name=guest_session_id"`
	SessionID      *string `queryParam:"style=form,explode=true,name=session_id"`
}

func (o *MovieDeleteRatingRequest) GetMovieID() int {
	if o == nil {
		return 0
	}
	return o.MovieID
}

func (o *MovieDeleteRatingRequest) GetGuestSessionID() *string {
	if o == nil {
		return nil
	}
	return o.GuestSessionID
}

func (o *MovieDeleteRatingRequest) GetSessionID() *string {
	if o == nil {
		return nil
	}
	return o.SessionID
}

// MovieDeleteRating200ApplicationJSON - 200
type MovieDeleteRating200ApplicationJSON struct {
	StatusCode    *int64  `default:"0" json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
}

func (m MovieDeleteRating200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MovieDeleteRating200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MovieDeleteRating200ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

func (o *MovieDeleteRating200ApplicationJSON) GetStatusMessage() *string {
	if o == nil {
		return nil
	}
	return o.StatusMessage
}

type MovieDeleteRatingResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	MovieDeleteRating200ApplicationJSONObject *MovieDeleteRating200ApplicationJSON
}

func (o *MovieDeleteRatingResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MovieDeleteRatingResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MovieDeleteRatingResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MovieDeleteRatingResponse) GetMovieDeleteRating200ApplicationJSONObject() *MovieDeleteRating200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.MovieDeleteRating200ApplicationJSONObject
}
