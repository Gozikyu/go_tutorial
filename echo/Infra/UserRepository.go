package Infra

import (
	"errors"
	"go_tutorial/echo/Domain"
	"go_tutorial/echo/Infra/ExternalServices"
	"strconv"
)

// UserRepository は IUserRepository インターフェースを実装するデータアクセスクラスです。
type UserRepository struct {
	Database *ExternalServices.Database
}

// NewUserRepository は新しい UserRepository インスタンスを作成します。
func NewUserRepository(db *ExternalServices.Database) *UserRepository {
	return &UserRepository{
		Database: db,
	}
}

// GetUserByID は指定されたユーザーIDのユーザー情報を取得します。
func (ur *UserRepository) GetUserById(userId Domain.UserId) (Domain.User, error) {
	userMap, err := ur.Database.GetUserByID(int(userId))
	if err != nil {
		return Domain.User{}, err
	}

	user, err := convertToUser(map[string]string{
		"userId":           strconv.Itoa(int(userId)),
		"email":            userMap["email"],
		"userType":         userMap["userType"],
		"isEmailConfirmed": userMap["isEmailConfirmed"],
	})

	if err != nil {
		return Domain.User{}, errors.New("Userへの変換に失敗しました")
	}

	return user, nil
}

// SaveUser は新しいユーザー情報を保存します。
func (ur *UserRepository) SaveUser(newUser Domain.User) {
	userMap := convertToUserMap(newUser)
	ur.Database.SaveUser(userMap)
}

// convertToUser はユーザーマップを User 構造体に変換します。
func convertToUser(userMap map[string]string) (Domain.User, error) {

	userId, err := strconv.Atoi(userMap["userId"])
	if err != nil {
		return Domain.User{}, errors.New("userIdをintに変換できませんでした")
	}

	id, err := Domain.NewUserId(userId)
	if err != nil {
		return Domain.User{}, errors.New("ドメインのuserId型への変換に失敗しました")
	}

	return Domain.User{
		UserId:           id,
		Email:            Domain.Email(userMap["email"]),
		UserType:         userMap["userType"],
		IsEmailConfirmed: convertStringToBool(userMap["isEmailConfirmed"]),
	}, nil
}

func convertStringToBool(isEmailConfirmed string) bool {
	return isEmailConfirmed == "true"
}

// convertToUserMap は User 構造体をユーザーマップに変換します。
func convertToUserMap(newUser Domain.User) map[string]string {
	userMap := make(map[string]string)
	userMap["userId"] = strconv.Itoa(int(newUser.UserId))
	userMap["email"] = string(newUser.Email)
	userMap["userType"] = newUser.UserType
	userMap["isEmailConfirmed"] = convertBoolToString(newUser.IsEmailConfirmed)
	return userMap
}

// convertBoolToString はブール値を文字列に変換します。
func convertBoolToString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}
