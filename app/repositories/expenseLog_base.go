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

type ExpenseLogRepository struct {
	getDB func(testId *string) (*gorm.DB, error)
}

func InitialiseExpenseLogRepository(getDB func(testId *string) (*gorm.DB, error)) IExpenseLogRepository {
	return ExpenseLogRepository{
		getDB: getDB,
	}
}

/*
	Get all expenseLogs

	page - The page number
	size - The size of the page
	q - The search query
*/
func (repository ExpenseLogRepository) GetAllExpenseLogs(context *contexts.Context, page *int, size *int, q *string) ([]entities.ExpenseLogEntity, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	var expenseLogs []models.ExpenseLog

	query := db

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

/*
	Count the number of expenseLogs
*/
func (repository ExpenseLogRepository) CountAllExpenseLogs(context *contexts.Context, q *string) (*int64, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	var count int64

	query := db.Model(&models.ExpenseLog{})

	

	if err := query.Count(&count).Error; err != nil {
		return nil, errors.New(UNABLE_TO_RETRIEVE_RESOURCE + "expense_logs")
	}

	return &count, nil
}

/*
	Get a expenseLog by id

	id - The id of the expenseLog
*/
func (repository ExpenseLogRepository) GetExpenseLogById(context *contexts.Context, id uuid.UUID) (*entities.ExpenseLogEntity, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	var expenseLog models.ExpenseLog

	query := db.
		// Get correct expenseLog
		Where("id = ?", id)

	if err := query.Take(&expenseLog).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(UNABLE_TO_FIND_RESOURCE + "expense_log")
		}
		return nil, errors.New(UNABLE_TO_RETRIEVE_RESOURCE + "expense_log")
	}

	expenseLogEntity := mappers.ToExpenseLogEntity(expenseLog)

	return &expenseLogEntity, nil
}

/*
	Add a new expenseLog

	expenseLogEntity - The expenseLog to create
*/
func (repository ExpenseLogRepository) AddExpenseLog(context *contexts.Context, expenseLogEntity entities.ExpenseLogEntity) (*entities.ExpenseLogEntity, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	expenseLog := mappers.ToExpenseLog(expenseLogEntity)

	expenseLog.Id = uuid.New()
	now := time.Now().UTC()
expenseLog.CreatedAt = now
expenseLog.ModifiedAt = now

	if err := db.Create(&expenseLog).Error; err != nil {
		return nil, errors.New(UNABLE_TO_CREATE_RESOURCE + "expenseLog")
	}

	expenseLogEntity = mappers.ToExpenseLogEntity(expenseLog)

	return &expenseLogEntity, nil
}

/*
	Modify a expenseLog

	expenseLogEntity - The new expenseLog data
*/
func (repository ExpenseLogRepository) ModifyExpenseLog(context *contexts.Context, expenseLogEntity entities.ExpenseLogEntity) (*entities.ExpenseLogEntity, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	expenseLog := mappers.ToExpenseLog(expenseLogEntity)
	expenseLog.ModifiedAt = time.Now().UTC()

	query := db.Model(&models.ExpenseLog{}).
		// Get correct expenseLog
		Where("id = ?", expenseLog.Id)

	// The fields to update
	newData := map[string]interface{}{
"modified_at": expenseLog.ModifiedAt,
"action": expenseLog.Action,
"author": expenseLog.Author,
"field": expenseLog.Field,
"new_value": expenseLog.NewValue,
"old_value": expenseLog.OldValue,
"role": expenseLog.Role,
	}

	if err := query.Updates(&newData).Error; err != nil {
		return nil, errors.New(UNABLE_TO_UPDATE_RESOURCE + "expense_log")
	}

	expenseLogEntity = mappers.ToExpenseLogEntity(expenseLog)

	return &expenseLogEntity, nil
}



/*
	Get all expenseLogs by expenseId

	ExpenseId - The expenseId
	page - The page number
	size - The size of the page
	q - The search query
*/
func (repository ExpenseLogRepository) GetAllExpenseLogsByExpenseId(context *contexts.Context, expenseId uuid.UUID, page *int, size *int, q *string) ([]entities.ExpenseLogEntity, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	var expenseLogs []models.ExpenseLog

	query := db.
		// Get correct ExpenseLogs
		Where("expense_id = ?", expenseId)

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

/*
	Count the number of expenseLogs

	ExpenseId - The expenseId
*/
func (repository ExpenseLogRepository) CountAllExpenseLogsByExpenseId(context *contexts.Context, expenseId uuid.UUID, q *string) (*int64, error) {
	// Get database
	db, err := repository.getDB(context.TestId)
	if err != nil {
		return nil, err
	}

	var count int64

	query := db.Model(&models.ExpenseLog{}).
		// Get correct projects
		Where("expense_id = ?", expenseId)

	

	if err := query.Count(&count).Error; err != nil {
		return nil, errors.New(UNABLE_TO_RETRIEVE_RESOURCE + "expense_logs")
	}

	return &count, nil
}