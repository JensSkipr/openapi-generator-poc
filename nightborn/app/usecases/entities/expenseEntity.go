/* This file is auto-generated, manual edits in this file will be overwritten! */
package entities

import (
"github.com/google/uuid"
"time"
)

type ExpenseEntity struct {
Id uuid.UUID
CreatedAt time.Time
ModifiedAt time.Time
Categorization ExpenseCategory
ExpenseAt time.Time
ParentExpenseId *uuid.UUID
ProgramId uuid.UUID
RefundStatus RefundStatus
ReviewStatus ReviewStatus
TotalAmount int
}