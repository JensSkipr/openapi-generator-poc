/* This file is auto-generated, manual edits in this file will be overwritten! */
package entities

import (
"github.com/google/uuid"
"time"
)

type ExpenseLogEntity struct {
Id uuid.UUID
CreatedAt time.Time
ModifiedAt time.Time
Action ExpenseLogAction
Author string
ExpenseId uuid.UUID
Field string
NewValue *string
OldValue *string
Role UserRole
}