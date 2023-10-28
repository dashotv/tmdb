// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

// PersonLatestID200ApplicationJSON - 200
type PersonLatestID200ApplicationJSON struct {
	Adult              *bool         `default:"true" json:"adult"`
	AlsoKnownAs        []interface{} `json:"also_known_as,omitempty"`
	Biography          *string       `json:"biography,omitempty"`
	Birthday           interface{}   `json:"birthday,omitempty"`
	Deathday           interface{}   `json:"deathday,omitempty"`
	Gender             *int64        `default:"0" json:"gender"`
	Homepage           interface{}   `json:"homepage,omitempty"`
	ID                 *int64        `default:"0" json:"id"`
	ImdbID             interface{}   `json:"imdb_id,omitempty"`
	KnownForDepartment interface{}   `json:"known_for_department,omitempty"`
	Name               *string       `json:"name,omitempty"`
	PlaceOfBirth       interface{}   `json:"place_of_birth,omitempty"`
	Popularity         *int64        `default:"0" json:"popularity"`
	ProfilePath        interface{}   `json:"profile_path,omitempty"`
}

func (p PersonLatestID200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonLatestID200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PersonLatestID200ApplicationJSON) GetAdult() *bool {
	if o == nil {
		return nil
	}
	return o.Adult
}

func (o *PersonLatestID200ApplicationJSON) GetAlsoKnownAs() []interface{} {
	if o == nil {
		return nil
	}
	return o.AlsoKnownAs
}

func (o *PersonLatestID200ApplicationJSON) GetBiography() *string {
	if o == nil {
		return nil
	}
	return o.Biography
}

func (o *PersonLatestID200ApplicationJSON) GetBirthday() interface{} {
	if o == nil {
		return nil
	}
	return o.Birthday
}

func (o *PersonLatestID200ApplicationJSON) GetDeathday() interface{} {
	if o == nil {
		return nil
	}
	return o.Deathday
}

func (o *PersonLatestID200ApplicationJSON) GetGender() *int64 {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *PersonLatestID200ApplicationJSON) GetHomepage() interface{} {
	if o == nil {
		return nil
	}
	return o.Homepage
}

func (o *PersonLatestID200ApplicationJSON) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PersonLatestID200ApplicationJSON) GetImdbID() interface{} {
	if o == nil {
		return nil
	}
	return o.ImdbID
}

func (o *PersonLatestID200ApplicationJSON) GetKnownForDepartment() interface{} {
	if o == nil {
		return nil
	}
	return o.KnownForDepartment
}

func (o *PersonLatestID200ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PersonLatestID200ApplicationJSON) GetPlaceOfBirth() interface{} {
	if o == nil {
		return nil
	}
	return o.PlaceOfBirth
}

func (o *PersonLatestID200ApplicationJSON) GetPopularity() *int64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *PersonLatestID200ApplicationJSON) GetProfilePath() interface{} {
	if o == nil {
		return nil
	}
	return o.ProfilePath
}

type PersonLatestIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	PersonLatestID200ApplicationJSONObject *PersonLatestID200ApplicationJSON
}

func (o *PersonLatestIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PersonLatestIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PersonLatestIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PersonLatestIDResponse) GetPersonLatestID200ApplicationJSONObject() *PersonLatestID200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PersonLatestID200ApplicationJSONObject
}
