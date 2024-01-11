package handler

import (
	"fmt"
	"net/http"
	"sample/pkg/helper"
	interfaces "sample/pkg/usecase/interface"
	"sample/pkg/utils/models"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUseCase interfaces.AdminUseCase
}

func NewAdminHandler(useCase interfaces.AdminUseCase) *AdminHandler {
	return &AdminHandler{adminUseCase: useCase}
}

func (u *AdminHandler) AdminLogin(c *gin.Context) {
	_, TokenExist := helper.CheckCookie(c)
	if TokenExist {
		c.Redirect(http.StatusFound, "/admin/")
	} else {
		c.HTML(http.StatusOK, "adminlogin.html", nil)
	}
}

func (u *AdminHandler) HandlerPostAdminLogin(c *gin.Context) {
	_, TokenExist := helper.CheckCookie(c)
	if TokenExist {
		c.Redirect(http.StatusFound, "/admin/")
	} else {
		err := c.Request.ParseForm()
		if err != nil {
			fmt.Println(err, "at Admin PostLogin")
		}

		name := c.Request.FormValue("name")
		password := c.Request.FormValue("password")
		LoginData := models.AdminLoginDetails{Name: name, Password: password}

		TokenData := models.GenerateToken{Email: name}

		error := u.adminUseCase.UseAdminLogin(LoginData)
		if error != nil {
			fmt.Println(err)
			c.HTML(http.StatusOK, "adminlogin.html", error)
		} else {
			helper.SetToken(TokenData, c)
			c.Redirect(http.StatusFound, "/admin/")
		}
	}
}

func (u *AdminHandler) HandlerAdminPage(c *gin.Context) {
	var User []models.UserData
	_, TokenExist := helper.CheckCookie(c)
	
	if TokenExist {
		UserCollection := u.adminUseCase.FullUserData()
		for _,data := range *UserCollection {
			User=append(User, data)
		}
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"user": User,
		})
	} else {
		c.Redirect(http.StatusFound, "/admin/login")
	}

}

func (u *AdminHandler) CreateUser(c *gin.Context) {

	_, TokenExist := helper.CheckCookie(c)

	if TokenExist {
		c.HTML(http.StatusOK, "createUser.html", nil)
	} else {
		c.Redirect(http.StatusFound, "/admin/login")
	}
}

func (u *AdminHandler) CreateUserPost(c *gin.Context) {

	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println(err, "at UserSignup")
	}

	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")
	phone := c.Request.FormValue("phone")
	password := c.Request.FormValue("password")
	confirmPassword := c.Request.FormValue("confirmpassword")

	SignupData := models.UserDetails{Name: name, Email: email, Phone: phone, Password: password, ConfirmPassword: confirmPassword}

	fmt.Println(SignupData)
	IsMatch := u.adminUseCase.CreateUser(SignupData)
	if IsMatch != nil {
		c.HTML(http.StatusOK, "createUser.html", IsMatch)
	} else {
		c.Redirect(http.StatusFound, "/admin/")
	}
}

func (u *AdminHandler) AdminLogout(c *gin.Context) {
	helper.DeleteToken(c)
	c.Redirect(http.StatusFound, "/admin/login")
}

func (u *AdminHandler)UserDelete(c *gin.Context){
	err:=c.Request.ParseForm()
	if err != nil {
		fmt.Println(err, "at Delete user")
	}
	email:=c.Request.FormValue("email")
	
	UserDeleteData:=models.UserDelete{Email: email}
	fmt.Println(UserDeleteData)

	u.adminUseCase.DeleteUser(UserDeleteData)
	c.Redirect(http.StatusFound,"/admin/")
}



func (u *AdminHandler) SingleUserData(c *gin.Context){
	email:=c.Request.FormValue("email")
	usermail:=models.UserMail{Email: email}
	SingleUserData:=u.adminUseCase.SingleUserData(usermail)
	fmt.Println(SingleUserData)

	c.HTML(http.StatusOK,"editUser.html", SingleUserData)
}

func (u *AdminHandler) EditUser(c *gin.Context){
	email:=c.Request.FormValue("email")
	name:=c.Request.FormValue("name")
	phone:=c.Request.FormValue("phone")

	UserData:=models.UserData{Name: name, Phone: phone, Email: email}

	u.adminUseCase.EditUser(UserData)
	c.Redirect(http.StatusFound, "/admin/")
}