// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodel

type Department struct {
	ID        int64       `json:"id"`
	Name      string      `json:"name"`
	Provinces []*Province `json:"provinces"`
}

type District struct {
	ID         int64  `json:"id"`
	ProvinceID int64  `json:"provinceId"`
	Name       string `json:"name"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewDepartment struct {
	Name string `json:"name"`
}

type NewDistrict struct {
	Name       string `json:"name"`
	ProvinceID int64  `json:"provinceId"`
}

type NewProvince struct {
	Name         string `json:"name"`
	DepartmentID int64  `json:"departmentId"`
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Province struct {
	ID           int64       `json:"id"`
	DepartmentID int64       `json:"departmentId"`
	Name         string      `json:"name"`
	Districts    []*District `json:"districts"`
}

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}