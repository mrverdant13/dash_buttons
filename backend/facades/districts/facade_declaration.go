package districts

import "github.com/mrverdant13/dash_buttons/backend/graph/model"

// Repo manages districts.
type Repo interface {
	Create(newDistrictData model.NewDistrict) (*model.District, error)
	GetByID(id string) (*model.District, error)
	GetAll() ([]*model.District, error)
}
