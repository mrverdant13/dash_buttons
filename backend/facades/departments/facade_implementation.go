package departments

import (
	"strconv"

	"github.com/mrverdant13/dash_buttons/backend/graph/model"
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

func (r *repo) Create(name string) (*model.Department, error) {
	department := Department{
		Name: name,
	}

	result := r.gormDB.Create(
		&department,
	)
	if result.Error != nil {
		return nil, result.Error
	}

	_department := model.Department{
		ID:   strconv.FormatInt(int64(department.ID), 10),
		Name: department.Name,
	}

	return &_department, nil
}

func (r *repo) GetAll() ([]*model.Department, error) {
	var departments []*Department

	result := r.gormDB.Find(&departments)
	if result.Error != nil {
		return nil, result.Error
	}

	var _departments []*model.Department
	for _, department := range departments {
		_department := model.Department{
			ID:   strconv.FormatInt(int64(department.ID), 10),
			Name: department.Name,
		}
		_departments = append(_departments, &_department)
	}

	return _departments, nil
}
