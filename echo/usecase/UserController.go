package UserController

import (
	"errors"
	"fmt"
	"go_tutorial/echo/Domain"
)

type userRepository Domain.IUserRepository
type companyRepository Domain.ICompanyRepository

type UserControllerType struct {
	userRepository    Domain.IUserRepository
	companyRepository Domain.ICompanyRepository
}

func NewUserController(ur Domain.IUserRepository, cr Domain.ICompanyRepository) UserControllerType {
	return UserControllerType{userRepository: ur, companyRepository: cr}
}

func (u *UserControllerType) ChangeEmail(userId int, newEmail string) error {
	fmt.Println("usecase開始")

	validatedNewEmail, error := Domain.NewEmail(newEmail)

	if error != nil {
		return errors.New("メールアドレスとして適切な値ではありません")
	}

	DomainUserId, err := Domain.NewUserId(userId)
	if err != nil {
		return errors.New("userIdの変換に失敗しました")
	}

	//チェックしていないerrorをどう気づくか
	user, err := u.userRepository.GetUserById(DomainUserId)
	if err != nil {
		return errors.New("userの取得に失敗しました")
	}

	//現在のメールアドレスと更新するメールアドレスが同じ場合は何もせずに終了
	if user.Email == validatedNewEmail {
		return nil
	}

	//本来はDBから取得する
	company, err := u.companyRepository.GetCompany()
	if err != nil {
		return errors.New("companyの取得に失敗しました")
	}

	if error != nil {
		return errors.New("カンパニーの取得に失敗しました")
	}

	//ロジックがあっているか確認
	error = user.UpdateProfile(Domain.Email(newEmail), &company)

	if error != nil {
		return errors.New("情報の更新に失敗しました")
	}

	//DBに更新後の保存をする
	u.userRepository.SaveUser(user)
	u.companyRepository.SaveCompany(company)

	//メールを送信する

	fmt.Println("usecase終了")

	return nil

}

func NewUserControllerType(ur userRepository, cr companyRepository) *UserControllerType {
	return &UserControllerType{userRepository: ur, companyRepository: cr}
}
