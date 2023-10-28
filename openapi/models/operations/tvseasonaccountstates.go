// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvSeasonAccountStatesRequest struct {
	SeasonNumber   int     `pathParam:"style=simple,explode=false,name=season_number"`
	SeriesID       int     `pathParam:"style=simple,explode=false,name=series_id"`
	GuestSessionID *string `queryParam:"style=form,explode=true,name=guest_session_id"`
	SessionID      *string `queryParam:"style=form,explode=true,name=session_id"`
}

func (o *TvSeasonAccountStatesRequest) GetSeasonNumber() int {
	if o == nil {
		return 0
	}
	return o.SeasonNumber
}

func (o *TvSeasonAccountStatesRequest) GetSeriesID() int {
	if o == nil {
		return 0
	}
	return o.SeriesID
}

func (o *TvSeasonAccountStatesRequest) GetGuestSessionID() *string {
	if o == nil {
		return nil
	}
	return o.GuestSessionID
}

func (o *TvSeasonAccountStatesRequest) GetSessionID() *string {
	if o == nil {
		return nil
	}
	return o.SessionID
}

type TvSeasonAccountStates200ApplicationJSONResultsRated struct {
	Value *int64 `default:"0" json:"value"`
}

func (t TvSeasonAccountStates200ApplicationJSONResultsRated) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeasonAccountStates200ApplicationJSONResultsRated) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeasonAccountStates200ApplicationJSONResultsRated) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type TvSeasonAccountStates200ApplicationJSONResults struct {
	EpisodeNumber *int64                                               `default:"0" json:"episode_number"`
	ID            *int64                                               `default:"0" json:"id"`
	Rated         *TvSeasonAccountStates200ApplicationJSONResultsRated `json:"rated,omitempty"`
}

func (t TvSeasonAccountStates200ApplicationJSONResults) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeasonAccountStates200ApplicationJSONResults) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeasonAccountStates200ApplicationJSONResults) GetEpisodeNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodeNumber
}

func (o *TvSeasonAccountStates200ApplicationJSONResults) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeasonAccountStates200ApplicationJSONResults) GetRated() *TvSeasonAccountStates200ApplicationJSONResultsRated {
	if o == nil {
		return nil
	}
	return o.Rated
}

// TvSeasonAccountStates200ApplicationJSON - 200
type TvSeasonAccountStates200ApplicationJSON struct {
	ID      *int64                                           `default:"0" json:"id"`
	Results []TvSeasonAccountStates200ApplicationJSONResults `json:"results,omitempty"`
}

func (t TvSeasonAccountStates200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeasonAccountStates200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeasonAccountStates200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeasonAccountStates200ApplicationJSON) GetResults() []TvSeasonAccountStates200ApplicationJSONResults {
	if o == nil {
		return nil
	}
	return o.Results
}

type TvSeasonAccountStatesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvSeasonAccountStates200ApplicationJSONObject *TvSeasonAccountStates200ApplicationJSON
}

func (o *TvSeasonAccountStatesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvSeasonAccountStatesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvSeasonAccountStatesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvSeasonAccountStatesResponse) GetTvSeasonAccountStates200ApplicationJSONObject() *TvSeasonAccountStates200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvSeasonAccountStates200ApplicationJSONObject
}