// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type PersonDetailsRequest struct {
	PersonID int `pathParam:"style=simple,explode=false,name=person_id"`
	// comma separated list of endpoints within this namespace, 20 items max
	AppendToResponse *string `queryParam:"style=form,explode=true,name=append_to_response"`
	Language         *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
}

func (p PersonDetailsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonDetailsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PersonDetailsRequest) GetPersonID() int {
	if o == nil {
		return 0
	}
	return o.PersonID
}

func (o *PersonDetailsRequest) GetAppendToResponse() *string {
	if o == nil {
		return nil
	}
	return o.AppendToResponse
}

func (o *PersonDetailsRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

// PersonDetails200ApplicationJSON - 200
type PersonDetails200ApplicationJSON struct {
	Adult              *bool       `default:"true" json:"adult"`
	AlsoKnownAs        []string    `json:"also_known_as,omitempty"`
	Biography          *string     `json:"biography,omitempty"`
	Birthday           *string     `json:"birthday,omitempty"`
	Deathday           interface{} `json:"deathday,omitempty"`
	Gender             *int64      `default:"0" json:"gender"`
	Homepage           interface{} `json:"homepage,omitempty"`
	ID                 *int64      `default:"0" json:"id"`
	ImdbID             *string     `json:"imdb_id,omitempty"`
	KnownForDepartment *string     `json:"known_for_department,omitempty"`
	Name               *string     `json:"name,omitempty"`
	PlaceOfBirth       *string     `json:"place_of_birth,omitempty"`
	Popularity         *float64    `default:"0" json:"popularity"`
	ProfilePath        *string     `json:"profile_path,omitempty"`
}

func (p PersonDetails200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonDetails200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PersonDetails200ApplicationJSON) GetAdult() *bool {
	if o == nil {
		return nil
	}
	return o.Adult
}

func (o *PersonDetails200ApplicationJSON) GetAlsoKnownAs() []string {
	if o == nil {
		return nil
	}
	return o.AlsoKnownAs
}

func (o *PersonDetails200ApplicationJSON) GetBiography() *string {
	if o == nil {
		return nil
	}
	return o.Biography
}

func (o *PersonDetails200ApplicationJSON) GetBirthday() *string {
	if o == nil {
		return nil
	}
	return o.Birthday
}

func (o *PersonDetails200ApplicationJSON) GetDeathday() interface{} {
	if o == nil {
		return nil
	}
	return o.Deathday
}

func (o *PersonDetails200ApplicationJSON) GetGender() *int64 {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *PersonDetails200ApplicationJSON) GetHomepage() interface{} {
	if o == nil {
		return nil
	}
	return o.Homepage
}

func (o *PersonDetails200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PersonDetails200ApplicationJSON) GetImdbID() *string {
	if o == nil {
		return nil
	}
	return o.ImdbID
}

func (o *PersonDetails200ApplicationJSON) GetKnownForDepartment() *string {
	if o == nil {
		return nil
	}
	return o.KnownForDepartment
}

func (o *PersonDetails200ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PersonDetails200ApplicationJSON) GetPlaceOfBirth() *string {
	if o == nil {
		return nil
	}
	return o.PlaceOfBirth
}

func (o *PersonDetails200ApplicationJSON) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *PersonDetails200ApplicationJSON) GetProfilePath() *string {
	if o == nil {
		return nil
	}
	return o.ProfilePath
}

type PersonDetailsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	PersonDetails200ApplicationJSONObject *PersonDetails200ApplicationJSON
}

func (o *PersonDetailsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PersonDetailsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PersonDetailsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PersonDetailsResponse) GetPersonDetails200ApplicationJSONObject() *PersonDetails200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PersonDetails200ApplicationJSONObject
}
