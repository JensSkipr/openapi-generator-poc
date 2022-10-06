package usecases

import (
	"github.com/nightborn-be/blink/skipr-test/app/database"
	"github.com/nightborn-be/blink/skipr-test/app/gateways"
	"github.com/nightborn-be/blink/skipr-test/app/repositories"
	"github.com/nightborn-be/blink/skipr-test/app/utils"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

// Setup test database
func setupTestUsecase() (*Usecase, *gorm.DB, error) {
	// Initialise faker generators
	utils.InitialiseFakerGenerators()

	db, err := database.InitialiseTest()
	if err != nil {
		return nil, nil, err
	}

	// Setup in memory db for repositories & in-memory gateways
	testRepository := repositories.InitialiseTest(db)
	testGateway, _ := gateways.InitialiseTest()

	// Setup usecases with repositories and gateways
	return lo.ToPtr(Initialise(&testRepository, &testGateway)), db, nil
}

// func createContext(sub string) contexts.Context {
// 	return contexts.Context{
// 		ContextBase: contexts.ContextBase{
//			Sub:    &sub,
//			Roles:  nil,
//			TestId: nil,
//		},
//	}
// }

// func createAdminContext(sub string) contexts.Context {
//	return contexts.Context{
//		ContextBase: contexts.ContextBase{
//			Sub:    &sub,
//			Roles:  []string{"admin"},
//			TestId: nil,
//		},
//	}
// }
