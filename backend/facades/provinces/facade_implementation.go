package provinces

import (
	"log"
	"strconv"

	"github.com/mrverdant13/dash_buttons/backend/facades/departments"
	"github.com/mrverdant13/dash_buttons/backend/graph/model"
	"gorm.io/gorm"
)

type repo struct {
	gormDB          *gorm.DB
	departmentsRepo departments.Repo
}

// NewRepo creates a new provinces repo.
func NewRepo(
	gormDB *gorm.DB,
	departmentsRepo departments.Repo,
) Repo {
	return &repo{
		gormDB:          gormDB,
		departmentsRepo: departmentsRepo,
	}
}

func (r *repo) Create(newProvinceData model.NewProvince) (*model.Province, error) {
	departmentID, err := strconv.ParseUint(newProvinceData.DepartmentID, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	province := Province{
		Name:         newProvinceData.Name,
		DepartmentID: departmentID,
	}

	result := r.gormDB.Create(
		&province,
	)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	return r.GetByID(strconv.FormatInt(int64(province.ID), 10))
}

func (r *repo) GetByID(id string) (*model.Province, error) {
	var province Province

	result := r.gormDB.First(&province, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	department, err := r.departmentsRepo.GetByID(strconv.FormatInt(int64(province.DepartmentID), 10))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	_province := model.Province{
		ID:         id,
		Name:       province.Name,
		Department: department,
	}

	return &_province, nil
}

func (r *repo) GetAll() ([]*model.Province, error) {
	var provinces []*Province

	result := r.gormDB.Find(&provinces)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	var _provinces []*model.Province
	for _, province := range provinces {
		department, err := r.departmentsRepo.GetByID(strconv.FormatInt(int64(province.DepartmentID), 10))
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		_province := model.Province{
			ID:         strconv.FormatInt(int64(province.ID), 10),
			Name:       province.Name,
			Department: department,
		}
		_provinces = append(_provinces, &_province)
	}

	return _provinces, nil
}

func (r *repo) DeleteByID(id string) (*model.Province, error) {
	var province Province

	result := r.gormDB.Delete(&province, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	department, err := r.departmentsRepo.GetByID(strconv.FormatInt(int64(province.DepartmentID), 10))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	_province := model.Province{
		ID:         id,
		Name:       province.Name,
		Department: department,
	}

	return &_province, nil
}
