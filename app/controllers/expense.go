/* This file is auto-generated, manual edits in this file will be overwritten! */
package controllers

import (
	"net/http"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/nightborn-be/blink/skipr-test/app/usecases"
"errors"
"github.com/google/uuid"
"github.com/nightborn-be/blink/skipr-test/app/controllers/contracts"
"github.com/nightborn-be/blink/skipr-test/app/contexts"
)

type ExpenseController struct {
	usecases usecases.Usecase
}

func InitialiseExpenseController(usecases usecases.Usecase) ExpenseController {
	return ExpenseController{usecases: usecases}
}


func (controller ExpenseController) GetExpenses(c *gin.Context) {
context, err := contexts.GetContext(c)
if err != nil {
	sentry.CaptureException(err)
	c.IndentedJSON(http.StatusBadRequest, err.Error())
	return
}








// Call the usecase
response, err := controller.usecases.ExpenseUsecase.GetExpenses(context)
if err != nil {
    sentry.CaptureException(err)
    c.IndentedJSON(http.StatusBadRequest, err.Error())
    return
}

// Success
c.IndentedJSON(http.StatusOK, response)
}


func (controller ExpenseController) CreateExpense(c *gin.Context) {
context, err := contexts.GetContext(c)
if err != nil {
	sentry.CaptureException(err)
	c.IndentedJSON(http.StatusBadRequest, err.Error())
	return
}






var createExpenseDTO contracts.CreateExpenseDTO
if err := c.BindJSON(&createExpenseDTO); err != nil {
	sentry.CaptureException(err)
	c.IndentedJSON(http.StatusBadRequest, err.Error())
	return
}
if createExpenseDTO.ExpenseAt.IsZero() {
	c.IndentedJSON(http.StatusBadRequest, errors.New(MISSING_BODY_PARAM+"expenseAt").Error())
	return
}
if createExpenseDTO.ProgramId == uuid.Nil {
	c.IndentedJSON(http.StatusBadRequest, errors.New(MISSING_BODY_PARAM+"programId").Error())
	return
}
if createExpenseDTO.TotalAmount == 0 {
	c.IndentedJSON(http.StatusBadRequest, errors.New(MISSING_BODY_PARAM+"totalAmount").Error())
	return
}

// Call the usecase
response, err := controller.usecases.ExpenseUsecase.CreateExpense(context, createExpenseDTO)
if err != nil {
    sentry.CaptureException(err)
    c.IndentedJSON(http.StatusBadRequest, err.Error())
    return
}

// Success
c.IndentedJSON(http.StatusOK, *response)
}


func (controller ExpenseController) GetExpense(c *gin.Context) {
context, err := contexts.GetContext(c)
if err != nil {
	sentry.CaptureException(err)
	c.IndentedJSON(http.StatusBadRequest, err.Error())
	return
}


expenseId := c.Params.ByName("expenseId")
if expenseId == "" {
	c.IndentedJSON(http.StatusBadRequest, errors.New(MISSING_PATH_PARAM+"expenseId").Error())
	return
}
_expenseId, err := uuid.Parse(expenseId)
if err != nil {
    c.IndentedJSON(http.StatusBadRequest, errors.New(UNABLE_TO_PARSE_PARAM+"expenseId").Error())
    return
}






// Call the usecase
response, err := controller.usecases.ExpenseUsecase.GetExpense(context, _expenseId)
if err != nil {
    sentry.CaptureException(err)
    c.IndentedJSON(http.StatusBadRequest, err.Error())
    return
}

// Success
c.IndentedJSON(http.StatusOK, *response)
}


func (controller ExpenseController) UpdateExpense(c *gin.Context) {
context, err := contexts.GetContext(c)
if err != nil {
	sentry.CaptureException(err)
	c.IndentedJSON(http.StatusBadRequest, err.Error())
	return
}


expenseId := c.Params.ByName("expenseId")
if expenseId == "" {
	c.IndentedJSON(http.StatusBadRequest, errors.New(MISSING_PATH_PARAM+"expenseId").Error())
	return
}
_expenseId, err := uuid.Parse(expenseId)
if err != nil {
    c.IndentedJSON(http.StatusBadRequest, errors.New(UNABLE_TO_PARSE_PARAM+"expenseId").Error())
    return
}




var updateExpenseDTO contracts.UpdateExpenseDTO
if err := c.BindJSON(&updateExpenseDTO); err != nil {
	sentry.CaptureException(err)
	c.IndentedJSON(http.StatusBadRequest, err.Error())
	return
}
if updateExpenseDTO.ExpenseAt.IsZero() {
	c.IndentedJSON(http.StatusBadRequest, errors.New(MISSING_BODY_PARAM+"expenseAt").Error())
	return
}
if updateExpenseDTO.ProgramId == uuid.Nil {
	c.IndentedJSON(http.StatusBadRequest, errors.New(MISSING_BODY_PARAM+"programId").Error())
	return
}
if updateExpenseDTO.TotalAmount == 0 {
	c.IndentedJSON(http.StatusBadRequest, errors.New(MISSING_BODY_PARAM+"totalAmount").Error())
	return
}

// Call the usecase
response, err := controller.usecases.ExpenseUsecase.UpdateExpense(context, _expenseId, updateExpenseDTO)
if err != nil {
    sentry.CaptureException(err)
    c.IndentedJSON(http.StatusBadRequest, err.Error())
    return
}

// Success
c.IndentedJSON(http.StatusOK, *response)
}