package interfaces

import "sample/pkg/utils/models"

type UserUseCase interface {
	UseUserSignUp(models.UserDetails) error
	UseUserLogin(models.UserLoginDetails) (error)
	UseUserName(string)string
}
