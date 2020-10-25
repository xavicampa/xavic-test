package myapi

import (
	openapi "myapi/go"
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockItemStore -
type MockItemStore struct {
	mock.Mock
}

// RetrieveItems - implementation
func (is *MockItemStore) RetrieveItems() []openapi.PersistedItem {
	args := is.Called()
	return args.Get(0).([]openapi.PersistedItem)
}

func (is *MockItemStore) CreateItem(s string) string {
	args := is.Called(s)
	return args.Get(0).(string)
}

func (is *MockItemStore) DeleteItem(s string) bool {
	args := is.Called(s)
	return args.Get(0).(bool)
}

func TestRetrieveItems(t *testing.T) {
	mockedStore := new(MockItemStore)
	mockedStore.On("RetrieveItems").Return([]openapi.PersistedItem{{Id: "1", Name: "test"}})

	myapi := &ItemAPIService{itemStore: mockedStore}
	myapi.ItemsGet()

	mockedStore.AssertExpectations(t)
	mockedStore.AssertCalled(t, "RetrieveItems")
}

func TestPostItem(t *testing.T) {
	mockedStore := new(MockItemStore)
	mockedStore.On("CreateItem", "haha").Return("123")

	myapi := &ItemAPIService{itemStore: mockedStore}
	myapi.ItemsPost(openapi.Item{Name: "haha"})

	mockedStore.AssertExpectations(t)
	mockedStore.AssertCalled(t, "CreateItem", "haha")
}

func TestDeleteItem(t *testing.T) {
	mockedStore := new(MockItemStore)
	mockedStore.On("DeleteItem", "haha").Return(true)

	myapi := &ItemAPIService{itemStore: mockedStore}
	myapi.ItemsItemIdDelete("haha")

	mockedStore.AssertExpectations(t)
	mockedStore.AssertCalled(t, "DeleteItem", "haha")
}
