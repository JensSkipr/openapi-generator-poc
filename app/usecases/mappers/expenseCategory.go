package mappers

import (
	"errors"

	"github.com/nightborn-be/blink/skipr-test/app/controllers/contracts"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
)

func ToExpenseCategoryDTO(entity entities.ExpenseCategory) (*contracts.ExpenseCategory, error) {
	types := map[entities.ExpenseCategory]contracts.ExpenseCategory{
		entities.EXPENSE_CATEGORY_PRODUCT:  contracts.EXPENSE_CATEGORY_PRODUCT,
		entities.EXPENSE_CATEGORY_PROVIDER: contracts.EXPENSE_CATEGORY_PROVIDER,
		entities.EXPENSE_CATEGORY_SERVICE:  contracts.EXPENSE_CATEGORY_SERVICE,
	}

	result, ok := types[entity]
	if !ok {
		return nil, errors.New(UNABLE_TO_PARSE_ENUM + "expense_category")
	}
	return &result, nil
}

func ToExpenseCategoryEntity(dto contracts.ExpenseCategory) (*entities.ExpenseCategory, error) {
	types := map[contracts.ExpenseCategory]entities.ExpenseCategory{
		contracts.EXPENSE_CATEGORY_PRODUCT:  entities.EXPENSE_CATEGORY_PRODUCT,
		contracts.EXPENSE_CATEGORY_PROVIDER: entities.EXPENSE_CATEGORY_PROVIDER,
		contracts.EXPENSE_CATEGORY_SERVICE:  entities.EXPENSE_CATEGORY_SERVICE,
	}

	result, ok := types[dto]
	if !ok {
		return nil, errors.New(UNABLE_TO_PARSE_ENUM + "expense_category")
	}
	return &result, nil
}
