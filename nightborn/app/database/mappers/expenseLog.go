/* This file is auto-generated, manual edits in this file will be overwritten! */
package mappers

import (
	"github.com/nightborn-be/blink/skipr-test/app/database/models"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
	"github.com/samber/lo"
)

func ToExpenseLogEntity(model models.ExpenseLog) entities.ExpenseLogEntity {
	return entities.ExpenseLogEntity{
Id: model.Id,
CreatedAt: model.CreatedAt,
ModifiedAt: model.ModifiedAt,
Action: model.Action,
Author: model.Author,
ExpenseId: model.ExpenseId,
Field: model.Field,
NewValue: model.NewValue,
OldValue: model.OldValue,
Role: model.Role,
	}
}

func ToExpenseLogEntities(modelArray []models.ExpenseLog) []entities.ExpenseLogEntity {
	return lo.Map(modelArray, func(model models.ExpenseLog, _ int) entities.ExpenseLogEntity{
		return ToExpenseLogEntity(model)
	})
}

func ToExpenseLog(entity entities.ExpenseLogEntity) models.ExpenseLog {
	return models.ExpenseLog{
Id: entity.Id,
CreatedAt: entity.CreatedAt,
ModifiedAt: entity.ModifiedAt,
Action: entity.Action,
Author: entity.Author,
ExpenseId: entity.ExpenseId,
Field: entity.Field,
NewValue: entity.NewValue,
OldValue: entity.OldValue,
Role: entity.Role,
	}
}

func ToExpenseLogs(entityArray []entities.ExpenseLogEntity) []models.ExpenseLog {
	return lo.Map(entityArray, func(entity entities.ExpenseLogEntity, _ int) models.ExpenseLog {
		return ToExpenseLog(entity)
	})
}
