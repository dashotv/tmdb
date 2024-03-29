// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ConfigurationDetails200ApplicationJSONImages struct {
	BackdropSizes []string `json:"backdrop_sizes,omitempty"`
	BaseURL       *string  `json:"base_url,omitempty"`
	LogoSizes     []string `json:"logo_sizes,omitempty"`
	PosterSizes   []string `json:"poster_sizes,omitempty"`
	ProfileSizes  []string `json:"profile_sizes,omitempty"`
	SecureBaseURL *string  `json:"secure_base_url,omitempty"`
	StillSizes    []string `json:"still_sizes,omitempty"`
}

func (o *ConfigurationDetails200ApplicationJSONImages) GetBackdropSizes() []string {
	if o == nil {
		return nil
	}
	return o.BackdropSizes
}

func (o *ConfigurationDetails200ApplicationJSONImages) GetBaseURL() *string {
	if o == nil {
		return nil
	}
	return o.BaseURL
}

func (o *ConfigurationDetails200ApplicationJSONImages) GetLogoSizes() []string {
	if o == nil {
		return nil
	}
	return o.LogoSizes
}

func (o *ConfigurationDetails200ApplicationJSONImages) GetPosterSizes() []string {
	if o == nil {
		return nil
	}
	return o.PosterSizes
}

func (o *ConfigurationDetails200ApplicationJSONImages) GetProfileSizes() []string {
	if o == nil {
		return nil
	}
	return o.ProfileSizes
}

func (o *ConfigurationDetails200ApplicationJSONImages) GetSecureBaseURL() *string {
	if o == nil {
		return nil
	}
	return o.SecureBaseURL
}

func (o *ConfigurationDetails200ApplicationJSONImages) GetStillSizes() []string {
	if o == nil {
		return nil
	}
	return o.StillSizes
}

// ConfigurationDetails200ApplicationJSON - 200
type ConfigurationDetails200ApplicationJSON struct {
	ChangeKeys []string                                      `json:"change_keys,omitempty"`
	Images     *ConfigurationDetails200ApplicationJSONImages `json:"images,omitempty"`
}

func (o *ConfigurationDetails200ApplicationJSON) GetChangeKeys() []string {
	if o == nil {
		return nil
	}
	return o.ChangeKeys
}

func (o *ConfigurationDetails200ApplicationJSON) GetImages() *ConfigurationDetails200ApplicationJSONImages {
	if o == nil {
		return nil
	}
	return o.Images
}

type ConfigurationDetailsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	ConfigurationDetails200ApplicationJSONObject *ConfigurationDetails200ApplicationJSON
}

func (o *ConfigurationDetailsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigurationDetailsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigurationDetailsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ConfigurationDetailsResponse) GetConfigurationDetails200ApplicationJSONObject() *ConfigurationDetails200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigurationDetails200ApplicationJSONObject
}
