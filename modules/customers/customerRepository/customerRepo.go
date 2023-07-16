package customerRepository

import (
	"bookingtogo-test-be/models/customerModel"
	"bookingtogo-test-be/modules/customers"
	"database/sql"
	"errors"
)

type sqlRepository struct {
	Conn *sql.DB
}

func NewCustomerRepository(Conn *sql.DB) customers.CustomerRepository {
	return &sqlRepository{Conn}
}

func (sql *sqlRepository) GetCustomerByID(id uint64) (*customerModel.Customer, error) {
	query := `SELECT customer.cst_name, customer.cst_dob, customer.cst_phoneNum, customer.cst_email, 
	CONCAT(nationality.nationality_name, ' (', nationality.nationality_code, ')') AS country
	FROM customer
	LEFT JOIN nationality ON nationality.nationality_id = customer.nationality_id
	WHERE customer.cst_id = ?`

	stat, errPrepare := sql.Conn.Prepare(query)
	if errPrepare != nil {
		return nil, errors.New("GetCustomerByID errPrepare = " + errPrepare.Error())
	}
	defer stat.Close()

	var resp customerModel.Customer
	errQueryRow := stat.QueryRow(id).Scan(&resp.CstName, &resp.CstDob, &resp.CstPhoneNum, &resp.CstEmail, &resp.Country)
	if errQueryRow != nil {

		return nil, errors.New("GetCustomerByID errQueryRow = " + errQueryRow.Error())
	}

	return &resp, nil
}

func (config *sqlRepository) GetFamilyBYCustId(id uint64) (*[]customerModel.Family, error) {
	rows, err := config.Conn.Query(`SELECT fl_relation, fl_name, fl_dob FROM family_list WHERE cst_id = ?`, id)

	if err != nil {
		return nil, errors.New("GetFamilyBYCustId err = " + err.Error())
	}
	defer rows.Close()

	var data []customerModel.Family
	for rows.Next() {
		var each customerModel.Family

		err = rows.Scan(&each.FlRelation, &each.FlName, &each.FlDob)
		if err != nil {
			return nil, errors.New("GetFamilyBYCustId err = " + err.Error())
		}

		data = append(data, each)
	}

	return &data, nil

}
