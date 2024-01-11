package interfaces

import "sample/pkg/utils/models"

type UserRepository interface {
	SaveUserData(models.UserDetails) error
	GetUserData(models.UserLoginDetails)(models.UserFeatchData, error)
	RepoGetUserName(string)string
}
