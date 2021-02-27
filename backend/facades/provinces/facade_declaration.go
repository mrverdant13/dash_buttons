package provinces

import "github.com/mrverdant13/dash_buttons/backend/graph/model"

// Repo manages provinces.
type Repo interface {
	Create(newProvinceData model.NewProvince) (*model.Province, error)
	GetByID(id string) (*model.Province, error)
	GetAll() ([]*model.Province, error)
}
