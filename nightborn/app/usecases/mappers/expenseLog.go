package mappers

import (
	"github.com/nightborn-be/blink/skipr-test/app/controllers/contracts"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
	"github.com/samber/lo"
)

func ToExpenseLogDTO(entity entities.ExpenseLogEntity) (*contracts.ExpenseLogDTO, error) {
	output := &contracts.ExpenseLogDTO{
		Id:         entity.Id,
		CreatedAt:  entity.CreatedAt,
		ModifiedAt: entity.ModifiedAt,
		Author:     entity.Author,
		ExpenseId:  entity.ExpenseId,
		Field:      entity.Field,
		NewValue:   lo.FromPtr(entity.NewValue),
		OldValue:   lo.FromPtr(entity.OldValue),
	}

	// Action
	action, err := ToExpenseLogActionDTO(entity.Action)
	if err != nil {
		return nil, err
	}
	output.Action = *action

	// Role
	role, err := ToUserRoleDTO(entity.Role)
	if err != nil {
		return nil, err
	}
	output.Role = *role

	return output, nil
}

func ToExpenseLogDTOs(values []entities.ExpenseLogEntity) ([]contracts.ExpenseLogDTO, error) {
	output := make([]contracts.ExpenseLogDTO, 0, len(values))
	for _, value := range values {
		dto, err := ToExpenseLogDTO(value)
		if err != nil {
			return nil, err
		}
		output = append(output, *dto)
	}
	return output, nil
}
