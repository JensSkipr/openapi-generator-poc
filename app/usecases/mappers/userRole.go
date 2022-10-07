package mappers

import (
	"errors"

	"github.com/nightborn-be/blink/skipr-test/app/controllers/contracts"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
)

func ToUserRoleDTO(entity entities.UserRole) (*contracts.UserRole, error) {
	types := map[entities.UserRole]contracts.UserRole{
		entities.USER_ROLE_ADMIN:    contracts.USER_ROLE_ADMIN,
		entities.USER_ROLE_EMPLOYEE: contracts.USER_ROLE_EMPLOYEE,
		entities.USER_ROLE_OPERATOR: contracts.USER_ROLE_OPERATOR,
		entities.USER_ROLE_REVIEWER: contracts.USER_ROLE_REVIEWER,
	}

	result, ok := types[entity]
	if !ok {
		return nil, errors.New(UNABLE_TO_PARSE_ENUM + "user_role")
	}
	return &result, nil
}
