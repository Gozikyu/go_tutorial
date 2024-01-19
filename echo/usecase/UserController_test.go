package UserController_test

import (
	UserController "go_tutorial/echo/usecase"
	"testing"
)

func TestUseController(t *testing.T) {
	err := UserController.UserController(1, "update@employee.com")
	if err != nil {
		t.Errorf("エラーが発生しました。got: %v, expected: nil", err)
	}
}
