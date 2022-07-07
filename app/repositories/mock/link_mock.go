package mock_repositories

import (
	models "ozon_test/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLinkRepository is a mock of LinkRepository interface.
type MockLinkRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLinkRepositoryMockRecorder
}

// MockLinkRepositoryMockRecorder is the mock recorder for MockLinkRepository.
type MockLinkRepositoryMockRecorder struct {
	mock *MockLinkRepository
}

// NewMockLinkRepository creates a new mock instance.
func NewMockLinkRepository(ctrl *gomock.Controller) *MockLinkRepository {
	mock := &MockLinkRepository{ctrl: ctrl}
	mock.recorder = &MockLinkRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLinkRepository) EXPECT() *MockLinkRepositoryMockRecorder {
	return m.recorder
}

// CreateLink mocks base method.
func (m *MockLinkRepository) CreateLink(link *models.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLink", link)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLink indicates an expected call of CreateLink.
func (mr *MockLinkRepositoryMockRecorder) CreateLink(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLink", reflect.TypeOf((*MockLinkRepository)(nil).CreateLink), link)
}

// GetLink mocks base method.
func (m *MockLinkRepository) GetLink(shortLink string) (*models.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLink", shortLink)
	ret0, _ := ret[0].(*models.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLink indicates an expected call of GetLink.
func (mr *MockLinkRepositoryMockRecorder) GetLink(shortLink interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLink", reflect.TypeOf((*MockLinkRepository)(nil).GetLink), shortLink)
}

// IsExistOriginal mocks base method.
func (m *MockLinkRepository) IsExistOriginal(originalLink string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExistOriginal", originalLink)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExistOriginal indicates an expected call of IsExistOriginal.
func (mr *MockLinkRepositoryMockRecorder) IsExistOriginal(originalLink interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExistOriginal", reflect.TypeOf((*MockLinkRepository)(nil).IsExistOriginal), originalLink)
}

// IsExistShort mocks base method.
func (m *MockLinkRepository) IsExistShort(shortLink string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExistShort", shortLink)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExistShort indicates an expected call of IsExistShort.
func (mr *MockLinkRepositoryMockRecorder) IsExistShort(shortLink interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExistShort", reflect.TypeOf((*MockLinkRepository)(nil).IsExistShort), shortLink)
}
