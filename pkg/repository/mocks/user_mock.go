// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\ahdaa\GOLANG\Leuse_Ecommerce\pkg\repository\interface\user.go

// Package mock_interfaces is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	domain "github.com/ahdaan98/pkg/domain"
	models "github.com/ahdaan98/pkg/utils/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserRepository) AddAddress(userID int, address models.AddAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", userID, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserRepositoryMockRecorder) AddAddress(userID, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserRepository)(nil).AddAddress), userID, address)
}

// ChangePassword mocks base method.
func (m *MockUserRepository) ChangePassword(id int, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", id, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserRepositoryMockRecorder) ChangePassword(id, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserRepository)(nil).ChangePassword), id, password)
}

// CheckBlockStatus mocks base method.
func (m *MockUserRepository) CheckBlockStatus(email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBlockStatus", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckBlockStatus indicates an expected call of CheckBlockStatus.
func (mr *MockUserRepositoryMockRecorder) CheckBlockStatus(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBlockStatus", reflect.TypeOf((*MockUserRepository)(nil).CheckBlockStatus), email)
}

// CheckIfFirstAddress mocks base method.
func (m *MockUserRepository) CheckIfFirstAddress(id int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfFirstAddress", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckIfFirstAddress indicates an expected call of CheckIfFirstAddress.
func (mr *MockUserRepositoryMockRecorder) CheckIfFirstAddress(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfFirstAddress", reflect.TypeOf((*MockUserRepository)(nil).CheckIfFirstAddress), id)
}

// CheckUserExist mocks base method.
func (m *MockUserRepository) CheckUserExist(email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserExist", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserExist indicates an expected call of CheckUserExist.
func (mr *MockUserRepositoryMockRecorder) CheckUserExist(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserExist", reflect.TypeOf((*MockUserRepository)(nil).CheckUserExist), email)
}

// CheckUserExistByID mocks base method.
func (m *MockUserRepository) CheckUserExistByID(id int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserExistByID", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserExistByID indicates an expected call of CheckUserExistByID.
func (mr *MockUserRepositoryMockRecorder) CheckUserExistByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserExistByID", reflect.TypeOf((*MockUserRepository)(nil).CheckUserExistByID), id)
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(user models.UserSignUp) (models.UserDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(models.UserDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), user)
}

// EditDetails mocks base method.
func (m *MockUserRepository) EditDetails(details models.EditUserDetails, id int) (models.UserDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditDetails", details, id)
	ret0, _ := ret[0].(models.UserDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditDetails indicates an expected call of EditDetails.
func (mr *MockUserRepositoryMockRecorder) EditDetails(details, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditDetails", reflect.TypeOf((*MockUserRepository)(nil).EditDetails), details, id)
}

// FindBrand mocks base method.
func (m *MockUserRepository) FindBrand(inventory_id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBrand", inventory_id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBrand indicates an expected call of FindBrand.
func (mr *MockUserRepositoryMockRecorder) FindBrand(inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBrand", reflect.TypeOf((*MockUserRepository)(nil).FindBrand), inventory_id)
}

// FindBrandName mocks base method.
func (m *MockUserRepository) FindBrandName(brand_id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBrandName", brand_id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBrandName indicates an expected call of FindBrandName.
func (mr *MockUserRepositoryMockRecorder) FindBrandName(brand_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBrandName", reflect.TypeOf((*MockUserRepository)(nil).FindBrandName), brand_id)
}

// FindCartQuantity mocks base method.
func (m *MockUserRepository) FindCartQuantity(cart_id, inventory_id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCartQuantity", cart_id, inventory_id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCartQuantity indicates an expected call of FindCartQuantity.
func (mr *MockUserRepositoryMockRecorder) FindCartQuantity(cart_id, inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCartQuantity", reflect.TypeOf((*MockUserRepository)(nil).FindCartQuantity), cart_id, inventory_id)
}

// FindCategory mocks base method.
func (m *MockUserRepository) FindCategory(inventory_id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCategory", inventory_id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCategory indicates an expected call of FindCategory.
func (mr *MockUserRepositoryMockRecorder) FindCategory(inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCategory", reflect.TypeOf((*MockUserRepository)(nil).FindCategory), inventory_id)
}

// FindCategoryName mocks base method.
func (m *MockUserRepository) FindCategoryName(category_id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCategoryName", category_id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCategoryName indicates an expected call of FindCategoryName.
func (mr *MockUserRepositoryMockRecorder) FindCategoryName(category_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCategoryName", reflect.TypeOf((*MockUserRepository)(nil).FindCategoryName), category_id)
}

// FindPrice mocks base method.
func (m *MockUserRepository) FindPrice(inventory_id int) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPrice", inventory_id)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPrice indicates an expected call of FindPrice.
func (mr *MockUserRepositoryMockRecorder) FindPrice(inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPrice", reflect.TypeOf((*MockUserRepository)(nil).FindPrice), inventory_id)
}

// FindProductNames mocks base method.
func (m *MockUserRepository) FindProductNames(inventory_id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductNames", inventory_id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductNames indicates an expected call of FindProductNames.
func (mr *MockUserRepositoryMockRecorder) FindProductNames(inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductNames", reflect.TypeOf((*MockUserRepository)(nil).FindProductNames), inventory_id)
}

// FindStock mocks base method.
func (m *MockUserRepository) FindStock(id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindStock", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindStock indicates an expected call of FindStock.
func (mr *MockUserRepositoryMockRecorder) FindStock(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindStock", reflect.TypeOf((*MockUserRepository)(nil).FindStock), id)
}

// GetAddresses mocks base method.
func (m *MockUserRepository) GetAddresses(id int) ([]domain.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddresses", id)
	ret0, _ := ret[0].([]domain.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddresses indicates an expected call of GetAddresses.
func (mr *MockUserRepositoryMockRecorder) GetAddresses(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddresses", reflect.TypeOf((*MockUserRepository)(nil).GetAddresses), id)
}

// GetCartID mocks base method.
func (m *MockUserRepository) GetCartID(id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartID", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartID indicates an expected call of GetCartID.
func (mr *MockUserRepositoryMockRecorder) GetCartID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartID", reflect.TypeOf((*MockUserRepository)(nil).GetCartID), id)
}

// GetProductsInCart mocks base method.
func (m *MockUserRepository) GetProductsInCart(cart_id int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsInCart", cart_id)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsInCart indicates an expected call of GetProductsInCart.
func (mr *MockUserRepositoryMockRecorder) GetProductsInCart(cart_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsInCart", reflect.TypeOf((*MockUserRepository)(nil).GetProductsInCart), cart_id)
}

// GetUserByEmail mocks base method.
func (m *MockUserRepository) GetUserByEmail(email string) (models.UserDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(models.UserDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserRepositoryMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetUserByEmail), email)
}

// GetUserDetails mocks base method.
func (m *MockUserRepository) GetUserDetails(id int) (models.UserDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDetails", id)
	ret0, _ := ret[0].(models.UserDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDetails indicates an expected call of GetUserDetails.
func (mr *MockUserRepositoryMockRecorder) GetUserDetails(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDetails", reflect.TypeOf((*MockUserRepository)(nil).GetUserDetails), id)
}

// GetUserPassword mocks base method.
func (m *MockUserRepository) GetUserPassword(email string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPassword", email)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPassword indicates an expected call of GetUserPassword.
func (mr *MockUserRepositoryMockRecorder) GetUserPassword(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPassword", reflect.TypeOf((*MockUserRepository)(nil).GetUserPassword), email)
}

// RemoveFromCart mocks base method.
func (m *MockUserRepository) RemoveFromCart(cart, inventory int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromCart", cart, inventory)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromCart indicates an expected call of RemoveFromCart.
func (mr *MockUserRepositoryMockRecorder) RemoveFromCart(cart, inventory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromCart", reflect.TypeOf((*MockUserRepository)(nil).RemoveFromCart), cart, inventory)
}

// UpdateQuantity mocks base method.
func (m *MockUserRepository) UpdateQuantity(id, inv_id, qty int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuantity", id, inv_id, qty)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQuantity indicates an expected call of UpdateQuantity.
func (mr *MockUserRepositoryMockRecorder) UpdateQuantity(id, inv_id, qty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuantity", reflect.TypeOf((*MockUserRepository)(nil).UpdateQuantity), id, inv_id, qty)
}
