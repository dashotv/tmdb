// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type AccountAddToWatchlistRequestBody struct {
	RawBody string `json:"RAW_BODY"`
}

func (o *AccountAddToWatchlistRequestBody) GetRawBody() string {
	if o == nil {
		return ""
	}
	return o.RawBody
}

type AccountAddToWatchlistRequest struct {
	AccountID   int                               `default:"null" pathParam:"style=simple,explode=false,name=account_id"`
	RequestBody *AccountAddToWatchlistRequestBody `request:"mediaType=application/json"`
	SessionID   *string                           `queryParam:"style=form,explode=true,name=session_id"`
}

func (a AccountAddToWatchlistRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountAddToWatchlistRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountAddToWatchlistRequest) GetAccountID() int {
	if o == nil {
		return 0
	}
	return o.AccountID
}

func (o *AccountAddToWatchlistRequest) GetRequestBody() *AccountAddToWatchlistRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *AccountAddToWatchlistRequest) GetSessionID() *string {
	if o == nil {
		return nil
	}
	return o.SessionID
}

// AccountAddToWatchlist200ApplicationJSON - 200
type AccountAddToWatchlist200ApplicationJSON struct {
	StatusCode    *int64  `default:"0" json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
}

func (a AccountAddToWatchlist200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountAddToWatchlist200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountAddToWatchlist200ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

func (o *AccountAddToWatchlist200ApplicationJSON) GetStatusMessage() *string {
	if o == nil {
		return nil
	}
	return o.StatusMessage
}

type AccountAddToWatchlistResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	AccountAddToWatchlist200ApplicationJSONObject *AccountAddToWatchlist200ApplicationJSON
}

func (o *AccountAddToWatchlistResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AccountAddToWatchlistResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountAddToWatchlistResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AccountAddToWatchlistResponse) GetAccountAddToWatchlist200ApplicationJSONObject() *AccountAddToWatchlist200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.AccountAddToWatchlist200ApplicationJSONObject
}
