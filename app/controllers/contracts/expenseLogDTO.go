/* This file is auto-generated, manual edits in this file will be overwritten! */
package contracts

import (
"time"
"github.com/google/uuid"
)


type ExpenseLogDTO struct {
Action ExpenseLogAction `json:"action"`
Author string `json:"author"`
CreatedAt time.Time `json:"createdAt" faker:"utcTime"`
ExpenseId uuid.UUID `json:"expenseId"`
Field string `json:"field"`
Id uuid.UUID `json:"id"`
ModifiedAt time.Time `json:"modifiedAt" faker:"utcTime"`
NewValue string `json:"newValue"`
OldValue string `json:"oldValue"`
Role UserRole `json:"role"`
}