// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodel

// Department data.
type Department struct {
	// Department ID.
	ID int64 `json:"id"`
	// Department name.
	Name string `json:"name"`
	// Department provinces. Provinces located in the department territory.
	Provinces []*Province `json:"provinces"`
}

// District data.
type District struct {
	// District ID.
	ID int64 `json:"id"`
	// Parent province ID.
	ProvinceID int64 `json:"provinceId"`
	// District name.
	Name string `json:"name"`
}

// Login credentials.
type Login struct {
	// User email
	Email string `json:"email"`
	// User password
	Password string `json:"password"`
}

// New department data.
type NewDepartment struct {
	// New department name.
	Name string `json:"name"`
}

// New district data.
type NewDistrict struct {
	// New district name.
	Name string `json:"name"`
	// New district parent department ID.
	ProvinceID int64 `json:"provinceId"`
}

// New province data.
type NewProvince struct {
	// New province name.
	Name string `json:"name"`
	// New province parent department ID.
	DepartmentID int64 `json:"departmentId"`
}

// New user data.
type NewUser struct {
	// New user email.
	Email string `json:"email"`
	// New user password.
	Password string `json:"password"`
	// New user admin indicator.
	IsAdmin *bool `json:"isAdmin"`
}

// Province data.
type Province struct {
	// Province ID.
	ID int64 `json:"id"`
	// Parent department ID.
	DepartmentID int64 `json:"departmentId"`
	// Province name
	Name string `json:"name"`
	// Province districts. Districts located in the province territory.
	Districts []*District `json:"districts"`
}

// User data.
type User struct {
	// User ID.
	ID int64 `json:"id"`
	// User email.
	Email string `json:"email"`
	// User admin indicator.
	IsAdmin bool `json:"isAdmin"`
}
