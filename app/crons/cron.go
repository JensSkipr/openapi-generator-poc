package crons

import "github.com/nightborn-be/blink/skipr-test/app/usecases"

type CronJob struct{}

func Initialise(usecase usecases.Usecase) CronJob {
	return CronJob{}
}

func (jobs CronJob) Run() error {
	return nil
}
