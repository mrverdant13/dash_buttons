package districts

import (
	"log"

	"github.com/mrverdant13/dash_buttons/backend/facades/provinces"
	"github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"
	"github.com/mrverdant13/dash_buttons/backend/internal/pkg/database/dbmodel"
	"gorm.io/gorm"
)

type repo struct {
	gormDB        *gorm.DB
	provincesRepo provinces.Repo
}

// NewRepo creates a new districts repo.
func NewRepo(
	gormDB *gorm.DB,
	provincesRepo provinces.Repo,
) Repo {
	return &repo{
		gormDB:        gormDB,
		provincesRepo: provincesRepo,
	}
}

func (r *repo) Create(newDistrictData gqlmodel.NewDistrict) (*gqlmodel.District, error) {
	district := dbmodel.District{
		Name:       newDistrictData.Name,
		ProvinceID: uint64(newDistrictData.ProvinceID),
	}

	result := r.gormDB.Create(
		&district,
	)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	// TODO: Test if "district" can be directly returned after conversion.
	return r.GetByID(uint64(newDistrictData.ProvinceID))
}

func (r *repo) GetByID(id uint64) (*gqlmodel.District, error) {
	var district dbmodel.District

	result := r.gormDB.First(&district, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	_district := district.ToGQL()
	return &_district, nil
}

func (r *repo) GetAll() ([]*gqlmodel.District, error) {
	return r.GetAllByProvinceID(0)
}

func (r *repo) GetAllByProvinceID(provinceID uint64) ([]*gqlmodel.District, error) {
	var districts dbmodel.Districts

	result := r.gormDB.Find(&districts, &dbmodel.District{ProvinceID: provinceID})
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	return districts.ToGQL(), nil
}

func (r *repo) DeleteByID(id uint64) (*gqlmodel.District, error) {
	var district dbmodel.District

	result := r.gormDB.Delete(&district, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	_district := district.ToGQL()
	return &_district, nil
}
