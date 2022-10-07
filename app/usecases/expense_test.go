package usecases

import (
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/nightborn-be/blink/skipr-test/app/controllers/contracts"
	"github.com/nightborn-be/blink/skipr-test/app/database/models"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/entities"
	"github.com/nightborn-be/blink/skipr-test/app/usecases/mappers"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func setupExpenseMock(t *testing.T) *models.Expense {

	var expense models.Expense
	err := faker.FakeData(&expense)
	if err != nil {
		t.Fatal(err)
	}
	expense.ParentExpenseId = nil
	expense.Categorization = entities.EXPENSE_CATEGORY_PRODUCT
	expense.RefundStatus = entities.REFUND_STATUS_ACCEPTED
	expense.ReviewStatus = entities.REVIEW_STATUS_APPROVED

	return &expense
}

func setupExpenseLogMock(t *testing.T) *models.ExpenseLog {

	var expenseLog models.ExpenseLog
	err := faker.FakeData(&expenseLog)

	if err != nil {
		t.Fatal(err)
	}
	expenseLog.Action = entities.EXPENSE_LOG_ACTION_EDIT_EXPENSE
	expenseLog.Role = entities.USER_ROLE_EMPLOYEE

	return &expenseLog
}

func Test_GetExpenseLogs_Success(t *testing.T) {
	usecase, db, err := setupTestUsecase()
	if err != nil {
		t.Fatal(err)
	}
	context := createContext("")

	expense1 := setupExpenseMock(t)
	if err := db.Create(&expense1).Error; err != nil {
		t.Fatal(err)
	}

	expense2 := setupExpenseMock(t)
	expense2.ParentExpenseId = &expense1.Id
	if err := db.Create(&expense2).Error; err != nil {
		t.Fatal(err)
	}

	expense3 := setupExpenseMock(t)
	expense3.ParentExpenseId = &expense1.Id
	if err := db.Create(&expense3).Error; err != nil {
		t.Fatal(err)
	}

	log1 := setupExpenseLogMock(t)
	log1.ExpenseId = expense2.Id
	if err := db.Create(&log1).Error; err != nil {
		t.Fatal(err)
	}

	log2 := setupExpenseLogMock(t)
	log2.ExpenseId = expense2.Id
	if err := db.Create(&log2).Error; err != nil {
		t.Fatal(err)
	}

	log3 := setupExpenseLogMock(t)
	log3.ExpenseId = expense2.Id
	db.Create(&log3)

	log4 := setupExpenseLogMock(t)
	log4.ExpenseId = expense3.Id
	db.Create(&log4)

	log5 := setupExpenseLogMock(t)
	log5.ExpenseId = expense3.Id
	db.Create(&log5)

	log6 := setupExpenseLogMock(t)
	log6.ExpenseId = expense3.Id
	db.Create(&log6)

	result, err := usecase.ExpenseUsecase.GetExpenseLogs(&context, expense1.Id, nil, nil, lo.ToPtr(0), nil, lo.ToPtr(10))

	assert.Nil(t, err)
	assert.Len(t, result, 6)
}

func Test_GetExpenseLogs_DateFromTo_Success(t *testing.T) {
	usecase, db, err := setupTestUsecase()
	if err != nil {
		t.Fatal(err)
	}
	context := createContext("")

	expense1 := setupExpenseMock(t)
	expense2 := setupExpenseMock(t)
	expense2.ParentExpenseId = &expense1.Id

	timeConfig := time.Now().AddDate(0, 0, -7)
	expense1.CreatedAt = timeConfig
	expense1.ModifiedAt = timeConfig
	if err := db.Create(&expense1).Error; err != nil {
		t.Fatal(err)
	}

	expense2.CreatedAt = timeConfig
	expense2.ModifiedAt = timeConfig
	if err := db.Create(&expense2).Error; err != nil {
		t.Fatal(err)
	}

	log1 := setupExpenseLogMock(t)
	log1.ExpenseId = expense2.Id
	log1.CreatedAt = timeConfig
	log1.ModifiedAt = timeConfig
	if err := db.Create(&log1).Error; err != nil {
		t.Fatal(err)
	}

	log2 := setupExpenseLogMock(t)
	log2.ExpenseId = expense2.Id
	timeConfig = time.Now().AddDate(0, 0, -6)
	log2.CreatedAt = timeConfig
	log2.ModifiedAt = timeConfig
	if err := db.Create(&log2).Error; err != nil {
		t.Fatal(err)
	}

	log3 := setupExpenseLogMock(t)
	log3.ExpenseId = expense2.Id
	timeConfig = time.Now().AddDate(0, 0, -5)
	log3.CreatedAt = timeConfig
	log3.ModifiedAt = timeConfig
	if err := db.Create(&log3).Error; err != nil {
		t.Fatal(err)
	}

	log4 := setupExpenseLogMock(t)
	log4.ExpenseId = expense2.Id
	timeConfig = time.Now().AddDate(0, 0, -4)
	log4.CreatedAt = timeConfig
	log4.ModifiedAt = timeConfig
	if err := db.Create(&log4).Error; err != nil {
		t.Fatal(err)
	}

	log5 := setupExpenseLogMock(t)
	log5.ExpenseId = expense2.Id
	timeConfig = time.Now().AddDate(0, 0, -3)
	log5.CreatedAt = timeConfig
	log5.ModifiedAt = timeConfig
	if err := db.Create(&log5).Error; err != nil {
		t.Fatal(err)
	}

	log6 := setupExpenseLogMock(t)
	log6.ExpenseId = expense2.Id
	timeConfig = time.Now().AddDate(0, 0, -2)
	log6.CreatedAt = timeConfig
	log6.ModifiedAt = timeConfig
	if err := db.Create(&log6).Error; err != nil {
		t.Fatal(err)
	}

	log7 := setupExpenseLogMock(t)
	log7.ExpenseId = expense2.Id
	timeConfig = time.Now().AddDate(0, 0, -1)
	log7.CreatedAt = timeConfig
	log7.ModifiedAt = timeConfig
	if err := db.Create(&log7).Error; err != nil {
		t.Fatal(err)
	}

	result, err := usecase.ExpenseUsecase.GetExpenseLogs(&context, expense1.Id, lo.ToPtr(time.Now().AddDate(0, 0, -4)), lo.ToPtr(time.Now().AddDate(0, 0, -1)), lo.ToPtr(0), nil, lo.ToPtr(10))

	assert.Nil(t, err)
	assert.Len(t, result, 3)
}

func Test_UpdateExpense_Success(t *testing.T) {
	usecase, db, err := setupTestUsecase()
	if err != nil {
		t.Fatal(err)
	}
	context := createContext("")

	expense1 := setupExpenseMock(t)
	if err := db.Create(&expense1).Error; err != nil {
		t.Fatal(err)
	}

	newRefundStatus, err := mappers.ToExpenseRefundStatusDTO(entities.REFUND_STATUS_PENDING)
	assert.Nil(t, err)

	expenseCategory, err := mappers.ToExpenseCategoryDTO(expense1.Categorization)
	assert.Nil(t, err)

	reviewStatus, err := mappers.ToExpenseReviewStatusDTO(expense1.ReviewStatus)
	assert.Nil(t, err)

	updateExpenseDTO := contracts.UpdateExpenseDTO{
		Categorization: *expenseCategory,
		ExpenseAt:      expense1.ExpenseAt,
		ProgramId:      expense1.ProgramId,
		RefundStatus:   *newRefundStatus,
		ReviewStatus:   *reviewStatus,
		TotalAmount:    666,
	}

	result, err := usecase.ExpenseUsecase.UpdateExpense(&context, expense1.Id, updateExpenseDTO)
	assert.Nil(t, err)

	expenctedDTO := contracts.ExpenseDTO{
		Categorization: *expenseCategory,
		CreatedAt:      expense1.CreatedAt,
		ExpenseAt:      expense1.ExpenseAt,
		Id:             expense1.Id,
		ProgramId:      expense1.ProgramId,
		RefundStatus:   *newRefundStatus,
		ReviewStatus:   *reviewStatus,
		TotalAmount:    updateExpenseDTO.TotalAmount,
	}

	assert.Empty(t, cmp.Diff(expenctedDTO, *result, cmpopts.IgnoreFields(contracts.ExpenseDTO{}, "ModifiedAt")))

	logs, err := usecase.ExpenseUsecase.GetExpenseLogs(&context, expense1.Id, nil, nil, lo.ToPtr(0), lo.ToPtr(""), lo.ToPtr(10))
	assert.Nil(t, err)
	assert.Len(t, logs, 2)
}
