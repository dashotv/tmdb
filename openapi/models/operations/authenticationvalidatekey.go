// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

// AuthenticationValidateKey401ApplicationJSON - 401
type AuthenticationValidateKey401ApplicationJSON struct {
	StatusCode    *int64  `default:"0" json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
	Success       *bool   `default:"true" json:"success"`
}

func (a AuthenticationValidateKey401ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthenticationValidateKey401ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AuthenticationValidateKey401ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

func (o *AuthenticationValidateKey401ApplicationJSON) GetStatusMessage() *string {
	if o == nil {
		return nil
	}
	return o.StatusMessage
}

func (o *AuthenticationValidateKey401ApplicationJSON) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}

// AuthenticationValidateKey200ApplicationJSON - 200
type AuthenticationValidateKey200ApplicationJSON struct {
	StatusCode    *int64  `default:"0" json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
	Success       *bool   `default:"true" json:"success"`
}

func (a AuthenticationValidateKey200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthenticationValidateKey200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AuthenticationValidateKey200ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

func (o *AuthenticationValidateKey200ApplicationJSON) GetStatusMessage() *string {
	if o == nil {
		return nil
	}
	return o.StatusMessage
}

func (o *AuthenticationValidateKey200ApplicationJSON) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}

type AuthenticationValidateKeyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	AuthenticationValidateKey200ApplicationJSONObject *AuthenticationValidateKey200ApplicationJSON
	// 401
	AuthenticationValidateKey401ApplicationJSONObject *AuthenticationValidateKey401ApplicationJSON
}

func (o *AuthenticationValidateKeyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AuthenticationValidateKeyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AuthenticationValidateKeyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AuthenticationValidateKeyResponse) GetAuthenticationValidateKey200ApplicationJSONObject() *AuthenticationValidateKey200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.AuthenticationValidateKey200ApplicationJSONObject
}

func (o *AuthenticationValidateKeyResponse) GetAuthenticationValidateKey401ApplicationJSONObject() *AuthenticationValidateKey401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.AuthenticationValidateKey401ApplicationJSONObject
}
