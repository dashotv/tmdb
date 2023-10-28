// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type GenreMovieListRequest struct {
	Language *string `default:"en" queryParam:"style=form,explode=true,name=language"`
}

func (g GenreMovieListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GenreMovieListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GenreMovieListRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

type GenreMovieList200ApplicationJSONGenres struct {
	ID   *int64  `default:"0" json:"id"`
	Name *string `json:"name,omitempty"`
}

func (g GenreMovieList200ApplicationJSONGenres) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GenreMovieList200ApplicationJSONGenres) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GenreMovieList200ApplicationJSONGenres) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GenreMovieList200ApplicationJSONGenres) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GenreMovieList200ApplicationJSON - 200
type GenreMovieList200ApplicationJSON struct {
	Genres []GenreMovieList200ApplicationJSONGenres `json:"genres,omitempty"`
}

func (o *GenreMovieList200ApplicationJSON) GetGenres() []GenreMovieList200ApplicationJSONGenres {
	if o == nil {
		return nil
	}
	return o.Genres
}

type GenreMovieListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	GenreMovieList200ApplicationJSONObject *GenreMovieList200ApplicationJSON
}

func (o *GenreMovieListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GenreMovieListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GenreMovieListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GenreMovieListResponse) GetGenreMovieList200ApplicationJSONObject() *GenreMovieList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GenreMovieList200ApplicationJSONObject
}