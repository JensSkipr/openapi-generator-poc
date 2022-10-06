/* This file is auto-generated, manual edits in this file will be overwritten! */
package usecases_base

import (
"github.com/nightborn-be/blink/skipr-test/app/controllers/contracts"
"github.com/google/uuid"
"github.com/nightborn-be/blink/skipr-test/app/contexts"
)

type IExpenseUsecaseBase interface {
GetExpenses(context *contexts.Context) ([]contracts.ExpenseDTO, error)
CreateExpense(context *contexts.Context, createExpenseDTO contracts.CreateExpenseDTO) (*contracts.ExpenseDTO, error)
GetExpense(context *contexts.Context, expenseId uuid.UUID) (*contracts.ExpenseDTO, error)
UpdateExpense(context *contexts.Context, expenseId uuid.UUID, updateExpenseDTO contracts.UpdateExpenseDTO) (*contracts.ExpenseDTO, error)
}