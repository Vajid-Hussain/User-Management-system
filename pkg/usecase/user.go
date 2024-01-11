package usecase

import (
	"errors"
	"fmt"
	interfaces "sample/pkg/repository/interface"
	interfacesUseCase "sample/pkg/usecase/interface"
	"sample/pkg/utils/models"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewuserUseCase(repo interfaces.UserRepository) interfacesUseCase.UserUseCase {
	return &userUseCase{userRepo: repo}
}


func (c *userUseCase) UseUserSignUp(userData models.UserDetails) error {
	if userData.ConfirmPassword == userData.Password {

		HashedPassword,err:=bcrypt.GenerateFromPassword([]byte(userData.Password),bcrypt.DefaultCost)
		if err!=nil{
			fmt.Println(err,"problem at hashing signup")
		}

		userData.Password=string(HashedPassword)
		exist:=c.userRepo.SaveUserData(userData)
		if exist!=nil{
			fmt.Println(exist,"at in usecase exist")
			return exist
		}
	} else {
		return errors.New("confirm password is not match")
	}
	return nil
} 

func (c *userUseCase) UseUserLogin(LoginData models.UserLoginDetails) (error){

	LoginFeatchData, err:=c.userRepo.GetUserData(LoginData)

	if err!=nil{
		return errors.New("no user exist")
	}else{

		err:=bcrypt.CompareHashAndPassword([]byte(LoginFeatchData.Password),[]byte(LoginData.Password))
		if err!=nil{
			return errors.New("password is not matched")
		}else{
			return nil
		}

	}
}

func (c *userUseCase) UseUserName(UserId string)string{
	name:=c.userRepo.RepoGetUserName(UserId)
	return name
}