package usecase

import (
	"errors"
	"go_tutorial/echo/Domain"
)

func usecase(userId int, newEmail string) error {

	//本来はDBから取得する
	user, error := Domain.GenUser(1, "hoge@customer.com", "CUSTOMER", true)

	if error != nil {
		return errors.New("ユーザーの取得に失敗しました")
	}

	//現在のメールアドレスと更新するメールアドレスが同じ場合は何もせずに終了
	if user.Email == Domain.Email(newEmail) {
		return nil
	}

	//本来はDBから取得する
	company := Domain.Company{NumberOfEmployees: 3, CompanyDomainName: "@employee"}

	//ロジックがあっているか確認
	error = user.UpdateProfile(Domain.Email(newEmail), &company)

	if error != nil {
		return errors.New("情報の更新に失敗しました")
	}

	//DBに更新後の保存をする

	//メールを送信する
	return nil
}
