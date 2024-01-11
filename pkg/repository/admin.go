package repository

import (
	"errors"
	"fmt"
	interfaces "sample/pkg/repository/interface"
	"sample/pkg/utils/models"

	"gorm.io/gorm"
)

type AdminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &AdminDatabase{DB: db}
}

func (c *AdminDatabase) GetAdminData(LoginData models.AdminLoginDetails) (models.AdminLoginDetails, error) {

	var AdminFeatchDetails models.AdminLoginDetails
	query := `SELECT name,password FROM admins WHERE name=$1`

	row := c.DB.Raw(query, LoginData.Name).Row()
	err := row.Scan(&AdminFeatchDetails.Name, &AdminFeatchDetails.Password)
	if err != nil {
		fmt.Println(err, "error at featching  data from database `GetUserData`")
	}

	if AdminFeatchDetails.Name == "" {
		return AdminFeatchDetails, errors.New("no admin")
	}

	return AdminFeatchDetails, nil
}

func (c *AdminDatabase) SaveUserData(userData models.UserDetails) error{
	var name string
	
	query1:="SELECT name FROM users WHERE email=$1"
	row:=c.DB.Raw(query1,userData.Email).Row()
	err:=row.Scan(&name)
	if err != nil {
		fmt.Println(err, "error at inserting of data to database `SaveUserData`")
	}

	if name!=""{
		fmt.Println("alrady account exist from gmail")
		return errors.New("email contain a account")
	}else{
		query:=`INSERT INTO Users (name, email, phone, password) VALUES($1, $2, $3, $4)`
		result:=c.DB.Exec(query, userData.Name, userData.Email, userData.Phone, userData.Password)
		if result != nil {
			fmt.Println(result, "error at inserting of data to database `SaveUserData`")
		}
	}
	return nil
}

func (c *AdminDatabase) AllUserData() *[]models.UserData{

	var user []models.UserData

	query:="SELECT name, email, phone FROM users"
	rows, err:=c.DB.Raw(query).Rows()
	if err!=nil{
		fmt.Println(err,"error at fetchin user data")
	}

	for rows.Next() {
		var u models.UserData
		err:=rows.Scan(&u.Name, &u.Email, &u.Phone)
		if err!=nil{
			fmt.Println(err, "error at rows scan ")
		}
		user = append(user, u )
	}
	return &user
}

func (c *AdminDatabase) UserDelete(UserMail models.UserDelete){
	query:="DELETE FROM users WHERE email=?"
	c.DB.Raw(query,UserMail.Email).Row()
}

func (c *AdminDatabase) SingleUserData(UserMail models.UserMail)models.UserData{
	var userData models.UserData
	query:="SELECT name,email,phone FROM users WHERE email=?"
	row:=c.DB.Raw(query, UserMail.Email).Row()
	err:=row.Scan(&userData.Name, &userData.Email, &userData.Phone)
	if err!=nil{
		fmt.Println(err, "error at rows scan single user data")
	}
	fmt.Println(userData)
	return userData
}

func (c *AdminDatabase) UserEdit( EditedUser models.UserData){
	query := "UPDATE users SET name=?, phone=? WHERE email=?"
	c.DB.Exec(query, EditedUser.Name, EditedUser.Phone, EditedUser.Email)
}