package mappers

import (
	"errors"

	"github.com/nightborn-be/blink/skipr-test/app/controllers/contracts"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
)

func ToExpenseRefundStatusDTO(entity entities.RefundStatus) (*contracts.RefundStatus, error) {
	types := map[entities.RefundStatus]contracts.RefundStatus{
		entities.REFUND_STATUS_ACCEPTED: contracts.REFUND_STATUS_ACCEPTED,
		entities.REFUND_STATUS_PENDING:  contracts.REFUND_STATUS_PENDING,
		entities.REFUND_STATUS_REFUSED:  contracts.REFUND_STATUS_REFUSED,
	}

	result, ok := types[entity]
	if !ok {
		return nil, errors.New(UNABLE_TO_PARSE_ENUM + "expense_refund_status")
	}
	return &result, nil
}

func ToExpenseRefundStatusEntity(dto contracts.RefundStatus) (*entities.RefundStatus, error) {
	types := map[contracts.RefundStatus]entities.RefundStatus{
		contracts.REFUND_STATUS_ACCEPTED: entities.REFUND_STATUS_ACCEPTED,
		contracts.REFUND_STATUS_PENDING:  entities.REFUND_STATUS_PENDING,
		contracts.REFUND_STATUS_REFUSED:  entities.REFUND_STATUS_REFUSED,
	}

	result, ok := types[dto]
	if !ok {
		return nil, errors.New(UNABLE_TO_PARSE_ENUM + "expense_refund_status")
	}
	return &result, nil
}
