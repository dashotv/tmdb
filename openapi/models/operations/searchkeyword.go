// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type SearchKeywordRequest struct {
	Query string `queryParam:"style=form,explode=true,name=query"`
	Page  *int   `default:"1" queryParam:"style=form,explode=true,name=page"`
}

func (s SearchKeywordRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SearchKeywordRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SearchKeywordRequest) GetQuery() string {
	if o == nil {
		return ""
	}
	return o.Query
}

func (o *SearchKeywordRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

type SearchKeyword200ApplicationJSONResults struct {
	ID   *int64  `default:"0" json:"id"`
	Name *string `json:"name,omitempty"`
}

func (s SearchKeyword200ApplicationJSONResults) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SearchKeyword200ApplicationJSONResults) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SearchKeyword200ApplicationJSONResults) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SearchKeyword200ApplicationJSONResults) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// SearchKeyword200ApplicationJSON - 200
type SearchKeyword200ApplicationJSON struct {
	Page         *int64                                   `default:"0" json:"page"`
	Results      []SearchKeyword200ApplicationJSONResults `json:"results,omitempty"`
	TotalPages   *int64                                   `default:"0" json:"total_pages"`
	TotalResults *int64                                   `default:"0" json:"total_results"`
}

func (s SearchKeyword200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SearchKeyword200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SearchKeyword200ApplicationJSON) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *SearchKeyword200ApplicationJSON) GetResults() []SearchKeyword200ApplicationJSONResults {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *SearchKeyword200ApplicationJSON) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}

func (o *SearchKeyword200ApplicationJSON) GetTotalResults() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalResults
}

type SearchKeywordResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	SearchKeyword200ApplicationJSONObject *SearchKeyword200ApplicationJSON
}

func (o *SearchKeywordResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SearchKeywordResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SearchKeywordResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SearchKeywordResponse) GetSearchKeyword200ApplicationJSONObject() *SearchKeyword200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.SearchKeyword200ApplicationJSONObject
}
