package mappers

import (
	"github.com/nightborn-be/blink/skipr-test/app/controllers/contracts"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
)

func ToExpenseDTO(entity entities.ExpenseEntity) (*contracts.ExpenseDTO, error) {
	output := &contracts.ExpenseDTO{
		Id:          entity.Id,
		CreatedAt:   entity.CreatedAt,
		ModifiedAt:  entity.ModifiedAt,
		ExpenseAt:   entity.ExpenseAt,
		ProgramId:   entity.ProgramId,
		TotalAmount: entity.TotalAmount,
	}

	// Categorization
	category, err := ToExpenseCategoryDTO(entity.Categorization)
	if err != nil {
		return nil, err
	}
	output.Categorization = *category

	// Refund status
	refundStatus, err := ToExpenseRefundStatusDTO(entity.RefundStatus)
	if err != nil {
		return nil, err
	}
	output.RefundStatus = *refundStatus

	// Review status
	reviewStatus, err := ToExpenseReviewStatusDTO(entity.ReviewStatus)
	if err != nil {
		return nil, err
	}
	output.ReviewStatus = *reviewStatus

	return output, nil
}

func ToExpenseEntity(dto contracts.ExpenseDTO) (*entities.ExpenseEntity, error) {
	output := &entities.ExpenseEntity{
		Id:          dto.Id,
		CreatedAt:   dto.CreatedAt,
		ModifiedAt:  dto.ModifiedAt,
		ExpenseAt:   dto.ExpenseAt,
		ProgramId:   dto.ProgramId,
		TotalAmount: dto.TotalAmount,
	}

	// Categorization
	category, err := ToExpenseCategoryEntity(dto.Categorization)
	if err != nil {
		return nil, err
	}
	output.Categorization = *category

	// Refund status
	refundStatus, err := ToExpenseRefundStatusEntity(dto.RefundStatus)
	if err != nil {
		return nil, err
	}
	output.RefundStatus = *refundStatus

	// Review status
	reviewStatus, err := ToExpenseReviewStatusEntity(dto.ReviewStatus)
	if err != nil {
		return nil, err
	}
	output.ReviewStatus = *reviewStatus

	return output, nil
}

func ToExpenseEntityFromCreateExpenseDTO(dto contracts.CreateExpenseDTO) (*entities.ExpenseEntity, error) {
	output := &entities.ExpenseEntity{
		ExpenseAt:   dto.ExpenseAt,
		ProgramId:   dto.ProgramId,
		TotalAmount: dto.TotalAmount,
	}

	// Categorization
	category, err := ToExpenseCategoryEntity(dto.Categorization)
	if err != nil {
		return nil, err
	}
	output.Categorization = *category

	return output, nil
}

func ToExpenseEntityFromUpdateExpenseDTO(dto contracts.UpdateExpenseDTO) (*entities.ExpenseEntity, error) {
	output := &entities.ExpenseEntity{
		ExpenseAt:   dto.ExpenseAt,
		ProgramId:   dto.ProgramId,
		TotalAmount: dto.TotalAmount,
	}

	// Categorization
	category, err := ToExpenseCategoryEntity(dto.Categorization)
	if err != nil {
		return nil, err
	}
	output.Categorization = *category

	// Refund status
	refundStatus, err := ToExpenseRefundStatusEntity(dto.RefundStatus)
	if err != nil {
		return nil, err
	}
	output.RefundStatus = *refundStatus

	// Review status
	reviewStatus, err := ToExpenseReviewStatusEntity(dto.ReviewStatus)
	if err != nil {
		return nil, err
	}
	output.ReviewStatus = *reviewStatus

	return output, nil
}

func ToExpenseDTOs(values []entities.ExpenseEntity) ([]contracts.ExpenseDTO, error) {
	output := make([]contracts.ExpenseDTO, 0, len(values))
	for _, value := range values {
		dto, err := ToExpenseDTO(value)
		if err != nil {
			return nil, err
		}
		output = append(output, *dto)
	}
	return output, nil
}
