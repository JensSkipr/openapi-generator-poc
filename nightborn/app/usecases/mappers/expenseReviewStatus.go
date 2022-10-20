package mappers

import (
	"errors"

	"github.com/nightborn-be/blink/skipr-test/app/controllers/contracts"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
)

func ToExpenseReviewStatusDTO(entity entities.ReviewStatus) (*contracts.ReviewStatus, error) {
	types := map[entities.ReviewStatus]contracts.ReviewStatus{
		entities.REVIEW_STATUS_APPROVED:      contracts.REVIEW_STATUS_APPROVED,
		entities.REVIEW_STATUS_INFO_REQUIRED: contracts.REVIEW_STATUS_INFO_REQUIRED,
		entities.REVIEW_STATUS_PENDING:       contracts.REVIEW_STATUS_PENDING,
		entities.REVIEW_STATUS_REFUSED:       contracts.REVIEW_STATUS_REFUSED,
	}

	result, ok := types[entity]
	if !ok {
		return nil, errors.New(UNABLE_TO_PARSE_ENUM + "expense_review_status")
	}
	return &result, nil
}

func ToExpenseReviewStatusEntity(dto contracts.ReviewStatus) (*entities.ReviewStatus, error) {
	types := map[contracts.ReviewStatus]entities.ReviewStatus{
		contracts.REVIEW_STATUS_APPROVED:      entities.REVIEW_STATUS_APPROVED,
		contracts.REVIEW_STATUS_INFO_REQUIRED: entities.REVIEW_STATUS_INFO_REQUIRED,
		contracts.REVIEW_STATUS_PENDING:       entities.REVIEW_STATUS_PENDING,
		contracts.REVIEW_STATUS_REFUSED:       entities.REVIEW_STATUS_REFUSED,
	}

	result, ok := types[dto]
	if !ok {
		return nil, errors.New(UNABLE_TO_PARSE_ENUM + "expense_review_status")
	}
	return &result, nil
}
