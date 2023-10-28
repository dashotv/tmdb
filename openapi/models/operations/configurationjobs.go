// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ConfigurationJobs200ApplicationJSON struct {
	Department *string  `json:"department,omitempty"`
	Jobs       []string `json:"jobs,omitempty"`
}

func (o *ConfigurationJobs200ApplicationJSON) GetDepartment() *string {
	if o == nil {
		return nil
	}
	return o.Department
}

func (o *ConfigurationJobs200ApplicationJSON) GetJobs() []string {
	if o == nil {
		return nil
	}
	return o.Jobs
}

type ConfigurationJobsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	ConfigurationJobs200ApplicationJSONObjects []ConfigurationJobs200ApplicationJSON
}

func (o *ConfigurationJobsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigurationJobsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigurationJobsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ConfigurationJobsResponse) GetConfigurationJobs200ApplicationJSONObjects() []ConfigurationJobs200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ConfigurationJobs200ApplicationJSONObjects
}