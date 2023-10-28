// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type ConfigurationCountriesRequest struct {
	Language *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
}

func (c ConfigurationCountriesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConfigurationCountriesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConfigurationCountriesRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

type ConfigurationCountries200ApplicationJSON struct {
	EnglishName *string `json:"english_name,omitempty"`
	Iso31661    *string `json:"iso_3166_1,omitempty"`
	NativeName  *string `json:"native_name,omitempty"`
}

func (o *ConfigurationCountries200ApplicationJSON) GetEnglishName() *string {
	if o == nil {
		return nil
	}
	return o.EnglishName
}

func (o *ConfigurationCountries200ApplicationJSON) GetIso31661() *string {
	if o == nil {
		return nil
	}
	return o.Iso31661
}

func (o *ConfigurationCountries200ApplicationJSON) GetNativeName() *string {
	if o == nil {
		return nil
	}
	return o.NativeName
}

type ConfigurationCountriesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	ConfigurationCountries200ApplicationJSONObjects []ConfigurationCountries200ApplicationJSON
}

func (o *ConfigurationCountriesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigurationCountriesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigurationCountriesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ConfigurationCountriesResponse) GetConfigurationCountries200ApplicationJSONObjects() []ConfigurationCountries200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigurationCountries200ApplicationJSONObjects
}
