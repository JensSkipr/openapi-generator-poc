package repositories

import (
	"github.com/nightborn-be/blink/skipr-test/app/contexts"
	repositories_base "github.com/nightborn-be/blink/skipr-test/app/repositories/base"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
)

type IExpenseRepository interface {
	repositories_base.IExpenseRepositoryBase
	GetAllExpensesWithoutParentExpenseId(context *contexts.Context, page *int, size *int, q *string) ([]entities.ExpenseEntity, error)
}
