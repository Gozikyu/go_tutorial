package Domain

import (
	"errors"
	"strings"
)

type UserId int

func NewUserId(userId int) (UserId, error) {
	if userId <= 0 {
		return 0, errors.New("ユーザーIDは1以上である必要があります")
	}
	return UserId(userId), nil
}

type Email string

func NewEmail(email string) (Email, error) {
	if len(email) == 0 {
		return "", errors.New("メールアドレスには値が必要です")
	}
	return Email(email), nil
}

func NewUser(userId int, email string, userType string, isEmailConfirmed bool) (*User, error) {

	validatedUserId, error := NewUserId(userId)
	if error != nil {
		return nil, errors.New("ユーザーIDとして適切な値ではありません")
	}

	validatedEmail, error := NewEmail(email)
	if error != nil {
		return nil, errors.New("メールアドレスとして適切な値ではありません")
	}

	return &User{
		UserId:           validatedUserId,
		Email:            validatedEmail,
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

	newType, error := company.getUserTypeByEmail(email)
	if error != nil {
		return errors.New("アップデートに失敗しました")
	}

	u.Email = email

	if newType == u.UserType {
		return nil
	}

	u.UserType = newType

	if newType == "EMPLOYEE" {
		company.NumberOfEmployees++
	}
	if newType == "CUSTOMER" {
		company.NumberOfEmployees--
	}
	return nil
}

type NumberOfEmployees int
type CompanyDomainName string

func NewNumberOfEmployees(number int) (NumberOfEmployees, error) {
	if number < 0 {
		return 0, errors.New("従業員数として適切な値ではありません")
	}
	return NumberOfEmployees(number), nil
}

func NewCompanyDomainName(name string) (CompanyDomainName, error) {
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

func NewCompany(numberOfEmployees int, companyDomainName string) (*Company, error) {
	validatedNumberOfEmployees, error := NewNumberOfEmployees(numberOfEmployees)
	if error != nil {
		return nil, errors.New("従業員数として適切な値ではありません")
	}

	validatedCompanyDomainName, error := NewCompanyDomainName(companyDomainName)
	if error != nil {
		return nil, errors.New("ドメイン名として適切な値ではありません")
	}

	return &Company{NumberOfEmployees: validatedNumberOfEmployees, CompanyDomainName: validatedCompanyDomainName}, nil
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

// インターフェイス
type IUserRepository interface {
	GetUserById(userId UserId) (User, error)
	SaveUser(user User)
}

type ICompanyRepository interface {
	GetCompany() (Company, error)
	SaveCompany(company Company)
}
