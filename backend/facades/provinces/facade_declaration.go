package provinces

import "github.com/mrverdant13/dash_buttons/backend/graph/model"

// Repo manages provinces.
type Repo interface {
	Create(newProvinceData model.NewProvince) (*model.Province, error)
	GetByID(id uint64) (*model.Province, error)
	GetAll() ([]*model.Province, error)
	GetAllByDepartmentID(departmentID uint64) ([]*model.Province, error)
	DeleteByID(id uint64) (*model.Province, error)
}
