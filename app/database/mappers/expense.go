/* This file is auto-generated, manual edits in this file will be overwritten! */
package mappers

import (
	"github.com/nightborn-be/blink/skipr-test/app/database/models"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
	"github.com/samber/lo"
)

func ToExpenseEntity(model models.Expense) entities.ExpenseEntity {
	return entities.ExpenseEntity{
Id: model.Id,
CreatedAt: model.CreatedAt,
ModifiedAt: model.ModifiedAt,
Categorization: model.Categorization,
ExpenseAt: model.ExpenseAt,
ProgramId: model.ProgramId,
RefundStatus: model.RefundStatus,
ReviewStatus: model.ReviewStatus,
TotalAmount: model.TotalAmount,
	}
}

func ToExpenseEntities(modelArray []models.Expense) []entities.ExpenseEntity {
	return lo.Map(modelArray, func(model models.Expense, _ int) entities.ExpenseEntity{
		return ToExpenseEntity(model)
	})
}

func ToExpense(entity entities.ExpenseEntity) models.Expense {
	return models.Expense{
Id: entity.Id,
CreatedAt: entity.CreatedAt,
ModifiedAt: entity.ModifiedAt,
Categorization: entity.Categorization,
ExpenseAt: entity.ExpenseAt,
ProgramId: entity.ProgramId,
RefundStatus: entity.RefundStatus,
ReviewStatus: entity.ReviewStatus,
TotalAmount: entity.TotalAmount,
	}
}

func ToExpenses(entityArray []entities.ExpenseEntity) []models.Expense {
	return lo.Map(entityArray, func(entity entities.ExpenseEntity, _ int) models.Expense {
		return ToExpense(entity)
	})
}
