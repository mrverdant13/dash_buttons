package districts

import (
	"log"

	"github.com/mrverdant13/dash_buttons/backend/facades/provinces"
	"github.com/mrverdant13/dash_buttons/backend/graph/model"
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

func (r *repo) Create(newDistrictData model.NewDistrict) (*model.District, error) {
	district := District{
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

	return r.GetByID(uint64(newDistrictData.ProvinceID))
}

func (r *repo) GetByID(id uint64) (*model.District, error) {
	var district District

	result := r.gormDB.First(&district, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	province, err := r.provincesRepo.GetByID(uint64(district.ProvinceID))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	_district := model.District{
		ID:       int64(id),
		Name:     district.Name,
		Province: province,
	}

	return &_district, nil
}

func (r *repo) GetAll() ([]*model.District, error) {
	var districts []*District

	result := r.gormDB.Find(&districts)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	var _districts []*model.District
	for _, district := range districts {
		province, err := r.provincesRepo.GetByID(uint64(district.ProvinceID))
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		_district := model.District{
			ID:       int64(district.ID),
			Name:     district.Name,
			Province: province,
		}
		_districts = append(_districts, &_district)
	}

	return _districts, nil
}

func (r *repo) DeleteByID(id uint64) (*model.District, error) {
	var district District

	result := r.gormDB.Delete(&district, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	province, err := r.provincesRepo.GetByID(uint64(district.ProvinceID))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	_district := model.District{
		ID:       int64(id),
		Name:     district.Name,
		Province: province,
	}

	return &_district, nil
}
