/* This file is auto-generated, manual edits in this file will be overwritten! */
package repositories

import (
	"gorm.io/gorm"
"github.com/nightborn-be/blink/skipr-test/app/contexts"
"errors"
"github.com/nightborn-be/blink/skipr-test/app/database/mappers"
"github.com/nightborn-be/blink/skipr-test/app/database/models"
"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
"github.com/google/uuid"
"time"
)

type ExpenseRepository struct {
	getDB func(testId *string) (*gorm.DB, error)
}

func InitialiseExpenseRepository(getDB func(testId *string) (*gorm.DB, error)) IExpenseRepository {
	return ExpenseRepository{
		getDB: getDB,
	}
}

/*
	Get all expenses

	page - The page number
	size - The size of the page
	q - The search query
*/
func (repository ExpenseRepository) GetAllExpenses(context *contexts.Context, page *int, size *int, q *string) ([]entities.ExpenseEntity, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	var expenses []models.Expense

	query := db

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

/*
	Count the number of expenses
*/
func (repository ExpenseRepository) CountAllExpenses(context *contexts.Context, q *string) (*int64, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	var count int64

	query := db.Model(&models.Expense{})

	

	if err := query.Count(&count).Error; err != nil {
		return nil, errors.New(UNABLE_TO_RETRIEVE_RESOURCE + "expenses")
	}

	return &count, nil
}

/*
	Get a expense by id

	id - The id of the expense
*/
func (repository ExpenseRepository) GetExpenseById(context *contexts.Context, id uuid.UUID) (*entities.ExpenseEntity, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	var expense models.Expense

	query := db.
		// Get correct expense
		Where("id = ?", id)

	if err := query.Take(&expense).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(UNABLE_TO_FIND_RESOURCE + "expense")
		}
		return nil, errors.New(UNABLE_TO_RETRIEVE_RESOURCE + "expense")
	}

	expenseEntity := mappers.ToExpenseEntity(expense)

	return &expenseEntity, nil
}

/*
	Add a new expense

	expenseEntity - The expense to create
*/
func (repository ExpenseRepository) AddExpense(context *contexts.Context, expenseEntity entities.ExpenseEntity) (*entities.ExpenseEntity, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	expense := mappers.ToExpense(expenseEntity)

	expense.Id = uuid.New()
	now := time.Now().UTC()
expense.CreatedAt = now
expense.ModifiedAt = now

	if err := db.Create(&expense).Error; err != nil {
		return nil, errors.New(UNABLE_TO_CREATE_RESOURCE + "expense")
	}

	expenseEntity = mappers.ToExpenseEntity(expense)

	return &expenseEntity, nil
}

/*
	Modify a expense

	expenseEntity - The new expense data
*/
func (repository ExpenseRepository) ModifyExpense(context *contexts.Context, expenseEntity entities.ExpenseEntity) (*entities.ExpenseEntity, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	expense := mappers.ToExpense(expenseEntity)
	expense.ModifiedAt = time.Now().UTC()

	query := db.Model(&models.Expense{}).
		// Get correct expense
		Where("id = ?", expense.Id)

	// The fields to update
	newData := map[string]interface{}{
"modified_at": expense.ModifiedAt,
"categorization": expense.Categorization,
"expense_at": expense.ExpenseAt,
"program_id": expense.ProgramId,
"refund_status": expense.RefundStatus,
"review_status": expense.ReviewStatus,
"total_amount": expense.TotalAmount,
	}

	if err := query.Updates(&newData).Error; err != nil {
		return nil, errors.New(UNABLE_TO_UPDATE_RESOURCE + "expense")
	}

	expenseEntity = mappers.ToExpenseEntity(expense)

	return &expenseEntity, nil
}



/*
	Get all expenses by parentExpenseId

	ParentExpenseId - The parentExpenseId
	page - The page number
	size - The size of the page
	q - The search query
*/
func (repository ExpenseRepository) GetAllExpensesByParentExpenseId(context *contexts.Context, parentExpenseId uuid.UUID, page *int, size *int, q *string) ([]entities.ExpenseEntity, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	var expenses []models.Expense

	query := db.
		// Get correct Expenses
		Where("parent_expense_id = ?", parentExpenseId)

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

/*
	Count the number of expenses

	ParentExpenseId - The parentExpenseId
*/
func (repository ExpenseRepository) CountAllExpensesByParentExpenseId(context *contexts.Context, parentExpenseId uuid.UUID, q *string) (*int64, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	var count int64

	query := db.Model(&models.Expense{}).
		// Get correct projects
		Where("parent_expense_id = ?", parentExpenseId)

	

	if err := query.Count(&count).Error; err != nil {
		return nil, errors.New(UNABLE_TO_RETRIEVE_RESOURCE + "expenses")
	}

	return &count, nil
}