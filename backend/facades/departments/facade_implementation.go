package departments

import (
	"log"

	"github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"
	"github.com/mrverdant13/dash_buttons/backend/internal/pkg/database/dbmodel"
	"gorm.io/gorm"
)

type repo struct {
	gormDB *gorm.DB
}

// NewRepo creates a new departments repo.
func NewRepo(gormDB *gorm.DB) Repo {
	return &repo{
		gormDB: gormDB,
	}
}

func (r *repo) Create(newDepartmentData gqlmodel.NewDepartment) (*gqlmodel.Department, error) {
	department := dbmodel.Department{
		Name: newDepartmentData.Name,
	}

	result := r.gormDB.Create(
		&department,
	)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	// TODO: Test if "departmen" can be directly returned after conversion.
	return r.GetByID(uint64(department.ID))
}

func (r *repo) GetByID(id uint64) (*gqlmodel.Department, error) {
	var department dbmodel.Department

	result := r.gormDB.First(&department, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	_department := department.ToGQL()
	return &_department, nil
}

func (r *repo) GetAll() ([]*gqlmodel.Department, error) {
	var departments dbmodel.Departments

	result := r.gormDB.Find(&departments)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	return departments.ToGQL(), nil
}

func (r *repo) DeleteByID(id uint64) (*gqlmodel.Department, error) {
	var department dbmodel.Department

	result := r.gormDB.Delete(&department, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	_department := department.ToGQL()
	return &_department, nil
}
