package myapi

import (
	"fmt"
	openapi "myapi/go"

	"github.com/google/uuid"
)

// MemoryItemStore -
type MemoryItemStore struct {
	items []openapi.PersistedItem
}

// RetrieveItems - implementation
func (is *MemoryItemStore) RetrieveItems() []openapi.PersistedItem {
	return is.items
}

// CreateItem -
func (is *MemoryItemStore) CreateItem(s string) string {
	uuid := fmt.Sprint(uuid.New())
	item := openapi.PersistedItem{Id: uuid, Name: s}
	var newItems []openapi.PersistedItem
	if is.items == nil {
		newItems = make([]openapi.PersistedItem, 1)
		newItems[0] = item
	} else {
		newItems = is.items
		newItems = append(newItems, item)
	}
	is.items = newItems
	return uuid
}

// DeleteItem -
func (is *MemoryItemStore) DeleteItem(s string) bool {
	for i := range is.items {
		if is.items[i].Id == s {
			if i == len(is.items)-1 {
				is.items = is.items[:i]
			} else if i == 0 {
				is.items = is.items[i+1:]
			} else {
				is.items = append(is.items[:i], is.items[i+1:]...)
			}
			return true
		}
	}
	return false
}
