package models

type AdminLoginDetails struct{
	Name 		string
	Password 	string
}

type UserData struct{
	Name 	string
	Phone 	string
	Email 	string
}

type UserDelete struct{
	Email 	string
}

type UserMail struct{
	Email string
}

