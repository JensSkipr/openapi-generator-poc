package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/nightborn-be/blink/skipr-test/app/contexts"
	repositories_base "github.com/nightborn-be/blink/skipr-test/app/repositories/base"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
)

type IExpenseLogRepository interface {
	repositories_base.IExpenseLogRepositoryBase
	GetAllExpenseLogsByParentId(context *contexts.Context, id uuid.UUID, page *int, size *int, q *string, dateFrom *time.Time, dateTo *time.Time) ([]entities.ExpenseLogEntity, error)
}
