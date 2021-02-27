package districts

import (
	"log"
	"strconv"

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
	provinceID, err := strconv.ParseUint(newDistrictData.ProvinceID, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	district := District{
		Name:       newDistrictData.Name,
		ProvinceID: provinceID,
	}

	result := r.gormDB.Create(
		&district,
	)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	return r.GetByID(strconv.FormatInt(int64(district.ID), 10))
}

func (r *repo) GetByID(id string) (*model.District, error) {
	var district District

	result := r.gormDB.First(&district, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	province, err := r.provincesRepo.GetByID(strconv.FormatInt(int64(district.ProvinceID), 10))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	_district := model.District{
		ID:       id,
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
		province, err := r.provincesRepo.GetByID(strconv.FormatInt(int64(district.ProvinceID), 10))
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		_district := model.District{
			ID:       strconv.FormatInt(int64(district.ID), 10),
			Name:     district.Name,
			Province: province,
		}
		_districts = append(_districts, &_district)
	}

	return _districts, nil
}

func (r *repo) DeleteByID(id string) (*model.District, error) {
	var district District

	result := r.gormDB.Delete(&district, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	province, err := r.provincesRepo.GetByID(strconv.FormatInt(int64(district.ProvinceID), 10))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	_district := model.District{
		ID:       id,
		Name:     district.Name,
		Province: province,
	}

	return &_district, nil
}
