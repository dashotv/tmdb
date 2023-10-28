// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvSeriesVideosRequest struct {
	SeriesID int `pathParam:"style=simple,explode=false,name=series_id"`
	// filter the list results by language, supports more than one value by using a comma
	IncludeVideoLanguage *string `queryParam:"style=form,explode=true,name=include_video_language"`
	Language             *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
}

func (t TvSeriesVideosRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesVideosRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesVideosRequest) GetSeriesID() int {
	if o == nil {
		return 0
	}
	return o.SeriesID
}

func (o *TvSeriesVideosRequest) GetIncludeVideoLanguage() *string {
	if o == nil {
		return nil
	}
	return o.IncludeVideoLanguage
}

func (o *TvSeriesVideosRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

type TvSeriesVideos200ApplicationJSONResults struct {
	ID          *string `json:"id,omitempty"`
	Iso31661    *string `json:"iso_3166_1,omitempty"`
	Iso6391     *string `json:"iso_639_1,omitempty"`
	Key         *string `json:"key,omitempty"`
	Name        *string `json:"name,omitempty"`
	Official    *bool   `default:"true" json:"official"`
	PublishedAt *string `json:"published_at,omitempty"`
	Site        *string `json:"site,omitempty"`
	Size        *int64  `default:"0" json:"size"`
	Type        *string `json:"type,omitempty"`
}

func (t TvSeriesVideos200ApplicationJSONResults) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesVideos200ApplicationJSONResults) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesVideos200ApplicationJSONResults) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeriesVideos200ApplicationJSONResults) GetIso31661() *string {
	if o == nil {
		return nil
	}
	return o.Iso31661
}

func (o *TvSeriesVideos200ApplicationJSONResults) GetIso6391() *string {
	if o == nil {
		return nil
	}
	return o.Iso6391
}

func (o *TvSeriesVideos200ApplicationJSONResults) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *TvSeriesVideos200ApplicationJSONResults) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TvSeriesVideos200ApplicationJSONResults) GetOfficial() *bool {
	if o == nil {
		return nil
	}
	return o.Official
}

func (o *TvSeriesVideos200ApplicationJSONResults) GetPublishedAt() *string {
	if o == nil {
		return nil
	}
	return o.PublishedAt
}

func (o *TvSeriesVideos200ApplicationJSONResults) GetSite() *string {
	if o == nil {
		return nil
	}
	return o.Site
}

func (o *TvSeriesVideos200ApplicationJSONResults) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *TvSeriesVideos200ApplicationJSONResults) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

// TvSeriesVideos200ApplicationJSON - 200
type TvSeriesVideos200ApplicationJSON struct {
	ID      *int64                                    `default:"0" json:"id"`
	Results []TvSeriesVideos200ApplicationJSONResults `json:"results,omitempty"`
}

func (t TvSeriesVideos200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesVideos200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesVideos200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeriesVideos200ApplicationJSON) GetResults() []TvSeriesVideos200ApplicationJSONResults {
	if o == nil {
		return nil
	}
	return o.Results
}

type TvSeriesVideosResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvSeriesVideos200ApplicationJSONObject *TvSeriesVideos200ApplicationJSON
}

func (o *TvSeriesVideosResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvSeriesVideosResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvSeriesVideosResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvSeriesVideosResponse) GetTvSeriesVideos200ApplicationJSONObject() *TvSeriesVideos200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvSeriesVideos200ApplicationJSONObject
}
