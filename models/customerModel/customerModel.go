package customerModel

import "time"

type Customer struct {
	CstName     string    `json:"nama"`
	CstDob      time.Time `json:"tanggal_lahir"`
	CstPhoneNum string    `json:"telepon"`
	CstEmail    string    `json:"email"`
	Country     string    `json:"kewarganegaraan"`
}

type Family struct {
	FlRelation string    `json:"hubungan"`
	FlName     string    `json:"nama"`
	FlDob      time.Time `json:"tanggal_lahir"`
}

type Response struct {
	CstName     string    `json:"nama"`
	CstDob      time.Time `json:"tanggal_lahir"`
	CstPhoneNum string    `json:"telepon"`
	CstEmail    string    `json:"email"`
	Country     string    `json:"kewarganegaraan"`
	Family      []Family  `json:"keluarga"`
}
