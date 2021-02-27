package provinces

import (
	"log"

	"github.com/mrverdant13/dash_buttons/backend/facades/departments"
	"github.com/mrverdant13/dash_buttons/backend/graph/model"
	"github.com/mrverdant13/dash_buttons/backend/internal/pkg/database/dbmodel"
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
	province := dbmodel.Province{
		Name:         newProvinceData.Name,
		DepartmentID: uint64(newProvinceData.DepartmentID),
	}

	result := r.gormDB.Create(
		&province,
	)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	// TODO: Test if "province" can be directly returned after conversion.
	return r.GetByID(uint64(province.ID))
}

func (r *repo) GetByID(id uint64) (*model.Province, error) {
	var province dbmodel.Province

	result := r.gormDB.First(&province, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	_province := province.ToGQL()
	return &_province, nil
}

func (r *repo) GetAll() ([]*model.Province, error) {
	return r.GetAllByDepartmentID(0)
}

func (r *repo) GetAllByDepartmentID(departmentID uint64) ([]*model.Province, error) {
	var provinces dbmodel.Provinces

	result := r.gormDB.Find(&provinces, &dbmodel.Province{DepartmentID: departmentID})
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	return provinces.ToGQL(), nil
}

func (r *repo) DeleteByID(id uint64) (*model.Province, error) {
	var province dbmodel.Province

	result := r.gormDB.Delete(&province, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	_province := province.ToGQL()
	return &_province, nil
}
