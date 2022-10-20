/* This file is auto-generated, manual edits in this file will be overwritten! */
package controllers

import "github.com/nightborn-be/blink/skipr-test/app/usecases"

type Controller struct {
ExpenseController ExpenseController
}

func Initialise(usecases usecases.Usecase) Controller {
	return Controller{
ExpenseController: InitialiseExpenseController(usecases),
	}
}