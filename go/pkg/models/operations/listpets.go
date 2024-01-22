// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/triple/v13/pkg/models/shared"
	"net/http"
)

type ListPetsRequest struct {
	// How many items to return at one time (max 100)
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
}

func (o *ListPetsRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

type ListPetsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// unexpected error
	Error   *shared.Error
	Headers map[string][]string
	// A paged array of pets
	Pets []shared.Pet
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListPetsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPetsResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ListPetsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *ListPetsResponse) GetPets() []shared.Pet {
	if o == nil {
		return nil
	}
	return o.Pets
}

func (o *ListPetsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPetsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
