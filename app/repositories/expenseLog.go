package repositories

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nightborn-be/blink/skipr-test/app/contexts"
	"github.com/nightborn-be/blink/skipr-test/app/database/mappers"
	"github.com/nightborn-be/blink/skipr-test/app/database/models"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
)

func (repository ExpenseLogRepository) GetAllExpenseLogsByParentId(context *contexts.Context, id uuid.UUID, page *int, size *int, q *string, dateFrom *time.Time, dateTo *time.Time) ([]entities.ExpenseLogEntity, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	var expenseLogs []models.ExpenseLog

	query := db

	// Filtering
	query = query.
		Model(&models.ExpenseLog{}).
		Joins("JOIN expenses on expenses.parent_expense_id = ? and expenses.id = expense_logs.expense_id", id)

	// Temporality
	if dateFrom != nil && dateTo != nil {
		query = query.Where("expense_logs.created_at <= ? and expense_logs.created_at >= ?", dateTo, dateFrom)
	}

	// Paging
	if page != nil && size != nil {
		query = query.Limit(*size).Offset(*page * *size)
	}

	if err := query.Find(&expenseLogs).Error; err != nil {
		return nil, errors.New(UNABLE_TO_RETRIEVE_RESOURCE + "expense_logs")
	}

	expenseLogEntities := mappers.ToExpenseLogEntities(expenseLogs)

	return expenseLogEntities, nil
}
