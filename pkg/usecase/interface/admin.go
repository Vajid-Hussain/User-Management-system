package interfaces

import "sample/pkg/utils/models"

type AdminUseCase interface{
	UseAdminLogin(models.AdminLoginDetails) (error)
	CreateUser(models.UserDetails) error
	FullUserData()*[]models.UserData
	DeleteUser(models.UserDelete)
	SingleUserData(models.UserMail)models.UserData
	EditUser(models.UserData)
}