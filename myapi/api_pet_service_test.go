/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package myapi

import (
	"testing"
)

func TestGetPetById(t *testing.T) {
	myapi := NewPetAPIService()
	ans, _ := myapi.GetPetById(1)
	if ans != "hello 1. This is new!" {
		t.Errorf("GetPetById('test') = '%s'; want 'hello 1. This is new!'", ans)
	}
}
