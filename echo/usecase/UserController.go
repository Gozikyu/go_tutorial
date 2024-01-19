package UserController

import (
	"errors"
	"fmt"
	"go_tutorial/echo/Domain"
	"go_tutorial/echo/Infra"
)

func UserController(userId int, newEmail string) error {
	fmt.Println("usecase開始")

	validatedNewEmail, error := Domain.NewEmail(newEmail)

	if error != nil {
		return errors.New("メールアドレスとして適切な値ではありません")
	}

	database := Infra.NewDatabase()

	//チェックしていないerrorをどう気づくか
	registeredUser, error := database.GetUserById(userId)
	if error != nil {
		return errors.New("ユーザーの取得に失敗しました")
	}

	user, error := Domain.NewUser(registeredUser.UserId, registeredUser.Email, registeredUser.UserType, registeredUser.IsEmailConfirmed)

	//現在のメールアドレスと更新するメールアドレスが同じ場合は何もせずに終了
	if user.Email == validatedNewEmail {
		return nil
	}

	//本来はDBから取得する
	registeredCompany := database.GetCompany()

	company, error := Domain.NewCompany(registeredCompany.NumberOfEmployees, registeredCompany.CompanyDomainName)

	if error != nil {
		return errors.New("カンパニーの取得に失敗しました")
	}

	//ロジックがあっているか確認
	error = user.UpdateProfile(Domain.Email(newEmail), company)

	if error != nil {
		return errors.New("情報の更新に失敗しました")
	}

	//DBに更新後の保存をする
	database.SaveUser(*user)
	database.SaveCompany(*company)

	//メールを送信する

	fmt.Println("usecase終了")

	return nil
}
