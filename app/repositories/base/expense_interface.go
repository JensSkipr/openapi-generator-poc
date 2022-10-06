/* This file is auto-generated, manual edits in this file will be overwritten! */
package repositories_base

import (
"github.com/nightborn-be/blink/skipr-test/app/contexts"
"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
"github.com/google/uuid"
)

type IExpenseRepositoryBase interface {
	GetAllExpenses(context *contexts.Context, page *int, size *int, q *string) ([]entities.ExpenseEntity, error)
	CountAllExpenses(context *contexts.Context, q *string) (*int64, error)
	GetExpenseById(context *contexts.Context, id uuid.UUID) (*entities.ExpenseEntity, error)
	AddExpense(context *contexts.Context, expenseEntity entities.ExpenseEntity) (*entities.ExpenseEntity, error)
	ModifyExpense(context *contexts.Context, expenseEntity entities.ExpenseEntity) (*entities.ExpenseEntity, error)
	

}
