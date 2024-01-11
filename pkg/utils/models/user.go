package models

type UserDetails struct{
	Name 			string `json:"name"`
	Email 			string `json:"email"`
	Phone 			string `json:"phone"`
	Password		string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}

type UserLoginDetails struct{
	Email 		string	
	Password 	string
}

type UserFeatchData struct{
	Name	 	string
	Email 		string	
	Password 	string
}

type GenerateToken struct{
	Email string
}