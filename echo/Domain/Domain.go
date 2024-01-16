package Domain

import (
	"errors"
	"strings"
)

type UserId int

func GenUserId(userId int) (UserId, error) {
	if userId <= 0 {
		return 0, errors.New("ユーザーIDは1以上である必要があります")
	}
	return UserId(userId), nil
}

type Email string

func GenEmail(email string) (Email, error) {
	if len(email) == 0 {
		return "", errors.New("メールアドレスには値が必要です")
	}
	return Email(email), nil
}

func GenUser(userId UserId, email Email, userType string, isEmailConfirmed bool) (User, error) {
	return User{
		UserId:           userId,
		Email:            email,
		UserType:         userType,
		IsEmailConfirmed: isEmailConfirmed,
	}, nil
}

type User struct {
	UserId           UserId
	Email            Email
	UserType         string
	IsEmailConfirmed bool
}

func (u *User) UpdateProfile(email Email, company *Company) error {
	if !u.IsEmailConfirmed {
		return errors.New("メールアドレスが認証済みである必要があります")
	}

	newType, error := company.getUserTypeByEmail(u.Email)
	if error != nil {
		return errors.New("アップデートに失敗しました")
	}

	u.UserType = newType

	if newType == u.UserType {
		return nil
	}

	if newType == "EMPLOYEE" {
		company.NumberOfEmployees++
	}
	if newType == "CUSTOMER" {
		company.NumberOfEmployees++
	}
	return nil
}

type NumberOfEmployees int
type CompanyDomainName string

func GenNumberOfEmployees(number int) (NumberOfEmployees, error) {
	if number < 0 {
		return 0, errors.New("従業員数として適切な値ではありません")
	}
	return NumberOfEmployees(number), nil
}

func GenCompanyDomainName(name string) (CompanyDomainName, error) {
	if !strings.Contains(name, "@") {
		return "", errors.New("ドメイン名には@が含まれている必要があります")
	}
	return CompanyDomainName(name), nil
}

// Companyドメイン
type Company struct {
	NumberOfEmployees NumberOfEmployees
	CompanyDomainName CompanyDomainName
}

func (c *Company) IncreaseEmployeeCount() {
	c.NumberOfEmployees++
}

func (c *Company) DecreaseEmployeeCount() {
	c.NumberOfEmployees--
}

func (c Company) getUserTypeByEmail(email Email) (string, error) {
	if email == "" {
		return "", errors.New("メールアドレスには値が必要です")
	}

	emailDomain := CompanyDomainName(strings.Split(string(email), "@")[1])
	if emailDomain == c.CompanyDomainName {
		return "EMPLOYEE", nil
	}

	return "CUSTOMER", nil
}
