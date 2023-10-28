// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type ListClearRequest struct {
	Confirm   bool   `default:"false" queryParam:"style=form,explode=true,name=confirm"`
	ListID    int    `pathParam:"style=simple,explode=false,name=list_id"`
	SessionID string `queryParam:"style=form,explode=true,name=session_id"`
}

func (l ListClearRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListClearRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListClearRequest) GetConfirm() bool {
	if o == nil {
		return false
	}
	return o.Confirm
}

func (o *ListClearRequest) GetListID() int {
	if o == nil {
		return 0
	}
	return o.ListID
}

func (o *ListClearRequest) GetSessionID() string {
	if o == nil {
		return ""
	}
	return o.SessionID
}

// ListClear200ApplicationJSON - 200
type ListClear200ApplicationJSON struct {
	StatusCode    *int64  `default:"0" json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
}

func (l ListClear200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListClear200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListClear200ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

func (o *ListClear200ApplicationJSON) GetStatusMessage() *string {
	if o == nil {
		return nil
	}
	return o.StatusMessage
}

type ListClearResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	ListClear200ApplicationJSONObject *ListClear200ApplicationJSON
}

func (o *ListClearResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListClearResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListClearResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListClearResponse) GetListClear200ApplicationJSONObject() *ListClear200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListClear200ApplicationJSONObject
}
