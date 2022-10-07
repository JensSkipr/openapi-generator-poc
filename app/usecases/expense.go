package usecases

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/nightborn-be/blink/skipr-test/app/contexts"
	"github.com/nightborn-be/blink/skipr-test/app/controllers/contracts"
	"github.com/nightborn-be/blink/skipr-test/app/gateways"
	"github.com/nightborn-be/blink/skipr-test/app/repositories"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/mappers"
	"github.com/nightborn-be/blink/skipr-test/app/utils"
	"github.com/samber/lo"
)

type ExpenseUsecase struct {
	repository *repositories.Repository
	gateway    *gateways.Gateway
}

func InitialiseExpenseUsecase(repository *repositories.Repository, gateway *gateways.Gateway) IExpenseUsecase {
	return ExpenseUsecase{
		repository: repository,
		gateway:    gateway,
	}
}

func (usecase ExpenseUsecase) GetExpenses(context *contexts.Context, page *int, q *string, size *int) ([]contracts.ExpenseDTO, error) {

	// Paging parsing
	page, size = utils.ConvertQueryPaging(page, size)

	// Get the expenses from the database
	expenses, err := usecase.repository.ExpenseRepository.GetAllExpensesWithoutParentExpenseId(context, page, size, q)
	if err != nil {
		return nil, err
	}

	// Map to DTOs
	expenseDTOs, err := mappers.ToExpenseDTOs(expenses)
	if err != nil {
		return nil, err
	}

	return expenseDTOs, nil
}

func (usecase ExpenseUsecase) CreateExpense(context *contexts.Context, createExpenseDTO contracts.CreateExpenseDTO) (*contracts.ExpenseDTO, error) {

	// Map expense to entity
	expense, err := mappers.ToExpenseEntityFromCreateExpenseDTO(createExpenseDTO)
	if err != nil {
		return nil, err
	}

	// Inject expense into DB
	insertedExpense, err := usecase.repository.ExpenseRepository.AddExpense(context, *expense)
	if err != nil {
		return nil, err
	}

	// Map expense to DTO
	expenseDTO, err := mappers.ToExpenseDTO(*insertedExpense)
	if err != nil {
		return nil, err
	}

	return expenseDTO, nil
}

func (usecase ExpenseUsecase) GetExpense(context *contexts.Context, id uuid.UUID) (*contracts.ExpenseDTO, error) {
	// Get the expenses from the database
	expense, err := usecase.repository.ExpenseRepository.GetExpenseById(context, id)
	if err != nil {
		return nil, err
	}

	// We don't let the user ask for an expense that has a parent
	if expense.ParentExpenseId != nil {
		return nil, errors.New(NOT_ALLOWED + "expense_has_parent")
	}

	// Map expense to DTO
	expenseDTO, err := mappers.ToExpenseDTO(*expense)
	if err != nil {
		return nil, err
	}

	return expenseDTO, nil
}

func (usecase ExpenseUsecase) UpdateExpense(context *contexts.Context, expenseId uuid.UUID, updateExpenseDTO contracts.UpdateExpenseDTO) (*contracts.ExpenseDTO, error) {
	// We map the updated expense
	updatedExpense, err := mappers.ToExpenseEntityFromUpdateExpenseDTO(updateExpenseDTO)
	if err != nil {
		return nil, err
	}

	// Get the expenses from the database
	existingExpense, err := usecase.repository.ExpenseRepository.GetExpenseById(context, expenseId)
	if err != nil {
		return nil, err
	}

	// We don't let the user ask for an expense that has a parent
	if existingExpense.ParentExpenseId != nil {
		return nil, errors.New(NOT_ALLOWED + "expense_has_parent")
	}

	// We will create a copy and link it to the master expense
	existingExpense.ParentExpenseId = &existingExpense.Id

	// Add the copy in the database
	existingExpense, err = usecase.repository.ExpenseRepository.AddExpense(context, *existingExpense)
	if err != nil {
		return nil, err
	}

	// Add back the expenseId to the Entity
	updatedExpense.Id = expenseId

	// Update expense in the DB
	updatedExpense, err = usecase.repository.ExpenseRepository.ModifyExpense(context, *updatedExpense)
	if err != nil {
		return nil, err
	}

	// Log every field changes
	err = usecase.logUpdateFields(context, entities.EXPENSE_LOG_ACTION_EDIT_EXPENSE, *existingExpense, *updatedExpense)
	if err != nil {
		return nil, err
	}

	// Map expense to DTO
	expenseDTO, err := mappers.ToExpenseDTO(*updatedExpense)
	if err != nil {
		return nil, err
	}

	return expenseDTO, nil
}

func (usecase ExpenseUsecase) GetExpenseLogs(context *contexts.Context, expenseId uuid.UUID, dateFrom *time.Time, dateTo *time.Time, page *int, q *string, size *int) ([]contracts.ExpenseLogDTO, error) {

	// Paging parsing
	page, size = utils.ConvertQueryPaging(page, size)

	// Check dateFrom should be before dateTo
	if dateFrom != nil && dateTo != nil {
		if dateFrom.After(*dateTo) {
			return nil, errors.New(INCORRECT_PARAMETERS + "dateFrom, dateTo")
		}
	}

	// Check expense exists
	expense, err := usecase.repository.ExpenseRepository.GetExpenseById(context, expenseId)
	if err != nil {
		return nil, err
	}

	// We don't let the user ask for an expense that has a parent
	if expense.ParentExpenseId != nil {
		return nil, errors.New(NOT_ALLOWED + "expense_has_parent")
	}

	// Get expense logs
	logs, err := usecase.repository.ExpenseLogRepository.GetAllExpenseLogsByParentId(context, expenseId, page, size, q, dateFrom, dateTo)
	if err != nil {
		return nil, err
	}

	// Map to DTOs
	logDTOs, err := mappers.ToExpenseLogDTOs(logs)
	if err != nil {
		return nil, err
	}

	return logDTOs, nil
}

func (usecase ExpenseUsecase) logUpdateFields(context *contexts.Context, action entities.ExpenseLogAction, existingExpense entities.ExpenseEntity, newExpense entities.ExpenseEntity) error {
	// Get the reflect value of the expenses
	existingExpenseValue := reflect.ValueOf(existingExpense)
	newExpenseValue := reflect.ValueOf(newExpense)

	for i := 0; i < newExpenseValue.NumField(); i++ {
		// Get each field
		fieldName := newExpenseValue.Type().Field(i).Name

		// Don't add log for the basic fields
		if lo.Contains([]string{"Id", "CreatedAt", "ModifiedAt", "ParentExpenseId"}, fieldName) {
			continue
		}

		existingField := existingExpenseValue.Field(i)
		newField := newExpenseValue.Field(i)

		// Check the values change
		if existingField.Interface() == newField.Interface() {
			continue
		}

		// Get old and new value
		var oldValue *string
		if existingField.Type().Kind() != reflect.Ptr || !existingField.IsNil() {
			oldValue = lo.ToPtr(fmt.Sprintf("%v", existingField))
		}
		var newValue *string
		if newField.Type().Kind() != reflect.Ptr || !newField.IsNil() {
			newValue = lo.ToPtr(fmt.Sprintf("%v", newField))
		}

		// Add a log if the value change
		log := entities.ExpenseLogEntity{
			Action: action,
			// Should be replaced by the user retrieved by the sub
			// Don't have this scope in my little repository...
			// --- START
			Author: "Maxime DENUIT",
			Role:   entities.USER_ROLE_EMPLOYEE,
			// --- END
			Field:     newExpenseValue.Type().Field(i).Name,
			ExpenseId: existingExpense.Id,
			OldValue:  oldValue,
			NewValue:  newValue,
		}
		_, err := usecase.repository.ExpenseLogRepository.AddExpenseLog(context, log)
		if err != nil {
			return err
		}
	}

	return nil
}
