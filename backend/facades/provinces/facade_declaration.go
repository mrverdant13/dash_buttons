package provinces

import "github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"

// Repo manages provinces.
type Repo interface {
	Create(newProvinceData gqlmodel.NewProvince) (*gqlmodel.Province, error)
	GetByID(id uint64) (*gqlmodel.Province, error)
	GetAll() ([]*gqlmodel.Province, error)
	GetAllByDepartmentID(departmentID uint64) ([]*gqlmodel.Province, error)
	DeleteByID(id uint64) (*gqlmodel.Province, error)
}
