package repositories

import (
	"errors"

	"github.com/nightborn-be/blink/skipr-test/app/contexts"
	"github.com/nightborn-be/blink/skipr-test/app/database/mappers"
	"github.com/nightborn-be/blink/skipr-test/app/database/models"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
)

/*
Get all expenses without a parent AKA the masters

page - The page number
size - The size of the page
q - The search query
*/
func (repository ExpenseRepository) GetAllExpensesWithoutParentExpenseId(context *contexts.Context, page *int, size *int, q *string) ([]entities.ExpenseEntity, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	var expenses []models.Expense

	query := db

	// Filtering
	query = query.Where("parent_expense_id is null")

	// Paging
	if page != nil && size != nil {
		query = query.Limit(*size).Offset(*page * *size)
	}

	if err := query.Find(&expenses).Error; err != nil {
		return nil, errors.New(UNABLE_TO_RETRIEVE_RESOURCE + "expenses")
	}

	expenseEntities := mappers.ToExpenseEntities(expenses)

	return expenseEntities, nil
}
