package Infra

import (
	"errors"
	"go_tutorial/echo/Domain"
	"go_tutorial/echo/Infra/ExternalServices"
	"strconv"
)

type CompanyRepository struct {
	Database *ExternalServices.Database
}

func NewCompanyRepository(db *ExternalServices.Database) *CompanyRepository {
	return &CompanyRepository{
		Database: db,
	}
}

func (cr *CompanyRepository) GetCompany() (Domain.Company, error) {
	companyMap := cr.Database.GetCompany()

	company, err := convertToCompany(companyMap)
	if err != nil {
		return Domain.Company{}, errors.New("Companyへの変換に失敗しました")
	}

	return company, nil
}

func (cr *CompanyRepository) SaveCompany(company Domain.Company) {
	cr.Database.SaveCompany(convertToCompanyMap(company))
}

func convertToCompany(company map[string]string) (Domain.Company, error) {

	numberOfEmployee, err := strconv.Atoi(company["numberOfEmployee"])
	if err != nil {
		return Domain.Company{}, errors.New("numberOfEmployeeをintに変換できませんでした")
	}

	domainCompany, err := Domain.NewCompany(numberOfEmployee, company["companyDomainName"])
	if err != nil {
		return Domain.Company{}, errors.New("ドメインへの変換に失敗しました")
	}

	return *domainCompany, nil
}

func convertToCompanyMap(company Domain.Company) map[string]string {
	companyMap := make(map[string]string)
	companyMap["numberOfEmployee"] = strconv.Itoa(int(company.NumberOfEmployees))
	companyMap["companyDomainName"] = string(company.CompanyDomainName)

	return companyMap
}
