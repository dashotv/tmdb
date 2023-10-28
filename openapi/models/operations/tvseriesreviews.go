// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type TvSeriesReviewsRequest struct {
	SeriesID int     `pathParam:"style=simple,explode=false,name=series_id"`
	Language *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
	Page     *int    `default:"1" queryParam:"style=form,explode=true,name=page"`
}

func (t TvSeriesReviewsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesReviewsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesReviewsRequest) GetSeriesID() int {
	if o == nil {
		return 0
	}
	return o.SeriesID
}

func (o *TvSeriesReviewsRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *TvSeriesReviewsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

type TvSeriesReviews200ApplicationJSONResultsAuthorDetails struct {
	AvatarPath *string `json:"avatar_path,omitempty"`
	Name       *string `json:"name,omitempty"`
	Rating     *int64  `default:"0" json:"rating"`
	Username   *string `json:"username,omitempty"`
}

func (t TvSeriesReviews200ApplicationJSONResultsAuthorDetails) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesReviews200ApplicationJSONResultsAuthorDetails) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesReviews200ApplicationJSONResultsAuthorDetails) GetAvatarPath() *string {
	if o == nil {
		return nil
	}
	return o.AvatarPath
}

func (o *TvSeriesReviews200ApplicationJSONResultsAuthorDetails) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TvSeriesReviews200ApplicationJSONResultsAuthorDetails) GetRating() *int64 {
	if o == nil {
		return nil
	}
	return o.Rating
}

func (o *TvSeriesReviews200ApplicationJSONResultsAuthorDetails) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type TvSeriesReviews200ApplicationJSONResults struct {
	Author        *string                                                `json:"author,omitempty"`
	AuthorDetails *TvSeriesReviews200ApplicationJSONResultsAuthorDetails `json:"author_details,omitempty"`
	Content       *string                                                `json:"content,omitempty"`
	CreatedAt     *string                                                `json:"created_at,omitempty"`
	ID            *string                                                `json:"id,omitempty"`
	UpdatedAt     *string                                                `json:"updated_at,omitempty"`
	URL           *string                                                `json:"url,omitempty"`
}

func (o *TvSeriesReviews200ApplicationJSONResults) GetAuthor() *string {
	if o == nil {
		return nil
	}
	return o.Author
}

func (o *TvSeriesReviews200ApplicationJSONResults) GetAuthorDetails() *TvSeriesReviews200ApplicationJSONResultsAuthorDetails {
	if o == nil {
		return nil
	}
	return o.AuthorDetails
}

func (o *TvSeriesReviews200ApplicationJSONResults) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *TvSeriesReviews200ApplicationJSONResults) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TvSeriesReviews200ApplicationJSONResults) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeriesReviews200ApplicationJSONResults) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *TvSeriesReviews200ApplicationJSONResults) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// TvSeriesReviews200ApplicationJSON - 200
type TvSeriesReviews200ApplicationJSON struct {
	ID           *int64                                     `default:"0" json:"id"`
	Page         *int64                                     `default:"0" json:"page"`
	Results      []TvSeriesReviews200ApplicationJSONResults `json:"results,omitempty"`
	TotalPages   *int64                                     `default:"0" json:"total_pages"`
	TotalResults *int64                                     `default:"0" json:"total_results"`
}

func (t TvSeriesReviews200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TvSeriesReviews200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TvSeriesReviews200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TvSeriesReviews200ApplicationJSON) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *TvSeriesReviews200ApplicationJSON) GetResults() []TvSeriesReviews200ApplicationJSONResults {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *TvSeriesReviews200ApplicationJSON) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

func (o *TvSeriesReviews200ApplicationJSON) GetTotalResults() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalResults
}

type TvSeriesReviewsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	TvSeriesReviews200ApplicationJSONObject *TvSeriesReviews200ApplicationJSON
}

func (o *TvSeriesReviewsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TvSeriesReviewsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TvSeriesReviewsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TvSeriesReviewsResponse) GetTvSeriesReviews200ApplicationJSONObject() *TvSeriesReviews200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.TvSeriesReviews200ApplicationJSONObject
}
