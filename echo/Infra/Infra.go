package Infra

import (
	"go_tutorial/echo/Domain"
	"strconv"
	"strings"
	"sync"
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

type Database struct {
	mu      sync.Mutex
	users   []NotValidatedUser
	company NotValidatedCompany
}

func NewDatabase() *Database {

	users := []NotValidatedUser{
		{UserId: 1, Email: "taro@customer.com", UserType: "CUSTOMER", IsEmailConfirmed: true},
		{UserId: 2, Email: "jiro@employee.com", UserType: "EMPLOYEE", IsEmailConfirmed: false},
		{UserId: 3, Email: "saburo@customer.com", UserType: "CUSTOMER", IsEmailConfirmed: true},
	}

	company := NotValidatedCompany{NumberOfEmployees: 2, CompanyDomainName: "@company.com"}

	return &Database{
		users:   users,
		company: company,
	}
}

func (db *Database) GetUserById(userId int) (NotValidatedUser, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for _, user := range db.users {
		if userId == int(user.UserId) {

			return user, nil
		}
	}
	return NotValidatedUser{}, nil
}

func (db *Database) SaveUser(newUser Domain.User) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, user := range db.users {
		if int(newUser.UserId) == user.UserId {
			db.users[i] = NotValidatedUser{UserId: int(newUser.UserId), Email: string(newUser.Email), UserType: newUser.UserType, IsEmailConfirmed: newUser.IsEmailConfirmed}
			return
		}
	}
}

func (db *Database) GetCompany() NotValidatedCompany {
	db.mu.Lock()
	defer db.mu.Unlock()

	return db.company
}

func (db *Database) SaveCompany(newCompany Domain.Company) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.company = NotValidatedCompany{NumberOfEmployees: int(newCompany.NumberOfEmployees), CompanyDomainName: string(newCompany.CompanyDomainName)}
}

func ConvertStringToBool(value string) bool {
	boolValue, _ := strconv.ParseBool(value)
	return boolValue
}

func ConvertStringToInt(value string) int {
	intValue, _ := strconv.Atoi(value)
	return intValue
}

func SplitEmailDomain(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]

}
