// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

// AuthenticationCreateGuestSession200ApplicationJSON - 200
type AuthenticationCreateGuestSession200ApplicationJSON struct {
	ExpiresAt      *string `json:"expires_at,omitempty"`
	GuestSessionID *string `json:"guest_session_id,omitempty"`
	Success        *bool   `default:"true" json:"success"`
}

func (a AuthenticationCreateGuestSession200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthenticationCreateGuestSession200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AuthenticationCreateGuestSession200ApplicationJSON) GetExpiresAt() *string {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *AuthenticationCreateGuestSession200ApplicationJSON) GetGuestSessionID() *string {
	if o == nil {
		return nil
	}
	return o.GuestSessionID
}

func (o *AuthenticationCreateGuestSession200ApplicationJSON) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}

type AuthenticationCreateGuestSessionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	AuthenticationCreateGuestSession200ApplicationJSONObject *AuthenticationCreateGuestSession200ApplicationJSON
}

func (o *AuthenticationCreateGuestSessionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AuthenticationCreateGuestSessionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AuthenticationCreateGuestSessionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AuthenticationCreateGuestSessionResponse) GetAuthenticationCreateGuestSession200ApplicationJSONObject() *AuthenticationCreateGuestSession200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.AuthenticationCreateGuestSession200ApplicationJSONObject
}
