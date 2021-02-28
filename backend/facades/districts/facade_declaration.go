package districts

import "github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"

// Repo manages districts.
type Repo interface {
	Create(newDistrictData gqlmodel.NewDistrict) (*gqlmodel.District, error)
	GetByID(id uint64) (*gqlmodel.District, error)
	GetAll() ([]*gqlmodel.District, error)
	GetAllByProvinceID(provinceID uint64) ([]*gqlmodel.District, error)
	DeleteByID(id uint64) (*gqlmodel.District, error)
}
