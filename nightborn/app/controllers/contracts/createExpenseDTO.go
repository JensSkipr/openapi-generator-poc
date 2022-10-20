/* This file is auto-generated, manual edits in this file will be overwritten! */
package contracts

import (
"time"
"github.com/google/uuid"
)


type CreateExpenseDTO struct {
Categorization ExpenseCategory `json:"categorization"`
// The date of the expense
ExpenseAt time.Time `json:"expenseAt" faker:"utcTime"`
// The ID of the allocated program
ProgramId uuid.UUID `json:"programId"`
// The total amount
TotalAmount int `json:"totalAmount"`
}