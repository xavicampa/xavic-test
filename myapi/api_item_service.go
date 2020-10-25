package myapi

import (
	"errors"
	openapi "myapi/go"
)

// ItemAPIService - object implementing interface
type ItemAPIService struct {
	itemStore ItemStore
}

// NewItemAPIService - returns object implementing interface
func NewItemAPIService(itemStore ItemStore) openapi.DefaultApiServicer {
	return &ItemAPIService{itemStore: itemStore}
}

// ItemsGet - Returns a list of items
func (s *ItemAPIService) ItemsGet() (interface{}, error) {
	// TODO - update ItemsGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	//return nil, errors.New("service method 'ItemsGet' not implemented")
	return s.itemStore.RetrieveItems(), nil
}

// ItemsPost -
func (s *ItemAPIService) ItemsPost(item openapi.Item) (interface{}, error) {
	// TODO - update PingGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	//return nil, errors.New("service method 'PingGet' not implemented")
	//result := openapi.HeartBeatResponse{Status: "OK", Message: "All good!"}
	id := s.itemStore.CreateItem(item.Name)
	result := openapi.PersistedItem{Id: id, Name: item.Name}
	return result, nil
}

// ItemsItemIdDelete - Removes an item from the list
func (s *ItemAPIService) ItemsItemIdDelete(itemId string) (interface{}, error) {
	// TODO - update ItemsItemIdDelete with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	//return nil, errors.New("service method 'ItemsItemIdDelete' not implemented")
	deleted := s.itemStore.DeleteItem(itemId)
	if deleted {
		return nil, nil
	} else {
		return nil, errors.New("NotFound")
	}
}

// PingGet - Heartbeat
func (s *ItemAPIService) PingGet() (interface{}, error) {
	// TODO - update PingGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	//return nil, errors.New("service method 'PingGet' not implemented")
	result := openapi.HeartBeatResponse{Status: "OK", Message: "All good!"}
	return result, nil
}
