package UserController_test

import (
	"go_tutorial/echo/Infra"
	"go_tutorial/echo/Infra/ExternalServices"
	"go_tutorial/echo/UserController"
	"reflect"
	"testing"
)

func SetupTestDatabase() *ExternalServices.Database {
	db := ExternalServices.Database{
		Users:   make(map[int]map[string]string),
		Company: make(map[string]string),
	}

	db.Users[1] = ExternalServices.NewDummyUser("user1@customer.com", "CUSTOMER", true)
	db.Users[2] = ExternalServices.NewDummyUser("user2@company.com", "EMPLOYEE", true)
	db.Users[3] = ExternalServices.NewDummyUser("user3@customer.com", "CUSTOMER", false)

	db.Company = ExternalServices.NewDummyCompany(3, "company.com")

	return &db
}

func TestMain(t *testing.T) {
	t.Run("customerからcompanyへの変更", func(t *testing.T) {
		db := SetupTestDatabase()

		userRepository := Infra.NewUserRepository(db)
		companyRepository := Infra.NewCompanyRepository(db)
		userController := UserController.NewUserController(userRepository, companyRepository)

		err := userController.ChangeEmail(1, "updated@company.com")
		gotUser := db.Users[1]
		gotCompany := db.Company

		wantUser := map[string]string{
			"email":            "updated@company.com",
			"userType":         "EMPLOYEE",
			"isEmailConfirmed": "false",
		}
		wantCompany := map[string]string{
			"numberOfEmployee":  "4",
			"companyDomainName": "company.com",
		}

		if err != nil {
			t.Errorf("エラーが発生しました。err: %v", err)
		}

		if !reflect.DeepEqual(gotUser, wantUser) {
			t.Errorf("ユーザーが正しく更新できていません。got: %v, want: %v", gotUser, wantUser)
		}
		if !reflect.DeepEqual(gotCompany, wantCompany) {
			t.Errorf("カンパニーが正しく更新できていません。got: %v, want: %v", gotCompany, wantCompany)
		}
	})

	t.Run("customerから変化なし", func(t *testing.T) {
		db := SetupTestDatabase()

		userRepository := Infra.NewUserRepository(db)
		companyRepository := Infra.NewCompanyRepository(db)
		userController := UserController.NewUserController(userRepository, companyRepository)

		err := userController.ChangeEmail(1, "updated@hoge.com")
		gotUser := db.Users[1]
		wantUser := map[string]string{
			"email":            "updated@hoge.com",
			"userType":         "CUSTOMER",
			"isEmailConfirmed": "false",
		}
		gotCompany := db.Company
		wantCompany := map[string]string{
			"numberOfEmployee":  "3",
			"companyDomainName": "company.com",
		}

		if err != nil {
			t.Errorf("エラーが発生しました。err: %v", err)
		}

		if !reflect.DeepEqual(gotUser, wantUser) {
			t.Errorf("ユーザーが正しく更新できていません。got: %v, want: %v", gotUser, wantUser)
		}

		if !reflect.DeepEqual(gotCompany, wantCompany) {
			t.Errorf("カンパニーが正しく更新できていません。got: %v, want: %v", gotUser, wantUser)
		}
	})

	t.Run("companyからcustomerへの変化", func(t *testing.T) {
		db := SetupTestDatabase()

		userRepository := Infra.NewUserRepository(db)
		companyRepository := Infra.NewCompanyRepository(db)
		userController := UserController.NewUserController(userRepository, companyRepository)

		err := userController.ChangeEmail(2, "updated@hoge.com")
		gotUser := db.Users[2]
		wantUser := map[string]string{
			"email":            "updated@hoge.com",
			"userType":         "CUSTOMER",
			"isEmailConfirmed": "false",
		}
		gotCompany := db.Company
		wantCompany := map[string]string{
			"numberOfEmployee":  "2",
			"companyDomainName": "company.com",
		}

		if err != nil {
			t.Errorf("エラーが発生しました。err: %v", err)
		}

		if !reflect.DeepEqual(gotUser, wantUser) {
			t.Errorf("ユーザーが正しく更新できていません。got: %v, want: %v", gotUser, wantUser)
		}

		if !reflect.DeepEqual(gotCompany, wantCompany) {
			t.Errorf("カンパニーが正しく更新できていません。got: %v, want: %v", gotCompany, wantCompany)
		}
	})

	t.Run("メールアドレスが既存と同じ", func(t *testing.T) {
		db := SetupTestDatabase()

		userRepository := Infra.NewUserRepository(db)
		companyRepository := Infra.NewCompanyRepository(db)
		userController := UserController.NewUserController(userRepository, companyRepository)

		err := userController.ChangeEmail(1, "user1@customer.com")
		gotUser := db.Users[1]
		wantUser := map[string]string{
			"email":            "user1@customer.com",
			"userType":         "CUSTOMER",
			"isEmailConfirmed": "true",
		}
		gotCompany := db.Company
		wantCompany := map[string]string{
			"numberOfEmployee":  "3",
			"companyDomainName": "company.com",
		}

		if err != nil {
			t.Errorf("エラーが発生しました。err: %v", err)
		}

		if !reflect.DeepEqual(gotUser, wantUser) {
			t.Errorf("ユーザーが正しく更新できていません。got: %v, want: %v", gotUser, wantUser)
		}

		if !reflect.DeepEqual(gotCompany, wantCompany) {
			t.Errorf("カンパニーが正しく更新できていません。got: %v, want: %v", gotCompany, wantCompany)
		}
	})

}
