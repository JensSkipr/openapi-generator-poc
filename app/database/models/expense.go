/* This file is auto-generated, manual edits in this file will be overwritten! */
package models

import (
"github.com/google/uuid"
"time"
"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
)

type Expense struct {
Id uuid.UUID `gorm:"not null;primaryKey"`
CreatedAt time.Time `gorm:"not null" faker:"utcTime"`
ModifiedAt time.Time `gorm:"not null" faker:"utcTime"`
Categorization entities.ExpenseCategory `gorm:"not null;default:SERVICE"`
ExpenseAt time.Time `gorm:"not null" faker:"utcTime"`
ProgramId uuid.UUID `gorm:"not null"`
RefundStatus entities.RefundStatus `gorm:"not null;default:PENDING"`
ReviewStatus entities.ReviewStatus `gorm:"not null;default:PENDING"`
TotalAmount int `gorm:"not null"`
}