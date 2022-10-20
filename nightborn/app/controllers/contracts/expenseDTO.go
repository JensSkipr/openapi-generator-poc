/* This file is auto-generated, manual edits in this file will be overwritten! */
package contracts

import (
"time"
"github.com/google/uuid"
)


type ExpenseDTO struct {
Categorization ExpenseCategory `json:"categorization"`
// The creation date of the expense
CreatedAt time.Time `json:"createdAt" faker:"utcTime"`
// The date of the expense
ExpenseAt time.Time `json:"expenseAt" faker:"utcTime"`
Id uuid.UUID `json:"id"`
// The modification date of the expense
ModifiedAt time.Time `json:"modifiedAt" faker:"utcTime"`
// The ID of the allocated program
ProgramId uuid.UUID `json:"programId"`
RefundStatus RefundStatus `json:"refundStatus"`
ReviewStatus ReviewStatus `json:"reviewStatus"`
// The total amount
TotalAmount int `json:"totalAmount"`
}