// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvEpisodeVideosRequest struct {
	EpisodeNumber int `pathParam:"style=simple,explode=false,name=episode_number"`
	// filter the list results by language, supports more than one value by using a comma
	IncludeVideoLanguage *string `queryParam:"style=form,explode=true,name=include_video_language"`
	Language             *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
	SeasonNumber         int     `pathParam:"style=simple,explode=false,name=season_number"`
	SeriesID             int     `pathParam:"style=simple,explode=false,name=series_id"`
}

func (t TvEpisodeVideosRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvEpisodeVideosRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvEpisodeVideosRequest) GetEpisodeNumber() int {
	if o == nil {
		return 0
	}
	return o.EpisodeNumber
}

func (o *TvEpisodeVideosRequest) GetIncludeVideoLanguage() *string {
	if o == nil {
		return nil
	}
	return o.IncludeVideoLanguage
}

func (o *TvEpisodeVideosRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *TvEpisodeVideosRequest) GetSeasonNumber() int {
	if o == nil {
		return 0
	}
	return o.SeasonNumber
}

func (o *TvEpisodeVideosRequest) GetSeriesID() int {
	if o == nil {
		return 0
	}
	return o.SeriesID
}

type TvEpisodeVideos200ApplicationJSONResults struct {
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

func (t TvEpisodeVideos200ApplicationJSONResults) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvEpisodeVideos200ApplicationJSONResults) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvEpisodeVideos200ApplicationJSONResults) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvEpisodeVideos200ApplicationJSONResults) GetIso31661() *string {
	if o == nil {
		return nil
	}
	return o.Iso31661
}

func (o *TvEpisodeVideos200ApplicationJSONResults) GetIso6391() *string {
	if o == nil {
		return nil
	}
	return o.Iso6391
}

func (o *TvEpisodeVideos200ApplicationJSONResults) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *TvEpisodeVideos200ApplicationJSONResults) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TvEpisodeVideos200ApplicationJSONResults) GetOfficial() *bool {
	if o == nil {
		return nil
	}
	return o.Official
}

func (o *TvEpisodeVideos200ApplicationJSONResults) GetPublishedAt() *string {
	if o == nil {
		return nil
	}
	return o.PublishedAt
}

func (o *TvEpisodeVideos200ApplicationJSONResults) GetSite() *string {
	if o == nil {
		return nil
	}
	return o.Site
}

func (o *TvEpisodeVideos200ApplicationJSONResults) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *TvEpisodeVideos200ApplicationJSONResults) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

// TvEpisodeVideos200ApplicationJSON - 200
type TvEpisodeVideos200ApplicationJSON struct {
	ID      *int64                                     `default:"0" json:"id"`
	Results []TvEpisodeVideos200ApplicationJSONResults `json:"results,omitempty"`
}

func (t TvEpisodeVideos200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvEpisodeVideos200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvEpisodeVideos200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvEpisodeVideos200ApplicationJSON) GetResults() []TvEpisodeVideos200ApplicationJSONResults {
	if o == nil {
		return nil
	}
	return o.Results
}

type TvEpisodeVideosResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvEpisodeVideos200ApplicationJSONObject *TvEpisodeVideos200ApplicationJSON
}

func (o *TvEpisodeVideosResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvEpisodeVideosResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvEpisodeVideosResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvEpisodeVideosResponse) GetTvEpisodeVideos200ApplicationJSONObject() *TvEpisodeVideos200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvEpisodeVideos200ApplicationJSONObject
}
