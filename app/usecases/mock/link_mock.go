package mock_usecases

import (
	models "ozon_test/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLinkUseCase is a mock of LinkUseCase interface.
type MockLinkUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockLinkUseCaseMockRecorder
}

// MockLinkUseCaseMockRecorder is the mock recorder for MockLinkUseCase.
type MockLinkUseCaseMockRecorder struct {
	mock *MockLinkUseCase
}

// NewMockLinkUseCase creates a new mock instance.
func NewMockLinkUseCase(ctrl *gomock.Controller) *MockLinkUseCase {
	mock := &MockLinkUseCase{ctrl: ctrl}
	mock.recorder = &MockLinkUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLinkUseCase) EXPECT() *MockLinkUseCaseMockRecorder {
	return m.recorder
}

// CreateLink mocks base method.
func (m *MockLinkUseCase) CreateLink(link *models.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLink", link)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLink indicates an expected call of CreateLink.
func (mr *MockLinkUseCaseMockRecorder) CreateLink(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLink", reflect.TypeOf((*MockLinkUseCase)(nil).CreateLink), link)
}

// GetLink mocks base method.
func (m *MockLinkUseCase) GetLink(link string) (models.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLink", link)
	ret0, _ := ret[0].(models.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLink indicates an expected call of GetLink.
func (mr *MockLinkUseCaseMockRecorder) GetLink(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLink", reflect.TypeOf((*MockLinkUseCase)(nil).GetLink), link)
}
