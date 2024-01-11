package interfaces

import "sample/pkg/utils/models"

type AdminRepository interface{
	GetAdminData(models.AdminLoginDetails) (models.AdminLoginDetails, error)
	SaveUserData(models.UserDetails) error
	AllUserData()*[]models.UserData
	UserDelete(models.UserDelete)
	SingleUserData(models.UserMail)models.UserData
	UserEdit(models.UserData)
}