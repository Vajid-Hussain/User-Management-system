package usecase

import (
	"errors"
	"fmt"
	userCase_interfaces "sample/pkg/repository/interface"
	interfaces "sample/pkg/usecase/interface"
	"sample/pkg/utils/models"

	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository userCase_interfaces.AdminRepository
}

func NewAdminUseCase(repo userCase_interfaces.AdminRepository) interfaces.AdminUseCase {
	return &adminUseCase{adminRepository: repo}
}

func (c *adminUseCase) UseAdminLogin(LoginData models.AdminLoginDetails) error {

	LoginFeatchData, err := c.adminRepository.GetAdminData(LoginData)

	if err != nil {
		return errors.New("no admin exist")
	} else {

		if LoginData.Password != LoginFeatchData.Password {
			return errors.New("password is not matched")
		} else {
			return nil
		}

	}
}

func (c *adminUseCase) CreateUser(userData models.UserDetails) error{

	if userData.ConfirmPassword == userData.Password {

		HashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err, "problem at hashing signup")
		}

		userData.Password = string(HashedPassword)
		exist := c.adminRepository.SaveUserData(userData)
		if exist != nil {
			fmt.Println(exist, "at in usecase exist")
			return exist
		}
	} else {
		return errors.New("confirm password is not match")
	}
	return nil

}

func (c *adminUseCase) FullUserData() *[]models.UserData{
	UserCollection:=c.adminRepository.AllUserData()
	return UserCollection
}


func (c *adminUseCase) DeleteUser(UserMail models.UserDelete){
	c.adminRepository.UserDelete(UserMail)
}

func (c *adminUseCase) SingleUserData(userMail models.UserMail)models.UserData{
	return c.adminRepository.SingleUserData(userMail)
}

func (c *adminUseCase)EditUser(Userdata models.UserData){
	c.adminRepository.UserEdit(Userdata)
}