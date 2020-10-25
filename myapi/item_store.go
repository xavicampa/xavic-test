package myapi

import (
	openapi "myapi/go"
)

// ItemStore -
type ItemStore interface {
	RetrieveItems() []openapi.PersistedItem
	CreateItem(string) string
	DeleteItem(string) bool
}
