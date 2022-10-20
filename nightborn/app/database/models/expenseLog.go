/* This file is auto-generated, manual edits in this file will be overwritten! */
package models

import (
"github.com/google/uuid"
"time"
"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
)

type ExpenseLog struct {
Id uuid.UUID `gorm:"not null;primaryKey"`
CreatedAt time.Time `gorm:"not null" faker:"utcTime"`
ModifiedAt time.Time `gorm:"not null" faker:"utcTime"`
Action entities.ExpenseLogAction `gorm:"not null;default:EDIT_EXPENSE"`
Author string `gorm:"not null"`
ExpenseId uuid.UUID `gorm:"not null"`
Field string `gorm:"not null"`
NewValue *string
OldValue *string
Role entities.UserRole `gorm:"not null;default:EMPLOYEE"`

// Foreign keys
ExpenseFK *Expense `gorm:"foreignKey:expense_id;references:id" faker:"-"`
}