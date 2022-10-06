package gateways

// ========================================
// =          PROVIDER AND MOCKS          =
// ========================================

type Mocks struct{}

func InitialiseTest() (Gateway, *Mocks) {
	mocks := Mocks{}

	return Gateway{}, &mocks
}
