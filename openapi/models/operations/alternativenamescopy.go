// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type AlternativeNamesCopyRequest struct {
	NetworkID int `pathParam:"style=simple,explode=false,name=network_id"`
}

func (o *AlternativeNamesCopyRequest) GetNetworkID() int {
	if o == nil {
		return 0
	}
	return o.NetworkID
}

type AlternativeNamesCopy200ApplicationJSONLogos struct {
	AspectRatio *float64 `default:"0" json:"aspect_ratio"`
	FilePath    *string  `json:"file_path,omitempty"`
	FileType    *string  `json:"file_type,omitempty"`
	Height      *int64   `default:"0" json:"height"`
	ID          *string  `json:"id,omitempty"`
	VoteAverage *float64 `default:"0" json:"vote_average"`
	VoteCount   *int64   `default:"0" json:"vote_count"`
	Width       *int64   `default:"0" json:"width"`
}

func (a AlternativeNamesCopy200ApplicationJSONLogos) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AlternativeNamesCopy200ApplicationJSONLogos) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AlternativeNamesCopy200ApplicationJSONLogos) GetAspectRatio() *float64 {
	if o == nil {
		return nil
	}
	return o.AspectRatio
}

func (o *AlternativeNamesCopy200ApplicationJSONLogos) GetFilePath() *string {
	if o == nil {
		return nil
	}
	return o.FilePath
}

func (o *AlternativeNamesCopy200ApplicationJSONLogos) GetFileType() *string {
	if o == nil {
		return nil
	}
	return o.FileType
}

func (o *AlternativeNamesCopy200ApplicationJSONLogos) GetHeight() *int64 {
	if o == nil {
		return nil
	}
	return o.Height
}

func (o *AlternativeNamesCopy200ApplicationJSONLogos) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AlternativeNamesCopy200ApplicationJSONLogos) GetVoteAverage() *float64 {
	if o == nil {
		return nil
	}
	return o.VoteAverage
}

func (o *AlternativeNamesCopy200ApplicationJSONLogos) GetVoteCount() *int64 {
	if o == nil {
		return nil
	}
	return o.VoteCount
}

func (o *AlternativeNamesCopy200ApplicationJSONLogos) GetWidth() *int64 {
	if o == nil {
		return nil
	}
	return o.Width
}

// AlternativeNamesCopy200ApplicationJSON - 200
type AlternativeNamesCopy200ApplicationJSON struct {
	ID    *int64                                        `default:"0" json:"id"`
	Logos []AlternativeNamesCopy200ApplicationJSONLogos `json:"logos,omitempty"`
}

func (a AlternativeNamesCopy200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AlternativeNamesCopy200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AlternativeNamesCopy200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AlternativeNamesCopy200ApplicationJSON) GetLogos() []AlternativeNamesCopy200ApplicationJSONLogos {
	if o == nil {
		return nil
	}
	return o.Logos
}

type AlternativeNamesCopyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	AlternativeNamesCopy200ApplicationJSONObject *AlternativeNamesCopy200ApplicationJSON
}

func (o *AlternativeNamesCopyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AlternativeNamesCopyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AlternativeNamesCopyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AlternativeNamesCopyResponse) GetAlternativeNamesCopy200ApplicationJSONObject() *AlternativeNamesCopy200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.AlternativeNamesCopy200ApplicationJSONObject
}
