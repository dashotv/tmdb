// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type WatchProvidersAvailableRegionsRequest struct {
	Language *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
}

func (w WatchProvidersAvailableRegionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WatchProvidersAvailableRegionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WatchProvidersAvailableRegionsRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

type WatchProvidersAvailableRegions200ApplicationJSONResults struct {
	EnglishName *string `json:"english_name,omitempty"`
	Iso31661    *string `json:"iso_3166_1,omitempty"`
	NativeName  *string `json:"native_name,omitempty"`
}

func (o *WatchProvidersAvailableRegions200ApplicationJSONResults) GetEnglishName() *string {
	if o == nil {
		return nil
	}
	return o.EnglishName
}

func (o *WatchProvidersAvailableRegions200ApplicationJSONResults) GetIso31661() *string {
	if o == nil {
		return nil
	}
	return o.Iso31661
}

func (o *WatchProvidersAvailableRegions200ApplicationJSONResults) GetNativeName() *string {
	if o == nil {
		return nil
	}
	return o.NativeName
}

// WatchProvidersAvailableRegions200ApplicationJSON - 200
type WatchProvidersAvailableRegions200ApplicationJSON struct {
	Results []WatchProvidersAvailableRegions200ApplicationJSONResults `json:"results,omitempty"`
}

func (o *WatchProvidersAvailableRegions200ApplicationJSON) GetResults() []WatchProvidersAvailableRegions200ApplicationJSONResults {
	if o == nil {
		return nil
	}
	return o.Results
}

type WatchProvidersAvailableRegionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	WatchProvidersAvailableRegions200ApplicationJSONObject *WatchProvidersAvailableRegions200ApplicationJSON
}

func (o *WatchProvidersAvailableRegionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WatchProvidersAvailableRegionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WatchProvidersAvailableRegionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *WatchProvidersAvailableRegionsResponse) GetWatchProvidersAvailableRegions200ApplicationJSONObject() *WatchProvidersAvailableRegions200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.WatchProvidersAvailableRegions200ApplicationJSONObject
}