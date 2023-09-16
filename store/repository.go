package store

import (
	"fmt"
	"path"

	"github.com/shoet/expense-manage-api-mock/entity"
	"github.com/shoet/expense-manage-api-mock/util"
)

type Repository struct{}

func (r *Repository) ListExpenses() (entity.Expenses, error) {
	data_file := "/data/expense.json"
	cwd, err := util.GetCurrentDirectory()
	if err != nil {
		return nil, fmt.Errorf("failed load Expenses: %v", err)
	}
	e := entity.Expenses{}
	if err := util.DeserializeFromJSON(path.Join(cwd, data_file), &e); err != nil {
		return nil, fmt.Errorf("failed Unmarshal Expenses: %v", err)
	}
	return e, nil
}
