package usecases

import (
	"github.com/nightborn-be/blink/skipr-test/app/gateways"
	"github.com/nightborn-be/blink/skipr-test/app/repositories"
)

type Usecase struct {
	UsecaseBase
}

func Initialise(repository *repositories.Repository, gateway *gateways.Gateway) Usecase {
	return Usecase{
		UsecaseBase: UsecaseBase{
			ExpenseUsecase: InitialiseExpenseUsecase(repository, gateway),
		},
	}
}
