package mappers

import (
	"errors"

	"github.com/nightborn-be/blink/skipr-test/app/controllers/contracts"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
)

func ToExpenseLogActionDTO(entity entities.ExpenseLogAction) (*contracts.ExpenseLogAction, error) {
	types := map[entities.ExpenseLogAction]contracts.ExpenseLogAction{
		entities.EXPENSE_LOG_ACTION_EDIT_EXPENSE: contracts.EXPENSE_LOG_ACTION_EDIT_EXPENSE,
	}

	result, ok := types[entity]
	if !ok {
		return nil, errors.New(UNABLE_TO_PARSE_ENUM + "expense_log_action")
	}
	return &result, nil
}
