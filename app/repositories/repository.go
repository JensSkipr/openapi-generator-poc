package repositories

import (
	"github.com/nightborn-be/blink/skipr-test/app/database"
	"gorm.io/gorm"
)

type Repository struct {
	RepositoryBase
}

func Initialise(db *gorm.DB) Repository {
	// Create get database function
	getDb := func(testId *string) (*gorm.DB, error) {
		if testId == nil || *testId == "" {
			return db, nil
		}
		return database.InitialiseTestFromTestId(*testId)
	}

	return Repository{
		RepositoryBase: RepositoryBase{
			ExpenseRepository: InitialiseExpenseRepository(getDb),
		},
	}
}

func InitialiseTest(db *gorm.DB) Repository {
	// Create get database function
	getDb := func(_ *string) (*gorm.DB, error) {
		return db, nil
	}

	return Repository{
		RepositoryBase: RepositoryBase{
			ExpenseRepository: InitialiseExpenseRepository(getDb),
		},
	}
}
