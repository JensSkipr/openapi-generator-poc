/* This file is auto-generated, manual edits in this file will be overwritten! */
package repositories_base

import (
"github.com/nightborn-be/blink/skipr-test/app/contexts"
"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
"github.com/google/uuid"
)

type IExpenseLogRepositoryBase interface {
	GetAllExpenseLogs(context *contexts.Context, page *int, size *int, q *string) ([]entities.ExpenseLogEntity, error)
	CountAllExpenseLogs(context *contexts.Context, q *string) (*int64, error)
	GetExpenseLogById(context *contexts.Context, id uuid.UUID) (*entities.ExpenseLogEntity, error)
	AddExpenseLog(context *contexts.Context, expenseLogEntity entities.ExpenseLogEntity) (*entities.ExpenseLogEntity, error)
	ModifyExpenseLog(context *contexts.Context, expenseLogEntity entities.ExpenseLogEntity) (*entities.ExpenseLogEntity, error)
	
GetAllExpenseLogsByExpenseId(context *contexts.Context, id uuid.UUID, page *int, size *int, q *string) ([]entities.ExpenseLogEntity, error)
CountAllExpenseLogsByExpenseId(context *contexts.Context, id uuid.UUID, q *string) (*int64, error)
}
