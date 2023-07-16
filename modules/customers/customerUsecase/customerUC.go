package customerUsecase

import (
	"bookingtogo-test-be/models/customerModel"
	"bookingtogo-test-be/modules/customers"
)

type customerUsecase struct {
	CustomerRepo customers.CustomerRepository
}

func NewCustomerUsecase(CustomerRepo customers.CustomerRepository) customers.CustomerUsecase {
	return &customerUsecase{CustomerRepo}
}

func (custUC *customerUsecase) FindCustomerById(id uint64) (*customerModel.Response, error) {
	customer, err := custUC.CustomerRepo.GetCustomerByID(id)

	if err != nil {
		return nil, err
	}

	family, errFamily := custUC.CustomerRepo.GetFamilyBYCustId(id)

	if err != nil {
		return nil, errFamily
	}

	resp := customerModel.Response{
		CstName:     customer.CstName,
		CstDob:      customer.CstDob,
		CstPhoneNum: customer.CstPhoneNum,
		CstEmail:    customer.CstEmail,
		Country:     customer.Country,
		Family:      *family,
	}

	return &resp, nil
}
