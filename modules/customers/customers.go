package customers

import "bookingtogo-test-be/models/customerModel"

type CustomerRepository interface {
	GetCustomerByID(id uint64) (*customerModel.Customer, error)
	GetFamilyBYCustId(id uint64) (*[]customerModel.Family, error)
}

type CustomerUsecase interface {
	FindCustomerById(id uint64) (*customerModel.Response, error)
}
