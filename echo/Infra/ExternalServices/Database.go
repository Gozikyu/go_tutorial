package ExternalServices

import (
	"errors"
	"fmt"
	"strconv"
)

type NotValidatedUser struct {
	UserId           int
	Email            string
	UserType         string
	IsEmailConfirmed bool
}

type NotValidatedCompany struct {
	NumberOfEmployees int
	CompanyDomainName string
}

// Database は外部サービス（Database）のダミーを表す構造体です。
type Database struct {
	Users   map[int]map[string]string
	Company map[string]string
}

// DBに登録するダミーユーザーを作成するヘルパー関数
func NewDummyUser(email string, userType string, isEmailConfirmed bool) map[string]string {
	user := make(map[string]string)

	var stringBool string

	if isEmailConfirmed {
		stringBool = "true"
	} else {
		stringBool = "false"
	}

	user["email"] = email
	user["userType"] = userType
	user["isEmailConfirmed"] = stringBool

	return user
}

// DBに登録するダミーカンパニーを作成するヘルパー関数
func NewDummyCompany(numberOfEmployee int, companyDomainName string) map[string]string {
	company := make(map[string]string)

	company["numberOfEmployee"] = strconv.Itoa(numberOfEmployee)
	company["companyDomainName"] = companyDomainName

	return company
}

// GetUserByID は指定されたユーザーIDのユーザー情報を取得します。
func (db *Database) GetUserByID(userId int) (map[string]string, error) {

	user, ok := db.Users[userId]

	if !ok {
		return nil, errors.New("ユーザーが見つかりませんでした。ユーザーID: " + fmt.Sprint(userId))
	}
	return user, nil
}

// SaveUser は新しいユーザー情報を保存します。
func (db *Database) SaveUser(user map[string]string) {
	if db.Users == nil {
		db.Users = make(map[int]map[string]string)
	}
	db.Users[userID(user)] = map[string]string{
		"email":            user["email"],
		"userType":         user["userType"],
		"isEmailConfirmed": user["isEmailConfirmed"],
	}
}

// userID はユーザーマップからユーザーIDを取得します。
func userID(user map[string]string) int {
	id, _ := strconv.Atoi(user["userId"])
	return id
}

func (db *Database) GetCompany() map[string]string {
	company := db.Company
	return company
}

func (db *Database) SaveCompany(newCompany map[string]string) {
	db.Company = newCompany
}

// インメモリのダミーデータベースを作成するヘルパー関数
func SetupTestDatabase() *Database {
	db := Database{
		Users:   make(map[int]map[string]string),
		Company: make(map[string]string),
	}

	db.Users[1] = NewDummyUser("user1@customer.com", "CUSTOMER", true)
	db.Users[2] = NewDummyUser("user2@company.com", "EMPLOYEE", true)
	db.Users[3] = NewDummyUser("user3@customer.com", "CUSTOMER", false)

	db.Company = NewDummyCompany(3, "company.com")

	return &db
}
