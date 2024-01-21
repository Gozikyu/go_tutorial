package UserController_test

import (
	"go_tutorial/echo/Infra"
	"go_tutorial/echo/Infra/ExternalServices"
	UserController "go_tutorial/echo/usecase"
	"testing"
)

func TestUseController(t *testing.T) {
	db := ExternalServices.Database{
		Users:   make(map[int]map[string]string),
		Company: make(map[string]string),
	}

	db.Users[1] = ExternalServices.NewDummyUser("user1@customer.com", "CUSTOMER", true)
	db.Users[2] = ExternalServices.NewDummyUser("user2@company.com", "EMPLOYEE", true)
	db.Users[3] = ExternalServices.NewDummyUser("user3@customer.com", "CUSTOMER", false)

	db.Company = ExternalServices.NewDummyCompany(3, "@company.com")

	userRepository := Infra.NewUserRepository(db)
	companyRepository := Infra.NewCompanyRepository(db)

	userController := UserController.NewUserController(userRepository, companyRepository)

	err := userController.ChangeEmail(1, "updated@company.com")
	if err != nil {
		t.Errorf("エラーが発生しました。got: %v, err:", err)
	}
}
